/*
 * @Description: 备注
 * @Author: wangcong0918
 * @Date: 2019-11-29 18:08:04
 * @LastEditTime: 2020-07-16 15:47:23
 * @LastEditors: wangcong0918
 */
// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sunrise

import (
	"io"
	"os"

	"github.com/wangcong0918/sunrise/binding"
)

// ENV_SUNRISE_MODE indicates environment name for gin mode.
const ENV_SUNRISE_MODE = "Sunrise_MODE"

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is relase.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)
const (
	debugCode = iota
	releaseCode
	testCode
)

// DefaultWriter is the default io.Writer used the Gin for debug output and
// middleware output like Logger() or Recovery().
// Note that both Logger and Recovery provides custom ways to configure their
// output io.Writer.
// To support coloring in Windows use:
// 		import "github.com/mattn/go-colorable"
// 		gin.DefaultWriter = colorable.NewColorableStdout()
var DefaultWriter io.Writer = os.Stdout
var DefaultErrorWriter io.Writer = os.Stderr

var ginMode = debugCode
var modeName = DebugMode

func init() {
	mode := os.Getenv(ENV_SUNRISE_MODE)
	SetMode(mode)
}

// SetMode sets gin mode according to input string.
func SetMode(value string) {
	switch value {
	case DebugMode, "":
		ginMode = debugCode
	case ReleaseMode:
		ginMode = releaseCode
	case TestMode:
		ginMode = testCode
	default:
		panic("gin mode unknown: " + value)
	}
	if value == "" {
		value = DebugMode
	}
	modeName = value
}

// DisableBindValidation closes the default validator.
func DisableBindValidation() {
	binding.Validator = nil
}

// EnableJsonDecoderUseNumber sets true for binding.EnableDecoderUseNumberto to
// call the UseNumber method on the JSON Decoder instance.
func EnableJsonDecoderUseNumber() {
	binding.EnableDecoderUseNumber = true
}

// Mode returns currently gin mode.
func Mode() string {
	return modeName
}
