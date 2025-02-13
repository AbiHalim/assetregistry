package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateAsset{}

func NewMsgCreateAsset(creator string, name string, description string, owner string, value uint64) *MsgCreateAsset {
	return &MsgCreateAsset{
		Creator:     creator,
		Name:        name,
		Description: description,
		Owner:       owner,
		Value:       value,
	}
}

func (msg *MsgCreateAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
