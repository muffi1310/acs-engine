package armhelpers

import (
	"github.com/Azure/azure-sdk-for-go/arm/resources/resources"
	"github.com/prometheus/common/log"
)

// EnsureResourceGroup ensures the named resouce group exists in the given location.
func (az *AzureClient) EnsureResourceGroup(name, location string) (resourceGroup *resources.Group, err error) {
	log.Debugf("Ensuring resource group exists. resourcegroup=%q", name)
	response, err := az.groupsClient.CreateOrUpdate(name, resources.Group{
		Name:     &name,
		Location: &location,
	})
	if err != nil {
		return &response, err
	}

	return &response, nil
}
