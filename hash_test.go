package wphash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckWordPressPasswordHash(t *testing.T) {
	// true case 123456 -> $P$BmIaPlVaAl6kEsffVZGdASCVH.i1cZ0
	ret := CheckWordPressPasswordHash("123456", "$P$BmIaPlVaAl6kEsffVZGdASCVH.i1cZ0")
	assert.Equal(t, true, ret)
	// false case 123456 -> $P$B4VXOnAaJ9nC10J0bJ8jnBxcP2L6Iv0
	ret2 := CheckWordPressPasswordHash("123456", "$P$B4VXOnAaJ9nC10J0bJ8jnBxcP2L6Iv0")
	assert.Equal(t, false, ret2)
}

func TestHashPassword(t *testing.T) {
	// case 123456 -> <dynamic hash> -> check match
	ret := HashPassword("123456")
	ret2 := CheckWordPressPasswordHash("123456", ret)
	assert.Equal(t, true, ret2)
}
