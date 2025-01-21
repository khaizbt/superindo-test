package helper

import "crypto/rand"

func GenerateNumber(length int) (string, error) {
	var chars = "0123456789"
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "0", err
	}

	charsLength := len(chars)
	for i := 0; i < length; i++ {
		buffer[i] = chars[int(buffer[i])%charsLength]
	}

	return string(buffer), nil
}
