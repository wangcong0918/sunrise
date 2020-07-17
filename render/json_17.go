/*
 * @Description: 备注
 * @Author: wangcong0918
 * @Date: 2019-11-29 18:08:04
 * @LastEditTime: 2020-07-16 15:39:20
 * @LastEditors: wangcong0918
 */
// Copyright 2018 Gin Core Team.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build go1.7

package render

import (
	"net/http"

	"github.com/wangcong0918/sunrise/internal/json"
)

// PureJSON contains the given interface object.
type PureJSON struct {
	Data interface{}
}

// Render (PureJSON) writes custom ContentType and encodes the given interface object.
func (r PureJSON) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	return encoder.Encode(r.Data)
}

// WriteContentType (PureJSON) writes custom ContentType.
func (r PureJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}
