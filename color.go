package goUIToolKit

type Color int

const (
	//White
	DEFAULT = iota
	//Blue
	PRIMARY
	//Green
	SUCCESS
	//Light-Blue
	INFO
	//Yellow
	WARNING
	//Red
	DANGER
	LINK
)
const (
	WHITE = iota
	BLUE
	GREEN
	LIGHTBLUE
	YELLOW
	RED
)

var colors = [...]string{
	"default",
	"primary",
	"success",
	"info",
	"warning",
	"danger",
	"link",
}

func (c Color) String() string {
	return colors[c]
}

type Colored interface {
	Color() Color
	SetColor(Color)
}
