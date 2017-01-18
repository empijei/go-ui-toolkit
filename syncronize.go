package goUIToolKit

var unmarshals = make(map[string]Unmarshal)

type Unmarshal func(buf []byte) Component

func addUnmarshal(tagname string, unmarshal Unmarshal) {
	unmarshals[tagname] = unmarshal
}

func addMarshal(tagname string, js string) {}
