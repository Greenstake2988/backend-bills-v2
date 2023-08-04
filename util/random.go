package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabhet = "abcdefghijklmnijklpoqrstuvwxyz"

// Create a new random generator with the specific seed value
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Generar
func RandomInt(min, max int64) int64 {

	return min + r.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabhet)

	for i := 0; i < n; i++ {
		c := alphabhet[r.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomMail() string {
	mails := []string{"@gmail.com", "@hotmail.com", "@outlook.com", "@valladolid.tecnm.mx"}
	n := len(mails)
	return mails[r.Intn(n)]
}

func RandomUserEmail() string {
	return RandomString(6) + RandomMail()
}

func RandomUserPassword() string {
	return RandomString(8)
}
