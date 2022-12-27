package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/nuclearIgor/go-hexagonal/adapters/cli"
	mock_application "github.com/nuclearIgor/go-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 210
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Crate(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	expectedResult := fmt.Sprintf("Product Id %s with the name %s has been created with price %d and status %s",
		productId, productName, productPrice, productStatus)

	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product %s has been enabled", productName)
	result, err = cli.Run(service, "enable", productId, productName, 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product %s has been disabled", productName)
	result, err = cli.Run(service, "disable", productId, productName, 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product Id: %s\nName: %s\nPrice: %d\nStatus: %s",
		productId, productName, productPrice, productStatus)
	result, err = cli.Run(service, "get", productId, productName, 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}
