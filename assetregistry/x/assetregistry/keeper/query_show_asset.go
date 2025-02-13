package keeper

import (
    "context"

    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    "assetregistry/x/assetregistry/types"
)

func (k Keeper) ShowAsset(goCtx context.Context, req *types.QueryShowAssetRequest) (*types.QueryShowAssetResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

    ctx := sdk.UnwrapSDKContext(goCtx)
    asset, found := k.GetAsset(ctx, req.Id)
    if !found {
        return nil, sdkerrors.ErrKeyNotFound
    }

    return &types.QueryShowAssetResponse{Asset: asset}, nil
}
