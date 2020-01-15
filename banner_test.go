package go_util

import "testing"

func TestBanner_Print(t *testing.T) {
	NewBanner("s","1.0","debug").Print()
}