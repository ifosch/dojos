package main

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"./dojos"
)

func TestInit(t *testing.T) {
	const pythonTestContent = "import unittest\nclass Test1(unittest.TestCase):\n  pass\n\nif __name__ == \"__main__\":\n  unittest.main()"
	var args []string
	var dir string
	var test_filename string
	var test_content string
	dojos.GetSessName = func(args []string) string {
		if len(args) > 0 {
			return args[0]
		}
		return "20140827"
	}
	dojos.GetCurDir = func() (string, error) { return "/tmp", nil }
	dojos.MkDir = func(name string, perm os.FileMode) error {
		dir = name
		return nil
	}
	dojos.WriteFile = func(name, content string) (int, error) {
		test_filename = name
		test_content = content
		return len(content), nil
	}
	Convey("Given no entry", t, func() {
		Convey("When InitAction is called", func() {
			InitAction(args)
			Convey("Directory with current date is created", func() {
				So(dir, ShouldEqual, "/tmp/20140827")
			})
			Convey("test.py file is created into session directory", func() {
				So(test_filename, ShouldEqual, "/tmp/20140827/tests.py")
				So(test_content, ShouldEqual, pythonTestContent)
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
