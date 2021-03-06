package autorest

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/mocks"
)

const (
	TestTenantID                = "TestTenantID"
	TestActiveDirectoryEndpoint = "https://login/test.com/"
)

func TestWithAuthorizer(t *testing.T) {
	r1 := mocks.NewRequest()

	na := &NullAuthorizer{}
	r2, err := Prepare(r1,
		na.WithAuthorization())
	if err != nil {
		t.Fatalf("autorest: NullAuthorizer#WithAuthorization returned an unexpected error (%v)", err)
	} else if !reflect.DeepEqual(r1, r2) {
		t.Fatalf("autorest: NullAuthorizer#WithAuthorization modified the request -- received %v, expected %v", r2, r1)
	}
}

func TestTokenWithAuthorization(t *testing.T) {
	token := &adal.Token{
		AccessToken: "TestToken",
		Resource:    "https://azure.microsoft.com/",
		Type:        "Bearer",
	}

	ba := NewBearerAuthorizer(token)
	req, err := Prepare(&http.Request{}, ba.WithAuthorization())
	if err != nil {
		t.Fatalf("azure: BearerAuthorizer#WithAuthorization returned an error (%v)", err)
	} else if req.Header.Get(http.CanonicalHeaderKey("Authorization")) != fmt.Sprintf("Bearer %s", token.AccessToken) {
		t.Fatal("azure: BearerAuthorizer#WithAuthorization failed to set Authorization header")
	}
}

func TestServicePrincipalTokenWithAuthorizationNoRefresh(t *testing.T) {
	oauthConfig, err := adal.NewOAuthConfig(TestActiveDirectoryEndpoint, TestTenantID)
	if err != nil {
		t.Fatalf("azure: BearerAuthorizer#WithAuthorization returned an error (%v)", err)
	}
	spt, err := adal.NewServicePrincipalToken(*oauthConfig, "id", "secret", "resource", nil)
	if err != nil {
		t.Fatalf("azure: BearerAuthorizer#WithAuthorization returned an error (%v)", err)
	}
	spt.SetAutoRefresh(false)
	s := mocks.NewSender()
	spt.SetSender(s)

	ba := NewBearerAuthorizer(spt)
	req, err := Prepare(mocks.NewRequest(), ba.WithAuthorization())
	if err != nil {
		t.Fatalf("azure: BearerAuthorizer#WithAuthorization returned an error (%v)", err)
	} else if req.Header.Get(http.CanonicalHeaderKey("Authorization")) != fmt.Sprintf("Bearer %s", spt.AccessToken) {
		t.Fatal("azure: BearerAuthorizer#WithAuthorization failed to set Authorization header")
	}
}

func TestServicePrincipalTokenWithAuthorizationRefresh(t *testing.T) {

	oauthConfig, err := adal.NewOAuthConfig(TestActiveDirectoryEndpoint, TestTenantID)
	if err != nil {
		t.Fatalf("azure: BearerAuthorizer#WithAuthorization returned an error (%v)", err)
	}
	refreshed := false
	spt, err := adal.NewServicePrincipalToken(*oauthConfig, "id", "secret", "resource", func(t adal.Token) error {
		refreshed = true
		return nil
	})
	if err != nil {
		t.Fatalf("azure: BearerAuthorizer#WithAuthorization returned an error (%v)", err)
	}

	jwt := `{
		"access_token" : "accessToken",
		"expires_in"   : "3600",
		"expires_on"   : "test",
		"not_before"   : "test",
		"resource"     : "test",
		"token_type"   : "Bearer"
	}`
	body := mocks.NewBody(jwt)
	resp := mocks.NewResponseWithBodyAndStatus(body, http.StatusOK, "OK")
	c := mocks.NewSender()
	s := DecorateSender(c,
		(func() SendDecorator {
			return func(s Sender) Sender {
				return SenderFunc(func(r *http.Request) (*http.Response, error) {
					return resp, nil
				})
			}
		})())
	spt.SetSender(s)

	ba := NewBearerAuthorizer(spt)
	req, err := Prepare(mocks.NewRequest(), ba.WithAuthorization())
	if err != nil {
		t.Fatalf("azure: BearerAuthorizer#WithAuthorization returned an error (%v)", err)
	} else if req.Header.Get(http.CanonicalHeaderKey("Authorization")) != fmt.Sprintf("Bearer %s", spt.AccessToken) {
		t.Fatal("azure: BearerAuthorizer#WithAuthorization failed to set Authorization header")
	}

	if !refreshed {
		t.Fatal("azure: BearerAuthorizer#WithAuthorization must refresh the token")
	}
}

func TestServicePrincipalTokenWithAuthorizationReturnsErrorIfConnotRefresh(t *testing.T) {
	oauthConfig, err := adal.NewOAuthConfig(TestActiveDirectoryEndpoint, TestTenantID)
	if err != nil {
		t.Fatalf("azure: BearerAuthorizer#WithAuthorization returned an error (%v)", err)
	}
	spt, err := adal.NewServicePrincipalToken(*oauthConfig, "id", "secret", "resource", nil)
	if err != nil {
		t.Fatalf("azure: BearerAuthorizer#WithAuthorization returned an error (%v)", err)
	}

	s := mocks.NewSender()
	s.AppendResponse(mocks.NewResponseWithStatus("400 Bad Request", http.StatusBadRequest))
	spt.SetSender(s)

	ba := NewBearerAuthorizer(spt)
	_, err = Prepare(mocks.NewRequest(), ba.WithAuthorization())
	if err == nil {
		t.Fatal("azure: BearerAuthorizer#WithAuthorization failed to return an error when refresh fails")
	}
}
