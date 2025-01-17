// Auto generated code - DO NOT EDIT.

package network

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
)

func TestNetworkInterfaces(t *testing.T) {
	client.MockTestHelper(t, Interfaces(), createInterfacesMock)
}

func createInterfacesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkInterfacesClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			Interfaces: mockClient,
		},
	}

	data := network.Interface{}
	require.Nil(t, faker.FakeObject(&data))

	result := network.NewInterfaceListResultPage(network.InterfaceListResult{Value: &[]network.Interface{data}}, func(ctx context.Context, result network.InterfaceListResult) (network.InterfaceListResult, error) {
		return network.InterfaceListResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any()).Return(result, nil)
	return s
}
