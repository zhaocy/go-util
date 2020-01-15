package go_util

import "testing"

func TestBanner_Print(t *testing.T) {
	newBanner := NewBanner("server_name", "1.0.0", "debug")
	newBanner.Print()
}