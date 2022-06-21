package types

import (
	glob "github.com/ryanuber/go-glob"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// SubModuleName defines the interchain accounts host module name
	SubModuleName = "icahost"

	// StoreKey is the store key string for the interchain accounts host module
	StoreKey = SubModuleName
)

// ContainsMsgType returns true if the sdk.Msg TypeURL is present in allowMsgs, otherwise false
func ContainsMsgType(allowMsgs []string, msg sdk.Msg) bool {
	typeUrl := sdk.MsgTypeURL(msg)
	for _, v := range allowMsgs {
		if glob.Glob(v, typeUrl) {
			return true
		}
		if v == typeUrl {
			return true
		}
	}

	return false
}
