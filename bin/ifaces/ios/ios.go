package ios

import "os"

var Getwd = func() (string, error) {
	return os.Getwd()
}

var Mkdir = func(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

var Create = func(name string) (*os.File, error) {
	return os.Create(name)
}
