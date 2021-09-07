package ali_ts_wrapper

import(
	ts "github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)


func NewPutRowChange(tableName string,primaryKey *ts.PrimaryKey,columns []ts.AttributeColumn,rowCondition *ts.RowCondition)*ts.PutRowChange{
	return &ts.PutRowChange{
		TableName  :tableName,
		PrimaryKey :primaryKey,
		Columns    :columns,
		Condition : rowCondition,
	}
}

