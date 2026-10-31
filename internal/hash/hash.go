package hash

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/json"

	"golang.org/x/crypto/argon2"
)

func GetHasher(hashConfig string) (Hasher, error) {
	hasher := Hasher{}

	err := json.Unmarshal([]byte(hashConfig), &hasher)

	return hasher, err
}

func (h Hasher) GenerateHash(password []byte, salt []byte) (Hash, error) {
	if len(salt) == 0 {
		var err error
		salt, err = randomSalt(h.SaltLen)

		if err != nil {
			return Hash{}, err
		}
	}

	_hash := argon2.IDKey(password, salt, h.Time, h.Memory, h.Threads, h.KeyLen)

	return Hash{Hash: _hash, Salt: salt}, nil
}

func (h Hasher) CompareHash(hash Hash, password string) (bool, error) {
	_passwordHash, err := h.GenerateHash([]byte(password), hash.Salt)
	if err != nil {
		return false, err
	}

	passwordHash := append(_passwordHash.Hash, _passwordHash.Salt...)
	targetHash := append(hash.Hash, hash.Salt...)

	if subtle.ConstantTimeCompare(targetHash, passwordHash) != 1 {
		return false, nil
	} else {
		return true, nil
	}
}

func randomSalt(length uint32) ([]byte, error) {
	secret := make([]byte, length)

	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}
