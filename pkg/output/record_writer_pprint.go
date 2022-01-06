package output

import (
	"bufio"
	"container/list"
	"strings"

	"github.com/johnkerl/miller/pkg/cli"
	"github.com/johnkerl/miller/pkg/mlrval"
	"github.com/johnkerl/miller/pkg/types"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type RecordWriterPPRINT struct {
	writerOptions *cli.TWriterOptions
	// Input:
	records *list.List

	// State:
	lastJoinedHeader *string
	batch            *list.List
}

func NewRecordWriterPPRINT(writerOptions *cli.TWriterOptions) (*RecordWriterPPRINT, error) {
	return &RecordWriterPPRINT{
		writerOptions: writerOptions,
		records:       list.New(),

		lastJoinedHeader: nil,
		batch:            list.New(),
	}, nil
}

// ----------------------------------------------------------------
func (writer *RecordWriterPPRINT) Write(
	outrec *mlrval.Mlrmap,
	_ *types.Context,
	bufferedOutputStream *bufio.Writer,
	outputIsStdout bool,
) error {
	// Group records by have-same-schema or not. Pretty-print each
	// homoegeneous sublist, or "batch".
	//
	// No output until end of a homogeneous batch of records, since we need to
	// find out max width down each column.

	if outrec != nil { // Not end of record stream
		if writer.lastJoinedHeader == nil {
			// First output record:
			// * New batch
			// * No old batch to print
			writer.batch.PushBack(outrec)
			temp := strings.Join(outrec.GetKeys(), ",")
			writer.lastJoinedHeader = &temp
		} else {
			// May or may not continue the same homogeneous batch
			joinedHeader := strings.Join(outrec.GetKeys(), ",")
			if *writer.lastJoinedHeader != joinedHeader {
				// Print and free old batch
				nonEmpty := writer.writeHeterogenousList(
					writer.batch,
					writer.writerOptions.BarredPprintOutput,
					bufferedOutputStream,
					outputIsStdout,
				)
				if nonEmpty {
					// Print a newline
					bufferedOutputStream.WriteString(writer.writerOptions.ORS)
				}
				// Start a new batch
				writer.batch = list.New()
				writer.batch.PushBack(outrec)
				writer.lastJoinedHeader = &joinedHeader
			} else {
				// Continue the batch
				writer.batch.PushBack(outrec)
			}
		}

	} else { // End of record stream
		if writer.batch.Front() != nil {
			writer.writeHeterogenousList(writer.batch, writer.writerOptions.BarredPprintOutput,
				bufferedOutputStream, outputIsStdout)
		}
	}

	return nil
}

// ----------------------------------------------------------------
// Returns false if there was nothing but empty record(s), e.g. 'mlr gap -n 10'.
func (writer *RecordWriterPPRINT) writeHeterogenousList(
	records *list.List,
	barred bool,
	bufferedOutputStream *bufio.Writer,
	outputIsStdout bool,
) bool {
	var maxNR int64 = 0

	for e := records.Front(); e != nil; e = e.Next() {
		outrec := e.Value.(*mlrval.Mlrmap)
		nr := outrec.FieldCount
		if maxNR < nr {
			maxNR = nr
		}
	}

	if maxNR == 0 {
		return false
	}

	onFirst := true
	t := getWriter(barred, bufferedOutputStream)
	for e := records.Front(); e != nil; e = e.Next() {
		outrec := e.Value.(*mlrval.Mlrmap)

		// Print header line
		if onFirst && !writer.writerOptions.HeaderlessOutput {
			headers := table.Row{}
			for pe := outrec.Head; pe != nil; pe = pe.Next {
				headers = append(headers, pe.Key)
			}
			t.AppendHeader(headers)
			onFirst = false
		}

		// Print data lines
		cols := table.Row{}
		for pe := outrec.Head; pe != nil; pe = pe.Next {
			cols = append(cols, pe.Value.String())
		}

		t.AppendRow(cols)
	}

	t.Render()

	return true
}

func getWriter(barred bool, bufferedOutputStream *bufio.Writer) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(bufferedOutputStream)

	if barred {
		t.SetStyle(table.StyleRounded)
		t.Style().Format.Header = text.FormatDefault
		t.Style().Format.Footer = text.FormatDefault
		return t
	}

	t.SetStyle(table.Style{
		Name: "NonBarred",
		Box: table.BoxStyle{
			BottomLeft:       "",
			BottomRight:      "",
			BottomSeparator:  "",
			Left:             "",
			LeftSeparator:    "",
			MiddleHorizontal: " ",
			MiddleSeparator:  "",
			MiddleVertical:   "",
			PaddingLeft:      "",
			PaddingRight:     " ",
			Right:            "",
			RightSeparator:   "",
			TopLeft:          "",
			TopRight:         "",
			TopSeparator:     "",
			UnfinishedRow:    "",
		},
		Format: table.FormatOptions{
			Header: text.FormatDefault,
			Row:    text.FormatDefault,
			Footer: text.FormatDefault,
		},
		Options: table.Options{
			DrawBorder:      false,
			SeparateColumns: false,
			SeparateFooter:  false,
			SeparateHeader:  false,
			SeparateRows:    false,
		},
	})

	return t
}
