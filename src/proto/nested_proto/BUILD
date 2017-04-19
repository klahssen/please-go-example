build_rule(
      name = 'protoc_go_debug',
      srcs = [
          'kitten2.proto',
      ],

      outs = [
          'some.txt'
      ],
      cmd = 'echo $TOOL1 --go_out=$PKG --plugin=protoc-gen-go=$TOOL2  ${SRCS} > some.txt',
      needs_transitive_deps = True,
      deps = [
          '//third_party/go:protoc-gen-go',
      ],
      tools = [
          '/usr/local/bin/protoc',
          '//third_party/go:protoc-gen-go',
      ],
      labels = [
          'grpc',
          'proto:go-map: proto/kitten.proto=proto/kittens',
      ],
      requires = [
          'nested_proto',
      ],
      visibility = [
          'PUBLIC',
      ],
  )


sh_binary(
    name  = 'generate_protos.sh',
    visibility=['PUBLIC'],
    main = 'generate_protos.sh',
)

build_rule(
    name = 'exec_shellscript',
    cmd = '$TOOL1',
     tools = [
          '//proto/nested_proto:generate_protos.sh',
      ],

)