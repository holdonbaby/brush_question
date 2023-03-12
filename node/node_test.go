package node

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNode(t *testing.T) {
	Convey("node test new and newWithContext", t, func() {
		node := NewNode()
		nodeWithContext := NewNodeWithContent(node, nil, "爱你一万年", 10000)
		node.Right = nodeWithContext
		So(node.ValueStr, ShouldEqual, "")
		So(nodeWithContext.Left, ShouldEqual, node)
	})
}
