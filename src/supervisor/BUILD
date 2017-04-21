# A go_binary builds a self-contained Go executable.
# Note that unlike normal 'go build', there is no requirement to
# necessarily have these in their own packages.
go_binary(
    name = 'go.supervisor',
    srcs = ['main.go'],
    deps = [
        '//third_party/go:supervisor',
    ],
    visibility = ['PUBLIC'],
)

go_test(
    name = 'supervisor_test',
    srcs = ['main_test.go'],
    deps = [
        ':go.supervisor',
        '//third_party/go:testify',
    ],
)
