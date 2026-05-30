package database

type DB struct {
	Storage map[string]any
}

func NewDatabase() *DB {
	return &DB{
		Storage: make(map[string]any),
	}
}

func (db *DB) Write(key string, value any) {

}

func (db *DB) Update(key string, value any) {

}

func (db *DB) Delete(key string) {

}

func (db *DB) Get(key string) {

}

func (db *DB) Exists(key string) bool {
	return false
}
