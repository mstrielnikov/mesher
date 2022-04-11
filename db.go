package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	"lukechampine.com/blake3"
)

type Storage struct {
	DB   *leveldb.DB
	Hash [32]byte
}

func NewStorage(levelDb string) Storage {
	db, err := leveldb.OpenFile(levelDb, nil)
	Ok(err)
	return Storage{
		DB: db,
	}
}

func (s *Storage) GetValueByKey(key []byte) []byte {
	bytes, err := s.DB.Get(key, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	return bytes
}

func (s *Storage) FetchKeys() []byte {
	var keys []byte
	iter := s.DB.NewIterator(nil, nil)
	for iter.Next() {
		keys = append(keys, iter.Key()...)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Println(err.Error())
	}
	return keys
}

func (s *Storage) FetchKeyValues() ([]byte, []byte) {
	var keys, values []byte
	iter := s.DB.NewIterator(nil, nil)
	for iter.Next() {
		keys = append(keys, iter.Key()...)
		values = append(values, iter.Value()...)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Println(err.Error())
	}
	return keys, values
}

func (s *Storage) UpdateByKey(key, value []byte) {
	err := s.DB.Put(key, value, nil)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		s.hash(value)
	}
}

func (s *Storage) hash(data []byte) {
	buf := append(s.Hash[:], data...)
	s.Hash = blake3.Sum256(buf)
}
