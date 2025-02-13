package keeper

import (
  "context"
  "fmt"

  "assetregistry/x/assetregistry/types"

  sdk "github.com/cosmos/cosmos-sdk/types"
  errorsmod "cosmossdk.io/errors"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateAsset(goCtx context.Context, msg *types.MsgUpdateAsset) (*types.MsgUpdateAssetResponse, error) {
  ctx := sdk.UnwrapSDKContext(goCtx)

  // Retrieve the existing asset by ID
  existingAsset, found := k.GetAsset(ctx, msg.Id)
  if !found {
    return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("asset with ID %d doesn't exist", msg.Id))
  }

  // Ensure the caller is the owner
  if msg.Owner != existingAsset.Owner {
    return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only the owner can update the asset")
  }

  // Create an updated asset instance
  updatedAsset := types.Asset{
    Id:          msg.Id,
    Name:        msg.Name,
    Description: msg.Description,
    Owner:       msg.Owner, // Ownership remains the same
    Value:       msg.Value,
  }

  // Save the updated asset
  k.SetAsset(ctx, updatedAsset)

  return &types.MsgUpdateAssetResponse{}, nil
}
