go_binary(
    name = 'sunserver',
    srcs = ['sunserver.go'],
    deps = [
        '//sunprotos:protos',
        '//proto_out/nested_proto:kitten2',
        '//third_party/go:grpc',
        '//third_party/go:go-flags',
    ],
    visibility = ['PUBLIC'],
)
