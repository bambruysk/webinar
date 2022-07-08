package postgres

import "fmt"

type pgstorage struct {
}

func New() *pgstorage {
	return &pgstorage{}
}

func (m pgstorage) Get(id int) string {
	return fmt.Sprintf("storage: %d", id)
}
