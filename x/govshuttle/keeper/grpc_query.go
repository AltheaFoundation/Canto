package keeper

import (
	"github.com/Canto-Network/Canto/v4/x/govshuttle/types"
)

var _ types.QueryServer = Keeper{}
