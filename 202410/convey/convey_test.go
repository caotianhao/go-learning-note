package convey

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_add(t *testing.T) {
	Convey("test add", t, func() {
		res := add(1, 2)
		So(res, ShouldEqual, 3)
		So(res, ShouldBeGreaterThan, 2)
	})
}
