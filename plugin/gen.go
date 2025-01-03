package plugin

import (
	goproto "github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

type PluginGen func()

func Run(req *pluginpb.CodeGeneratorRequest, args []string, pluginGen PluginGen) (*pluginpb.CodeGeneratorResponse, error) {
	data, err := goproto.Marshal(req)
	if err != nil {
		return nil, err
	}

	result, err := pluginRun(data, args, pluginGen)
	if err != nil {
		return nil, err
	}

	res := pluginpb.CodeGeneratorResponse{}
	if err = goproto.Unmarshal(result, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
