go_library(
    name = 'protos',
    srcs = glob(['protos/*.pb.go']),
    deps = [
        '//third_party/go:grpc',
        '//third_party/go:protobuf',
        '//third_party/go:protoc-gen-go',
    ],
    visibility = ['PUBLIC'],
)
