package utils

func GenerateToken(...any) string {
	return "Hello World"
}

func CheckToken(data []byte) bool {
	// fmt.Printf("CheckToken: %s\n", string(data))
	return string(data) == "Hello World"
}
