module miller

go 1.15

require (
	github.com/goccmack/gocc v0.0.0-20210201103733-1bd198f09019 // indirect
	github.com/jedib0t/go-pretty/v6 v6.1.0 // indirect
	golang.org/x/term v0.0.0-20201210144234-2321bbc49cbf // indirect
	miller/auxents v0.0.0-00010101000000-000000000000
	miller/cli v0.0.0-00010101000000-000000000000
	miller/cliutil v0.0.0-00010101000000-000000000000 // indirect
	miller/dsl v0.0.0-00010101000000-000000000000
	miller/input v0.0.0-00010101000000-000000000000 // indirect
	miller/lib v0.0.0-00010101000000-000000000000 // indirect
	miller/output v0.0.0-00010101000000-000000000000 // indirect
	miller/parsing v0.0.0-00010101000000-000000000000 // indirect
	miller/runtime v0.0.0-00010101000000-000000000000 // indirect
	miller/stream v0.0.0-00010101000000-000000000000
	miller/transformers v0.0.0-00010101000000-000000000000 // indirect
	miller/transforming v0.0.0-00010101000000-000000000000 // indirect
	miller/types v0.0.0-00010101000000-000000000000 // indirect
	miller/version v0.0.0-00010101000000-000000000000 // indirect
)

replace miller/cli => ./src/miller/cli

replace miller/cliutil => ./src/miller/cliutil

replace miller/dsl => ./src/miller/dsl

replace miller/input => ./src/miller/input

replace miller/lib => ./src/miller/lib

replace miller/output => ./src/miller/output

replace miller/parsing => ./src/miller/parsing

replace miller/stream => ./src/miller/stream

replace miller/transformers => ./src/miller/transformers

replace miller/transforming => ./src/miller/transforming

replace miller/types => ./src/miller/types

replace miller/version => ./src/miller/version

replace miller/runtime => ./src/miller/runtime

replace miller/auxents => ./src/miller/auxents

replace github.com/mattn/go-runewidth => github.com/waltarix/go-runewidth v0.0.10-custom
