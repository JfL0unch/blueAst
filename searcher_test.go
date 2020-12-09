package blueAst

import (
	"bytes"
	"github.com/JfL0unch/dst"
	"github.com/JfL0unch/dst/decorator"
	"go/format"
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


func TestSearcher_Replace(t *testing.T) {
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
func NewAstCase2() *Ast{
	// todo
	return &Ast{}
}

func NewAstCase1() *Ast{
	// todo
	return &Ast{}
}`


	stmts := make([]dst.Stmt,0)
	lhs := make([]dst.Expr,0)
	lhs = append(lhs,&dst.Ident{
		Name: "x",
	})
	rhs := make([]dst.Expr,0)
	rhs = append(rhs,&dst.UnaryExpr{
		X: &dst.CompositeLit{
			Type:&dst.Ident{
				Name: "Ast",
			},
		},
		Op: token.AND,
	})

	stmts = append(stmts,&dst.AssignStmt{
		Lhs:lhs,
		Rhs: rhs,
		Tok: token.DEFINE,
	})

	stmts = append(stmts,&dst.EmptyStmt{
		Decs: dst.EmptyStmtDecorations{
			NodeDecs:dst.NodeDecs{
				End: []string{"\n"},
			},
		},})

	results := make([]dst.Expr,0)
	results = append(results,&dst.Ident{
		Name: "x",
	})

	stmts = append(stmts,&dst.ReturnStmt{
		Results:results,
	})



	blockStmt := &dst.BlockStmt{
		List: stmts,
	}

	fieldList := make([]*dst.Field,0)
	fieldList = append(fieldList,&dst.Field{
		Type:&dst.StarExpr{
			X: &dst.Ident{
				Name: "Ast",
			},
		},
	})

	targetNode := &dst.FuncDecl{
		Name: &dst.Ident{
			Name: "NewAstCase1",
		},
		Type: &dst.FuncType{
			Results: &dst.FieldList{
				List: fieldList,
			},
		},
	}

	replaceNode := &dst.FuncDecl{
		Name: &dst.Ident{
			Name: "NewAst",
		},
		Type: &dst.FuncType{
			Results: &dst.FieldList{
				List: fieldList,
			},
		},
		Body: blockStmt,
	}

	expected := ""

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

	newNode,err := searcher.Replace(targetNode,replaceNode)
	if err !=nil {
		t.Error(err)
		return
	}

	if dstFile,ok := newNode.(*dst.File);ok{
		restoredFset, restoredFile, err := decorator.RestoreFile(dstFile)
		if err != nil {
			t.Fatal(err)
		}
		var buf bytes.Buffer
		if err := format.Node(&buf, restoredFset, restoredFile); err != nil {
			t.Fatal(err)
		}

		got := buf.String()

		if got != expected {
			t.Errorf("got %s,expect %s",got,expected)
		}
	}else{
		t.Errorf("got %s,expect %s","",expected)
	}


}