package main

import "github.com/empijei/goUIToolKit"

func main() {
	v := goUIToolKit.NewRootView("Foo")
	b := goUIToolKit.NewButton("btn1")
	b.SetColor(goUIToolKit.BLUE)
	b.SetText("SimpleButton")
	v.AddComponent(b)
	r := goUIToolKit.GetRuntime()
	r.Start(v)
}
