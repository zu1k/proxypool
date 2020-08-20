package provider

type Provider interface {
	Provide() string
}

func checkInList(list []string, item string) bool {
	for _, i := range list {
		if item == i {
			return true
		}
	}
	return false
}
