package blueAst

import (
	"fmt"
	"reflect"

	"github.com/JfL0unch/dst"
	"github.com/JfL0unch/dst/dstutil"
)

type Searcher struct {
	ast Ast
}

func NewSearcher(ast Ast)(*Searcher,error){
	return &Searcher{ast:ast},nil
}

func (v Searcher) FuncDecl(fnc dst.FuncDecl)(*dst.FuncDecl,error){
	fn := func(c *dstutil.Cursor)bool{
		if sim, hit := c.Similarity(&fnc);sim >0 && sim==hit {

			x := reflect.ValueOf(c.Node())
			fmt.Printf("m=%s",x.Type())

			fmt.Printf("got %d,expect %d",sim,hit)
			return false
		}else{

			return true
		}

	}

	node := dstutil.Apply(v.ast.DstNode, nil, fn)



	if funcDeclNode,ok := node.(*dst.FuncDecl);ok{
		return funcDeclNode,nil
	}else{
		fmt.Println("hhh")
		return nil,nil
	}
}
