package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	var q Queue[int]
	assert.True(t, q.IsEmpty(), "queue should be empty at first")

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.False(t, q.IsEmpty(), "queue should not be empty after enqueue")
	assert.Equal(t, 1, q.Dequeue())
	assert.Equal(t, 2, q.Dequeue())
	assert.Equal(t, 3, q.Dequeue())
	assert.True(t, q.IsEmpty(), "queue should be empty after dequeue all")
}
