package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, ReturnValue(), "ozoohaih")
}
