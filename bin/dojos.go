package main

import (
	"fmt"
	"log"
	"path"

	"github.com/spf13/cobra"

	"./ifaces"
	"./ifaces/itime"
)

const defaultSessName = "20060102"

func WriteFile(name, content string) (int, error) {
	f, err := ifaces.Create(name)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	w := ifaces.NewWriter(f)
	bytes, err := w.WriteString(content)
	if err != nil {
		return bytes, err
	}
	w.Flush()
	return bytes, nil
}

func GetSessName(args []string) string {
	sessName := itime.Now(defaultSessName)
	if len(args) > 0 {
		sessName = args[0]
	}
	return sessName
}

func InitAction(args []string) {
	const pythonTestContent = "import unittest\nclass Test1(unittest.TestCase):\n  pass\n\nif __name__ == \"__main__\":\n  unittest.main()"
	sessionName := GetSessName(args)
	cwd, err := ifaces.GetCurDir()
	if err != nil {
		log.Fatal(err)
		return
	}
	dir := path.Join(cwd, sessionName)
	// TODO : Use a debug message for this. Or a way to filter with quiet and verbose flags.
	fmt.Println("Init: " + dir)
	ifaces.MkDir(dir, 0777)
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
