package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructEmbed(t *testing.T) {
	type structEmbedParent struct {
		ParentA string
		ParentB string
	}
	type structEmbedChild struct {
		structEmbedParent
		ChildA string
		ChildB string
	}

	// The following code will not compile
	// child := structEmbedChild{
	// 	ParentA: "ParentA",
	// 	ParentB: "ParentB",
	// 	ChildA:  "ChildA",
	// 	ChildB:  "ChildB",
	// }

	// The following code will compile
	child := structEmbedChild{
		structEmbedParent: structEmbedParent{
			ParentA: "ParentA",
			ParentB: "ParentB",
		},
		ChildA: "ChildA",
		ChildB: "ChildB",
	}
	assert.Equal(t, "ParentA", child.ParentA)
	assert.Equal(t, "ParentB", child.ParentB)
	assert.Equal(t, "ChildA", child.ChildA)
	assert.Equal(t, "ChildB", child.ChildB)
}
