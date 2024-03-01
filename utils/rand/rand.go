package rand

import (
	"math/rand"
	"time"
)

var randomGenerator *rand.Rand

func init() {
	randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// GetDelayDuration returns delay duration from 1 μs to 10ms.
func GetDelayDuration() time.Duration {
	return time.Duration(randomGenerator.Intn(10000)+1) * time.Microsecond
}
