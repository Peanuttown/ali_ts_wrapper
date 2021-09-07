package ali_ts_wrapper

import (
	ts "github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)


func NewAttributeColumn(columnName string,value interface{})ts.AttributeColumn{
	return ts.AttributeColumn{
		ColumnName:columnName,
		Value:value,
	}
}
