// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package plugin represents the exported interface "octahive:octabot/plugin@0.1.0".
package plugin

import (
	"go.bytecodealliance.org/cm"
)

// Action represents the record "octahive:octabot/plugin@0.1.0#action".
//
// Action
//
//	record action {
//		name: string,
//	}
type Action struct {
	_ cm.HostLayout
	// The name of the action
	Name string
}

// Metadata represents the record "octahive:octabot/plugin@0.1.0#metadata".
//
// The metadata for a plugin used for registration and setup
//
//	record metadata {
//		name: string,
//		version: string,
//		author: string,
//		description: string,
//	}
type Metadata struct {
	_ cm.HostLayout
	// The friendly name of the plugin
	Name string

	// The version of the plugin
	Version string

	// The author of the plugin
	Author string

	// The description of the plugin. This will be used as the top level help text for
	// the plugin
	Description string
}

// Error represents the variant "octahive:octabot/plugin@0.1.0#error".
//
// Errors related to interacting with Plugin
//
//	variant error {
//		other(string),
//	}
type Error cm.Variant[uint8, string, string]

// ErrorOther returns a [Error] of case "other".
//
// Some other error occurred.
func ErrorOther(data string) Error {
	return cm.New[Error](0, data)
}

// Other returns a non-nil *[string] if [Error] represents the variant case "other".
func (self *Error) Other() *string {
	return cm.Case[string](self, 0)
}

var stringsError = [1]string{
	"other",
}

// String implements [fmt.Stringer], returning the variant case name of v.
func (v Error) String() string {
	return stringsError[v.Tag()]
}
