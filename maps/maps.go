package maps

// GetOrDefault returns the value stored with key in map m, if it exists;
// otherwise returns defaultValue.
func GetOrDefault[TKey comparable, TValue any](m map[TKey]TValue, key TKey, defaultValue TValue) TValue {
	if value, ok := m[key]; ok {
		return value
	}
	return defaultValue
}

// HasKey determines if key exists in map m.
func HasKey[TKey comparable, TValue any](m map[TKey]TValue, key TKey) bool {
	_, ok := m[key]
	return ok
}

// Keys returns a slice of keys in map m.
func Keys[TKey comparable, TValue any](m map[TKey]TValue) []TKey {
	result := make([]TKey, len(m))
	i := 0
	for key := range m {
		result[i] = key
		i++
	}
	return result
}

// Values returns a slice of values in map m.
func Values[TKey comparable, TValue any](m map[TKey]TValue) []TValue {
	result := make([]TValue, len(m))
	i := 0
	for _, value := range m {
		result[i] = value
		i++
	}
	return result
}
