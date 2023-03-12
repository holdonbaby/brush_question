package queue

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQueue(t *testing.T) {
	Convey("queue test pull and push", t, func() {
		q := NewQueue()
		q.Push(1).Push(10).Push(30)
		So(q.Pull(), ShouldEqual, 1)
		So(q.Pull(), ShouldEqual, 10)
		So(q.Pull(), ShouldEqual, 30)
	})
}
