package main

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
func logError(format string, a ...any) {
	log.Print("[ERROR]: ", fmt.Sprintf(format, a...))
}

// logs formated string and exits
func logFatal(format string, a ...any) {
	log.Fatal("[ERROR]: ", fmt.Sprintf(format, a...))
}

// logs error and exits
func logFatalError(err error) {
	log.Fatal("[ERROR]: ", err)
}
