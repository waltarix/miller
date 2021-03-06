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
const verbNameSeqgen = "seqgen"

var SeqgenSetup = transforming.TransformerSetup{
	Verb:         verbNameSeqgen,
	ParseCLIFunc: transformerSeqgenParseCLI,
	UsageFunc:    transformerSeqgenUsage,
	IgnoresInput: true,
}

func transformerSeqgenParseCLI(
	pargi *int,
	argc int,
	args []string,
	_ *cliutil.TReaderOptions,
	__ *cliutil.TWriterOptions,
) transforming.IRecordTransformer {

	// Skip the verb name from the current spot in the mlr command line
	argi := *pargi
	verb := args[argi]
	argi++

	fieldName := "i"
	startString := "1"
	stopString := "100"
	stepString := "1"

	for argi < argc /* variable increment: 1 or 2 depending on flag */ {
		opt := args[argi]
		if !strings.HasPrefix(opt, "-") {
			break // No more flag options to process
		}
		argi++

		if opt == "-h" || opt == "--help" {
			transformerSeqgenUsage(os.Stdout, true, 0)

		} else if opt == "-f" {
			fieldName = cliutil.VerbGetStringArgOrDie(verb, opt, args, &argi, argc)

		} else if opt == "--start" {
			startString = cliutil.VerbGetStringArgOrDie(verb, opt, args, &argi, argc)

		} else if opt == "--stop" {
			stopString = cliutil.VerbGetStringArgOrDie(verb, opt, args, &argi, argc)

		} else if opt == "--step" {
			stepString = cliutil.VerbGetStringArgOrDie(verb, opt, args, &argi, argc)

		} else {
			transformerSeqgenUsage(os.Stderr, true, 1)
		}
	}

	transformer, err := NewTransformerSeqgen(
		fieldName,
		startString,
		stopString,
		stepString,
	)
	// TODO: put error return into this API
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	*pargi = argi
	return transformer
}

func transformerSeqgenUsage(
	o *os.File,
	doExit bool,
	exitCode int,
) {
	fmt.Fprintf(o, "Usage: %s %s [options]\n", lib.MlrExeName(), verbNameSeqgen)
	fmt.Fprintf(o, "Passes input records directly to output. Most useful for format conversion.\n")
	fmt.Fprintf(o, "Produces a sequence of counters.  Discards the input record stream. Produces\n")
	fmt.Fprintf(o, "output as specified by the options\n")
	fmt.Fprintf(o, "\n")
	fmt.Fprintf(o, "Options:\n")
	fmt.Fprintf(o, "-f {name} (default \"i\") Field name for counters.\n")
	fmt.Fprintf(o, "-start {value} (default 1) Inclusive start value.\n")
	fmt.Fprintf(o, "-step {value} (default 1) Step value.\n")
	fmt.Fprintf(o, "-stop {value} (default 100) Inclusive stop value.\n")
	fmt.Fprintf(o, "-h|--help Show this message.\n")

	fmt.Fprintf(o, "Start, stop, and/or step may be floating-point. Output is integer if start,\n")
	fmt.Fprintf(o, "stop, and step are all integers. Step may be negative. It may not be zero\n")
	fmt.Fprintf(o, "unless start == stop.\n")

	if doExit {
		os.Exit(exitCode)
	}
}

// ----------------------------------------------------------------
type TransformerSeqgen struct {
	fieldName      string
	start          types.Mlrval
	stop           types.Mlrval
	step           types.Mlrval
	doneComparator types.BinaryFunc
}

// ----------------------------------------------------------------
func NewTransformerSeqgen(
	fieldName string,
	startString string,
	stopString string,
	stepString string,
) (*TransformerSeqgen, error) {
	start := types.MlrvalFromInferredType(startString)
	stop := types.MlrvalFromInferredType(stopString)
	step := types.MlrvalFromInferredType(stepString)
	var doneComparator types.BinaryFunc = nil

	fstart, startIsNumeric := start.GetNumericToFloatValue()
	if !startIsNumeric {
		return nil, errors.New(
			fmt.Sprintf(
				"mlr seqgen: start value should be number; got \"%s\"",
				startString,
			),
		)
	}

	fstop, stopIsNumeric := stop.GetNumericToFloatValue()
	if !stopIsNumeric {
		return nil, errors.New(
			fmt.Sprintf(
				"mlr seqgen: stop value should be number; got \"%s\"",
				stopString,
			),
		)
	}

	fstep, stepIsNumeric := step.GetNumericToFloatValue()
	if !stepIsNumeric {
		return nil, errors.New(
			fmt.Sprintf(
				"mlr seqgen: step value should be number; got \"%s\"",
				stepString,
			),
		)
	}

	if fstep > 0 {
		doneComparator = types.MlrvalGreaterThan
	} else if fstep < 0 {
		doneComparator = types.MlrvalLessThan
	} else {
		if fstart == fstop {
			doneComparator = types.MlrvalEquals
		} else {
			return nil, errors.New(
				"mlr seqgen: step must not be zero unless start == stop.",
			)
		}
	}

	return &TransformerSeqgen{
		fieldName:      fieldName,
		start:          start,
		stop:           stop,
		step:           step,
		doneComparator: doneComparator,
	}, nil
}

// ----------------------------------------------------------------
func (this *TransformerSeqgen) Transform(
	inrecAndContext *types.RecordAndContext,
	outputChannel chan<- *types.RecordAndContext,
) {
	counter := this.start
	context := types.NewContext(nil)
	context.UpdateForStartOfFile("seqgen")

	for {
		mdone := this.doneComparator(&counter, &this.stop)
		done, _ := mdone.GetBoolValue()
		if done {
			break
		}

		outrec := types.NewMlrmapAsRecord()
		outrec.PutCopy(this.fieldName, &counter)

		context.UpdateForInputRecord()

		outrecAndContext := types.NewRecordAndContext(outrec, context)
		outputChannel <- outrecAndContext

		counter = types.MlrvalBinaryPlus(&counter, &this.step)
	}

	outputChannel <- types.NewEndOfStreamMarker(context)
}
