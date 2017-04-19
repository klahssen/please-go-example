#!/usr/bin/env bash
# we are in plz-out/bin/proto/nested_proto
RELATIVE_BUILD="../../../../.."
RELATIVE_PKG="../.."
cd $RELATIVE_BUILD/proto/nested_proto
PROTOC_GEN_GO="$RELATIVE_PKG/plz-out/bin/third_party/go/bin/protoc-gen-go"
mkdir -p kitten2
protoc --go_out=kitten2 --plugin=protoc-gen-go=$PROTOC_GEN_GO kitten2.proto

# protoc --go_out=. --plugin=protoc-gen-go=plz-out/bin/third_party/go/bin/protoc-gen-go proto/kitten.proto
# protoc --go_out=proto --plugin=protoc-gen-go=plz-out/bin/third_party/go/bin/protoc-gen-go proto/kitten.proto
