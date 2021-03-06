// ================================================================
// Just playing around -- nothing serious here.
// ================================================================

package repl

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"miller/cliutil"
	"miller/dsl/cst"
	"miller/input"
	"miller/output"
	"miller/runtime"
	"miller/types"
)

// ----------------------------------------------------------------
func NewRepl(
	exeName string,
	replName string,
	astPrintMode ASTPrintMode,
	options *cliutil.TOptions,
) (*Repl, error) {

	recordReader := input.Create(&options.ReaderOptions)
	if recordReader == nil {
		return nil, errors.New("Input format not found: " + options.ReaderOptions.InputFileFormat)
	}

	recordWriter := output.Create(&options.WriterOptions)
	if recordWriter == nil {
		return nil, errors.New("Output format not found: " + options.WriterOptions.OutputFileFormat)
	}

	// $* is the empty map {} until/unless the user opens a file and reads records from it.
	inrec := types.NewMlrmapAsRecord()
	// NR is 0, etc until/unless the user opens a file and reads records from it.
	context := types.NewContext(options)
	runtimeState := runtime.NewEmptyState()
	runtimeState.Update(inrec, context)
	// The filter expression for the main Miller DSL is any non-assignment
	// statment like 'true' or '$x > 0.5' etc. For the REPL, we re-use this for
	// interactive expressions to be printed to the terminal. For the main DSL,
	// the default is types.MlrvalFromTrue(); for the REPL, the default is
	// types.MlrvalFromVoid().
	runtimeState.FilterExpression = types.MlrvalFromVoid()

	return &Repl{
		exeName:  exeName,
		replName: replName,

		inputIsTerminal: getInputIsTerminal(),
		prompt1:         getPrompt1(),
		prompt2:         getPrompt2(),

		astPrintMode: ASTPrintNone,
		cstRootNode:  cst.NewEmptyRoot(&options.WriterOptions).WithRedefinableUDFS(),

		options:      options,
		inputChannel: nil,
		errorChannel: nil,
		recordReader: recordReader,
		recordWriter: recordWriter,

		runtimeState: runtimeState,
	}, nil
}

// ----------------------------------------------------------------
func (this *Repl) handleSession(istream *os.File) {
	this.printStartupBanner()

	lineReader := bufio.NewReader(istream)

	for {
		this.printPrompt1()

		line, err := lineReader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "%s %s: %w\n", this.exeName, this.replName, err)
			os.Exit(1)
		}

		// This trims the trailing newline, as well as leading/trailing whitespace:
		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "<" {
			this.handleMultiline(lineReader)
		} else if trimmedLine == ":quit" {
			break
		} else if this.handleNonDSLLine(trimmedLine) {
			// Handled in that method.
		} else {
			// We need the non-trimmed line here since the DSL syntax for comments is '#.*\n'.
			err = this.handleDSLStringImmediate(line)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}

// ----------------------------------------------------------------
// Context: the "<" has already been seen. we read until ">".

func (this *Repl) handleMultiline(lineReader *bufio.Reader) {
	var buffer bytes.Buffer
	for {
		this.printPrompt2()

		line, err := lineReader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "%s %s: %w\n", this.exeName, this.replName, err)
			os.Exit(1)
		}

		if strings.TrimSpace(line) == ">" {
			break
		}
		buffer.WriteString(line)
	}
	dslString := buffer.String()

	err := this.handleDSLStringBulk(dslString)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
