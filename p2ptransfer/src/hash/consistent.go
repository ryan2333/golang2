package consistent

import (
	"fmt"
	"hash/fnv"
	"log"
	"sort"
)

type Bucket struct {
	Name     string
	Position uint64
}

type ConsistentRing struct {
	Range   uint64
	Buckets []Bucket
}

// type ConsistentRing struct {  //单元测试数据1
// 	Range int64
// 	Nodes []int64
// }

// type ConsistentRange struct {
// 	Node  int64
// 	Start int64
// 	End   int64
// }

func (a ConsistentRing) Len() int {
	return len(a.Buckets)
}

func (a ConsistentRing) Swap(i, j int) {
	a.Buckets[i], a.Buckets[j] = a.Buckets[j], a.Buckets[i]
}

// func (a ConsistentRing) Less(i, j int) bool {  //单元测试数据1
// 	return a.Nodes[i]%a.Range < a.Nodes[j]%a.Range
// }

func (a ConsistentRing) Less(i, j int) bool {
	return (a.Buckets[i].Position) < (a.Buckets[j].Position)
}

func (c *ConsistentRing) AddNode(name string) {
	h := fnv.New64() //取hash值
	h.Write([]byte(name))
	c.Buckets = append(c.Buckets, Bucket{name, uint64(h.Sum64()) % c.Range})
	log.Println(c.Buckets)
}

// func (c ConsistentRing) AddNode(node int64) {  //单元测试数据1
// 	c.Nodes = append(c.Nodes, node)
// }

// func (c ConsistentRing) GetNodeRange(node int64) (r ConsistentRange) {
// 	return ConsistentRange{}
// }

func (c ConsistentRing) DumpNodesRange() ConsistentRing {
	sort.Sort(c)
	fmt.Printf("%v\n", c)
	return c
}

func (c ConsistentRing) findBucketByKey(key string) Bucket {
	h := fnv.New64()
	h.Write([]byte(key))
	keyHash := uinit64(h.Sum64()) % c.Range
	bkHash := c.Buckets
	for i := 0; i < len(bkHash); i++ {
		if bkHash[i].Position > keyHash {
			return bkHash[i-1]
		} else if bkHash[i].Position < keyHash && i == len(bkHash)-1 {
			return bkHash[len(bkHash)-1]
		}
	}
}
