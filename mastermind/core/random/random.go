package random

import (
	"math/rand"
	"time"
	"errors"
)

var (
	sizeLimitPositiveGreatThanZeroError = "Size and Limit should be positive and great than 0"
	sizeGreatThanLimitWithoutRepetitionsError = "Size great than limit without repetitions, will cause slice bounds out of range"
)

// Generates an array with amount of random numbers with or without repetitions
// size: size of array
// limit: between 0 and LIMIT
// repetition: allow repetitions
// Example: size=5, limit=5, repetition=false
// Result: [3][1][4][2][0]
// Example: size=5, limit=5, repetition=true
// Result: [3][1][0][2][0]
func Random(size, limit int, repetition bool) ([]int, error) {
	err := validateRandom(size, limit, repetition)
	if err != nil {
		return nil, err
	}

	rand.Seed(int64(time.Now().Nanosecond()))
	if repetition {
		randomArray := make([]int, size)

		for i := 0; i < size; i++ {
			randomArray[i] = rand.Intn(limit)
		}
		return randomArray, nil
	} else {
		return rand.Perm(limit)[:size], nil
	}
}

func validateRandom(size, limit int, repetition bool) (error) {
	if size <= 0 || limit <= 0 {
		return errors.New(sizeLimitPositiveGreatThanZeroError)
	}

	if !repetition && size > limit {
		return errors.New(sizeGreatThanLimitWithoutRepetitionsError)
	}
	return nil
}