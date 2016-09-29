package elements

import (
	"log"
	"math/rand"
	"time"
	"reflect"
	"github.com/iris-contrib/errors"
)

type Board struct {
	Codes      []int
	Columns    int
	Rows       int
	Matrix     map[int][]int
	Repetition bool
	checks     int
}

func (b *Board) Check(codes []int) (_, _, misses int, err error) {
	log.Println("Check codes:", codes, "Codes:", b.Codes)

	var hitsOutOfOrderCheck, hitsCheck []int
	if codes == nil || b.Codes == nil || len(codes) != len(b.Codes) {
		return 0, 0, 0, errors.New("Divergence of size between the codes to check and the board")
	}
	err = checkMatrix(b, codes)
	if err != nil {
		return 0, 0, 0, err
	}

	for i := range codes {
		for x := range b.Codes {
			if i == x && codes[i] == b.Codes[x] {
				hitsCheck = append(hitsCheck, i)
			} else if codes[i] == b.Codes[x] && !contains(hitsOutOfOrderCheck, i) {
				hitsOutOfOrderCheck = append(hitsOutOfOrderCheck, i)
			}
		}
	}
	for i := range codes {
		if !contains(hitsCheck, i) && !contains(hitsOutOfOrderCheck, i) {
			misses++
		}
	}
	return len(hitsCheck), len(hitsOutOfOrderCheck), misses, nil
}

func checkMatrix(b *Board, codes []int) (error) {
	if b.checks >= b.Rows {
		return errors.New("All checks already done.")
	}
	b.Matrix[b.checks] = codes
	b.checks++
	return nil
}

func Create(columns, rows, codeLength int, repetition bool) (Board) {
	b := Board{Codes: random(rows, codeLength, repetition), Columns: columns, Rows: rows, Repetition: repetition, Matrix: make(map[int][]int)}

	log.Println("Columns:", b.Columns, "Rows:", b.Rows, "Code Length:", codeLength, "Repetition:", b.Repetition, "Codes:", b.Codes)

	return b
}

func random(rows, codeLength int, repetition bool) []int {
	rand.Seed(int64(time.Now().Nanosecond()))
	if repetition {
		randWithRepetition := make([]int, rows)

		for i := 0; i < rows; i++ {
			randWithRepetition[i] = rand.Intn(codeLength)
		}

		return randWithRepetition
	} else {
		return rand.Perm(codeLength)[:rows]
	}
}

func contains(slice interface{}, val interface{}) bool {
	if slice != nil {
		sv := reflect.ValueOf(slice)

		for i := 0; i < sv.Len(); i++ {
			if sv.Index(i).Interface() == val {
				return true
			}
		}
	}
	return false
}