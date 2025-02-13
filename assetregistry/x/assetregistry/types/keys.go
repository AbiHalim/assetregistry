package types

const (
	// ModuleName defines the module name
	ModuleName = "assetregistry"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_assetregistry"

	AssetKey      = "Asset/value/"
	AssetCountKey = "Asset/count/"
)

var (
	ParamsKey = []byte("p_assetregistry")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
