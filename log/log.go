/*
 * @Description: 备注
 * @Author: wangcong0918
 * @Date: 2019-11-29 18:08:04
 * @LastEditTime: 2020-01-13 10:31:39
 * @LastEditors: wangcong0918
 */
package log

import (
	"os"

	"github.com/op/go-logging"
)

var (
	Logger = logging.MustGetLogger("sunrise_frame")
	format = logging.MustStringFormatter(
		"%{color} %{level:.4s} %{id:03x} %{time:2006-01-02 15:04:05.000} %{shortfile}\t%{shortfunc}\t> %{message}%{color:reset}",
	)
)

func init() {
	backend2 := logging.NewLogBackend(os.Stdout, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	logging.SetBackend(backend2Formatter)
}
