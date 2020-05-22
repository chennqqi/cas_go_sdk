module github.com/chennqqi/cas_go_sdk/cmd/cascmd

go 1.13

replace github.com/chennqqi/cas_go_sdk => ../../

replace github.com/chennqqi/cas_go_sdk/cas => ../../cas

require (
	github.com/antihax/optional v1.0.0
	github.com/araddon/dateparse v0.0.0-20200409225146-d820a6159ab1
	github.com/chennqqi/cas_go_sdk v0.0.0-00010101000000-000000000000
	github.com/chennqqi/cas_go_sdk/cas v0.0.0-00010101000000-000000000000
	github.com/chennqqi/cas_go_sdk/treehash v0.0.0-20200519132735-4657e9f88252
	github.com/chennqqi/goutils v0.1.5
	github.com/dustin/go-humanize v1.0.0
	github.com/fatih/color v1.9.0 // indirect
	github.com/google/subcommands v1.2.0
	github.com/mattn/go-colorable v0.1.6 // indirect
	golang.org/x/sys v0.0.0-20200519105757-fe76b779f299 // indirect
	gopkg.in/cheggaaa/pb.v1 v1.0.28
)
