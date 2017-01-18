package goUIToolKit

import (
	"bytes"
	"encoding/json"
	"io"
	"text/template"
)

const btnHtml = `
<button 
	type="button" 
	id="{{.Id}}" 
	class="btn btn-{{.Color}}"
	onclick="btn-clickhandler()"
	>
		{{.Text}}
</button>
`

var btnTmpl, _ = template.New("button").Parse(btnHtml)

func init() {
	addUnmarshal("button", func(buf []byte) Component {
		type btn struct {
			Color, Text, ID string
		}
		var tmpbtn btn
		err := json.Unmarshal(buf, tmpbtn)
		if err != nil {
			//TODO create a generic handler for this
		}
		button := &Button{
			//TODO color unserilization
			//color: tmpbtn.Color,
			text: tmpbtn.Text,
			id:   tmpbtn.ID,
		}
		return button
	})
	addMarshal("button", `
	function(e){
		return {
			//TODO color
			"Color": e.className,
			"Text": e.innerText()
			"ID": e.id
		}
	}
	`)
}

type Button struct {
	color     Color
	text      string
	finalized bool
	id        string
}

//Creates a new Button with the specified id and sets default values.
//if an empty string is passed, and id will be set by the view
//the button is added to.
func NewButton(id string) *Button {
	id = idchecker(id)
	return &Button{id: id, color: DEFAULT}
}

func (b *Button) Getcolor() Color {
	return b.color
}

func (b *Button) SetColor(c Color) {
	b.color = c
	if b.finalized {
		//Call to runtime applier!
	}
}

func (b *Button) Buffer() io.Reader {
	b.finalized = true
	buf := bytes.NewBuffer(nil)
	data := struct {
		Id    string
		Color string
		Text  string
	}{
		b.id,
		b.color.String(),
		b.text,
	}
	err := btnTmpl.Execute(buf, data)
	if err != nil {
		panic(err)
	}
	return buf
}

func (b *Button) ID() string {
	return b.id
}

func (b *Button) SetText(text string) {
	b.text = text
	if b.finalized {
		//Call to runtime applier!
	}
}

func (b *Button) OnClick(c CallBack) {
	addEventHandler("click", b.id, c)
}
