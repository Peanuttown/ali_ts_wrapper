package ali_ts_wrapper

import(
		
	ts "github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	 "github.com/aliyun/aliyun-tablestore-go-sdk/tablestore/search"
)



func NewSearchRequest(
	tableName string,
	indexName string,
	searchQuery search.SearchQuery,
	columnToGet *ts.ColumnsToGet,
)*ts.SearchRequest{
	return &ts.SearchRequest{
		TableName     :tableName,
		IndexName     :indexName,
		SearchQuery   :searchQuery,
		ColumnsToGet  :columnToGet,
	}
}




func NewColumnsToGet(returnAll bool,columnsToGet []string)*ts.ColumnsToGet{
	return &ts.ColumnsToGet{
		ReturnAll:returnAll,
		Columns:columnsToGet,
	}
}
