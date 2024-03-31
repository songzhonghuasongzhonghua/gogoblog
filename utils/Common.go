package utils

func InList(key string, list []string) bool {
	for _, l := range list {
		if key == l {
			return true
		}
	}

	return false

}
