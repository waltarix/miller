module github.com/johnkerl/miller

// The repo is 'miller' and the executable is 'mlr', going back many years and
// predating the Go port.
//
// If we had ./mlr.go then 'go build github.com/johnkerl/miller' then the
// executable would be 'miller' not 'mlr'.
//
// So we have cmd/mlr/main.go:
// * go build   github.com/johnkerl/miller/cmd/mlr
// * go install github.com/johnkerl/miller/cmd/mlr

go 1.15

require (
	github.com/jedib0t/go-pretty/v6 v6.2.4
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51
	github.com/lestrrat-go/strftime v1.0.5
	github.com/mattn/go-isatty v0.0.14
	github.com/mattn/go-runewidth v0.0.9
	github.com/pkg/profile v1.6.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c
	golang.org/x/term v0.0.0-20201210144234-2321bbc49cbf
)

replace github.com/mattn/go-runewidth => github.com/waltarix/go-runewidth v0.0.13-custom-r1
