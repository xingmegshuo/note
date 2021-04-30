/***************************
@File        : server.go
@Time        : 2021/04/08 11:56:38
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : tls Server
****************************/

package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type dataValue struct {
	Data   []string
	Result []bool
}

var value dataValue
var lock sync.RWMutex
var word = make(map[string]bool)

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "json")
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	} else {
		if err := json.Unmarshal([]byte(body), &value); err == nil {
			r := false
			result := []bool{}
			for _, s := range value.Data {
				lock.Lock()
				if _, ok := word[s]; ok {
					r = true
					lock.Unlock()
				} else {
					lock.Unlock()
					lock.Lock()
					word[s] = true
					lock.Unlock()
					r = false
				}
				result = append(result, r)
			}
			v := dataValue{
				Result: result,
			}
			jsonStr, e := json.Marshal(v)
			if e != nil {
				log.Println(e)
			}
			w.Write([]byte(jsonStr))

		} else {
			log.Println(err)
			jsonStr, e := json.Marshal(dataValue{})
			if e != nil {
				log.Println(e)
			}
			w.Write([]byte(jsonStr))
		}
	}
}

func Start() error {

	http.HandleFunc("/", handler)
	log.Println("监听 8000 端口成功，可以通过 https://127.0.0.1:8000/ 访问")
	err := http.ListenAndServeTLS(":8000", "server.crt", "server.key", nil)
	return err
}

func init() {
	go Start()
}
