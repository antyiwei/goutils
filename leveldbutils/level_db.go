package leveldbutils

import (
	"fmt"
	generic_sync "github.com/SaveTheRbtz/generic-sync-map-go"
	"github.com/syndtr/goleveldb/leveldb"
	"time"
)

type DBManager struct {
	dbName     string
	dbFilePath string
	db         *leveldb.DB
}

var syncMap = new(generic_sync.MapOf[string, DBManager])

func NewDB(dbName string, dbFilePath string) {
	localDB, err := initDB(dbFilePath)
	if err != nil {
		panic(dbFilePath + ",db file path is err")
	}
	syncMap.Store(dbName, DBManager{
		dbName:     dbName,
		dbFilePath: dbFilePath,
		db:         localDB,
	})
}

func initDB(dbFilePath string) (*leveldb.DB, error) {
	if dbFilePath == "" {
		return nil, fmt.Errorf("dbFilePath:%s is empty", dbFilePath)
	}

	db, err := leveldb.OpenFile(dbFilePath, nil)
	if err != nil {
		return reInit(dbFilePath)
	}

	return db, nil
}

func reInit(dbFilePath string) (*leveldb.DB, error) {

	time.Sleep(10 * time.Microsecond)

	return initDB(dbFilePath)
}

func Get(dbName string) *leveldb.DB {
	m, ok := syncMap.Load(dbName)
	if !ok {
		return nil
	}
	if m.db == nil {
		return nil
	}
	return m.db
}

func GetDBManager(dbName string) *DBManager {
	m, ok := syncMap.Load(dbName)
	if !ok {
		return nil
	}
	return &m
}

func (m *DBManager) SetKV(key string, value string) error {
	return m.db.Put([]byte(key), []byte(value), nil)
}

func (m *DBManager) SetV(value string) error {
	return m.SetKV("single", value)
}

func (m *DBManager) GetKV(key string) (string, error) {

	ok, err := m.db.Has([]byte(key), nil)
	if err != nil {
		return "", err
	}

	if !ok {
		return "", nil
	}

	v, err := m.db.Get([]byte(key), nil)
	if err != nil {
		return "", err
	}

	return string(v), nil

}

func (m *DBManager) GetV() (string, error) {
	return m.GetKV("single")
}

func GetKV(dbName string, key string) (string, error) {
	m := GetDBManager(dbName)
	if m == nil {
		return "", fmt.Errorf("db manager not init")
	}

	return m.GetKV(key)
}

func GetV(dbName string) (string, error) {
	return GetKV(dbName, "single")
}

func SetKV(dbName string, key string, value string) error {
	m := GetDBManager(dbName)
	if m == nil {
		return fmt.Errorf("db manager not init")
	}

	return m.SetKV(key, value)
}

func SetV(dbName string, value string) error {
	return SetKV(dbName, "single", value)
}
