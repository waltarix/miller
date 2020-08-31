package mapping

import (
	"miller/containers"
	"miller/runtime"
)

type IRecordMapper interface {
	Name() string
	Map(
		inrec *containers.Lrec,
		context *runtime.Context,
		outrecs chan<- *containers.Lrec,
	)
}