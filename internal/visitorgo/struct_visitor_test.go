package visitorgo

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestGetTypes(t *testing.T) {
	fileSet := token.NewFileSet()
	fixtureSourceFileContent := `package fixture
type FixtureStruct0 struct {}
type FixtureStruct1 FixtureStruct0
type FixtureStruct2 = FixtureStruct1
func FixtureFunc() interface{} {
	type FixtureStruct3 struct{}
	return &FixtureStruct3{}
}
type FixtureInterface0 interface {}
type FixtureInterface1 = FixtureInterface0
type FixtureInterface2 = FixtureInterface0
type FixtureTypeInt0 int
type FixtureTypeInt1 = int
`
	fileNode, err := parser.ParseFile(fileSet, "/fixture.go", fixtureSourceFileContent, parser.ParseComments)
	if err != nil {
		t.Fatalf("ParseFile: got %+v", err)
	}

	structTypes := GetStructTypes(fileNode)
	if want := 1; len(structTypes) != want {
		t.Errorf("structTypes length got %q\nwant  %q", len(structTypes), want)
	}
}

func TestGetTypesMap(t *testing.T) {
	fileSet := token.NewFileSet()
	fixtureSourceFileContent := `package fixture
type FixtureStruct0 struct {}
type FixtureStruct1 FixtureStruct0
type FixtureStruct2 = FixtureStruct1
type FixtureInterface0 interface {}
type FixtureInterface1 = FixtureInterface0
type FixtureInterface2 = FixtureInterface0
type FixtureTypeInt0 int
type FixtureTypeInt1 = int
`
	fileNode, err := parser.ParseFile(fileSet, "/fixture.go", fixtureSourceFileContent, parser.ParseComments)
	if err != nil {
		t.Fatalf("ParseFile: got %+v", err)
	}

	structTypeMap := GetStructTypeMap(fileNode)
	if want := 1; len(structTypeMap) != want {
		t.Errorf("structTypeMap length got %q\nwant  %q", len(structTypeMap), want)
	}

	{
		typ := structTypeMap["FixtureStruct0"]
		if typ == nil {
			t.Errorf("FixtureStruct0 not in structTypeMap")
		} else if want := "FixtureStruct0"; typ.Name.String() != want {
			t.Errorf("FixtureStruct0 name: got %q\nwant %q", typ.Name, want)
		}
	}
}
