package helper

import (
	"crypto/rand"
	"strconv"
)

func RandomNumbers(length int) (int, error) {
	const numbers = "1234567890"
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return 0, err
	}
	numberLength := len(numbers)
	for i := 0; i < length; i++ {
		buffer[i] = numbers[int(buffer[i])%numberLength]
	}
	return strconv.Atoi(string(buffer))
}
