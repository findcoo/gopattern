package gopattern

import "testing"

func TestDisplayBanner(t *testing.T) {
	banner := NewDisplayBanner("welcome to my story")
	banner.PrintStrong()
	banner.PrintWeak()
}
