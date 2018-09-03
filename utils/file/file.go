package file

import "os"

func MkdirIfNotExist(path string) error {
	exists := PathExist(path)
	if !exists {
		return os.MkdirAll(path, 0711)

	}
	return nil
}

func PathExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
