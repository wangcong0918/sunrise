package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/wangcong0918/sunrise/log"
)

func Post(url string, contentType string, body io.Reader) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
			log.Logger.Error("http request failed --> ", url)
			return
		}
	}()
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		log.Logger.Error("http response --> ", err.Error())
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Error("read http respBody --> ", err.Error())
		return "", err
	}
	return string(respBody), nil
}

// 返回请求的状态码
func GetResponseStatusByPost(url string, contentType string, body io.Reader) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
			log.Logger.Error("http request failed --> ", url)
			return
		}
	}()
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		log.Logger.Error("http response --> ", err.Error())
		return ""
	}
	defer resp.Body.Close()

	return resp.Status
}

func Get(url string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
			log.Logger.Error("http request failed --> ", url)
			return
		}
	}()

	resp, err := http.Get(url)
	if err != nil {
		log.Logger.Error("http response --> ", err.Error())
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Error("read http respBody --> ", err.Error())
		return "", err
	}
	return string(respBody), nil
}

func RequestByHeader(url string, method string, headerMap map[string]string, body io.Reader) (string, error) {
	var CLIENT http.Client
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}

	for k, v := range headerMap {
		req.Header.Set(k, v)
	}

	resp, err := CLIENT.Do(req)

	if err != nil {
		log.Logger.Error("http response --> ", err.Error())
		return "", err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Error("read http respBody --> ", err.Error())
		return "", err
	}

	return string(respBody), nil
}

func OtherMethod(method string, contentType string, url string, body io.Reader, timeOut time.Duration) (string, error) {
	var CLIENT http.Client
	// 设置超时时间
	CLIENT.Timeout = timeOut
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}

	if contentType == "" {
		req.Header.Del("Content-Type")
	} else {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := CLIENT.Do(req)

	if err != nil {
		log.Logger.Error("http response --> ", err.Error())
		return "", err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Error("read http respBody --> ", err.Error())
		return "", err
	}

	return string(respBody), nil

}

// url.Values{"key": {"Value"}, "id": {"123"}}
func PostForm(url string, data url.Values) (string, error) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		log.Logger.Error("http response --> ", err.Error())
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Error("read http respBody --> ", err.Error())
		return "", err
	}
	return string(respBody), nil
}

func RequestByHeaderAndTime(url string, method string, headerMap map[string]string, body io.Reader, timeOut time.Duration) (string, error) {
	var CLIENT http.Client
	// 设置超时时间
	CLIENT.Timeout = timeOut
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}

	for k, v := range headerMap {
		req.Header.Set(k, v)
	}

	resp, err := CLIENT.Do(req)

	if err != nil {
		log.Logger.Error("http response --> ", err.Error())
		return "", err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Error("read http respBody --> ", err.Error())
		return "", err
	}

	return string(respBody), nil
}
