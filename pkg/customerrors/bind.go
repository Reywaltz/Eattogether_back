package customerrors

import "fmt"

type DataNotBindable struct {
	Struct interface{}
}

func (d *DataNotBindable) Error() string {
	return fmt.Sprintf("can't bind to provided struct: %v\n", d.Struct)
}
