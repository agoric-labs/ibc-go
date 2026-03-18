package keeper

/*
	This file is to allow for unexported functions to be accessible to the testing package.
*/

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetRecvStartSequence is a wrapper around setRecvStartSequence to allow the function to be directly called in tests.
func (k *Keeper) SetRecvStartSequence(ctx sdk.Context, portID, channelID string, sequence uint64) {
	k.setRecvStartSequence(ctx, portID, channelID, sequence)
}
