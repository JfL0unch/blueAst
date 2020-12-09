package blueAst

import (
	"fmt"
	"github.com/JfL0unch/dst"
	"github.com/JfL0unch/dst/dstutil"
)

type Searcher struct {
	ast Ast
}

func NewSearcher(ast Ast)(*Searcher,error){
	return &Searcher{ast:ast},nil
}

func (v Searcher) Node(fnc dst.Node)(dst.Node,error){
	fn := func(c *dstutil.Cursor)bool{
		if sim, hit := c.Similarity(fnc); sim >0 &&sim==hit {
			fmt.Printf("sim %d,hit %d",sim,hit)
			return true
		}else{
			return false
		}
	}

	node,_ := dstutil.Find(v.ast.DstNode, fn)

	if node != nil{
		return node,nil
	}
	return nil,nil

}


func (v Searcher) Replace(targetNode,replaceNode dst.Node)(dst.Node,error){
	fn := func(c *dstutil.Cursor)bool{
		if sim, hit := c.Similarity(targetNode); sim >0 &&sim==hit {
			c.Replace(replaceNode)
			return true
		}else{
			return false
		}
	}

	newNode,_ := dstutil.Rewrite(v.ast.DstNode, fn)

	if newNode != nil{
		return newNode,nil
	}
	return nil,nil

}
