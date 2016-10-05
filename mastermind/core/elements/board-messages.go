package elements

import "errors"

var (
	MsgValidationFailed = "Validation failed!"
	MsgBoardCompleted = "Excelent, you found the code!"
	MsgChancesOver = "Your chances are over!"
)

var (
	ErrOnCheckUserAndBoardCodes = errors.New("Divergence of size between the codes to check and the board.")
	ErrColumnRowCodeInvalidSize = errors.New("Columns, rows and codeLength should be great than zero.")
	ErrCodeZeroOrNegative = errors.New("Code can't be zero or negative.")
)
