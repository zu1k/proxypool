package tool

func CheckInList(list []string, item string) bool {
	for _, i := range list {
		if item == i {
			return true
		}
	}
	return false
}
