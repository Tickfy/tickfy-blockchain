package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "tickfy-blockchain/testutil/keeper"
	"tickfy-blockchain/testutil/nullify"
	"tickfy-blockchain/x/event/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestEventQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EventKeeper(t)
	msgs := createNEvent(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetEventRequest
		response *types.QueryGetEventResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetEventRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetEventResponse{Event: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetEventRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetEventResponse{Event: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetEventRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Event(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestEventQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EventKeeper(t)
	msgs := createNEvent(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllEventRequest {
		return &types.QueryAllEventRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.EventAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Event), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Event),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.EventAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Event), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Event),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.EventAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Event),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.EventAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
