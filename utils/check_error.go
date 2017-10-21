package utils

import "fmt"

// CheckError ...
func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
