package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
)

// IsControllerEnabled retrieves the controller enabled boolean from the paramstore.
// True is returned if the controller submodule is enabled.
func (k Keeper) IsControllerEnabled(ctx sdk.Context) bool {
	var res bool
	k.paramSpace.Get(ctx, types.KeyControllerEnabled, &res)
	return res
}
