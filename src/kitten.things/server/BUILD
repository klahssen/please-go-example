go_binary(
    name = 'go.server',
    srcs = ['server.go'],
    deps = [
        '//kitten.things/kittenlib',
        '//proto_out/nested_proto:kitten2',
        '//third_party/go:grpc',
        '//third_party/go:go-flags',
    ],
    visibility = ['PUBLIC'],
)
