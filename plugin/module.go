package plugin

import (
	"io"
)

type Plugin interface {
	Run(io.Writer) error
	Stop() error
	Cannel() error
}
