package coalesce

// Coalesce is a simple function that returns the first non-nil value.
func Coalesce(values ...interface{}) interface{} {
	for _, val := range values {
		if val != nil {
			return val
		}
	}

	return nil
}
