package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
)

func GenerateMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func Randint64() (int64, error) {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(b[:])), nil
}

func RandUint64() uint64 {
	buf := make([]byte, 8)
	rand.Read(buf) // Always succeeds, no need to check error
	return uint64(binary.LittleEndian.Uint32(buf))
}
