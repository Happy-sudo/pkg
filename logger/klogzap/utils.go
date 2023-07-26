package klogzap

// inArray check if a string in a slice
func inArray(key ExtraKey, arr []ExtraKey) bool {
	for _, k := range arr {
		if k == key {
			return true
		}
	}
	return false
}
