package progreflect

import (
	"bytes"
	"encoding/gob"
	"go/build"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"text/template"
)

// GetFileDescriptorProtoRawByProg get descriptor raw by prog reflect
func GetFileDescriptorProtoRawByProg(goPackage string, descriptorFilename string) ([]byte, error) {
	program, err := writeProgram(progReflectArgs{
		GoPackage:          goPackage,
		DescriptorFilename: descriptorFilename,
	})
	if err != nil {
		return nil, err
	}

	wd, _ := os.Getwd()

	// Try to run the reflection program  in the current working directory.
	if p, err := runInDir(program, wd); err == nil {
		return p, nil
	}

	// Try to run the program in the same directory as the input package.
	if p, err := build.Import(goPackage, wd, build.FindOnly); err == nil {
		dir := p.Dir
		if p, err := runInDir(program, dir); err == nil {
			return p, nil
		}
	}

	// Try to run it in a standard temp directory.
	return runInDir(program, "")
}

func writeProgram(data progReflectArgs) ([]byte, error) {
	var program bytes.Buffer
	if err := reflectProgram.Execute(&program, &data); err != nil {
		return nil, err
	}

	return program.Bytes(), nil
}

// run the given program and parse the output as a model.Package.
func run(program string) ([]byte, error) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		return nil, err
	}

	filename := f.Name()
	defer os.Remove(filename)
	if err := f.Close(); err != nil {
		return nil, err
	}

	// Run the program.
	cmd := exec.Command(program, "-output", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	f, err = os.Open(filename)
	if err != nil {
		return nil, err
	}

	// Process output.
	var rawDesc []byte
	if err := gob.NewDecoder(f).Decode(&rawDesc); err != nil {
		return nil, err
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	return rawDesc, nil
}

// runInDir writes the given program into the given dir, runs it there, and
// parses the output as a model.Package.
func runInDir(program []byte, dir string) ([]byte, error) {
	// We use TempDir instead of TempFile so we can control the filename.
	tmpDir, err := os.MkdirTemp(dir, "pbconv_reflect_")
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			log.Printf("failed to remove temp directory: %s", err)
		}
	}()
	const progSource = "prog.go"
	var progBinary = "prog.bin"
	if runtime.GOOS == "windows" {
		// Windows won't execute a program unless it has a ".exe" suffix.
		progBinary += ".exe"
	}

	if err := os.WriteFile(filepath.Join(tmpDir, progSource), program, 0600); err != nil {
		return nil, err
	}

	cmdArgs := []string{}
	cmdArgs = append(cmdArgs, "build")
	cmdArgs = append(cmdArgs, "-o", progBinary, progSource)

	// Build the program.
	buf := bytes.NewBuffer(nil)
	cmd := exec.Command("go", cmdArgs...)
	cmd.Dir = tmpDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = io.MultiWriter(os.Stderr, buf)
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return run(filepath.Join(tmpDir, progBinary))
}

type progReflectArgs struct {
	GoPackage          string
	DescriptorFilename string
}

// This program reflects on an interface value, and prints the
// gob encoding of a model.Package to standard output.
// JSON doesn't work because of the model.Type interface.
var reflectProgram = template.Must(template.New("program").Parse(`
package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"os"

	goproto "github.com/golang/protobuf/proto"
	pbconv "github.com/josudoey/go-pbconv"
	_ {{printf "%q" .GoPackage}}
)

var output = flag.String("output", "", "The output file name, or empty to use stdout.")

func main() {
	flag.Parse()

	file, err := pbconv.GetFileDescriptorProtoByFilename({{printf "%q" .DescriptorFilename}})
	if err != nil {
		fmt.Fprintf(os.Stderr, "get file descriptor: %v\n", err)
		os.Exit(1)
	}

	rawDesc, err := goproto.Marshal(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get raw desc: %v\n", err)
		os.Exit(1)
	}

	outfile := os.Stdout
	if len(*output) != 0 {
		var err error
		outfile, err = os.Create(*output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open output file %q", *output)
		}
		defer func() {
			if err := outfile.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "failed to close output file %q", *output)
				os.Exit(1)
			}
		}()
	}

	if err := gob.NewEncoder(outfile).Encode(rawDesc); err != nil {
		fmt.Fprintf(os.Stderr, "gob encode: %v\n", err)
		os.Exit(1)
	}
}
`))
