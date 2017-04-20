# grpc_library compiles a .proto file with rpc endpoints into
# language-specific implementations. It creates a series of internal rules
# for each language so you don't need to add another go_library etc to make
# use of the output, and each language will automatically depend on the
# appropriate implementation.
#
# See https://developers.google.com/protocol-buffers and http://grpc.io for
# general information about gRPC and protobufs.
grpc_library(
    name = 'kitten',
    srcs = ['kitten.proto'],
    visibility = ['PUBLIC'],
    languages = [
        'go',
    ],
)


# subinclude('//build_defs:my_protos')

# build_rule(
#       name = 'protoc_go',
#       srcs = [
#           'kitten.proto',
#       ],
#       outs = [
#           'kitten.pb.go',
#       ],
#       cmd = '$TOOL1 --go_out=$TMP_DIR --plugin=protoc-gen-go=$TOOL2  ${SRCS} && (mv -f ${PKG}/* .; true)',
#       needs_transitive_deps = True,
#       deps = [
#           '//third_party/go:protoc-gen-go',
#       ],
#       tools = [
#           '/usr/local/bin/protoc',
#           '//third_party/go:protoc-gen-go',
#       ],
#       labels = [
#           'grpc',
#           'proto:go-map: proto/kitten.proto=proto/kittens',
#       ],
#       requires = [
#           'proto',
#       ],
#       visibility = [
#           'PUBLIC',
#       ],
#   )

# build_rule(
#       name = 'protoc_go_debug',
#       srcs = [
#           'kitten.proto',
#       ],

#       outs = [
#         #   'kittens.pb.go',
#           'some.txt'
#       ],
#       cmd = 'echo $TOOL1 --go_out=$PKG --plugin=protoc-gen-go=$TOOL2  ${SRCS} > some.txt',
#       needs_transitive_deps = True,
#       deps = [
#           '//third_party/go:protoc-gen-go',
#       ],
#       tools = [
#           '/usr/local/bin/protoc',
#           '//third_party/go:protoc-gen-go',
#       ],
#       labels = [
#           'grpc',
#           'proto:go-map: proto/kitten.proto=proto/kittens',
#       ],
#       requires = [
#           'proto',
#       ],
#       visibility = [
#           'PUBLIC',
#       ],
#   )
# # plz query completions "//proto:all"
# genrule(
#     name = 'library',
#     outs = ['lib.txt'],
#     cmd = 'touch $OUT',
# )


# genrule(
#     name = 'kitten_bin',
#     outs = ['bin.sh'],
#     # srcs = ['kitten.proto'],
#     binary = False,
#     cmd  = "touch $OUT",
#     # cmd = 'protoc $SRC --go_out=plugins=grpc:kitten',
#     # plz build "//third_party/go:protoc-gen-go"
#     # cmd = 'cat $SRC | $(exe //third_party/go:protoc-gen-go) > $OUT'
#     visibility = ['PUBLIC'],
# )
