package storage

import (
	"testing"
)

var (
	es = EncryptedStorage{
		crypt:   CryptMock{shouldFail: false},
		storage: NewStorageMock(),
	}
)

func TestInsert(t *testing.T) {
	uuid, err := es.Insert("example")
	if err != nil {
		t.Errorf("Insert failed with error: %s", err.Error())
	}

	if v, _ := es.storage.Get(uuid); v != "ciphertext" {
		t.Error(
			"expected", "ciphertext",
			"got", v,
		)
	}
}

func TestGet(t *testing.T) {
	uuid, err := es.Insert("example")
	if err != nil {
		t.Errorf("Insert failed with error: %s", err.Error())
	}

	if v, err := es.Get(uuid); err != nil {
		t.Errorf("Get failed with error: %s", err.Error())
	} else if v != "plaintext" {
		t.Error(
			"expected", "plaintext",
			"got", v,
		)
	}
}
