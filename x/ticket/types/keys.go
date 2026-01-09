package types

const (
	// ModuleName defines the module name
	ModuleName = "ticket"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ticket"
)

var (
	ParamsKey = []byte("p_ticket")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
