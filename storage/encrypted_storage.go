package storage

import (
	"crypto/rand"
)

type EncryptedStorage struct {
	crypt   Crypt
	storage Storage
}

type Crypt interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
}

type DefaultCrypt struct {
	secret []byte
}

func NewDefaultCrypt(secret []byte) DefaultCrypt {
	return DefaultCrypt{secret: secret}
}

func NewEncryptedStorage(secret []byte, storage Storage) EncryptedStorage {
	if random, err := randomBytes(256); err != nil {
		panic(err)
	} else {
		return EncryptedStorage{crypt: NewDefaultCrypt(random), storage: storage}
	}
}

func (s EncryptedStorage) Insert(data string) (string, error) {
	if cipherText, err := s.crypt.Encrypt(data); err != nil {
		return "", err
	} else {
		return s.storage.Insert(cipherText)
	}
}

func (s EncryptedStorage) Get(uuid string) (string, error) {
	if cipherText, err := s.storage.Get(uuid); err != nil {
		return "", err
	} else {
		return s.crypt.Decrypt(cipherText)
	}
}

func (s EncryptedStorage) Update(uuid string, data string) error {
	if cipherText, err := s.crypt.Encrypt(data); err != nil {
		return err
	} else {
		return s.storage.Update(uuid, cipherText)
	}
}

func (s EncryptedStorage) Remove(uuid string) error {
	return s.storage.Remove(uuid)
}

func (c DefaultCrypt) Encrypt(plainText string) (string, error) {
	return plainText, nil
}

func (c DefaultCrypt) Decrypt(cipherText string) (string, error) {
	return cipherText, nil
}

func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
