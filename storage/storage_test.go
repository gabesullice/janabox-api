package storage

import (
	"crypto/md5"
	"fmt"
)

type StorageMock struct {
	shouldFail bool
	KVStore    map[string]string
}

func NewStorageMock() *StorageMock {
	return &StorageMock{
		shouldFail: false,
		KVStore:    map[string]string{},
	}
}

func (s *StorageMock) ShouldFail(shouldFail bool) {
	s.shouldFail = shouldFail
}

func (s *StorageMock) Insert(data string) (string, error) {
	if s.shouldFail {
		return "", fmt.Errorf("Expected failure")
	} else {
		uuid := getFakeUUID(data)
		s.KVStore[uuid] = data
		return uuid, nil
	}
}

func (s *StorageMock) Get(uuid string) (string, error) {
	if s.shouldFail {
		return "", fmt.Errorf("Expected failure")
	} else {
		return s.KVStore[uuid], nil
	}
}

func (s *StorageMock) Update(uuid string, data string) error {
	if s.shouldFail {
		return fmt.Errorf("Expected failure")
	} else {
		s.KVStore[uuid] = data
		return nil
	}
}

func (s *StorageMock) Remove(uuid string) error {
	if s.shouldFail {
		return fmt.Errorf("Expected failure")
	} else {
		delete(s.KVStore, uuid)
		return nil
	}
}

func getFakeUUID(data string) string {
	uuid := md5.Sum([]byte(data))
	return string(uuid[:])
}
