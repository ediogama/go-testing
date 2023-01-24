package tax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tax := CalculateTax(1000.0)
	assert.Equal(t, 10.0, tax)
}
