package ali_ts_wrapper

import (
	"errors"

	ts "github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

const (
	ERR_CODE_OTSObjectAlreadyExist = "OTSObjectAlreadyExist"
)

func ErrIsOstErr(err error) *ts.OtsError {
	var errIns *ts.OtsError
	if errors.As(err, &errIns) {
		return errIns
	}
	return nil
}

func ErrIsTableExists(err error) bool {
	if ostErr := ErrIsOstErr(err); ostErr != nil {
		return ostErr.Code == "OTSObjectAlreadyExist"
	}
	return false
}

func ErrIsSearchIndexExists(err error) bool {
	if ostErr := ErrIsOstErr(err); ostErr != nil {
		return ostErr.Code == "OTSObjectAlreadyExist"
	}
	return false
}
