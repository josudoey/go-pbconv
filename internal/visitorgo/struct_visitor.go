package visitorgo

import (
	"go/ast"
)

type structVisitor struct {
	Items []*ast.TypeSpec
}

func (v *structVisitor) Visit(n ast.Node) ast.Visitor {
	switch node := n.(type) {
	case *ast.File, *ast.GenDecl:
		return v
	case *ast.TypeSpec:
		if _, ok := node.Type.(*ast.StructType); ok {
			v.Items = append(v.Items, node)
			return nil
		}

		return nil
	}

	return nil
}

func GetStructTypes(n ast.Node) []*ast.TypeSpec {
	v := &structVisitor{}
	ast.Walk(v, n)
	return v.Items
}

func GetStructTypeMap(n ast.Node) map[string]*ast.TypeSpec {
	nodes := GetStructTypes(n)
	nodeMap := map[string]*ast.TypeSpec{}
	for _, node := range nodes {
		nodeMap[node.Name.String()] = node
	}
	return nodeMap
}
