package goUIToolKit

import (
	"fmt"
	"math/rand"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

var runtimeinit sync.Once
var curRuntime *Runtime

func GetRuntime() *Runtime {
	runtimeinit.Do(func() {
		curRuntime = &Runtime{
			ids: make(map[string]struct{}),
		}
	})
	return curRuntime
}

type Runtime struct {
	started bool
	pattern string
	ids     map[string]struct{}
	ws      *websocket.Conn
}

func (r *Runtime) Start(rootView *View) {
	r.started = true
	rootView.finalize()
	//Generate listener with random path
	r.pattern = "/" + randString([]rune(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`), 16)

	onConnected := func(ws *websocket.Conn) {
		defer func() {
			//TODO handle this error?
			_ = ws.Close()
		}()
		//TODO thread-safe this
		r.ws = ws
		//TODO if already connected reject
		go runtimeSender(ws)
		go runtimeReceiver(ws)
		//client := NewClient(ws, s)
		//s.AddClient(client)
		//client.Listen()
	}

	http.Handle(r.pattern+"/ws", websocket.Handler(onConnected))

	http.HandleFunc(r.pattern+"/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(rootView.content)
	})
	http.Handle("/", http.FileServer(assetFS()))

	port := rand.Intn(1000) + 9000
	strport := ":" + strconv.Itoa(port)
	go func() {
		time.Sleep(1 * time.Second)
		var cmdStr, arg string
		arg = "http://localhost" + strport + r.pattern
		//TODO try to connect to test if it is the correct endpoint
		switch runtime.GOOS {
		case "linux":
			cmdStr = "xdg-"
			fallthrough
		case "darwin":
			cmdStr += "open"
		case "windows":
			cmdStr = "start"
		}
		cmd := exec.Command(cmdStr, arg)
		_ = cmd.Run()
	}()
	err := http.ListenAndServe("localhost"+strport, nil)
	for err != nil && strings.Contains(err.Error(), "address already in use") {
		port = rand.Intn(1000) + 9000
		strport = ":" + strconv.Itoa(port)
		err = http.ListenAndServe("localhost"+strport, nil)
	}

}

func (r *Runtime) genID() string {
	id := "goUIToolKitID-" + randString([]rune(`0123456789abcdef`), 10)
	//Make sure the generated ID is unique
	for _, ok := r.ids[id]; ok; {
		id = "goUIToolKitID-" + randString([]rune(`0123456789abcdef`), 10)
	}
	//Add ID to ids Set
	r.ids[id] = struct{}{}
	return id
}

func idchecker(id string) string {
	r := GetRuntime()
	if id == "" {
		//User did not provide ID, generate one
		id = r.genID()
	} else {
		if _, ok := r.ids[id]; !ok {
			//ID is unique, add it to ids set
			r.ids[id] = struct{}{}
		} else {
			//Duplicate user-defined ID, abort
			panic(fmt.Errorf("Duplicate ID: \"%s\" ", id))
		}
	}
	return id
}
