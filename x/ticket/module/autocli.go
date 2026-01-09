package ticket

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "tickfy-blockchain/api/tickfyblockchain/ticket"
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
					RpcMethod: "TicketAll",
					Use:       "list-ticket",
					Short:     "List all ticket",
				},
				{
					RpcMethod:      "Ticket",
					Use:            "show-ticket [id]",
					Short:          "Shows a ticket",
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
					RpcMethod:      "CreateTicket",
					Use:            "create-ticket [index] [event-index] [event-day-index] [owner] [price] [metadata]",
					Short:          "Send a createTicket tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "eventIndex"}, {ProtoField: "eventDayIndex"}, {ProtoField: "owner"}, {ProtoField: "price"}, {ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "UpdateTicket",
					Use:            "update-ticket [index] [metadata] [price]",
					Short:          "Send a updateTicket tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "metadata"}, {ProtoField: "price"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
