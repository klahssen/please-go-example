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




# A go_test is similar to a go_binary in that it builds a self-contained
# executable, but it's typically run through 'plz test' to run all its
# tests and report their results.
go_test(
    name = 'sleepy_test',
    srcs = ['sleepy_test.go'],
    deps = [
        '//third_party/go:testify',
    ],
)
