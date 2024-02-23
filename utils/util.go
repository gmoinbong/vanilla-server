package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func CityInputReader() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the city: ")

	city, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return "", errors.New("unexpected end of input, please try again")
		}
		return "", fmt.Errorf("error while read city value: %v", err)
	}

	city = strings.TrimSpace(city)
	return city, nil
}
