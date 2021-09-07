package ali_ts_wrapper

import(
	ts "github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func NewRowCondition(rowExistExpectation ts.RowExistenceExpectation )*ts.RowCondition{
	return &ts.RowCondition{
		RowExistenceExpectation:rowExistExpectation,
	}
}
