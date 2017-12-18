package consistent

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("Give some integer with a starting value", t, func() {
		// c := ConsistentRing{ //单元测试数据1
		// 	Range: 100,
		// 	Nodes: []int64{
		// 		10, 20, 3, 111, 270,
		// 	},
		// }
		c := ConsistentRing{
			Range: 100,
			Nodes: []Node{},
		}
		c.AddNode("aaa")
		c.AddNode("bbb")
		c.AddNode("ccc")
		c.DumpNodesRange()
		So(nil, ShouldEqual, nil)
	})
}
