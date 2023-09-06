package utils

import (
	"fmt"
	"log"
)

// func mapConcat[K comparable, T any](m1 *map[K]T, m2 map[K]T) {
// 	for key, value := range m2 {
// 		m := *m1
// 		m[key] = value
// 	}
// }

// log formated string
func LogNonFatal(format string, a ...any) {
	log.Print("[ERROR]: ", fmt.Sprintf(format, a...))
}

// logs formated string or error and exits
// accepts error as error.String()
func LogFatal(error string, a ...any) {
	log.Fatal("[ERROR]: ", fmt.Sprintf(error, a...))
}
