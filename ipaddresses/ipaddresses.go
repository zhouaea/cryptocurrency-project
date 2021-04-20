package ipaddresses

func GetMiners() []string {
	return []string{"127.0.0.1:1232", "127.0.0.1:1233", "127.0.0.1:1234"}
}

func GetClients() []string {
	return []string{"127.0.0.1:1235", "127.0.0.1:1236"}
}

func GetController() string {
	return "127.0.0.1:1236"
}
