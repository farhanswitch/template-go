package utilities

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"
)

func GenerateUUIDv7() (string, error) {
	var uuid [16]byte

	// Get timestamp in milliseconds
	timestamp := uint64(time.Now().UnixMilli())

	// Put timestamp in first 6 bytes (48 bits)
	binary.BigEndian.PutUint32(uuid[0:4], uint32(timestamp>>16))
	binary.BigEndian.PutUint16(uuid[4:6], uint16(timestamp&0xFFFF))

	// Generate random bytes for the remaining parts
	if _, err := rand.Read(uuid[6:]); err != nil {
		return "", err
	}

	// Set version 7 (0111) in bits 48-51
	uuid[6] = 0x70 | (uuid[6] & 0x0F)

	// Set variant 2 (10) in bits 64-65
	uuid[8] = 0x80 | (uuid[8] & 0x3F)

	// Format UUID string with hyphens
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		uuid[0], uuid[1], uuid[2], uuid[3],
		uuid[4], uuid[5],
		uuid[6], uuid[7],
		uuid[8], uuid[9],
		uuid[10], uuid[11], uuid[12], uuid[13], uuid[14], uuid[15]), nil
}
