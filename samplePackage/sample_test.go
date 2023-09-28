package samplePackage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("should always return 5", func(t *testing.T) {
		assert.Equal(t, 5, add())
	})
}
