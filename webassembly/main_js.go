//go:build js && wasm
// +build js,wasm

package main

//go:generate cp $GOROOT/misc/wasm/wasm_exec.js .

import (
	"bytes"
	"syscall/js"

	goMacho "github.com/blacktop/go-macho"
)

func readMacho_JavaScript(this js.Value, args []js.Value) interface{} {
	buffer := make([]byte, args[0].Length())

	js.CopyBytesToGo(buffer, args[0])

	reader := bytes.NewReader(buffer)
	m, err := goMacho.NewFile(reader)
	if err != nil {
		panic(err)
	}

	return m.FileTOC.String()
}

func main() {
	js.Global().Set("readMachoGo", js.FuncOf(readMacho_JavaScript))
	select {} // block the main thread forever
}
