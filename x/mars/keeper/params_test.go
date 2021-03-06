package keeper_test

import (
	"testing"

	testkeeper "github.com/auagho/mars/testutil/keeper"
	"github.com/auagho/mars/x/mars/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MarsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
