package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tal-tech/go-zero/core/jsonx"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type UserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	uploadUrl := "http://127.0.0.1:8888/video/upload/1/5"
	client := http.Client{}
	t := login(client)
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
		c := 0
		for _, err = w.Write(buf[0:n]); err == nil; _, err = w.Write(buf[0:n]) {
			c++
			fmt.Println(c, n)
			time.Sleep(time.Millisecond * 50)
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