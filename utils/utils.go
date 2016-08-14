package utils

// Check ...
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
