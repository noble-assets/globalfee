package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/errors"
	"github.com/noble-assets/globalfee/types"
)

var _ types.MsgServer = &msgServer{}

type msgServer struct {
	*Keeper
}

func NewMsgServer(keeper *Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) UpdateGasPrices(ctx context.Context, msg *types.MsgUpdateGasPrices) (*types.MsgUpdateGasPricesResponse, error) {
	if msg.Signer != k.authority {
		return nil, errors.Wrapf(types.ErrInvalidAuthority, "expected %s, got %s", k.authority, msg.Signer)
	}

	gasPrices := types.GasPrices{Value: msg.GasPrices}
	if err := gasPrices.Validate(); err != nil {
		return nil, errors.Wrap(err, "failed to validate gas prices")
	}

	if err := k.GasPrices.Set(ctx, gasPrices); err != nil {
		return nil, errors.Wrap(err, "failed to set gas prices in state")
	}

	return &types.MsgUpdateGasPricesResponse{}, nil
}

func (k msgServer) UpdateBypassMessages(ctx context.Context, msg *types.MsgUpdateBypassMessages) (*types.MsgUpdateBypassMessagesResponse, error) {
	if msg.Signer != k.authority {
		return nil, errors.Wrapf(types.ErrInvalidAuthority, "expected %s, got %s", k.authority, msg.Signer)
	}

	if err := k.BypassMessages.Clear(ctx, nil); err != nil {
		return nil, errors.Wrap(err, "failed to reset bypass messages in state")
	}

	for _, bypassMessage := range msg.BypassMessages {
		resolved, err := k.registry.Resolve(bypassMessage)
		if err != nil || resolved == nil {
			return nil, fmt.Errorf("invalid bypass message %s", bypassMessage)
		}

		if err := k.BypassMessages.Set(ctx, bypassMessage); err != nil {
			return nil, errors.Wrap(err, "failed to set bypass message in state")
		}
	}

	return &types.MsgUpdateBypassMessagesResponse{}, nil
}
