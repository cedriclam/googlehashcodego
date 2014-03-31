package utils

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestGetInput(t *testing.T){
	word := "foo"
	isfoo, err := GetInput(word)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t,true,isfoo)
}

