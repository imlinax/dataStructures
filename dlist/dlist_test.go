package dlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dlist(t *testing.T) {
	var dl DList

	assert.Equal(t, 0, dl.Len())
	assert.Equal(t, true, dl.IsEmpty())

	dl.Init()
	assert.Equal(t, 0, dl.Len())
	assert.Equal(t, true, dl.IsEmpty())

	dl.PushBack(1)
	assert.Equal(t, 1, dl.Len())
	assert.Equal(t, 1, dl.Front().value)
	assert.Equal(t, 1, dl.Back().value)
	assert.Equal(t, false, dl.IsEmpty())

	dl.PushFront(2)
	assert.Equal(t, 2, dl.Len())
	assert.Equal(t, false, dl.IsEmpty())
	assert.Equal(t, 2, dl.Front().value)
	assert.Equal(t, 1, dl.Back().value)
}
