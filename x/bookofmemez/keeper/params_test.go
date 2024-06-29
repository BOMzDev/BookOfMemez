package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "BookOfMemez/testutil/keeper"
	"BookOfMemez/x/bookofmemez/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BookofmemezKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
