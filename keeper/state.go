// Copyright 2024 NASD Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
