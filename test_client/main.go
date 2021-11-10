package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tal-tech/go-zero/core/jsonx"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

type UserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	client := http.Client{}
	t := login(client)
	upload(client, t)
	download(client, t)
}

func login(client http.Client) string {
	loginUrl := "http://127.0.0.1:8888/user/login"
	bs, err := jsonx.Marshal(&UserReq{
		Username: "fenghai",
		Password: "aca9ec30e0f3e40553580d01ab729de63cc30dc18e832a9f70c0062dc97bd44f",
	})
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", loginUrl, bytes.NewReader(bs))
	if err != nil {
		panic("login: " + err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	d, err := ioutil.ReadAll(resp.Body)
	m := make(map[string]interface{})
	err = json.Unmarshal(d, &m)
	if err != nil {
		panic(err)
	}
	return m["token"].(string)
}

func upload(client http.Client, t string) {
	uploadUrl := "http://127.0.0.1:8888/video/upload/1/5"
	f, err := os.Open("./test_client/api-tag.PNG")
	if err != nil {
		panic(err)
	}
	r, w := io.Pipe()
	req, err := http.NewRequest("POST", uploadUrl, r)
	if err != nil {
		panic(err)
	}
	req.Header.Set("authorization", t)
	go func() {
		buf := make([]byte, 1024)
		var n int
		n, err = f.Read(buf)
		if err != nil {
			panic(err)
		}
		for _, err = w.Write(buf[0:n]); err == nil; _, err = w.Write(buf[0:n]) {
			time.Sleep(time.Millisecond * 75)
			n, err = f.Read(buf)
			if err == io.EOF {
				fmt.Println(n)
				break
			}
			if err != nil {
				panic(err)
			}
		}
		err = w.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println(client.Do(req))
}

func download(client http.Client, t string) {
	uploadUrl := "http://127.0.0.1:8888/video/download/1/5"
	req, err := http.NewRequest("POST", uploadUrl, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("authorization", t)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println(resp.Status, err)
	f, err := os.OpenFile("./test_client/desFile.png", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	bufferR := httputil.NewChunkedReader(resp.Body)
	bufferW := bufio.NewWriter(f)
	buf := make([]byte, 1024 * 1024)
	var n int
	for n, err = bufferR.Read(buf); err == nil; n, err = bufferR.Read(buf) {
		_, err = bufferW.Write(buf[0:n])
		if err != nil {
			break
		}

		err = bufferW.Flush()
		if err != nil {
			break
		}

	}

	if err != nil && err != io.EOF && err.Error() != "unexpected EOF" {
		panic(err)
	}

}