package utils

import (
	"fmt"
	"log"
)

func LogNonFatal(format string, a ...any) {
	log.Print("[ERROR]: ", fmt.Sprintf(format, a...))
}

func LogFatal(error string, a ...any) {
	log.Fatal("[ERROR]: ", fmt.Sprintf(error, a...))
}
