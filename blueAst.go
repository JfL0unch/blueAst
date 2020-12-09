package blueAst

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/JfL0unch/dst"
	"github.com/JfL0unch/dst/decorator"
)

type Ast struct {
	FileSet  *token.FileSet // file info
	AstNode  *ast.File      // ast.Node
	DstNode  *dst.File      // dst.Node
}

func NewAst(fileName string,src string) (*Ast,error){

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, src, parser.ParseComments)
	if err != nil {
		return nil,err
	}

	d, err := decorator.DecorateFile(fset, f)
	if err != nil {
		return nil,err
	}

	a := &Ast{
		FileSet: fset,
		AstNode: f,
		DstNode: d,
	}

	return a,nil
}