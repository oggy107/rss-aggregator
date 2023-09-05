package main

func mapConcat[K comparable, T any](m1 *map[K]T, m2 map[K]T) {
	for key, value := range m2 {
		m := *m1
		m[key] = value
	}
}
