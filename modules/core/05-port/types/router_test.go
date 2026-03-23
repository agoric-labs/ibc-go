package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	channeltypes "github.com/cosmos/ibc-go/v10/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v10/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v10/modules/core/exported"
)

type mockIBCModule struct{}

func (mockIBCModule) OnChanOpenInit(sdk.Context, channeltypes.Order, []string, string, string, channeltypes.Counterparty, string) (string, error) {
	return "", nil
}

func (mockIBCModule) OnChanOpenTry(sdk.Context, channeltypes.Order, []string, string, string, channeltypes.Counterparty, string) (string, error) {
	return "", nil
}

func (mockIBCModule) OnChanOpenAck(sdk.Context, string, string, string, string) error {
	return nil
}

func (mockIBCModule) OnChanOpenConfirm(sdk.Context, string, string) error {
	return nil
}

func (mockIBCModule) OnChanCloseInit(sdk.Context, string, string) error {
	return nil
}

func (mockIBCModule) OnChanCloseConfirm(sdk.Context, string, string) error {
	return nil
}

func (mockIBCModule) OnRecvPacket(sdk.Context, string, channeltypes.Packet, sdk.AccAddress) exported.Acknowledgement {
	return nil
}

func (mockIBCModule) OnAcknowledgementPacket(sdk.Context, string, channeltypes.Packet, []byte, sdk.AccAddress) error {
	return nil
}

func (mockIBCModule) OnTimeoutPacket(sdk.Context, string, channeltypes.Packet, sdk.AccAddress) error {
	return nil
}

func TestAddRouteAcceptsValidPortIdentifier(t *testing.T) {
	router := porttypes.NewRouter()
	module := mockIBCModule{}

	require.NotPanics(t, func() {
		router.AddRoute("icacontroller-1", module)
	}, "expected a valid IBC port identifier to be accepted")

	require.NotPanics(t, func() {
		// Allows all alphanumerics.
		router.AddRoute("i", module)
	}, "accept a valid alphanumeric port identifier")

	require.NotPanics(t, func() {
		// Allows all alphanumerics.
		router.AddRoute("O", module)
	}, "accept a valid alphanumeric port identifier")

	require.NotPanics(t, func() {
		// Allows all alphanumerics.
		router.AddRoute("1", module)
	}, "accept a valid alphanumeric port identifier")

	route, ok := router.Route("icacontroller-1")
	require.True(t, ok)
	require.Equal(t, module, route)
}
