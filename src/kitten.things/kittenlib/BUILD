# A go_library is a collection of Go files compiled into one package.
# Typically you'd have one of these per directory, corresponding to the
# Go package structure. Please doesn't fundamentally require that you obey
# that structure but if you don't there's a reasonable chance of running
# afoul of obtuse compiler errors.

go_library(
    name = 'kittenlib',
    srcs = ['kitten.go'],
    deps = [
        '//proto:kitten',
        '//third_party/go:grpc',
    ],
    visibility = ['PUBLIC'],
)

# A go_test is similar to a go_binary in that it builds a self-contained
# executable, but it's typically run through 'plz test' to run all its
# tests and report their results.
go_test(
    name = 'kitten_test',
    srcs = ['kitten_test.go'],
    deps = [
        ':kittenlib',
        '//third_party/go:testify',
    ],
)
