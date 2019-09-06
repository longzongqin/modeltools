package models

type AuthInfo struct {
	Id int `json:"id"` 
	Bigint int64 `json:"bigint"` 
	Binary string `json:"binary"` 
	Bit int8 `json:"bit"` 
	Blob string `json:"blob"` 
	Char string `json:"char"` 
	Date string `json:"date"` 
	Datetime time.Time `json:"datetime"` 
	Decimal float32 `json:"decimal"` 
	Double float32 `json:"double"` 
	Enum string `json:"enum"` 
	Float float32 `json:"float"` 
	Geometry string `json:"geometry"` 
	Geometrycollection string `json:"geometrycollection"` 
	Int int `json:"int"` 
	Integer int `json:"integer"` 
	Json string `json:"json"` 
	Linestring string `json:"linestring"` 
	Longblob string `json:"longblob"` 
	Longtext string `json:"longtext"` 
	Mediumblob string `json:"mediumblob"` 
	Mediumint int `json:"mediumint"` 
	Mediumtext string `json:"mediumtext"` 
	Multilinestring string `json:"multilinestring"` 
	Multipoint string `json:"multipoint"` 
	Multipolygon string `json:"multipolygon"` 
	Numeric float32 `json:"numeric"` 
	Point string `json:"point"` 
	Polygon string `json:"polygon"` 
	Real float32 `json:"real"` 
	Set string `json:"set"` 
	Smallint int8 `json:"smallint"` 
	Text string `json:"text"` 
	Time time.Time `json:"time"` 
	Timestamp time.Time `json:"timestamp"` 
	Tinyblob string `json:"tinyblob"` 
	Tinyint int8 `json:"tinyint"` 
	Tinytext string `json:"tinytext"` 
	Varbinary string `json:"varbinary"` 
	Varchar string `json:"varchar"` 
	Year int8 `json:"year"` 
}