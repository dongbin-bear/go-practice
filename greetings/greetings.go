package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (message string, err error) {
	if name == "" {
		err = errors.New("empty name")
		return
	}
	message = fmt.Sprintf(randomFormat(), name)
	// message = fmt.Sprintf(randomFormat())
	return
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
