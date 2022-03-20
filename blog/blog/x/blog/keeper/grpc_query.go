package keeper

import (
	"github.com/alice/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
