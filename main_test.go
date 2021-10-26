package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	go main()
}

var (
	client = &http.Client{
		Timeout: 1 * time.Second,
	}
	u = []byte(`{
		"userID":   1,
		"name": "Bill"
	}`)
)

func TestEmptyGetRes(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://localhost:4000/user", nil)

	w, _ := client.Do(r)
	assert.Equal(t, http.StatusInternalServerError, w.StatusCode)
	body, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, []byte("no user data"), body)
}

func TestValidPostReq(t *testing.T) {
	r, _ := http.NewRequest("POST", "http://localhost:4000/user", bytes.NewBuffer(u))

	w, _ := client.Do(r)
	assert.Equal(t, http.StatusCreated, w.StatusCode)
}

func TestInvalidPostReq(t *testing.T) {
	r, _ := http.NewRequest("POST", "http://localhost:4000/user", bytes.NewBuffer(u))

	w, _ := client.Do(r)
	assert.Equal(t, http.StatusInternalServerError, w.StatusCode)
	body, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, []byte("user already exists with ID 1"), body)
}

func TestValidGetRes(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://localhost:4000/user", nil)

	w, _ := client.Do(r)
	assert.Equal(t, http.StatusOK, w.StatusCode)
	body, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, []byte(`[{"userID":1,"name":"Bill"}]`), body[:len(body)-1])
}
