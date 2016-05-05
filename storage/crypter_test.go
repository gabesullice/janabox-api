package storage

import (
	"fmt"
)

type CryptMock struct {
	shouldFail bool
}

func (c *CryptMock) ShouldFail(shouldFail bool) {
	c.shouldFail = shouldFail
}

func (c CryptMock) Encrypt(plaintext string) (string, error) {
	if c.shouldFail {
		return "", fmt.Errorf("Expected failure")
	} else {
		return "ciphertext", nil
	}
}

func (c CryptMock) Decrypt(ciphertext string) (string, error) {
	if c.shouldFail {
		return "", fmt.Errorf("Expected failure")
	} else {
		return "plaintext", nil
	}
}
