package goUIToolKit

import "io"

type Component interface {
	//A call to Buffer should finalize the component
	Buffer() io.Reader
	ID() string
}
