package hashid

import (
	"fmt"
	"os"
	"strconv"

	"github.com/speps/go-hashids/v2"
)

var (
	hashIDSalt      = env("HASHID_SALT", "!@#!@F234t23g2354rr2tg35")
	hashIDMinLength = envInt("HASHID_MIN_LENGTH", 8)
	hashIDAlphabet  = env("HASHID_ALPHABET", "qwertyuiopasdfghjklzxcvbnm1234567890")

	hashIDCoder *hashids.HashID
)

func init() {
	var err error
	hashIDCoder, err = hashids.NewWithData(&hashids.HashIDData{
		Alphabet:  hashIDAlphabet,
		MinLength: hashIDMinLength,
		Salt:      hashIDSalt,
	})
	if err != nil {
		panic(err)
	}
}

func encode(number int64) (result string, err error) {
	result, err = hashIDCoder.EncodeInt64([]int64{number})
	return
}

func decode(hash string) (number int64, err error) {
	result, err := hashIDCoder.DecodeInt64WithError(hash)
	if err == nil {
		if len(result) != 0 {
			number = result[0]
		}
	}
	return
}

func env(name string, defaultValue string) string {
	val, ok := os.LookupEnv(name)
	if !ok || val == "" {
		return defaultValue
	}

	return val
}

func envInt(name string, defaultValue int) int {
	val, ok := os.LookupEnv(name)
	if !ok || val == "" {
		return defaultValue
	}

	if intVal, err := strconv.Atoi(val); err == nil {
		return intVal
	} else {
		return defaultValue
	}
}

func print(args ...interface{}) {
	fmt.Println(args...)
}
