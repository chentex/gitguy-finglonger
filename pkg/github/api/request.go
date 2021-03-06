package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	token string
)

func request(method string, url string, payload interface{}) (int, []byte) {
	nb, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(nb))
	t := fmt.Sprintf(`token %s`, token)
	req.Header.Add("Accept", "application/vnd.github.inertia-preview+json")
	req.Header.Add("Authorization", t)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("request error: %s - %d - %s", resp.Status, resp.StatusCode, err)
		log.Fatalln(err)
		return resp.StatusCode, nil
	}
	bd, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read respose Body: %s", err)
		return http.StatusInternalServerError, nil
	}
	log.Printf("request %s status: %s - %d", url, resp.Status, resp.StatusCode)
	return resp.StatusCode, bd
}
