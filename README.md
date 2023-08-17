# pbconv

## protoc-gen-iface

### Usage

```
$ GOBIN=$(pwd)/bin go install github.com/josudoey/pbconv/protoc-gen-iface
$ PATH=$(pwd)/bin:$PATH protoc \
    -I=. \
    --iface_out=paths=source_relative:. \
    ./path/to/*.proto
```
