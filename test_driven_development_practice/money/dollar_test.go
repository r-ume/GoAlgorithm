package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(int64(5))
	assert.Equal(t, NewDollar(int64(10)), five.Times(2))
	assert.Equal(t, NewDollar(int64(15)), five.Times(3))
}

func TestEquality(t *testing.T) {
	assert.True(t, NewDollar(int64(5)).EqualsTo(NewDollar(int64(5))))
}
