package v4

import "strings"

type FieldType string

// columns he index逗号问题
const (
	ILL        FieldType = "ILL"
	tTINYINT   FieldType = "tinyint"
	tINT       FieldType = "int"
	tBIGINT    FieldType = "bigint"
	tSMALLINT  FieldType = "smallint"
	tMEDIUMINT FieldType = "mediumint"
	tFLOAT     FieldType = "float"
	tDOUBLE    FieldType = "double"
	tCHAR      FieldType = "char"
	tVARCHAR   FieldType = "varchar"
	tTEXT      FieldType = "text"
	tDate      FieldType = "date"
	tDateTime  FieldType = "datetime"
	tTimestamp FieldType = "timestamp"
)

func convertFieldType(s string) FieldType {
	return FieldType(s)
}

const (
	optBegin = 1 << iota
	optAutoIncrement
	optUnsigned
	optNull
	optNotNull
	optDefaultNull
	optDefaultCurrentTimestamp
	optOnUpdCurrentTs
	optCharacterSet
	optCollate
)

var keywords = map[string]int{

	"unsigned":       kwUnsigned,
	"not":            kwNOT,
	"null":           kwNull,
	"AUTO_INCREMENT": kwAutoIncrement,
	"CHARACTER":      kwCharacter,
	"charset":        kwCharset,
	"SET":            kwSet,

	"Comment":           kwComment,
	"Default":           kwDefault,
	"Table":             kwTable,
	"Create":            kwCreate,
	"Collate":           kwCollate,
	"Current_Timestamp": kwCurrentTimestamp,
	"on":                kwOn,
	"update":            kwUpdate,
	"InnoDB":            kwInnoDB,
	"Engine":            kwEngine,
	"Primary":           kwPrimary,
	"Key":               kwKey,
	"UNIQUE":            kwUnique,
	// type begin
	"tinyint":   tFieldType,
	"smallint":  tFieldType,
	"mediumint": tFieldType,
	"int":       tFieldType,
	"bigint":    tFieldType,
	"float":     tFieldType,
	"double":    tFieldType,
	"char":      tFieldType,
	"varchar":   tFieldType,
	"text":      tFieldType,
	"date":      tFieldType,
	"datetime":  tFieldType,
	"timestamp": tFieldType,
}

func findKeyword(s string) int {
	for k, v := range keywords {
		if strings.ToLower(s) == strings.ToLower(k) {
			return v
		}
	}
	return 0
}

type node interface {
	Pos()
	End()
}
type TokenNode interface {
	Tok() int
	node
}
type StmtNode interface {
	Format() string
	String() string
	Validate() error
}

type FieldTypeExpr struct {
	t      FieldType
	s      string
	length int
}

func NewFieldTypeExpr(s FieldType, l int) FieldTypeExpr {
	return FieldTypeExpr{
		t:      s,
		length: l,
	}
}

type FieldOptStmt struct {
	n         int
	d         FieldDefault
	opt       int
	character string
	collate   string
}

func NewFieldOptStmt() FieldOptStmt {
	return FieldOptStmt{}
}
func (s FieldOptStmt) SetOpt(opt int) {
	s.opt |= opt
}
func (s FieldOptStmt) SetDefault(d FieldDefault) {
	s.d = d
	s.opt |= d.opt
}

type FieldComment struct {
	s string
}
type FieldDefault struct {
	s   string
	opt int
}
type ColumnStmt struct {
	t    int
	name string
	Cm   FieldComment
	opts FieldOptStmt
	tp   FieldTypeExpr
}
type TableStmt struct {
	action    string
	tableName string
	columns   []ColumnStmt
	indexes   []IndexStmt
	opt       TableOptStmt
}
type IndexStmt struct {
	t    string
	name string
	cols []string
}
type TableOptStmt struct {
	engine        string
	autoIncrement int
	charset       string
	collate       string
	defaults      []string
}
