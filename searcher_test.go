package blueAst

import (
	"github.com/JfL0unch/dst"
	"testing"
)

func TestSearcher_FuncDecl(t *testing.T) {

	input := `package blueAst

import (
	"go/ast"
	"go/token"

	"github.com/JfL0unch/dst"
)

type Ast struct {
	FileSet  *token.FileSet // file info
	AstNode  *ast.File      // ast.Node
	DstNode  *dst.File      // dst.Node
}


func NewAst() *Ast{
	// todo
	return &Ast{}
}`
	targetName := "NewAst"

	expected := "NewAst"

	ast,err := NewAst("",input)
	if err != nil{
		t.Error(err)
		return
	}

	searcher,err := NewSearcher(*ast)
	if err != nil{
		t.Error(err)
		return
	}


	params := make([]*dst.Field,0)
	params = append(params,&dst.Field{
		Type: &dst.StarExpr{
			X: &dst.Ident{Name: "Ast" },
		},
	})
	fieldListParams := &dst.FieldList{
		List: params,
	}

	fnc := dst.FuncDecl{
		Name:&dst.Ident{
			Name: targetName,
		},
		Type: &dst.FuncType{
			Params: fieldListParams,
		},
	}
	funcDecl,err := searcher.FuncDecl(fnc)
	if err !=nil {
		t.Error(err)
		return
	}

	got := ""
	if funcDecl != nil&& funcDecl.Name !=nil{
		got = funcDecl.Name.Name
	}


	if got != expected {
		t.Errorf("got %s,expect %s",got,expected)
	}

}
