module github.com/johnkerl/miller/v6

// The repo is 'miller' and the executable is 'mlr', going back many years and
// predating the Go port.
//
// If we had ./mlr.go then 'go build github.com/johnkerl/miller' then the
// executable would be 'miller' not 'mlr'.
//
// So we have cmd/mlr/main.go:
// * go build   github.com/johnkerl/miller/v6/cmd/mlr
// * go install github.com/johnkerl/miller/v6/cmd/mlr

// go get github.com/johnkerl/lumin@v1.0.0
// Local development:
// replace github.com/johnkerl/lumin => /Users/kerl/git/johnkerl/lumin

go 1.24

require (
	github.com/facette/natsort v0.0.0-20181210072756-2cd4dd1e2dcb
	github.com/jedib0t/go-pretty/v6 v6.6.7
	github.com/johnkerl/lumin v1.0.0
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51
	github.com/klauspost/compress v1.17.10
	github.com/lestrrat-go/strftime v1.1.0
	github.com/mattn/go-isatty v0.0.20
	github.com/nine-lives-later/go-windows-terminal-sequences v1.0.4
	github.com/pkg/profile v1.7.0
	github.com/rivo/uniseg v0.4.7
	github.com/stretchr/testify v1.10.0
	golang.org/x/sys v0.30.0
	golang.org/x/term v0.29.0
	golang.org/x/text v0.22.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/fgprof v0.9.5 // indirect
	github.com/google/pprof v0.0.0-20240227163752-401108e1b7e7 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/rivo/uniseg => github.com/waltarix/uniseg v0.4.7-custom-r1
