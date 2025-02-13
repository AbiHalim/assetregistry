package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"assetregistry/x/assetregistry/types"
)

func (k msgServer) CreateAsset(goCtx context.Context, msg *types.MsgCreateAsset) (*types.MsgCreateAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var asset = types.Asset{
		Name:        msg.Name,
		Description: msg.Description,
		Owner:       msg.Creator,
		Value:       msg.Value,
	}
	id := k.AppendAsset(ctx, asset)
	return &types.MsgCreateAssetResponse{Id: id}, nil
}
