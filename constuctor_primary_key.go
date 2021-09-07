package ali_ts_wrapper



import (
	ts "github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func NewPrimaryKey(columns []*ts.PrimaryKeyColumn)*ts.PrimaryKey{

	return &ts.PrimaryKey{
		PrimaryKeys:columns,
	}
}

func NewPrimayKeyColumn(columnName string,value interface{},option ts.PrimaryKeyOption)*ts.PrimaryKeyColumn{

	return &ts.PrimaryKeyColumn{
		ColumnName:columnName,
		Value:value,
		PrimaryKeyOption:option,
	}

}
