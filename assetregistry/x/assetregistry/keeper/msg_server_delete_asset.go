package keeper

import (
	"context"
	"fmt"

	"assetregistry/x/assetregistry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteAsset(goCtx context.Context, msg *types.MsgDeleteAsset) (*types.MsgDeleteAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Fetch the asset to be deleted
	asset, found := k.GetAsset(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("asset with ID %d doesn't exist", msg.Id))
	}

	// Ensure the creator is the owner of the asset
	if msg.Creator != asset.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only the owner can delete the asset")
	}

	// Remove the asset from the store
	k.RemoveAsset(ctx, msg.Id)

	return &types.MsgDeleteAssetResponse{}, nil
}
