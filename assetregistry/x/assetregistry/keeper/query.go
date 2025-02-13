package keeper

import (
	"assetregistry/x/assetregistry/types"
)

var _ types.QueryServer = Keeper{}
