package utils

import (
	"log"
)

// CheckError ...
func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}
