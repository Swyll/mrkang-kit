package utils

import "testing"

func TestCapPanic(t *testing.T) {
	err := CapPanic("/root/go/src/mrkang-kit/utils/std_err.log")
	if err != nil {
		t.Error(err)
	}

	panic("llllssss===========")
}
