// getComic project main.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	. "github.com/soekchl/myUtils"
)

const (
	downloadUrl = "https://res.nbhbzl.com/"
)

func main() {
	for k, v := range reqList {
		getOne(k+1, v)
		time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	}
}

// 获取一话完整图片
func getOne(m int, urlReq string) {
	headers := make(map[string]string)
	headers["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3"
	headers["Cookie"] = "UM_distinctid=16be3fa014c951-0d0f815ec8a01a-1a201708-1fa400-16be3fa014dbfe; CNZZDATA1273814033=1260435515-1562896174-https%253A%252F%252Fwww.baidu.com%252F%7C1562896174"
	headers["User-Agent"] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"
	code, body, err := httpReq(urlReq, "get", nil, make(url.Values)) // 请求主网址内容
	if err != nil {
		Error(err)
		return
	}
	if code != 200 {
		Errorf("code=%v err=%s", code, body)
		return
	}
	list := getImages(body)	// 从主网址中分离目标图表
	Notice(m, "  ", len(list))
	dir := fmt.Sprint("./comic/", m, "/")
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		Error(err)
		return
	}
	n := 1
	for _, v := range list { // 一个一个下载
		err = downloadImages(downloadUrl+v, fmt.Sprintf("%vIMGE_%v.jpg", dir, n))
		if err != nil {
			Errorf("%v %v", n, err)
		}
		n++
	}
}

// get请求内容 下载到制定文件中
func downloadImages(reqUrl, fileName string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	fi, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer fi.Close()
	_, err = io.Copy(fi, response.Body)
	return err
}

// 从读取出来的 html中查找 chapterImages = [] 内容
func getImages(body string) (imgList []string) {
	n := strings.Index(body, "chapterImages")
	start := strings.Index(body[n+1:], "[")
	stop := strings.Index(body[n+1:], "]")
	jsonStr := body[start+n : n+stop+2]

	err := json.Unmarshal([]byte(jsonStr), &imgList)
	if err != nil {
		Error(err)
	}
	return
}

// http 请求通用
func httpReq(reqUrl, method string, headers map[string]string, value url.Values) (code int, body string, err error) {
	method = strings.ToUpper(method)
	client := &http.Client{}
	val := value.Encode()
	if method == "GET" {
		reqUrl = fmt.Sprintf("%s?%v", reqUrl, val)
		val = ""
	}
	req, err := http.NewRequest(method, reqUrl, strings.NewReader(val))
	if err != nil {
		return 0, "", err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	response, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, "", err
	}
	code = response.StatusCode
	body = string(b)
	return
}
