package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(int64(5))
	assert.Equal(t, NewFranc(int64(10)), five.Times(int64(2)))
	assert.Equal(t, NewFranc(int64(15)), five.Times(int64(3)))
}
