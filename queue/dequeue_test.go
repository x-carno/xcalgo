package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeQueue(t *testing.T) {
	tdq := New()

	err := tdq.AddFirst(nil)
	assert.EqualError(t, err, ErrNilItem.Error())
	err = tdq.AddLast(nil)
	assert.EqualError(t, err, ErrNilItem.Error())

	tdq.AddFirst("f1")
	tdq.AddLast("l1")
	iter, err := tdq.Iterator()
	assert.Nil(t, err, nil)
	assert.Equal(t, []interface{}{"f1", "l1"}, iter)

	first, err := tdq.RemoveFirst()
	assert.Nil(t, err, nil)
	assert.Equal(t, "f1", first)

	last, err := tdq.RemoveLast()
	assert.Nil(t, err, nil)
	assert.Equal(t, "l1", last)

	_, err = tdq.RemoveFirst()
	assert.EqualError(t, err, ErrDeQueueIsEmpty.Error())
	_, err = tdq.RemoveLast()
	assert.EqualError(t, err, ErrDeQueueIsEmpty.Error())

	_, err = tdq.Iterator()
	assert.EqualError(t, err, ErrDeQueueIsEmpty.Error())

	assert.Equal(t, true, tdq.IsEmpty())
}

func BenchmarkAddFirst(b *testing.B) {
	tdq := New()
	for i := 0; i < b.N; i++ {
		tdq.AddFirst(1)
	}
}

func BenchmarkAddLast(b *testing.B) {
	tdq := New()
	for i := 0; i < b.N; i++ {
		tdq.AddLast(1)
	}
}
