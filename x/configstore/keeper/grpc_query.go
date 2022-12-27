package keeper

import (
	"runos_chain/x/configstore/types"
)

var _ types.QueryServer = Keeper{}
