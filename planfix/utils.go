package planfix

func stringInSlice(slice []string, search string) bool {
	for _, elem := range slice {
		if search == elem {
			return true
		}
	}
	return false
}
