// Code generated by wit-bindgen-go. DO NOT EDIT.

package plugin

import (
	"go.bytecodealliance.org/cm"
)

// This file contains wasmimport and wasmexport declarations for "octahive:octabot@0.1.0".

//go:wasmexport octahive:octabot/plugin@0.1.0#init
//export octahive:octabot/plugin@0.1.0#init
func wasmexport_Init() (result *Metadata) {
	result_ := Exports.Init()
	result = &result_
	return
}

//go:wasmexport octahive:octabot/plugin@0.1.0#process
//export octahive:octabot/plugin@0.1.0#process
func wasmexport_Process(config0 *uint8, config1 uint32, payload0 *uint8, payload1 uint32) (result *cm.Result[ErrorShape, cm.List[Action], Error]) {
	config := cm.LiftString[string]((*uint8)(config0), (uint32)(config1))
	payload := cm.LiftString[string]((*uint8)(payload0), (uint32)(payload1))
	result_ := Exports.Process(config, payload)
	result = &result_
	return
}