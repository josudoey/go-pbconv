package plugin

import (
	"io"
	"os"
	"sync"
)

var lock sync.Mutex

func pluginRun(data []byte, args []string, pluginGen PluginGen) ([]byte, error) {
	pluginStdoutReader, pluginStdout, err := os.Pipe()
	if err != nil {
		return nil, err
	}

	pluginStdin, pluginStdinWriter, err := os.Pipe()
	if err != nil {
		return nil, err
	}

	lock.Lock()
	defer lock.Unlock()

	originArgs := os.Args
	os.Args = args
	defer func() {
		os.Args = originArgs
	}()

	originStdin := os.Stdin
	os.Stdin = pluginStdin
	defer func() {
		os.Stdin = originStdin
	}()

	originStdout := os.Stdout
	os.Stdout = pluginStdout
	defer func() {
		os.Stdout = originStdout
	}()

	go func() {
		defer pluginStdinWriter.Close()
		pluginStdinWriter.Write(data)
	}()

	go func() {
		defer pluginStdout.Close()
		pluginGen()
	}()

	result, err := io.ReadAll(pluginStdoutReader)
	if err != nil {
		return nil, err
	}

	return result, nil
}
