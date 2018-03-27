package store

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/golang/protobuf/proto"
	"reflect"
)

// Database represents a file-based database that acts as a simple key-value store for protobuf messages.
type Database struct {
	bolt *bolt.DB
}

var defaultOptions = &bolt.Options{
	Timeout:    0,
	NoGrowSync: false,
	ReadOnly:   false,
}
var readOnlyOptions = &bolt.Options{
	Timeout:    0,
	NoGrowSync: false,
	ReadOnly:   true,
}

// OpenDatabase opens an existing database or creates a new database at the specified location
func OpenDatabase(filename string, readOnly bool) (*Database, error) {
	options := defaultOptions
	if readOnly {
		options = readOnlyOptions
	}
	bolt, err := bolt.Open(filename, 0644, options)
	if err != nil {
		return nil, err
	}
	return &Database{bolt: bolt}, nil
}

// Get returns a single value from the persistence store
func (db *Database) Get(bucket string, key string, pb proto.Message) error {
	return db.bolt.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b != nil {
			v := b.Get([]byte(key))
			return proto.Unmarshal(v, pb)
		}
		return fmt.Errorf("missing bucket '%s'", bucket)
	})
}

// Put puts values into the persistence store, using the key derived from the property specified.
func (db *Database) Put(bucket string, propertyKey string, pbs []proto.Message) error {
	return db.bolt.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		for _, pb := range pbs {
			v, err := proto.Marshal(pb)
			if err != nil {
				return err
			}
			vv := reflect.ValueOf(pb).Elem()
			field := vv.FieldByName(propertyKey)
			key := field.String()
			if err = b.Put([]byte(key), v); err != nil {
				return err
			}
		}
		return nil
	})
}
