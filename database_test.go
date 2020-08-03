package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDatabase(t *testing.T) {
	Convey("Testing initialization of database", t, func() {
		conf := &Config{
			Database: configDatabase{
				Type: "postgres",
			},
		}
		_, err := initDatabase(conf)
		So(err, ShouldNotBeNil)
	})
}
