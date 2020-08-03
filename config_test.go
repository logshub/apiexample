package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfig(t *testing.T) {
	Convey("Testing initialization of config", t, func() {
		conf := &Config{}
		err := initConfig("/etc/non-existingfile.toml", conf)
		So(err, ShouldNotBeNil)

		err = initConfig("./config.example.toml", conf)
		So(err, ShouldBeNil)

		So(len(conf.Port), ShouldBeGreaterThan, 0)
		So(len(conf.Database.Type), ShouldBeGreaterThan, 0)
	})
}
