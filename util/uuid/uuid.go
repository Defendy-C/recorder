package uuid

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const Base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// UniqueFilename 文件名唯一表示 len:字符串长度
func UniqueFilename(length int) string {

	vb := make([]byte, length)
	for i, _ := range vb {
		n := rand.Intn(62)
		vb[i] = Base62[n]
	}

	return string(vb)
}
