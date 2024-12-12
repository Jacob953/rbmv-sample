package rbmv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReBalance_NoMove(t *testing.T) {
	saveNodes := []*Node{NewNode("node1")}
	delNodes := make([]*Node, 0)
	res, exist := ReBalance(saveNodes, delNodes)
	assert.Nil(t, res)
	assert.False(t, exist)
}

func TestReBalance_Move(t *testing.T) {
	toPeer, fromPeer := "node1", "node2"
	saveNodes := []*Node{NewNode(toPeer)}
	delNodes := []*Node{NewNode(fromPeer)}
	moveShard := NewShard("shard1", 1)
	delNodes[0].Push(moveShard)
	res, exist := ReBalance(saveNodes, delNodes)
	assert.Equal(t, moveShard, res.shard)
	assert.Equal(t, true, exist)
	assert.True(t, exist)
}
