package types

import "cosmossdk.io/errors"

var ErrInvalidAuthority = errors.Register(ModuleName, 1, "signer is not authority")
