package stringslice

// Contains ...
func Contains(slice []string, str string) bool {
	for _, value := range slice {
		if value == str {
			return true
		}
	}
	return false
}

// Delete ...
func Delete(slice *[]string, str string) bool {
	for i := 0; i < len(*slice); i++ {
		if (*slice)[i] == str {
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
			return true
		}
	}
	return false
}
