package utils

import "crypto/sha256"


func Sha256(password string) string {
	sum := sha256.Sum256([]byte(password))
	return string(sum[0:32])
}