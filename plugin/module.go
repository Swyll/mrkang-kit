package plugin

import (
	"io"
)

type Plugin interface {
	GetMsg() map[string]string
	SetArgs(...string)
	AddEnv(k, v string)

	Run(io.Writer) error
	Stop() error
	Cannel() error
}
