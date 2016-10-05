package elements

import (
	"../random"
)

type Board struct {
	Codes      []int `json:"-"`
	CodeLength int `json:"codeLength"`
	Columns    int `json:"columns"`
	Rows       int `json:"rows"`
	Matrix     map[int][]int `json:"matrix"`
	Repetition bool `json:"repetition"`
	checks     int
}

type BoardResponse struct {
	Codes          []int `json:"codes"`
	Message        string `json:"message"`
	Hits           int `json:"hits"`
	HitsOutOfOrder int `json:"hitsOutOfOrder"`
	Misses         int `json:"misses"`
	Error          error `json:"error"`
	RemainingMoves int `json:"remainingMoves"`
	Board          Board `json:"board"`
}

type BoardRequest struct {
	Columns    int `json:"columns"`
	Rows       int `json:"rows"`
	CodeLength int `json:"codeLength"`
	Repetition bool `json:"repetition"`
}

func (b *Board) Check(codes []int) BoardResponse {
	var misses, hitsCheck, hitsOutOfOrderCheck []int
	var err error

	err = validateCheck(codes, b.Codes)
	if err != nil {
		return BoardResponse{Message: MsgValidationFailed, Error: err}
	}

	boardComplete, msg, err := validateBoardComplete(b)

	if !boardComplete {
		checkMatrix(b, codes)

		misses, hitsCheck, hitsOutOfOrderCheck = calcHits(codes, b.Codes)

		boardComplete, msg, err = validateBoardComplete(b)
	}

	return BoardResponse{Board: *b, Codes: codes, Message: msg, Hits: len(hitsCheck), HitsOutOfOrder:len(hitsOutOfOrderCheck), Misses:len(misses), Error: err, RemainingMoves: calcRemainingMoves(b)}
}

func checkMatrix(b *Board, codes []int) {
	b.Matrix[b.checks] = codes
	b.checks++
}

func Create(br BoardRequest) (b Board, err error) {
	err = validadeCreate(br.Columns, br.Rows, br.CodeLength)
	if err != nil {
		return b, err
	}

	arrayRandom, err := random.Random(br.Columns, br.CodeLength, br.Repetition)
	if err != nil {
		return b, err
	}

	if containsNegative(arrayRandom) {
		return b, ErrCodeZeroOrNegative
	}

	b = Board{Codes: arrayRandom, Columns: br.Columns, Rows: br.Rows, Repetition: br.Repetition, Matrix: make(map[int][]int), CodeLength: br.CodeLength}

	return b, nil
}