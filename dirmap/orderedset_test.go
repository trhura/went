package dirmap

import "testing"
import "github.com/stretchr/testify/assert"

func TestNew(t *testing.T) {
	s := NewOrderedSet()
	assert.Equal(t, 0, s.Len(), "Newly Ringset Len must be zero.")
}

func TestPrepend(t *testing.T) {
	s := NewOrderedSet()
	s.Prepend(3)
	s.Prepend(2)
	s.Prepend(1)
	s.Prepend(0)

	assert.Equal(t, s.Get(0), 0)
	assert.Equal(t, s.Get(1), 1)
	assert.Equal(t, s.Get(2), 2)
	assert.Equal(t, s.Get(3), 3)
	assert.Equal(t, s.Len(), 4)
}

func TestPush(t *testing.T) {
	s := NewOrderedSet()
	s.Push(3)
	s.Push(0)
	s.Push(2)
	s.Push(1)
	s.Push(0)

	assert.Equal(t, s.Get(0), 0)
	assert.Equal(t, s.Get(1), 1)
	assert.Equal(t, s.Get(2), 2)
	assert.Equal(t, s.Get(3), 3)
	assert.Equal(t, s.Len(), 4)
}
