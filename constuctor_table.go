package ali_ts_wrapper

import (
	ts "github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func NewCreateTableReq(
	tableMeta *ts.TableMeta,
	tableOption *ts.TableOption,
	reservedThroughOutput *ts.ReservedThroughput,
) *ts.CreateTableRequest {
	return &ts.CreateTableRequest{
		TableMeta:          tableMeta,
		TableOption:        tableOption,
		ReservedThroughput: reservedThroughOutput,
	}
}

func NewTableUpdateReq(
	name string,
	option *ts.TableOption,
	reservedThroughOutput *ts.ReservedThroughput,
) *ts.UpdateTableRequest {
	return &ts.UpdateTableRequest{
		TableName:          name,
		TableOption:        option,
		ReservedThroughput: reservedThroughOutput,
	}
}

func NewTableMeta(tableName string, primaryKeySchemas []*ts.PrimaryKeySchema, columns []*ts.DefinedColumnSchema) *ts.TableMeta {
	return &ts.TableMeta{
		TableName:      tableName,
		SchemaEntry:    primaryKeySchemas,
		DefinedColumns: columns,
	}

}

func NewPrimaryKeySchema(name string, primaryKeyType ts.PrimaryKeyType, option ts.PrimaryKeyOption) *ts.PrimaryKeySchema {
	return &ts.PrimaryKeySchema{
		Name:   &name,
		Type:   &primaryKeyType,
		Option: &option,
	}
}

func NewColumnSchema(name string, tpe ts.DefinedColumnType) *ts.DefinedColumnSchema {
	return &ts.DefinedColumnSchema{
		Name:       name,
		ColumnType: tpe,
	}
}

func NewTableOption(extraTTLPlusToOneDayInSec uint, neverExpire bool, maxVersion int) *ts.TableOption {
	var ttlInSec = 86400 // minmal ttl
	ttlInSec += int(extraTTLPlusToOneDayInSec)
	if neverExpire {
		ttlInSec = -1
	}
	if maxVersion == 0 {
		maxVersion = 1
	}
	return &ts.TableOption{
		TimeToAlive: ttlInSec,
		MaxVersion:  maxVersion,
	}
}

func NewReservedThroughOutput(readCap, writeCap int) *ts.ReservedThroughput {
	return &ts.ReservedThroughput{
		Readcap:  readCap,
		Writecap: writeCap,
	}
}
