package go_ordered_id

import (
	"strconv"
	"sync"
	"time"

	nanoid "github.com/matoous/go-nanoid/v2"
)

const baseAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const randomLength = 4

type EncodedType string

const EncodedTypeBase62 EncodedType = "base62"
const EncodedTypeBase36 EncodedType = "base36"

var (
	lastTimestamp int64
	mu            sync.Mutex
)

func encodeBase62(base62Alphabet string, num int64) string {
	if num == 0 {
		return "0"
	}
	result := ""
	for num > 0 {
		remainder := num % 62
		result = string(base62Alphabet[remainder]) + result
		num /= 62
	}
	return result
}

func encodeBase36(num int64) string {
	if num == 0 {
		return "0"
	}
	return strconv.FormatInt(num, 36)
}

// Generate creates a globally ordered, short unique ID
func Generate() string {
	now := time.Now().UnixMicro()
	return GenerateCustom(now, EncodedTypeBase62, baseAlphabet, randomLength)
}

func GenerateBase62() string {
	now := time.Now().UnixMicro()
	return GenerateCustom(now, EncodedTypeBase62, baseAlphabet, randomLength)
}

func GenerateBase36() string {
	now := time.Now().UnixMicro()
	return GenerateCustom(now, EncodedTypeBase36, baseAlphabet, randomLength)
}

func GenerateCustom(timeUnix int64, encodedType EncodedType, base62Alphabet string, rndLength int) string {
	var timestampPart string
	mu.Lock()
	defer mu.Unlock()

	now := timeUnix
	if now <= lastTimestamp {
		now = lastTimestamp + 1
	}
	lastTimestamp = now

	switch encodedType {
	case EncodedTypeBase62:
		timestampPart = encodeBase62(base62Alphabet, now)
	case EncodedTypeBase36:
		timestampPart = encodeBase36(now)
	}

	randomPart, err := nanoid.Generate(base62Alphabet, rndLength)
	if err != nil {
		panic(err)
	}

	return timestampPart + randomPart
}
