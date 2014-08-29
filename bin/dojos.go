package main

import (
	"fmt"
	"log"
	"path"

	"github.com/spf13/cobra"

	"./ifaces/ios"
)

const defaultSessName = "20060102"

func InitAction(args []string) {
	// TODO : Use a template, instead of such an ugly constant like this one
	const pythonTestContent = "import unittest\nclass Test1(unittest.TestCase):\n  pass\n\nif __name__ == \"__main__\":\n  unittest.main()"
	sessionName := GetSessName(args)
	cwd, err := ios.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}
	dir := path.Join(cwd, sessionName)
	// TODO : Use a debug message for this. Or a way to filter with quiet and verbose flags.
	fmt.Println("Init: " + dir)
	ios.Mkdir(dir, 0777)
	WriteFile(path.Join(dir, "tests.py"), pythonTestContent)
}

var initCmd = &cobra.Command{
	Use:   "init [session name]",
	Short: "Initializes a directory for the dojo session, using current date or specified name.",
	Run: func(cmd *cobra.Command, args []string) {
		InitAction(args)
	},
}

var dojosCmd = &cobra.Command{
	Use:   "dojos",
	Short: "Dojos is a tool to manipulate dojo sessions",
}

func main() {
	dojosCmd.AddCommand(initCmd)
	dojosCmd.Execute()
}
