%{
package v5
var ftm map[string]int //field type map
func init() {
    ftm = map[string]int{
    		"tinyint":tpTinyint,
               	"smallint":tpSmallint,
               	"mediumint":tpMediumint,
               	"int":tpInt,
               	"bigint":tpBigint,
               	"float":tpFloat,
               	"doublee":tpDouble,
               	"char":tpChar,
               	"varchar":tpVarchar_,
               	"text":tpText,
               	"date":tpDate,
               	"datetime":tpDatetime,
               	"timestamp":tpTimestamp,
    }
}


%}

%union {
	pos int
	s 	string
	i 	int64
	express ExprNode
	statement StmtNode

}
%token	<s> tString
%token	<i> tNumber
%type <express> NameNode CommentNode  FieldDefaultNode FieldTypeNode TableOptDefaultNode

%type <statement> CreateTableStmt FieldOptsStmt ColumnsStmt TableOptsStmt
                  ColumnStmt TableIndexListStmt TableIndexStmt   NameListStmt


%token LPAREN "("
	LBRACK "["
	LBRACE "{"
	COMMA  ","
	PERIOD "."

	RPAREN    ")"
	RBRACK    "]"
	RBRACE    "}"
	SEMICOLON ";"
	EQ "="
	QUOTE "'"
	DQUOTE  "\""
	TILDE  "`"
%token
                kwUnsigned
             	kwNOT
             	kwNull
             	kwAutoIncrement
             	kwCharacter
             	kwCharset
             	kwSet
             	kwComment
             	kwDefault
             	kwTable
             	kwCreate
             	kwCollate
             	kwCurrentTimestamp
             	kwOn
             	kwUpdate
             	kwEngine
             	kwInnoDB
             	kwPrimary
             	kwKey
             	kwUnique
%token 	tpTinyint
       	tpSmallint
       	tpMediumint
       	tpInt
       	tpBigint
       	tpFloat
       	tpDouble
       	tpChar
       	tpVarchar_
       	tpText
       	tpDate
       	tpDatetime
       	tpTimestamp

%start main

%%

main: CreateTableStmt {
    yylex.(*lex).Stmt=$1.( *CreateTableStmt)
}

CreateTableStmt:
    CreateTableStmt SEMICOLON {
        $$=$1
    }

    | kwCreate  kwTable  NameNode LPAREN   ColumnsStmt  RPAREN TableOptsStmt  {
        $$=&CreateTableStmt{}
        $$.Receive($3,yylex)
        $$.Receive($5,yylex)
        $$.Receive($7,yylex)

	}
    | kwCreate  kwTable  NameNode LPAREN   ColumnsStmt TableIndexListStmt RPAREN TableOptsStmt {
        $$=&CreateTableStmt{}
        $$.Receive($3,yylex)
        $$.Receive($5,yylex)
        $$.Receive($6,yylex)
        $$.Receive($8,yylex)
    }
    ;
ColumnsStmt:

    ColumnStmt {
            $$=new(ColumnsStmt)
            $$.Receive($1,yylex)
        }
    | ColumnStmt COMMA {
            $$=new(ColumnsStmt)
            $$.Receive($1,yylex)
    }
    | ColumnStmt COMMA ColumnsStmt   {
            $$=$3
            $$.Receive($1,yylex)
        }
    ;
ColumnStmt:
     ColumnStmt CommentNode {
        $$=$1
        $$.Receive($2,yylex)
     }

     | NameNode FieldTypeNode FieldOptsStmt  {
        $$=&ColumnStmt{pos:yyVAL.pos}
        $$.Receive($1,yylex)
        $$.Receive($2,yylex)
        $$.Receive($3,yylex)
    }
    ;
FieldTypeNode :
    tpTinyint  {
         $$ = FieldTypeNode{t:tTINYINT}
    }
    | tpInt  {
         $$ = FieldTypeNode{t:tINT}
    }
    | tpSmallint  {
         $$ = FieldTypeNode{t:tSMALLINT}
    }
    | tpMediumint  {
         $$ = FieldTypeNode{t:tMEDIUMINT}
    }
    | tpBigint  {
         $$ = FieldTypeNode{t:tBIGINT}
    }
    | tpFloat  {
         $$ = FieldTypeNode{t:tFLOAT}
    }
    | tpDouble  {
         $$ = FieldTypeNode{t:tDOUBLE}
    }
    | tpDate  {
         $$ = FieldTypeNode{t:tDATE}
    }
    | tpDatetime  {
         $$ = FieldTypeNode{t:tDATETIME}
    }
    | tpTimestamp  {
         $$ = FieldTypeNode{t:tTIMESTAMP}
    }
    | tpText  {
         $$ = FieldTypeNode{t:tTEXT}
    }
    | tpVarchar_   LPAREN  tNumber  RPAREN {
        $$ = FieldTypeNode{t:tVARCHAR,length:$3}

    }
    | tpChar   LPAREN  tNumber  RPAREN {
        $$ = FieldTypeNode{t:tCHAR,length:$3}
    }


    ;

FieldOptsStmt :
    {
         $$=new(FieldOptsStmt)
    }
    |  kwUnsigned FieldOptsStmt{
         $2.Receive(FieldOptNode{opt:optUnsigned,pos:yyVAL.pos},yylex)
         $$=$2
    }
    |   kwNull  FieldOptsStmt {
            $2.Receive(FieldOptNode{opt:optNull,pos:yyVAL.pos},yylex)
            $$=$2
    }
    |  kwNOT kwNull FieldOptsStmt  {
        $3.Receive(FieldOptNode{opt:optNotNull,pos:yyVAL.pos},yylex)
        $$=$3
    }
    | kwAutoIncrement FieldOptsStmt  {
        $2.Receive(FieldOptNode{opt:optAutoIncrement,pos:yyVAL.pos},yylex)
        $$=$2
    }
    |  kwCharacter kwSet tString FieldOptsStmt{
        $4.Receive(FieldOptNode{opt:optCharacterSet,s:$3,pos:yyVAL.pos},yylex)

         $$=$4
    }
    |  kwCollate tString FieldOptsStmt {
        $3.Receive(FieldOptNode{opt:optCollate,s:$2,pos:yyVAL.pos},yylex)
         $$=$3
    }
    | FieldDefaultNode FieldOptsStmt {

    	$2.Receive($1,yylex)
    	$$=$2
    }

    ;
FieldDefaultNode:
    kwDefault  TILDE tString TILDE {
        $$=FieldOptNode{opt:optDefault,s:$3,pos:yyVAL.pos}
    }
    | kwDefault  tString  {
        $$=FieldOptNode{opt:optDefault,s:$2,pos:yyVAL.pos}
    }
    | kwDefault  kwNull  {
            $$=FieldOptNode{s:"null",opt:optDefaultNull,pos:yyVAL.pos}
    }
    | kwDefault  kwCurrentTimestamp {
        $$=FieldOptNode{s:"kwCurrentTimestamp",opt:optDefaultCurrentTimestamp,pos:yyVAL.pos}
    }
    | kwDefault  kwCurrentTimestamp kwOn kwUpdate kwCurrentTimestamp  {
            $$=FieldOptNode{s:"kwCurrentTimestamp on UPDATE",opt: optOnUpdCurrentTs,pos:yyVAL.pos}
     }
TableIndexListStmt :

     TableIndexStmt {
             $$=new(TableIndexListStmt)
             $$.Receive($1,yylex)
         }
    |  TableIndexStmt COMMA {
             $$=new(TableIndexListStmt)
             $$.Receive($1,yylex)
    }
    | TableIndexStmt  COMMA TableIndexListStmt {
        $3.Receive($1,yylex)
        $$=$3
    }
    ;
TableIndexStmt :
    kwPrimary  kwKey LPAREN  NameListStmt  RPAREN {
        $$=&TableIndexStmt{t:idxPrimary}
        $$.Receive($4,yylex)

    }

    | kwUnique  kwKey NameNode LPAREN  NameListStmt  RPAREN {
        $$=&TableIndexStmt{t:idxUnique}
        $$.Receive($3,yylex)
        $$.Receive($5,yylex)
    }
    | kwKey NameNode LPAREN  NameListStmt  RPAREN {
        $$=&TableIndexStmt{t:idxKey}
        $$.Receive($2,yylex)
        $$.Receive($4,yylex)
    }
    ;
NameListStmt:
    NameNode {
        $$=new(NameListStmt)
        $$.Receive($1,yylex)
    }
    | NameNode COMMA NameListStmt {
        $3.Receive($1,yylex)
        $$=$3
    }
TableOptsStmt :
     kwEngine EQ kwInnoDB TableOptsStmt {
        $4.Receive(TableOptNode{opt:optEngine,s:"InnoDB",pos:yyVAL.pos},yylex)
        $$=$4
     }
    |  kwAutoIncrement  EQ tNumber  TableOptsStmt {
        $4.Receive(TableOptNode{opt:optAutoIncrement,i:$3,pos:yyVAL.pos},yylex)
        $$=$4
    }

    |   kwDefault TableOptDefaultNode  TableOptsStmt{
        $3.Receive($2,yylex)
        $$=$3
    }
    |   kwDefault TableOptDefaultNode TableOptDefaultNode TableOptsStmt{
            $4.Receive($2,yylex)
            $4.Receive($3,yylex)
            $$=$4
        }
    | {
        $$=&TableOptsStmt{}
    }


TableOptDefaultNode:
    kwCharset EQ tString  {
        $$=TableOptNode{opt:optCharSet,s:$3,pos:yyVAL.pos}
    }
    | kwCollate EQ tString{
        $$=TableOptNode{opt:optCollate,s:$3,pos:yyVAL.pos}
    }
NameNode:
        TILDE tString TILDE {
            $$=NameNode{s:$2,pos:yyVAL.pos}
        }
        | tString {
            $$=NameNode{s:$1,pos:yyVAL.pos}
        }
    ;
CommentNode:
    kwComment  tString  {
            $$=CommentNode{s:$2}
    }
    ;