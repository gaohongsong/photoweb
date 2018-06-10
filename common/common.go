package common

import "os"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return os.IsExist(err)
}
