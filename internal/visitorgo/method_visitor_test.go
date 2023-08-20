package visitorgo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestGetFuncDecls(t *testing.T) {
	fileSet := token.NewFileSet()
	fixtureSourceFileContent := `package fixture
func FixtureFunc() {}
type FixtureStruct0 struct {}
func (t FixtureStruct0) GetFixtureField0() interface{}  { return nil }
func (t *FixtureStruct0) GetFixtureField1() interface{} { return nil }
`
	fileNode, err := parser.ParseFile(fileSet, "/fixture.go", fixtureSourceFileContent, parser.ParseComments)
	if err != nil {
		t.Fatalf("ParseFile: got %+v", err)
	}

	methodFuncs := GetMethodFuncs(fileNode)
	if want := 2; len(methodFuncs) != want {
		t.Errorf("methodFuncs length got %q\nwant  %q", len(methodFuncs), want)
	}
	{
		methodFunc := methodFuncs[0]
		if want := "GetFixtureField0"; methodFunc.Name.String() != want {
			t.Errorf("methodFuncs[0].Name got %q\nwant  %q", methodFunc.Name.String(), want)
		}

		if _, ok := methodFunc.Type.Results.List[0].Type.(*ast.InterfaceType); !ok {
			t.Errorf("methodFuncs[0] return type got %s", methodFunc.Type.Results.List[0].Type)
		}
	}

	if want := "GetFixtureField1"; methodFuncs[1].Name.String() != want {
		t.Errorf("methodFuncs[1].Name got %q\nwant  %q", methodFuncs[1].Name.String(), want)
	}

}
