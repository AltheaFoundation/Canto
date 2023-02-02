package keeper

import (
	"github.com/Canto-Network/Canto/v3/x/govshuttle/types"
)

var _ types.QueryServer = Keeper{}
