# go_library(
#     name = 'settings',
#     srcs = [ ':settings_files' ],
#     deps = [
#         '//third_party/go:grpc',
#         '//third_party/go:protobuf',
#         '//third_party/go:protoc-gen-go',
#     ],
#     visibility = ['PUBLIC',],
# )


# filegroup(
#     name = 'protos',
#     srcs = glob(['*.pb.go']),
#     test_only = False,
#     visibility = ['PUBLIC'],
# )

# "//third_party/python:_enum#install"
# file_group(
# glob(['**/*.pb.go'])
# )

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
