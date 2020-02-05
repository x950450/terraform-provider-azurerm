package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type DevspaceControllerId struct {
	ResourceGroup string
	Name          string
}

func DevspaceControllerID(input string) (*DevspaceControllerId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Devspace Controller ID %q: %+v", input, err)
	}

	controller := DevspaceControllerId{
		ResourceGroup: id.ResourceGroup,
	}

	if controller.Name, err = id.PopSegment("controllers"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &controller, nil
}
