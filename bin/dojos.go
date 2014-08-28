package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/spf13/cobra"
)

var GetSessName = func(args []string) string {
	const defaultSessName = "20060102"
	sessName := time.Now().Format(defaultSessName)
	if len(args) > 0 {
		sessName = args[0]
	}
	return sessName
}

var GetCurDir = func() (string, error) {
	return os.Getwd()
}

var MkDir = func(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

func InitAction(args []string) {
	sessionName := GetSessName(args)
	cwd, err := GetCurDir()
	if err != nil {
		log.Fatal(err)
		return
	}
	dir := path.Join(cwd, sessionName)
	// TODO : Use a debug message for this. Or a way to filter with quiet and verbose flags.
	fmt.Println("Init: " + dir)
	MkDir(dir, 0777)
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
