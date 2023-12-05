package xcrypto

import (
	"crypto/md5"
	"encoding/hex"
)

// 生成md5码
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
