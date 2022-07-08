package repo

import (
	"errors"
	"webinar/repo/memory"
	"webinar/repo/postgres"
)

type Options struct {
	Environment string
}

var ErrEmptyOptions = errors.New("empty options")
var ErrInvalidOptions = errors.New("invalid options")

type Getter interface {
	Get(id int) string
}

func NewGetter(opts *Options) (Getter, error) {
	if opts == nil {
		return nil, ErrEmptyOptions
	}

	switch opts.Environment {
	case "debug":
		return memory.New(), nil
	case "release":
		return postgres.New(), nil
	default:
		return nil, errors.New("invalid settings")
	}

}
