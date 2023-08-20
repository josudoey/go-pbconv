package visitorgo

import (
	"go/ast"
)

type methodFuncVisitor struct {
	Items []*ast.FuncDecl
}

func (v *methodFuncVisitor) Visit(n ast.Node) ast.Visitor {
	switch node := n.(type) {
	case *ast.File:
		return v
	case *ast.FuncDecl:
		if node.Recv == nil {
			return nil
		}
		v.Items = append(v.Items, node)
	}

	return nil
}

func GetMethodFuncs(n ast.Node) []*ast.FuncDecl {
	v := &methodFuncVisitor{}
	ast.Walk(v, n)
	return v.Items
}

func GetMethodFuncMap(n ast.Node) map[string]*ast.FuncDecl {
	nodes := GetMethodFuncs(n)
	nodeMap := map[string]*ast.FuncDecl{}
	for _, node := range nodes {
		nodeMap[node.Name.String()] = node
	}
	return nodeMap
}
