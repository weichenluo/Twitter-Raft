package util

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var muLock sync.Mutex

func RaftInit(ctxt context.Context, method string, url string, loading *strings.Reader, rChan chan *http.Response) {
	muLock.Lock()

	log.Println("infomation is shown below:")
	log.Println("url = ", url)
	log.Println("method = ", method)
	// log.Println("payload = ", loading)

	request, err := http.NewRequest(method, url, loading)

	if err != nil {
		log.Println("Can't create with current error: ", err)
		muLock.Unlock()
		return
	}

	select {
	case <-time.After(100 * time.Nanosecond):
		log.Println("Time exceeded on calling Raft server with : ", url)
	case <-ctxt.Done():
		log.Println("Error occured in context with url: ", url, " method: ", method, " error: ", ctxt.Err())
		muLock.Unlock()
		return
	}

	log.Println("Process called at: ", url)

	var response *http.Response
	response, err = http.DefaultClient.Do(request)

	if err != nil {
		log.Println("Error received from Raft: ", err)
		muLock.Unlock()
		return
	}

	log.Println("Response succesfully recieved: ", response)
	muLock.Unlock()
	rChan <- response

	log.Println("Response channal setup finished for: ", url)

}

func Raftstorage(method string, key string, value interface{}) (string, error) {
	var loadingVal string

	if method != "GET" {
		var buffer bytes.Buffer
		if err := gob.NewEncoder(&buffer).Encode(value); err != nil {
			log.Println("Error occured while encoding with key", key, " data =", err)
			return "", err
		}
		loadingVal = buffer.String()
	}

	rChan := make(chan *http.Response)

	ports := [3]string{"12380", "22380", "32380"}

	ctxt := context.Background()
	ctxt, cancel := context.WithTimeout(ctxt, 3*time.Second)
	defer cancel()

	for _, port := range ports {
		var payload *strings.Reader
		payload = nil
		if value != nil {
			payload = strings.NewReader(loadingVal)
		}
		// log.Println("Payload: ", payload)

		url := "http://127.0.0.1:" + port + "/" + key

		go RaftInit(ctxt, method, url, payload, rChan)
	}

	var response *http.Response
	select {
	case res := <-rChan:
		log.Println("response from channel is ", res)
		response = res
		cancel()
	case <-time.After(10 * time.Second):
		log.Println("Raft Server time out")
		return "", errors.New("Raft server time out!!")
	}

	content, err := io.ReadAll(response.Body)

	if err != nil {
		log.Println("Error occured while decoding response from Raft ", err)
		return "", err
	}

	log.Println("Data received from Raft after calling ", method, " with: ", string(content))

	return string(content), nil

}
