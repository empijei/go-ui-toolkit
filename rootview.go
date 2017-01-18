package goUIToolKit

import (
	"bytes"
	"io"
	"text/template"
)

const rootViewHtml = `
<!DOCTYPE html>
<html lang="en">
<head>
  <title>{{.Title}}</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <!-- TODO fix this -->
  <link rel="stylesheet" href="/bootstrap.min.css">
  <script src="/jquery.min.js"></script>
  <script src="/bootstrap.min.js"></script>
  <script src="/goUIToolKit.js"></script>
  <!-- 
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
  -->
</head>
<body>
<div class="container">
{{.Body}}
</div>
</body>
</html>
`

var rootViewTmpl, _ = template.New("New View").Parse(rootViewHtml)

type View struct {
	finalized  bool
	title      string
	components []Component
	body       *bytes.Buffer
	content    []byte
	template   *template.Template
}

func (v *View) finalize() {
	if v.finalized {
		return
	}
	v.finalized = true
	for _, c := range v.components {
		_, _ = io.Copy(v.body, c.Buffer())
	}
	data := struct {
		Title string
		Body  string
	}{
		v.title,
		v.body.String(),
	}
	tmpCnt := bytes.NewBuffer(nil)
	err := v.template.Execute(tmpCnt, data)
	if err != nil {
		panic(err)
	}
	v.content = tmpCnt.Bytes()
}

func NewRootView(title string) *View {
	return &View{
		body:     bytes.NewBuffer(nil),
		title:    title,
		template: rootViewTmpl,
	}
}

func (v *View) AddComponent(c Component) {
	v.components = append(v.components, c)
}
