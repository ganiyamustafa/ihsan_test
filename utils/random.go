package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func RandStringNumber(totalNum int) string {
	rand.Seed(time.Now().UnixNano())
	randString := ""

	for i := 0; i < totalNum; i++ {
		randString += fmt.Sprintf("%d", rand.Intn(9))
	}

	return randString
}
