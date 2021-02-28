package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"strconv"
	"time"
)

// Generate 返回给定参数生成的一个6位TOTP值。
// 参数 k：预共享的密钥
// 参数 t：时间
// 返回 code：一次性密码
func Generate(k string, t time.Time) (code string, err error) {

	bk, err := base32.StdEncoding.DecodeString(k)
	if err != nil {
		return
	}

	c := uint64(t.Unix()) / 30
	msg := make([]byte, 8)
	binary.BigEndian.PutUint64(msg, c)

	sha := hmac.New(sha1.New, bk)
	sha.Write(msg)
	hash := sha.Sum(nil)

	offset := hash[19] & 0xf
	oi := hash[offset : offset+4]
	oi[0] = oi[0] & 0x7f
	i := binary.BigEndian.Uint32(oi) % 1000000

	code = strconv.FormatUint(uint64(i), 10)
	if l := len(code); l < 6 {
		zero := []string{"00000", "0000", "000", "00", "0"}
		code = zero[l-1] + code
	}

	return
}

// Authenticate 返回验证一次性密码结果
// 参数 k：预共享的密钥
// 参数 t：时间
// 参数 code：一次性密码
func Authenticate(k string, t time.Time, code string) (success bool, err error) {
	c, err := Generate(k, t)
	if err != nil {
		return
	}
	return c == code, nil
}
