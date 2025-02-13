package assetregistry

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"assetregistry/testutil/sample"
	assetregistrysimulation "assetregistry/x/assetregistry/simulation"
	"assetregistry/x/assetregistry/types"
)

// avoid unused import issue
var (
	_ = assetregistrysimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateAsset = "op_weight_msg_create_asset"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAsset int = 100

	opWeightMsgUpdateAsset = "op_weight_msg_update_asset"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAsset int = 100

	opWeightMsgDeleteAsset = "op_weight_msg_delete_asset"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAsset int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	assetregistryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&assetregistryGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateAsset int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateAsset, &weightMsgCreateAsset, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAsset = defaultWeightMsgCreateAsset
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAsset,
		assetregistrysimulation.SimulateMsgCreateAsset(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAsset int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateAsset, &weightMsgUpdateAsset, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAsset = defaultWeightMsgUpdateAsset
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAsset,
		assetregistrysimulation.SimulateMsgUpdateAsset(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAsset int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteAsset, &weightMsgDeleteAsset, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAsset = defaultWeightMsgDeleteAsset
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAsset,
		assetregistrysimulation.SimulateMsgDeleteAsset(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateAsset,
			defaultWeightMsgCreateAsset,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				assetregistrysimulation.SimulateMsgCreateAsset(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateAsset,
			defaultWeightMsgUpdateAsset,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				assetregistrysimulation.SimulateMsgUpdateAsset(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteAsset,
			defaultWeightMsgDeleteAsset,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				assetregistrysimulation.SimulateMsgDeleteAsset(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
