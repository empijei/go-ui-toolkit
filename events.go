package goUIToolKit

var callbacks = make(map[string]map[string]CallBack)

type CallBack func(component *Component, event string, args ...map[string]string)

func addRuntimeEventHandler(eventname, componentId string, callback CallBack) {
	//TODO contact runtime applier
	addEventHandler(eventname, componentId, callback)
}

func addEventHandler(eventname, componentId string, callback CallBack) {
	if _, ok := callbacks[componentId]; !ok {
		callbacks[componentId] = make(map[string]CallBack)
	}
	callbacks[componentId][eventname] = callback
}
