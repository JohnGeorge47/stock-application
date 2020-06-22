package uuid

import (
	"encoding/hex"
	"github.com/gofrs/uuid"
)

// GetUUID generate and return unique id
// the id will be always 32 byte long unique string
func GetUUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(uuid.Bytes())
}
