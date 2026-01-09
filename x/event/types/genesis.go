package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		EventList:    []Event{},
		EventDayList: []EventDay{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in event
	eventIndexMap := make(map[string]struct{})

	for _, elem := range gs.EventList {
		index := string(EventKey(elem.Index))
		if _, ok := eventIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for event")
		}
		eventIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in eventDay
	eventDayIndexMap := make(map[string]struct{})

	for _, elem := range gs.EventDayList {
		index := string(EventDayKey(elem.Index))
		if _, ok := eventDayIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for eventDay")
		}
		eventDayIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
