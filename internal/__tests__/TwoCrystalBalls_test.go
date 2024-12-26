package internal

import (
	alg "github.com/vinewz/alGOs/src/_today"
	"math/rand"
	"testing"
	"time"
)

func TestTwoCrystalBalls(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(10000)
	data := make([]bool, 10000)

	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	result := alg.TwoCrystalBalls(data)
	if result != idx {
		t.Errorf("Expected %v, but got %v", idx, result)
	}

	result = alg.TwoCrystalBalls(make([]bool, 821))
	if result != -1 {
		t.Errorf("Expected %v, but got %v", -1, result)
	}
}
