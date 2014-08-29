package main

import (
	"./ifaces/ibufio"
	"./ifaces/ios"
	"./ifaces/itime"
)

func WriteFile(name, content string) (int, error) {
	f, err := ios.Create(name)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	w := ibufio.NewWriter(f)
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
