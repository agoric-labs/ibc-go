package keeper_test

import (
	"testing"

	testifysuite "github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	channeltypes "github.com/cosmos/ibc-go/v10/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v10/modules/core/05-port/keeper"
	porttypes "github.com/cosmos/ibc-go/v10/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v10/modules/core/exported"
	"github.com/cosmos/ibc-go/v10/testing/simapp"
)

type KeeperTestSuite struct {
	testifysuite.Suite

	ctx    sdk.Context
	keeper *keeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	isCheckTx := false
	app := simapp.Setup(suite.T(), isCheckTx)

	suite.ctx = app.NewContext(isCheckTx)
	suite.keeper = app.IBCKeeper.PortKeeper
}

func TestKeeperTestSuite(t *testing.T) {
	testifysuite.Run(t, new(KeeperTestSuite))
}

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

func (suite *KeeperTestSuite) TestRouteExactMatchPreferredOverFallback() {
	exactModule := &mockIBCModule{}
	fallbackModule := &mockIBCModule{}

	rtr := porttypes.NewRouter().
		AddRoute("transfer", fallbackModule).
		AddRoute("transferchannel0", exactModule)
	rtr.Seal()

	k := keeper.NewKeeper()
	k.Router = rtr

	route, ok := k.Route("transferchannel0")
	suite.Require().True(ok)
	suite.Require().Same(exactModule, route)
}

func (suite *KeeperTestSuite) TestRouteFallsBackToMatchingPrefix() {
	fallbackModule := &mockIBCModule{}

	rtr := porttypes.NewRouter().
		AddRoute("transfer", fallbackModule)
	rtr.Seal()

	k := keeper.NewKeeper()
	k.Router = rtr

	route, ok := k.Route("transferchannel0")
	suite.Require().True(ok)
	suite.Require().Same(fallbackModule, route)
}
