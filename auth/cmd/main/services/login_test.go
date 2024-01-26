package services

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	data := []struct {
		input LoginInput
		t     string
	}{
		{
			LoginInput{
				Email:    "unittest@gmail.com",
				Password: "mauvaisMdp",
			}, "d"},
		{LoginInput{
			Email:    "unittest@gmail.com",
			Password: "mauvaisMdp",
		}, "p"},
	}

	for _, tt := range data {
		testname := fmt.Sprintf("Login as %s", tt.t)
		t.Run(testname, func(t *testing.T) {
			_, err := Login(tt.input, tt.t)
			if err == nil {
				t.Errorf("expected error but got success")
			}
		})
	}
}
