package output

import (
	"container/list"
	"io"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"miller/cliutil"
	"miller/types"
)

type RecordWriterPPRINT struct {
	// Input:
	records *list.List
	// For detecting schema changes: we print a newline and the new header.
	barred bool

	// State:
	lastJoinedHeader *string
	batch            *list.List
}

func NewRecordWriterPPRINT(writerOptions *cliutil.TWriterOptions) *RecordWriterPPRINT {
	return &RecordWriterPPRINT{
		records: list.New(),
		barred:  writerOptions.BarredPprintOutput,

		lastJoinedHeader: nil,
		batch:            list.New(),
	}
}

// ----------------------------------------------------------------
func (this *RecordWriterPPRINT) Write(
	outrec *types.Mlrmap,
	ostream io.WriteCloser,
) {
	// Group records by have-same-schema or not. Pretty-print each
	// homoegeneous sublist, or "batch".
	//
	// No output until end of a homogeneous batch of records, since we need to
	// find out max width down each column.

	if outrec != nil { // Not end of record stream

		if this.lastJoinedHeader == nil {
			// First output record:
			// * New batch
			// * No old batch to print
			this.batch.PushBack(outrec)
			temp := strings.Join(outrec.GetKeys(), ",")
			this.lastJoinedHeader = &temp
		} else {
			// May or may not continue the same homogeneous batch
			joinedHeader := strings.Join(outrec.GetKeys(), ",")
			if *this.lastJoinedHeader != joinedHeader {
				// Print and free old batch
				if this.writeHeterogenousList(this.batch, this.barred, ostream) {
					// Print a newline
					ostream.Write([]byte("\n"))
				}
				// Start a new batch
				this.batch = list.New()
				this.batch.PushBack(outrec)
				this.lastJoinedHeader = &joinedHeader
			} else {
				// Continue the batch
				this.batch.PushBack(outrec)
			}
		}

	} else { // End of record stream

		if this.batch.Front() != nil {
			this.writeHeterogenousList(this.batch, this.barred, ostream)
		}
	}
}

// ----------------------------------------------------------------
// Returns false if there was nothing but empty record(s), e.g. 'mlr gap -n 10'.
func (this *RecordWriterPPRINT) writeHeterogenousList(
	records *list.List,
	barred bool,
	ostream io.WriteCloser,
) bool {
	var maxNR int = 0

	for e := records.Front(); e != nil; e = e.Next() {
		outrec := e.Value.(*types.Mlrmap)
		nr := outrec.FieldCount
		if maxNR < nr {
			maxNR = nr
		}
	}

	if maxNR == 0 {
		return false
	}

	onFirst := true
	t := getWriter(barred, ostream)
	for e := records.Front(); e != nil; e = e.Next() {
		outrec := e.Value.(*types.Mlrmap)

		// Print header line
		if onFirst {
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

// ----------------------------------------------------------------
func getWriter(barred bool, ostream io.WriteCloser) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(ostream)

	if barred {
		t.SetStyle(table.StyleRounded)
		t.Style().Format.Footer = text.FormatDefault
		t.Style().Format.Header = text.FormatDefault
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
			Footer: text.FormatDefault,
			Header: text.FormatDefault,
			Row:    text.FormatDefault,
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
