package goUIToolKit

import (
	"bytes"
	"io"
)

var divHtml = `
<div class="{{.Class}}">
{{.Body}}
</div>
`

//This is not implemented yet!
type Div struct {
	finalized  bool
	title      string
	components []Component
	body       *bytes.Buffer
}

func (d *Div) Buffer() io.Reader {
	panic("not implemented")
}

func (d *Div) ID() string {
	panic("not implemented")
}

func (d *Div) SetID(string) {
	panic("not implemented")
}

func (d *Div) AddComponent(c *Component) (id string) {
	panic("not implemented")
}
