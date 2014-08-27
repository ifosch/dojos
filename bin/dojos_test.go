package main

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInit(t *testing.T) {
	var args []string
	var DirName string
	getSessionName = func(args []string) string {
		if len(args) > 0 {
			return args[0]
		} else {
			return "20140827"
		}
	}
	getDirectory = func() (string, error) { return "/tmp", nil }
	makeDirectory = func(name string, perm os.FileMode) error {
		DirName = name
		return nil
	}
	Convey("Given no entry", t, func() {
		Convey("When InitAction is called", func() {
			initAction(args)
			Convey("Directory with current date is created", func() {
				So(DirName, ShouldEqual, "/tmp/20140827")
			})
		})
	})
	Convey("Given an entry", t, func() {
		args = []string{"session"}
		Convey("When InitAction is called", func() {
			initAction(args)
			Convey("Directory with current date is created", func() {
				So(DirName, ShouldEqual, "/tmp/session")
			})
		})
	})
}
