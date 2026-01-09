package event

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "tickfy-blockchain/api/tickfyblockchain/event"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "EventAll",
					Use:       "list-event",
					Short:     "List all event",
				},
				{
					RpcMethod:      "Event",
					Use:            "show-event [id]",
					Short:          "Shows a event",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "EventDayAll",
					Use:       "list-event-day",
					Short:     "List all eventDay",
				},
				{
					RpcMethod:      "EventDay",
					Use:            "show-event-day [id]",
					Short:          "Shows a eventDay",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateEvent",
					Use:            "create-event [index] [name] [description] [location] [image-url] [creator-fee]",
					Short:          "Send a createEvent tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "name"}, {ProtoField: "description"}, {ProtoField: "location"}, {ProtoField: "imageUrl"}, {ProtoField: "creatorFee"}},
				},
				{
					RpcMethod:      "UpdateEvent",
					Use:            "update-event [index] [name] [description] [location] [image-url] [creator-fee]",
					Short:          "Send a updateEvent tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "name"}, {ProtoField: "description"}, {ProtoField: "location"}, {ProtoField: "imageUrl"}, {ProtoField: "creatorFee"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
