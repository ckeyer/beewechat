package global

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

//  返回Get数据
func HttpGet(url string) (content string, statusCode int) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statusCode = -200
		return
	}
	statusCode = resp.StatusCode
	content = string(data)
	return
}

//  返回Get数据
func HttpGetToBytes(url string) (content []byte, statusCode int) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statusCode = -200
		return
	}
	statusCode = resp.StatusCode
	content = data
	return
}

// 返回Post数据
func HttpPost(url string, post_data string) (content string, statusCode int) {
	client := &http.Client{}
	postBytesReader := bytes.NewReader([]byte(post_data))
	resp, err1 := client.Post(url, "posttext", postBytesReader)
	// resp, err1 := http.NewRequest("POST", url, postBytesReader)
	// resp, err1 := http.Post(url, "bodyType", body)
	if err1 != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statusCode = -200
		return
	}
	content = string(data)
	return
}
