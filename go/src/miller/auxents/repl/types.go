// ================================================================
// Just playing around -- nothing serious here.
// ================================================================

package repl

import (
	"miller/cliutil"
	"miller/dsl/cst"
	"miller/input"
	"miller/output"
	"miller/runtime"
	"miller/types"
)

// ================================================================
type ASTPrintMode int

const (
	ASTPrintNone ASTPrintMode = iota
	ASTPrintParex
	ASTPrintParexOneLine
	ASTPrintIndent
)

// ================================================================
type Repl struct {
	// From os.Args[] as we were invoked. These are for printing error messages.
	exeName  string
	replName string

	// Prompt1 is the main prompt, like $PS1. Prompt2 is for "<" ... ">"
	// multi-line-input mode.
	inputIsTerminal     bool
	prompt1             string
	prompt2             string
	doingMultilineInput bool

	astPrintMode ASTPrintMode
	cstRootNode  *cst.RootNode

	options *cliutil.TOptions

	inputChannel chan *types.RecordAndContext
	errorChannel chan error
	recordReader input.IRecordReader
	recordWriter output.IRecordWriter

	runtimeState *runtime.State
}
