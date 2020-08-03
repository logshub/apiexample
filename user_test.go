package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUser(t *testing.T) {
	Convey("Testing user struct", t, func() {
		u := &User{}
		So(u.isValid(), ShouldEqual, false)
		So(u.isValidID(), ShouldEqual, false)
		u.UserID = "123"
		So(u.isValidID(), ShouldEqual, false)

		u.UserID = "79fe3975-52b3-4d8a-856e-5c326309654a"
		So(u.isValidID(), ShouldEqual, true)
		// name not set yet
		So(u.isValid(), ShouldEqual, false)
		u.Name = "Brzenczeszczykiewicz"
		So(u.isValid(), ShouldEqual, true)
	})
}
