package main

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInit(t *testing.T) {
	var args []string
	var dir string
	GetSessName = func(args []string) string {
		if len(args) > 0 {
			return args[0]
		}
		return "20140827"
	}
	GetCurDir = func() (string, error) { return "/tmp", nil }
	MkDir = func(name string, perm os.FileMode) error {
		dir = name
		return nil
	}
	Convey("Given no entry", t, func() {
		Convey("When InitAction is called", func() {
			InitAction(args)
			Convey("Directory with current date is created", func() {
				So(dir, ShouldEqual, "/tmp/20140827")
			})
		})
	})
	Convey("Given an entry", t, func() {
		args = []string{"session"}
		Convey("When InitAction is called", func() {
			InitAction(args)
			Convey("Directory with current date is created", func() {
				So(dir, ShouldEqual, "/tmp/session")
			})
		})
	})
}
