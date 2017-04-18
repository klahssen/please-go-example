# A go_binary builds a self-contained Go executable.
# Note that unlike normal 'go build', there is no requirement to
# necessarily have these in their own packages.
go_binary(
    name = 'go.client',
    srcs = ['client.go'],
    deps = [
        '//kitten.things/kittenlib',
        '//proto:kitten',
        '//third_party/go:go-flags',
    ],
    visibility = ['PUBLIC'],
)
