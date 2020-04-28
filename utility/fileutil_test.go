package opshelper

import "testing"

func TestCountLeadingSpace(t *testing.T) {
	err := InsertStringToFile("test.txt", "hello world1!\nhello world 2!\n", 2)
	if err != nil {
		t.Error(err)
	}
}
