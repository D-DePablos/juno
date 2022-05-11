package db

// KeyValueStore implement the Storer interface that use a Databaser
type KeyValueStore struct {
	db     *Transactioner
	prefix []byte
}

func NewKeyValueStore(db *Transactioner, prefix string) KeyValueStore {
	return KeyValueStore{db: db, prefix: []byte(prefix)}
}

func (k KeyValueStore) Delete(key []byte) {
	err := (*k.db).Delete(append(k.prefix, key...))
	if err != nil {
		return
	}
}

func (k KeyValueStore) Get(key []byte) ([]byte, bool) {
	get, err := (*k.db).Get(append(k.prefix, key...))
	if err != nil {
		return nil, false
	}
	return get, get != nil
}

func (k KeyValueStore) Put(key, val []byte) {
	err := (*k.db).Put(append(k.prefix, key...), val)
	if err != nil {
		return
	}
}

func (k KeyValueStore) Begin() {
	(*k.db).Begin()
}

func (k KeyValueStore) Rollback() {
	(*k.db).Rollback()
}

func (k KeyValueStore) Close() {
	(*k.db).Close()
}
