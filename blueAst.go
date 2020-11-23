package blueAst

import (
	"go/ast"
	"go/token"

	"github.com/dave/dst"
)

type Ast struct {
	Searcher Searcher
	FileSet  *token.FileSet // file info
	AstNode  *ast.File      // ast.Node
	Dst      *dst.File      // dst.Node
}
