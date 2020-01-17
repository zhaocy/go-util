package go_util

import (
	"fmt"
	"testing"
)

func TestElo_rating_system(t *testing.T) {
	A, B := Elo_rating_system(1000, 1000, 2, 2)
	fmt.Println(A," - ",B)
}