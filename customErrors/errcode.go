package customErrors

import "errors"

type ErrCode string

const (
	InsertDataFailed ErrCode = "D0001"
	GetDataFailed    ErrCode = "D0002"
	DataNotFound     ErrCode = "D0003"

	ReqBodyDecodeFailed ErrCode = "R0001"
)

var ErrNoData = errors.New("got no data from the database")
