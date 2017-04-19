

    $ plz query output "//proto:all"
    plz-out/gen/proto/kitten.pb.go
    plz-out/gen/proto/kitten.a
    plz-out/gen/proto/kitten.proto
    plz-out/gen/proto/kitten.pb.go
    plz-out/gen/proto/bin.sh
    plz-out/gen/proto/lib.txt


    $ plz query print "//proto:all"




E2E tests:
  - https://github.com/thought-machine/please/blob/dbfe20a55caac8209e48fadd225dd84e06b76b15/test/BUILD