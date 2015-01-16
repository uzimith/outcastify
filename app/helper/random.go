package helper

//url: http://qiita.com/suin/items/062dab1e6dc82c81c320

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func GenerateRandom(length int) string {
	const base = 36
	size := big.NewInt(base)
	n := make([]byte, length)
	for i, _ := range n {
		c, _ := rand.Int(rand.Reader, size)
		n[i] = strconv.FormatInt(c.Int64(), base)[0]
	}
	return string(n)
}
