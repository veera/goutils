package maps

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func Invert[K comparable, V comparable](m map[K]V) map[V]K {
	inv := make(map[V]K)

	for k, v := range m {
		inv[v] = k
	}

	return inv
}
