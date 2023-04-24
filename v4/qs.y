
%{
package v4

%}

%union {
    offset int // offset
    ident string

    fieldType FieldType
    i int
    columnStmt ColumnStmt
    columnsStmt []ColumnStmt
    fieldOptStmt FieldOptStmt
    tableStmt TableStmt
    fieldCommentExpr FieldComment
    fieldTypeExpr FieldTypeExpr
    fieldDefault FieldDefault
    fieldNameIdenList []string
    indexStmt IndexStmt
    indexListStmt []IndexStmt
    tableOptStmt TableOptStmt

}

%token	<ident> tString IndexName

%type <ident> TableNameIden FieldNameIden
%token	<i> tNumber
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





%token LPAREN '('
	LBRACK "["
	LBRACE "{"
	COMMA  ","
	PERIOD "."

	RPAREN    ')'
	RBRACK    "]"
	RBRACE    "}"
	SEMICOLON ";"
	EQ "="
	QUOTE
	DQUOTE
	TILDE  "`"

%token <fieldType> tFieldType


%type <tableStmt> TableStmt

%type	<fieldOptStmt>   FieldOptStmt
%type <columnStmt>  ColumnStmt
%type <indexStmt>   IndexStmt
%type <columnsStmt> ColumnsStmt
%type <indexListStmt> IndexListStmt
%type	<fieldTypeExpr>   FieldTypeExpr
%type	<fieldDefault>  FieldDefaultExpr
%type	<fieldOptStmt> FieldOptStmt
%type <fieldCommentExpr>  FieldCommentExpr
%type <fieldNameIdenList>  FieldNameIdenList  TableOptDefaultStmt
%type <tableOptStmt> TableOptStmt
%right COMMA
%start main

%%


main: TableStmt {
    yylex.(*lex).Stmt=$1
}

TableStmt:
    TableStmt SEMICOLON{
        $$=$1
    }
	| kwCreate  kwTable  TableNameIden LPAREN   ColumnsStmt RPAREN TableOptStmt {
        $$=TableStmt{action:"create",tableName:$3,columns:$5,opt:$7}
	}
	| kwCreate  kwTable  TableNameIden LPAREN  ColumnsStmt    IndexListStmt  RPAREN TableOptStmt  {
            $$=TableStmt{action:"create",tableName:$3,columns:$5,indexes:$6,opt:$8}
    	}

    ;
TableNameIden:
    TILDE tString TILDE {
        $$=$2
    }
    | tString {
        $$=$1
    }
    ;


ColumnsStmt:

    ColumnStmt {
            $$=[]ColumnStmt{$1}
        }
    | ColumnStmt COMMA {
                $$=[]ColumnStmt{$1}
            }
    | ColumnStmt COMMA ColumnsStmt   {
            $3 = append($3,$1)
            $$=$3
        }
    ;
ColumnStmt:
     ColumnStmt FieldCommentExpr  {

        $1.Cm = $2
        $$=$1

    }
    | FieldNameIden FieldTypeExpr FieldOptStmt  {

              $$=ColumnStmt{t:1,tp:$2,opts:$3,name:$1}

          }
    ;



FieldTypeExpr:

     tFieldType   LPAREN  tNumber  RPAREN {
        $$=NewFieldTypeExpr($1,$3)
   }

  |  tFieldType {
        $$=NewFieldTypeExpr($1,0)
    }


    ;

FieldOptStmt :
    {
         $$=NewFieldOptStmt()
    }
    | FieldOptStmt kwUnsigned {
         $1.SetOpt(optUnsigned)
         $$=$1
    }
    | FieldOptStmt  kwNull  {
            $1.SetOpt(optNull)
            $$=$1
        }
    | FieldOptStmt kwNOT kwNull  {
        $1.SetOpt(optNotNull)
        $$=$1
    }
    | FieldOptStmt kwAutoIncrement  {
        $1.SetOpt(optAutoIncrement)
        $$=$1
    }
    | FieldOptStmt kwCharacter kwSet tString {
        $1.SetOpt(optCharacterSet)
        $1.character = $4
         $$=$1
    }
    | FieldOptStmt kwCollate tString {
        $1.SetOpt(optCollate)
        $1.collate = $3
         $$=$1
    }
    | FieldOptStmt FieldDefaultExpr   {
        $1.SetDefault($2)
        $$=$1
    }
    ;

FieldDefaultExpr:
    kwDefault  TILDE tString TILDE {
        $$=FieldDefault{s:$3}
    }
    | kwDefault  tString  {
        $$=FieldDefault{s:$2}
    }
    | kwDefault  kwNull  {
            $$=FieldDefault{s:"null",opt:optDefaultNull}
    }
    | kwDefault  kwCurrentTimestamp {
        $$=FieldDefault{s:"kwCurrentTimestamp",opt:optDefaultCurrentTimestamp}
    }
    | kwDefault  kwCurrentTimestamp kwOn kwUpdate kwCurrentTimestamp  {
            $$=FieldDefault{s:"kwCurrentTimestamp on UPDATE",opt: optOnUpdCurrentTs}
     }


FieldCommentExpr:
kwComment  tString  {
        $$=FieldComment{s:$2}
     }
     ;
IndexListStmt :
     IndexStmt{
             $$=[]IndexStmt{$1}
         }
    |  IndexStmt COMMA {
             $$=[]IndexStmt{$1}
         }
    | IndexStmt  COMMA IndexListStmt{
        $$=append($3,$1)
    }
    ;
IndexStmt :
    kwPrimary  kwKey LPAREN  FieldNameIden  RPAREN {
        $$=IndexStmt{t:"primary",cols:[]string{$4}}
    }

    | kwUnique  kwKey FieldNameIden LPAREN  FieldNameIdenList  RPAREN {
        $$=IndexStmt{t:"unique",name:$3,cols:$5}
    }
    | kwKey FieldNameIden LPAREN  FieldNameIden  RPAREN {
        $$=IndexStmt{t:"key",name:$2,cols:[]string{$4}}
    }
    ;
FieldNameIdenList :
    FieldNameIden {
       $$=[]string{$1}
    }
    | FieldNameIden COMMA FieldNameIdenList {
        $3=append($3,$1)
    }
    ;
FieldNameIden:
        TILDE tString TILDE {
            $$=$2
        }
        | tString {
            $$=$1
        }
    ;
TableOptStmt :
     kwEngine EQ kwInnoDB TableOptStmt {
         $4.engine= "innodb"
                 $$=$4
     }
    |  kwAutoIncrement  EQ tNumber  TableOptStmt {
        $4.autoIncrement= $3
        $$=$4
    }
    | TableOptDefaultStmt {
         $$=TableOptStmt{defaults: $1}
    }
    |{
        $$=TableOptStmt{}
    }
TableOptDefaultStmt :
    kwDefault kwCharset EQ tString {
        $$=[]string{"charset",$4}
    }
    | kwDefault kwCollate EQ tString {
        $$=[]string{"collate",$4}
    }
    | TableOptDefaultStmt kwCharset EQ tString {
        $$=append($1,"charset",$4)
    }
    | TableOptDefaultStmt kwCollate EQ tString {
        $$=append($1,"collate",$4)
    }
