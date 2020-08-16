package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(in string) string {
	hash := md5.New()
	hash.Write([]byte(in))

	return hex.EncodeToString(hash.Sum(nil))
}
