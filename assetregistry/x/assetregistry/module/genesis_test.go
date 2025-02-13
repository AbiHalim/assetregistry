package assetregistry_test

import (
	"testing"

	keepertest "assetregistry/testutil/keeper"
	"assetregistry/testutil/nullify"
	assetregistry "assetregistry/x/assetregistry/module"
	"assetregistry/x/assetregistry/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AssetregistryKeeper(t)
	assetregistry.InitGenesis(ctx, k, genesisState)
	got := assetregistry.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
