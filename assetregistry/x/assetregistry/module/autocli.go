package assetregistry

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "assetregistry/api/assetregistry/assetregistry"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "ShowAsset",
					Use:            "show-asset [id]",
					Short:          "Query show-asset",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListAsset",
					Use:            "list-asset",
					Short:          "Query list-asset",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateAsset",
					Use:            "create-asset [name] [description] [owner] [value]",
					Short:          "Send a create-asset tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "description"}, {ProtoField: "owner"}, {ProtoField: "value"}},
				},
				{
					RpcMethod:      "UpdateAsset",
					Use:            "update-asset [id] [name] [description] [owner] [value]",
					Short:          "Send a update-asset tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "name"}, {ProtoField: "description"}, {ProtoField: "owner"}, {ProtoField: "value"}},
				},
				{
					RpcMethod:      "DeleteAsset",
					Use:            "delete-asset [id]",
					Short:          "Send a delete-asset tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
