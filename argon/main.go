package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

func main() {
	pw := uuid.New().String()

	fmt.Println(pw)

	var (
		memory     uint32 = 64 * 1024
		iterations uint32 = 3
		threads    uint8  = 4
		saltLength int    = 16
		keyLength  uint32 = 32
	)

	salt, err := generateRandomBytes(saltLength)
	if err != nil {
		panic(err)
	}

	hash := argon2.IDKey([]byte(pw), salt, iterations, memory, threads, keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, memory, iterations, threads, b64Salt, b64Hash)

	fmt.Println(encodedHash)
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
