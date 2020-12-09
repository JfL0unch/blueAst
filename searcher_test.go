package blueAst

import (
	"github.com/JfL0unch/dst"
	"go/token"
	"testing"
)

func TestSearcher_Node(t *testing.T) {

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
	targetName := "Ast"

	expected := "Ast"

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

	specs := make([]dst.Spec,0)
	specs = append(specs,&dst.TypeSpec{
		Name: &dst.Ident{
			Name: targetName,
		},
		Type: &dst.StructType{},
	})

	fnc := dst.GenDecl{
		Tok:token.TYPE,
		Specs: specs,
	}

	node,err := searcher.Node(&fnc)
	if err !=nil {
		t.Error(err)
		return
	}

	if genDecl,ok := node.(*dst.GenDecl);ok{
		got := ""
		if genDecl != nil&& len(genDecl.Specs) > 0{
			got = genDecl.Specs[0].(*dst.TypeSpec).Name.Name
		}

		if got != expected {
			t.Errorf("got %s,expect %s",got,expected)
		}
	}else{
		t.Errorf("got %s,expect %s","",expected)
	}

}
