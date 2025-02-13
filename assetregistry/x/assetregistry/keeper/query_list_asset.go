package keeper

import (
    "context"

    "cosmossdk.io/store/prefix"
    "github.com/cosmos/cosmos-sdk/runtime"
    "github.com/cosmos/cosmos-sdk/types/query"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    "assetregistry/x/assetregistry/types"
)

func (k Keeper) ListAsset(ctx context.Context, req *types.QueryListAssetRequest) (*types.QueryListAssetResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

    storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AssetKey))

    var assets []types.Asset
    pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
        var asset types.Asset
        if err := k.cdc.Unmarshal(value, &asset); err != nil {
            return err
        }

        assets = append(assets, asset)
        return nil
    })

    if err != nil {
        return nil, status.Error(codes.Internal, err.Error())
    }

    return &types.QueryListAssetResponse{Asset: assets, Pagination: pageRes}, nil
}
