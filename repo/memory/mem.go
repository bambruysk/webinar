package memory

import "fmt"

type memstorage struct {
}

func New() *memstorage {
	return &memstorage{}
}

func (m memstorage) Get(id int) string {
	return fmt.Sprintf("memory: %d", id)
}
