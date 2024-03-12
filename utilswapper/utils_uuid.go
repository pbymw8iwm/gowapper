package utilswapper

import (
	"github.com/rs/xid"
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
)

func GetUuids() string {
	id := xid.New()
	data := []byte(id.String())
    has := md5.Sum(data) 
    return fmt.Sprintf("%x", has)
} 
 
func GetUuid() string {
    return uuid.New().String()
}