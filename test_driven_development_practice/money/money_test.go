package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(int64(5))
	product := five.Times(2)
	assert.Equal(t, int64(10), product.Amount)
}

func TestEquality(t *testing.T) {
	assert.True(t, NewDollar(int64(5)).EqualsTo(NewDollar(int64(5))))
}
