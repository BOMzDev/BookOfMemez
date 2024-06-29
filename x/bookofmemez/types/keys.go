package types

const (
	// ModuleName defines the module name
	ModuleName = "bookofmemez"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bookofmemez"
)

var (
	ParamsKey = []byte("p_bookofmemez")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
