package vmrpc

import (
	"context"
	_ "embed"
	"fmt"
	"math/big"
	"strings"

	"github.com/NethermindEth/juno/pkg/store"
	"github.com/NethermindEth/juno/pkg/trie"
	"github.com/NethermindEth/juno/pkg/types"
)

type storageRPCServer struct {
	UnimplementedStorageAdapterServer
}

// DEBUG.
//go:embed test_contract_def.json
var def string

// DEBUG.
var db = map[string]string{
	// "patricia_node:0704dfcbc470377c68e6f5ffb83970ebd0d7c48d5b8d2f4ed61a24e795e034bd": "002e9723e54711aec56e3fb6ad1bb8272f64ec92e0a43a20feed943b1d4f73c5057dde83c18c0efe7123c36a52d704cf27d5c38cdf0b1e1edc3b0dae3ee4e374fb",
	"contract_state:002e9723e54711aec56e3fb6ad1bb8272f64ec92e0a43a20feed943b1d4f73c5": `{
			"storage_commitment_tree": {
				"root": "04fb440e8ca9b74fc12a22ebffe0bc0658206337897226117b985434c239c028", 
				"height": 251
			}, 
			"contract_hash": "050b2148c0d782914e0b12a1a32abe5e398930b7e914f82c65cb7afce0a0ab9b"
		}`,
	// "patricia_node:04fb440e8ca9b74fc12a22ebffe0bc0658206337897226117b985434c239c028":            "00000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000084fb",
	"starknet_storage_leaf:0000000000000000000000000000000000000000000000000000000000000003":    "0000000000000000000000000000000000000000000000000000000000000003",
	"contract_definition_fact:050b2148c0d782914e0b12a1a32abe5e398930b7e914f82c65cb7afce0a0ab9b": def,
}

// DEBUG.
// Databases that provide the back end to the contract storage and
// global state tries.
var storage, state store.Ephemeral

// DEBUG.
func init() {
	// STORAGE.
	old_storage := store.New()
	contract, _ := trie.New(old_storage, trie.EmptyNode.Hash(), 251)

	// Storage modifications.
	mods := []struct {
		key, val types.Felt
	}{
		{types.BigToFelt(big.NewInt(132)), types.BigToFelt(big.NewInt(3))},
		// TODO: Test for trie with a non-empty binary node i.e. with a
		// modification key = 133 which would create a parent with a node
		// that has the form (0, 0, h(H(left), H(right))) or higher up in
		// the tree e.g. key = 131.
	}

	for _, mod := range mods {
		contract.Put(&mod.key, &mod.val)
	}

	// STATE.
	old_state := store.New()
	global, _ := trie.New(old_state, trie.EmptyNode.Hash(), 251)

	// Contract address.
	key, _ := new(big.Int).SetString("57dde83c18c0efe7123c36a52d704cf27d5c38cdf0b1e1edc3b0dae3ee4e374", 16)
	// h(h(h(contract_hash, storage_root), 0), 0) where h is the Pedersen
	// hash function and storage_root is the commitment of the contract
	// storage trie above.
	val, _ := new(big.Int).SetString("002e9723e54711aec56e3fb6ad1bb8272f64ec92e0a43a20feed943b1d4f73c5", 16)

	// XXX: Why doesn't the following work?
	// global.Put(types.BigToFelt(key), types.BigToFelt(val))
	feltKey, feltVal := types.BigToFelt(key), types.BigToFelt(val)
	global.Put(&feltKey, &feltVal)

	// Simulate new node serialisation.
	state = store.New()
	nodes := old_state.List()
	for _, pair := range nodes {
		key := "patricia_node:" + fmt.Sprintf("%.64x", types.BytesToFelt(pair[0]).Big())
		val := new(trie.Node)
		val.UnmarshalJSON(pair[1])
		state.Put([]byte(key), []byte(val.CairoRepr()))
	}

	storage = store.New()
	nodes = old_storage.List()
	for _, pair := range nodes {
		key := "patricia_node:" + fmt.Sprintf("%.64x", types.BytesToFelt(pair[0]).Big())
		val := new(trie.Node)
		val.UnmarshalJSON(pair[1])
		storage.Put([]byte(key), []byte(val.CairoRepr()))
	}
}

func NewStorageRPCServer() *storageRPCServer {
	return &storageRPCServer{}
}

// GetValue calls the get_value method of the Storage adapter,
// StorageRPCClient, on the Cairo RPC server.
func (s *storageRPCServer) GetValue(ctx context.Context, request *GetValueRequest) (*GetValueResponse, error) {
	// XXX: Handle the following cases i.e. key prefixes. See the
	// following for details https://github.com/starkware-libs/cairo-lang/blob/167b28bcd940fd25ea3816204fa882a0b0a49603/src/starkware/starkware_utils/serializable.py#L25-L31.
	//
	//	1.	patricia_node. A node in a global state or contract storage
	//			trie. Note that cairo-lang does not make a distinction between
	//			whether the node being queried comes from the global state
	//			trie or contract storage trie. One idea on how to address this
	//			is to query the global state tree first and only if the key
	//			does not exit there, query the contract storage trie. See the
	//			following for details https://github.com/eqlabs/pathfinder/blob/82425d44d7aa148bd31a60a7823a3e42b8d613f4/py/src/call.py#L338-L353.
	//
	//	2.	contract_state. The key suffix is the result of
	//			h(h(h(contract_hash, storage_root), 0), 0) where h is the
	//			StarkNet Pedersen hash and the value is perhaps better
	// 			explained by the following reference and example:
	//				- https://github.com/eqlabs/pathfinder/blob/31a308709141cc0d0c0f5568a67e2c9aa89be959/py/src/call.py#L355-L380.
	//				- https://github.com/NethermindEth/juno/blob/42077622e5134e6835f05df0fac9dfd0a2505e9f/pkg/rpc/call.py#L27-L31.
	//			See also above db variable.
	//
	//	3.	contract_definition_fact. The key suffix is the contract hash
	//			(class hash) and the value is the compiled contract. Since
	//			this is not being stored locally, see getFullContract in
	//			internal/services/vm_utils.go that fetches it from the feeder
	//			gateway.
	//
	//	4.	starknet_storage_leaf. Here the key suffix *is* the value so
	//			that could be returned without any lookup on the Python side.
	parts := strings.Split(string(request.GetKey()), ":")
	switch prefix := parts[0]; prefix {
	case "patricia_node":
		// Check for node in state trie first.
		val, ok := state.Get(request.GetKey())
		if !ok {
			// If the key does not exist, check the storage trie.
			val, _ = storage.Get(request.GetKey())
		}
		return &GetValueResponse{Value: val}, nil
	default:
		return &GetValueResponse{Value: []byte(db[string(request.GetKey())])}, nil
	}
}
