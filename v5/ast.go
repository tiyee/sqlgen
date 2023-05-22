package v5

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type NodeType int

const (
	nodeBegin NodeType = iota
	ndCreateTableStmt
	ndNameListStmt
	ndNameNode
	ndColumnsNode
	ndColumnNode
	ndCommentNode
	ndFieldOptsStmt
	ndFieldOptNode
	ndFieldDefaultNode
	ndFieldTypeNode
	ndTableIndexListStmt
	ndTableIndexStmt
	ndTableOptNode
	ndTableOptsStmt
	ndEnd
)

type FieldType string

// columns he index逗号问题
const (
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
	tDATE      FieldType = "date"
	tDATETIME  FieldType = "datetime"
	tTIMESTAMP FieldType = "timestamp"
)
const (
	optBegin = 1 << iota
	optAutoIncrement
	optUnsigned
	optNull
	optNotNull
	optDefault
	optDefaultNull
	optDefaultCurrentTimestamp
	optOnUpdCurrentTs
	optEngine
	optCharacterSet
	optCharSet
	optCollate
	optEnd
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
}

type IndexType int

const (
	idxPrimary IndexType = 1 + iota
	idxKey
	idxUnique
)

type IPos interface {
	Pos() int
}
type node interface {
	Type() NodeType
	String() string
}

type ExprNode interface {
	node
}
type StmtNode interface {
	node
	Receive(nod node, lexer yyLexer)
}

type CreateTableStmt struct {
	NameNode
	ColumnsStmt        *ColumnsStmt
	TableIndexListStmt *TableIndexListStmt
	TableOptsStmt      *TableOptsStmt
}

func (c *CreateTableStmt) Receive(nod node, s yyLexer) {
	switch nod.Type() {
	case ndNameNode:
		c.NameNode = nod.(NameNode)
	case ndColumnsNode:
		c.ColumnsStmt = nod.(*ColumnsStmt)
	case ndTableIndexListStmt:
		c.TableIndexListStmt = nod.(*TableIndexListStmt)
	case ndTableOptsStmt:
		c.TableOptsStmt = nod.(*TableOptsStmt)
	default:
		unexpectedNode(nod, s)
	}

}
func (c *CreateTableStmt) Type() NodeType {
	return ndCreateTableStmt
}
func (c *CreateTableStmt) String() string {
	tableName := "create table " + tildeWrapper(c.NameNode.s)
	return fmt.Sprintf("%s  %s   %s ", tableName, c.specs(), c.TableOptsStmt.String())
}
func (c *CreateTableStmt) specs() string {
	cols := c.ColumnsStmt.String()
	idx := c.TableIndexListStmt.String()
	str := make([]string, 0, 4)
	if len(cols) > 0 {
		str = append(str, cols)
	}
	if len(idx) > 0 {
		str = append(str, idx)
	}
	return "(\n" + strings.Join(str, ",\n") + "\n)"
}

type NameNode struct {
	s   string
	pos int
}

func (nn NameNode) Type() NodeType {
	return ndNameNode
}
func (nn NameNode) String() string {
	return nn.s
}

type NameListStmt []NameNode

func (n *NameListStmt) Type() NodeType {
	return ndNameListStmt
}
func (n *NameListStmt) Receive(nod node, lexer yyLexer) {
	switch nod.Type() {
	case ndNameNode:
		*n = append(*n, nod.(NameNode))
	default:
		unexpectedNode(nod, lexer)
	}
}
func (n *NameListStmt) String() string {
	str := make([]string, 0)
	for _, s := range *n {
		str = append(str, tildeWrapper(s.String()))
	}
	return strings.Join(str, " ,")
}

type ColumnsStmt []*ColumnStmt

func (cs *ColumnsStmt) Type() NodeType {
	return ndColumnsNode
}
func (cs *ColumnsStmt) Receive(nod node, lexer yyLexer) {
	switch nod.Type() {
	case ndColumnNode:
		*cs = append(*cs, nod.(*ColumnStmt))
		cs.Sort()
	default:
		unexpectedNode(nod, lexer)
	}
}
func (cs *ColumnsStmt) ExistF(fn func(col *ColumnStmt) bool) bool {
	for _, s := range *cs {
		if fn(s) {
			return true
		}
	}
	return false
}
func (cs *ColumnsStmt) Sort() {
	arr := PosArray[*ColumnStmt](*cs)
	sort.Sort(arr)
}
func (cs *ColumnsStmt) String() string {
	str := make([]string, 0)
	for _, s := range *cs {
		str = append(str, s.String())
	}
	return strings.Join(str, ",\n")
}

type ColumnStmt struct {
	pos int
	NameNode
	FieldTypeNode
	FieldOptsStmt
	CommentNode //comment
}

func (cs *ColumnStmt) Receive(nod node, lexer yyLexer) {
	switch nod.Type() {
	case ndNameNode:
		cs.NameNode = nod.(NameNode)
	case ndFieldTypeNode:
		cs.FieldTypeNode = nod.(FieldTypeNode)
	case ndFieldOptsStmt:
		cs.FieldOptsStmt = *nod.(*FieldOptsStmt)
	case ndCommentNode:
		cs.CommentNode = nod.(CommentNode)

	default:
		unexpectedNode(nod, lexer)

	}
}
func (cs *ColumnStmt) String() string {
	var str []string
	str = append(str, tildeWrapper(cs.NameNode.s), string(cs.FieldTypeNode.t))
	if cs.FieldTypeNode.length > 0 {
		str = append(str, fmt.Sprintf("(%d)", cs.FieldTypeNode.length))
	}
	str = append(str, cs.FieldOptsStmt.String())
	return strings.Join(str, " ")
}

func (cs *ColumnStmt) Type() NodeType {
	return ndColumnNode
}
func (cs *ColumnStmt) Pos() int {
	return cs.pos
}

type FieldDefaultNode OptUnionNode

func (f FieldDefaultNode) Type() NodeType {
	return ndFieldDefaultNode
}
func (f FieldDefaultNode) String() string {
	return f.s
}

type CommentNode struct {
	s string
}

func (cn CommentNode) Type() NodeType {
	return ndCommentNode
}
func (cn CommentNode) String() string {
	return cn.s
}

type FieldTypeNode struct {
	t      FieldType
	length int64
}

func (f FieldTypeNode) Type() NodeType {
	return ndFieldTypeNode
}
func (f FieldTypeNode) String() string {
	if f.length == 0 {
		return string(f.t)
	}
	return fmt.Sprintf("%s(%d)", f.t, f.length)
}

type FieldOptsStmt []FieldOptNode

func (f *FieldOptsStmt) Type() NodeType {
	return ndFieldOptsStmt
}
func (f *FieldOptsStmt) Receive(nd node, lexer yyLexer) {
	switch nd.Type() {

	case ndFieldOptNode:
		*f = append(*f, nd.(FieldOptNode))
	default:
		unexpectedNode(nd, lexer)
	}

}
func (f *FieldOptsStmt) String() string {
	var str []string
	arr := PosArray[FieldOptNode](*f)
	sort.Sort(arr)
	for _, opt := range arr {
		str = append(str, opt.String())
	}
	return strings.Join(str, " ")

}

type FieldOptNode OptUnionNode

func (f FieldOptNode) Type() NodeType {
	return ndFieldOptNode
}
func (f FieldOptNode) String() string {
	switch f.opt {
	case optUnsigned:
		return "unsigned"
	case optNotNull:
		return "NOT NULL"
	case optNull:
		return "NULL"
	case optAutoIncrement:
		return "AUTO_INCREMENT"
	case optCharacterSet:
		return "CHARACTER SET " + f.s
	case optCollate:
		return "COLLATE " + f.s
	case optDefault:
		return "DEFAULT " + f.s
	case optDefaultCurrentTimestamp:
		return "DEFAULT CURRENT_TIMESTAMP"
	case optOnUpdCurrentTs:
		return "DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"
	case optDefaultNull:
		return "DEFAULT NULL"

	}
	return ""
}
func (f FieldOptNode) Pos() int {
	return f.pos
}

type TableIndexListStmt []TableIndexStmt

func (t *TableIndexListStmt) Type() NodeType {
	return ndTableIndexListStmt
}
func (t *TableIndexListStmt) Receive(nd node, lexer yyLexer) {
	switch nd.Type() {
	case ndTableIndexStmt:
		fon := *nd.(*TableIndexStmt)
		*t = append(*t, fon)
	default:
		unexpectedNode(nd, lexer)
	}
}
func (t *TableIndexListStmt) IsExist(idx IndexType) bool {
	for _, s := range *t {
		if s.t == idx {
			return true
		}
	}
	return false
}
func (t *TableIndexListStmt) String() string {
	str := make([]string, 0)
	for _, s := range *t {
		str = append(str, s.String())
	}
	return strings.Join(str, ",\n")
}
func (t *TableIndexListStmt) Find(it IndexType) (*TableIndexStmt, bool) {
	for _, s := range *t {
		if s.t == it {
			return &s, true
		}
	}
	return nil, false
}

type TableIndexStmt struct {
	t IndexType
	NameNode
	NameListStmt NameListStmt
}

func (t *TableIndexStmt) Type() NodeType {
	return ndTableIndexStmt
}
func (t *TableIndexStmt) Receive(nd node, lexer yyLexer) {
	switch nd.Type() {
	case ndNameNode:
		fon := nd.(NameNode)
		t.NameNode = fon
	case ndNameListStmt:
		nls := nd.(*NameListStmt)
		t.NameListStmt = *nls
	default:
		unexpectedNode(nd, lexer)
	}
}
func (t *TableIndexStmt) String() string {
	var pre string
	switch t.t {
	case idxPrimary:
		pre = "PRIMARY KEY"
	case idxUnique:
		pre = "UNIQUE KEY" + tildeWrapper(t.NameNode.s)
	case idxKey:
		pre = "KEY " + tildeWrapper(t.NameNode.s)
	}
	return pre + " (" + t.NameListStmt.String() + ")"
}

type TableOptNode OptUnionNode

func (t TableOptNode) Type() NodeType {
	return ndTableOptNode
}
func (t TableOptNode) String() string {
	switch t.opt {
	case optCharSet:
		return "CHARSET=" + t.s
	case optCollate:
		return "Collate=" + t.s
	case optEngine:
		return "ENGINE=" + t.s
	case optAutoIncrement:
		return "AUTO_INCREMENT=" + strconv.FormatInt(int64(t.i), 10)
	}
	return t.s
}
func (t TableOptNode) Pos() int {
	return t.pos
}

type TableOptsStmt []TableOptNode

func (t *TableOptsStmt) Type() NodeType {
	return ndTableOptsStmt
}
func (t *TableOptsStmt) Receive(nd node, lexer yyLexer) {
	switch nd.Type() {

	case ndTableOptNode:
		*t = append(*t, nd.(TableOptNode))
	case ndTableOptsStmt:
		*t = append(*t, *nd.(*TableOptsStmt)...)
	default:
		unexpectedNode(nd, lexer)

	}
}
func (t *TableOptsStmt) String() string {
	arr := PosArray[TableOptNode](*t)
	sort.Sort(arr)
	str := make([]string, 0, len(*t))
	defs := make([]string, 0, 2)
	for _, to := range arr {
		if to.opt&(optCharSet|optCollate) > 0 {
			defs = append(defs, to.String())
		} else {
			str = append(str, to.String())
		}

	}
	if len(defs) > 0 {
		str = append(str, "DEFAULT")
		str = append(str, defs...)
	}
	return strings.Join(str, " ")
}

type OptUnionNode struct {
	opt int
	s   string
	i   int64
	pos int
}
