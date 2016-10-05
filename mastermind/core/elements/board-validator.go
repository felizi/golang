package elements

import (
	"errors"
	"../validade"
)

func validateCheck(userCodes, boardCodes []int) error {
	if userCodes == nil || boardCodes == nil || len(userCodes) != len(boardCodes) || len(userCodes) <= 0 {
		return ErrOnCheckUserAndBoardCodes
	}

	if containsNegative(userCodes) || containsNegative(boardCodes) {
		return ErrCodeZeroOrNegative
	}
	return nil
}

func containsNegative(slice []int) bool {
	for _, vlr := range slice {
		if vlr < 0 {
			return true
		}
	}
	return false
}

func validateBoardComplete(b *Board) (boardComplete bool, msg string, err error) {
	if anyCheckAreEqualsToBoardCode(b) {
		boardComplete = true
		msg = MsgBoardCompleted
	}

	if !boardComplete && !anyCheckAreAvaliable(b) {
		boardComplete = true
		msg = MsgChancesOver
		err = errors.New(MsgChancesOver)
	}

	return boardComplete, msg, err
}

func validadeCreate(columns, rows, codeLength int) error {
	if columns <= 0 || rows <= 0 || codeLength <= 0 {
		return ErrColumnRowCodeInvalidSize
	}
	return nil
}

func anyCheckAreAvaliable(b *Board) (bool) {
	if b.checks < b.Rows {
		return true
	}
	return false
}

func anyCheckAreEqualsToBoardCode(b *Board) (bool) {
	for _, column := range b.Matrix {
		if validate.SliceEquals(column, b.Codes) {
			return true
		}
	}
	return false
}