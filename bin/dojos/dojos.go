package dojos

import (
	"bufio"
	"os"
	"time"
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

var WriteFile = func(name, content string) (int, error) {
	f, err := os.Create(name)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString(content)
	w.Flush()
	return 0, nil
}
