package dict

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Dict_Get(t *testing.T) {
	d := New()

	assert.Nil(t, d.Get(""))
	assert.Nil(t, d.Get("a"))
	assert.Nil(t, d.Get("abc"))
}

func Test_Dict_Set(t *testing.T) {
	d := New()

	d.Set("abc", 33)
	assert.Equal(t, 33, d.Get("abc"))

	d.Set("ab", 44)
	assert.Equal(t, 44, d.Get("ab"))

	d.Set("ba", 55)
	assert.Equal(t, 55, d.Get("ba"))

	d.Set("b", 66)
	assert.Equal(t, 66, d.Get("b"))

	d.Set("ac", 87)
	assert.Equal(t, 87, d.Get("ac"))

	d.Set("ad", 88)
	assert.Equal(t, 88, d.Get("ad"))

	d.Set("ac", 187)
	assert.Equal(t, 187, d.Get("ac"))

	assert.Nil(t, d.Get("abcd"))
	assert.Nil(t, d.Get("azcd"))
}