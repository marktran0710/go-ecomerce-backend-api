package random

import (
	"math/rand"
	"time"
)

func GenerateSixDigitOtp() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	opt := 100000 + rng.Intn(900000) // 100000 - 999999
	return opt
}
