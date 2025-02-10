package keeper

import "context"

// GetBypassMessages is a utility that returns all bypass messages from state.
func (k *Keeper) GetBypassMessages(ctx context.Context) (bypassMessages []string, err error) {
	err = k.BypassMessages.Walk(ctx, nil, func(bypassMessage string) (stop bool, err error) {
		bypassMessages = append(bypassMessages, bypassMessage)
		return false, nil
	})

	return
}
