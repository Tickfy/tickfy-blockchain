package types

const (
	// ModuleName defines the module name
	ModuleName = "event"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_event"
)

var (
	ParamsKey = []byte("p_event")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
