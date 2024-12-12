package rbmv

func NewNode(peer string) *Node {
	return &Node{
		shards:   make(ShardQueue, 0),
		shardMap: make(map[string]struct{}),
		mem:      0,
		peer:     peer,
	}
}

type Node struct {
	shards   ShardQueue
	shardMap map[string]struct{}
	mem      float64
	peer     string
}

func (n *Node) Push(shard *Shard) {
	n.shards = append(n.shards, shard)
	n.mem += shard.mem
	n.shardMap[shard.name] = struct{}{}
}

func (n *Node) Pop() *Shard {
	old := n.shards
	tail := len(old) - 1
	item := old[tail]
	old[tail] = nil // don't stop the GC from reclaiming the item eventually
	n.shards = old[0:tail]
	return item
}

func (n *Node) HasShards() bool {
	return len(n.shards) > 0
}

func (n *Node) ExistShard(name string) bool {
	_, ok := n.shardMap[name]
	return ok
}

type ShardQueue []*Shard

func NewShard(name string, mem float64) *Shard {
	return &Shard{
		name: name,
		mem:  mem,
	}
}

type Shard struct {
	name string
	mem  float64
}
