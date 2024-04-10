package keeper_test

import (
	"testing"

	"github.com/egeaybars123/checkers/testutil"
	keepertest "github.com/egeaybars123/checkers/testutil/keeper"
	"github.com/egeaybars123/checkers/x/checkers"
	"github.com/egeaybars123/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
	carol = testutil.Carol
)

func TestCreateGame1(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())

	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "", // TODO: update with a proper value when updated
	}, *createResponse)
}
