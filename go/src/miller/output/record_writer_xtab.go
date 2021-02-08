package output

import (
	"bytes"
	"io"

	"github.com/mattn/go-runewidth"

	"miller/cliutil"
	"miller/types"
)

type RecordWriterXTAB struct {
	onFirst bool
}

func NewRecordWriterXTAB(writerOptions *cliutil.TWriterOptions) *RecordWriterXTAB {
	return &RecordWriterXTAB{
		onFirst: true,
	}
}

func (this *RecordWriterXTAB) Write(
	outrec *types.Mlrmap,
	ostream io.WriteCloser,
) {
	// End of record stream: nothing special for this output format
	if outrec == nil {
		return
	}

	maxKeyLength := 1
	for pe := outrec.Head; pe != nil; pe = pe.Next {
		keyLength := runewidth.StringWidth(pe.Key)
		if keyLength > maxKeyLength {
			maxKeyLength = keyLength
		}
	}

	var buffer bytes.Buffer // 5x faster than fmt.Print() separately

	// Put a blank line between records, but not before the first or after the last
	if this.onFirst {
		this.onFirst = false
	} else {
		buffer.WriteString("\n")
	}

	for pe := outrec.Head; pe != nil; pe = pe.Next {
		keyLength := runewidth.StringWidth(pe.Key)
		padLength := maxKeyLength - keyLength

		buffer.WriteString(pe.Key)
		buffer.WriteString(" ")
		for i := 0; i < padLength; i++ {
			buffer.WriteString(" ")
		}
		buffer.WriteString(pe.Value.String())
		buffer.WriteString("\n")
	}
	ostream.Write(buffer.Bytes())
}
