package types

const (
	// ModuleName defines the module name
	ModuleName = "treasury"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_treasury"
)

var (
	ParamsKey = []byte("p_treasury")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
