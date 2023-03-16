package xvalidator

import (
	"errors"
	"strings"
)

func ConsistsOfSpaces(data []string) error {
	for _, str := range data {
		if len(strings.TrimSpace(str)) == 0 {
			return errors.New("string consists of spaces")
		}
	}
	return nil
}

func ValidateAndConvertAnimalTypes(animalTypes []*int64) ([]int64, error) {
	if len(animalTypes) == 0 {
		return nil, errors.New("empty slice of animal types")
	}

	convAnimalTypes := make([]int64, len(animalTypes))

	for i, animalType := range animalTypes {
		if animalType != nil && *animalType > 0 {
			convAnimalTypes[i] = *animalTypes[i]
		} else {
			return nil, errors.New("invalid animal type id")
		}
	}

	return convAnimalTypes, nil
}
