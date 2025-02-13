package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "assetregistry/testutil/keeper"
	"assetregistry/x/assetregistry/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AssetregistryKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
