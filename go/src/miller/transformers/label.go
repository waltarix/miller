package transformers

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"miller/cliutil"
	"miller/lib"
	"miller/transforming"
	"miller/types"
)

// ----------------------------------------------------------------
const verbNameLabel = "label"

var LabelSetup = transforming.TransformerSetup{
	Verb:         verbNameLabel,
	ParseCLIFunc: transformerLabelParseCLI,
	UsageFunc:    transformerLabelUsage,
	IgnoresInput: false,
}

func transformerLabelParseCLI(
	pargi *int,
	argc int,
	args []string,
	_ *cliutil.TReaderOptions,
	__ *cliutil.TWriterOptions,
) transforming.IRecordTransformer {

	// Skip the verb name from the current spot in the mlr command line
	argi := *pargi
	argi++

	for argi < argc /* variable increment: 1 or 2 depending on flag */ {
		opt := args[argi]
		if !strings.HasPrefix(opt, "-") {
			break // No more flag options to process
		}
		argi++

		if opt == "-h" || opt == "--help" {
			transformerLabelUsage(os.Stdout, true, 0)

		} else {
			transformerLabelUsage(os.Stderr, true, 1)
		}
	}

	// Get the label field names from the command line
	if argi >= argc {
		transformerLabelUsage(os.Stderr, true, 1)
	}
	newNames := lib.SplitString(args[argi], ",")
	argi++

	transformer, err := NewTransformerLabel(
		newNames,
	)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return nil
	}

	*pargi = argi
	return transformer
}

func transformerLabelUsage(
	o *os.File,
	doExit bool,
	exitCode int,
) {
	fmt.Fprintf(o, "Usage: %s %s [options] {new1,new2,new3,...}\n", lib.MlrExeName(), verbNameLabel)
	fmt.Fprintf(o, "Given n comma-separated names, renames the first n fields of each record to\n")
	fmt.Fprintf(o, "have the respective name. (Fields past the nth are left with their original\n")
	fmt.Fprintf(o, "names.) Particularly useful with --inidx or --implicit-csv-header, to give\n")
	fmt.Fprintf(o, "useful names to otherwise integer-indexed fields.\n")
	fmt.Fprintf(o, "\n")
	fmt.Fprintf(o, "Options:\n")
	fmt.Fprintf(o, "-h|--help Show this message.\n")

	if doExit {
		os.Exit(exitCode)
	}
}

// ----------------------------------------------------------------
type TransformerLabel struct {
	newNames []string
}

func NewTransformerLabel(
	newNames []string,
) (*TransformerLabel, error) {
	// TODO: make this a library function.
	uniquenessChecker := make(map[string]bool)
	for _, newName := range newNames {
		_, ok := uniquenessChecker[newName]
		if ok {
			return nil, errors.New(
				fmt.Sprintf(
					"mlr label: labels must be unique; got duplicate \"%s\"\n",
					newName,
				),
			)
		}
		uniquenessChecker[newName] = true
	}

	this := &TransformerLabel{
		newNames: newNames,
	}

	return this, nil
}

// ----------------------------------------------------------------
func (this *TransformerLabel) Transform(
	inrecAndContext *types.RecordAndContext,
	outputChannel chan<- *types.RecordAndContext,
) {
	if !inrecAndContext.EndOfStream {
		inrec := inrecAndContext.Record
		inrec.Label(this.newNames)
	}
	outputChannel <- inrecAndContext // including end-of-stream marker
}
