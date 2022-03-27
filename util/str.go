package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	STR_ALPHANUMERIC = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	STR_LOWER_ALPHA  = "abcdefghijklmnopqrstuvwxyz"
	STR_UPPER_ALPHA  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	STR_NUMERIC      = "1234567890"
	STR_STANDARD     = STR_ALPHANUMERIC + "-_"
)

func GenerateRandomString(length int, characters string) string {
	const letterIdxBits = 6
	const letterIdxMask = 1<<letterIdxBits - 1
	const letterIdxMax = 63 / letterIdxBits

	var src = rand.NewSource(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(characters) {
			sb.WriteByte(characters[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}
