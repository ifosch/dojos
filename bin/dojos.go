package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/spf13/cobra"
)

var getSessionName = func(args []string) string {
	const defaultSessionName = "20060102"
	SessionName := time.Now().Format(defaultSessionName)
	if len(args) > 0 {
		SessionName = args[0]
	}
	return SessionName
}

var getDirectory = func() (string, error) {
	return os.Getwd()
}

var makeDirectory = func(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

func initAction(args []string) {
	sessionName := getSessionName(args)
	cwd, err := getDirectory()
	if err != nil {
		log.Fatal(err)
	} else {
		dir := path.Join(cwd, sessionName)
		fmt.Println("Init: " + dir)
		makeDirectory(dir, 0777)
	}
}

var DojosCmd = &cobra.Command{
	Use:   "dojos",
	Short: "Dojos is a tool to manipulate dojo sessions",
}

var InitCmd = &cobra.Command{
	Use:   "init [session name]",
	Short: "Initializes a directory for the dojo session, using current date or specified name.",
	Run: func(cmd *cobra.Command, args []string) {
		initAction(args)
	},
}

func main() {
	DojosCmd.AddCommand(InitCmd)
	DojosCmd.Execute()
}
