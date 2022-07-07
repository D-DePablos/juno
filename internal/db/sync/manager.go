package sync

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NethermindEth/juno/internal/db"
)

var (
	DbError                        = errors.New("database error")
	UnmarshalError                 = errors.New("unmarshal error")
	MarshalError                   = errors.New("marshal error")
	latestBlockSyncKey             = []byte("latestBlockSync")
	blockOfLatestEventProcessedKey = []byte("blockOfLatestEventProcessed")
	latestStateRoot                = []byte("latestStateRoot")
	latestBlockInfoFetched         = []byte("latestBlockInfoFetched")
)

// Manager is a Block database manager to save and search the blocks.
type Manager struct {
	database db.Database
}

// NewSyncManager returns a new Block manager using the given database.
func NewSyncManager(database db.Database) *Manager {
	return &Manager{database: database}
}

// StoreLatestBlockSync stores the latest block sync.
func (m *Manager) StoreLatestBlockSync(latestBlockSync int64) {

	// Marshal the latest block sync
	value, err := json.Marshal(latestBlockSync)
	if err != nil {
		panic(any(fmt.Errorf("%w: %s", MarshalError, err)))
	}

	// Store the latest block sync
	err = m.database.Put(latestBlockSyncKey, value)
	if err != nil {
		panic(any(fmt.Errorf("%w: %s", DbError, err.Error())))
	}
}

// GetLatestBlockSync returns the latest block sync.
func (m *Manager) GetLatestBlockSync() int64 {
	// Query to database
	data, err := m.database.Get(latestBlockSyncKey)
	if err != nil {
		if db.ErrNotFound == err {
			return 0
		}
		// notest
		panic(any(fmt.Errorf("%w: %s", DbError, err)))
	}
	if data == nil {
		// notest
		return 0
	}
	// Unmarshal the data from database
	latestBlockSync := new(int64)
	if err := json.Unmarshal(data, latestBlockSync); err != nil {
		// notest
		panic(any(fmt.Errorf("%w: %s", UnmarshalError, err.Error())))
	}
	return *latestBlockSync
}

// StoreLatestStateRoot stores the latest state root.
func (m *Manager) StoreLatestStateRoot(stateRoot string) {
	// Store the latest state root
	err := m.database.Put(latestStateRoot, []byte(stateRoot))
	if err != nil {
		panic(any(fmt.Errorf("%w: %s", DbError, err.Error())))
	}
}

// GetLatestStateRoot returns the latest state root.
func (m *Manager) GetLatestStateRoot() string {
	// Query to database
	data, err := m.database.Get(latestStateRoot)
	if err != nil {
		if db.ErrNotFound == err {
			return ""
		}
		// notest
		panic(any(fmt.Errorf("%w: %s", DbError, err)))
	}
	if data == nil {
		// notest
		return ""
	}
	// Unmarshal the data from database
	return string(data)
}

// StoreBlockOfProcessedEvent stores the block of the latest event processed,
func (m *Manager) StoreBlockOfProcessedEvent(starknetFact, l1Block int64) {

	key := []byte(fmt.Sprintf("%s%d", blockOfLatestEventProcessedKey, starknetFact))
	// Marshal the latest block sync
	value, err := json.Marshal(l1Block)
	if err != nil {
		panic(any(fmt.Errorf("%w: %s", MarshalError, err)))
	}

	// Store the latest block sync
	err = m.database.Put(key, value)
	if err != nil {
		panic(any(fmt.Errorf("%w: %s", DbError, err.Error())))
	}
}

// GetBlockOfProcessedEvent returns the block of the latest event processed,
func (m *Manager) GetBlockOfProcessedEvent(starknetFact int64) int64 {
	// Query to database
	key := []byte(fmt.Sprintf("%s%d", blockOfLatestEventProcessedKey, starknetFact))
	data, err := m.database.Get(key)
	if err != nil {
		if db.ErrNotFound == err {
			return 0
		}
		// notest
		panic(any(fmt.Errorf("%w: %s", DbError, err)))
	}
	if data == nil {
		// notest
		return 0
	}
	// Unmarshal the data from database
	blockSync := new(int64)
	if err := json.Unmarshal(data, blockSync); err != nil {
		// notest
		panic(any(fmt.Errorf("%w: %s", UnmarshalError, err.Error())))
	}
	return *blockSync
}

// StoreLatestBlockInfoFetched stores the block number of the latest block used to fetch information
func (m *Manager) StoreLatestBlockInfoFetched(blockNumber int64) {

	value, err := json.Marshal(blockNumber)
	if err != nil {
		panic(any(fmt.Errorf("%w: %s", MarshalError, err)))
	}

	// Store the latest block sync
	err = m.database.Put(latestBlockInfoFetched, value)
	if err != nil {
		panic(any(fmt.Errorf("%w: %s", DbError, err.Error())))
	}
}

// GetLatestBlockInfoFetched returns the block number of the latest block used to fetch information
func (m *Manager) GetLatestBlockInfoFetched() int64 {
	// Query to database
	data, err := m.database.Get(latestBlockInfoFetched)
	if err != nil {
		if db.ErrNotFound == err {
			return 0
		}
		// notest
		panic(any(fmt.Errorf("%w: %s", DbError, err)))
	}
	if data == nil {
		// notest
		return 0
	}
	// Unmarshal the data from database
	blockNumber := new(int64)
	if err := json.Unmarshal(data, blockNumber); err != nil {
		// notest
		panic(any(fmt.Errorf("%w: %s", UnmarshalError, err.Error())))
	}
	return *blockNumber
}

// Close closes the Manager.
func (m *Manager) Close() {
	m.database.Close()
}
