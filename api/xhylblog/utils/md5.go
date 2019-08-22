package utils

import (
	"crypto/md5"
	"encoding/hex"
)
/**
 * 用于加密,解密,(包含MD5加密和base64加密/解密)以及GUID的生成
 * 时间:
 * zhifieya
 */

/**
 * 对一个字符串进行MD5加密,不可解密
 */
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s)) //使用名字做散列值，设定后不要变
	return hex.EncodeToString(h.Sum(nil))
}
