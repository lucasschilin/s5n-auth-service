package util

func InStringSlice(slice []string, target string) (bool, int) {
	for key, value := range slice {
		if value == target {
			return true, key
		}
	}

	return false, -1

}
