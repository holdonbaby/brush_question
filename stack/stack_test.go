package stack

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStack(t *testing.T) {
	Convey("Stack test pull and push", t, func() {
		q := NewStack()
		q.Push(1).Push(10).Push(30)
		So(q.Pull(), ShouldEqual, 30)
		So(q.Pull(), ShouldEqual, 10)
		So(q.Pull(), ShouldEqual, 1)
	})
}
