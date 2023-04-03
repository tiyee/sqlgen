package gen

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

type StructField struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}
type StructCreator struct {
	PackageName string
	TableName   string
	Pk          string
	Fields      []StructField
	HasTime     bool
}

func exportName(s string) string {
	if strings.Index(s, "_") > 0 {
		arr := strings.Split(s, "_")
		for i, item := range arr {
			if len(item) > 0 {
				arr[i] = strings.ToUpper(item[0:1]) + item[1:]
			}
		}
		return strings.Join(arr, "")
	}
	return strings.ToUpper(s[0:1]) + s[1:]

}

func (c *StructCreator) StructName() string {
	if len(c.TableName) == 1 {
		return strings.ToUpper(c.TableName)
	}
	return exportName(c.TableName)
}

func (c *StructCreator) Create() *ast.GenDecl {
	fields := make([]*ast.Field, 0, len(c.Fields))
	for _, field := range c.Fields {
		fields = append(fields, &ast.Field{
			Names: []*ast.Ident{
				&ast.Ident{
					Name: field.Name,
				},
			},
			Type: &ast.Ident{
				Name: field.Type,
			},
			Tag: &ast.BasicLit{
				Kind:  token.STRING,
				Value: field.Tag, // "`json:\"id\"`",
			},

			Comment: &ast.CommentGroup{
				List: []*ast.Comment{
					&ast.Comment{
						Text: field.Comment,
					},
				},
			},
		})
	}
	ret := &ast.GenDecl{

		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: &ast.Ident{
					Name: c.StructName(),
				},
				Type: &ast.StructType{
					Fields: &ast.FieldList{
						List: fields,
					},
				},
			},
		},
	}

	return ret
}
func (c *StructCreator) makeFunc(funcName, ret string) *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				&ast.Field{
					Names: []*ast.Ident{
						&ast.Ident{
							Name: strings.ToLower(c.TableName[0:1]),
						},
					},
					Type: &ast.StarExpr{
						X: &ast.Ident{
							Name: c.StructName(),
						},
					},
				},
			},
		},
		Name: &ast.Ident{
			Name: funcName,
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					&ast.Field{
						Type: &ast.Ident{
							Name: "string",
						},
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.BasicLit{
							Kind:  token.STRING,
							Value: fmt.Sprintf("\"%s\"", ret),
						},
					},
				},
			},
		},
	}
}
func (c *StructCreator) PkFunc() *ast.FuncDecl {
	return c.makeFunc("Pk", c.Pk)
}
func (c *StructCreator) TableFunc() *ast.FuncDecl {
	return c.makeFunc("TableName", c.TableName)
}
