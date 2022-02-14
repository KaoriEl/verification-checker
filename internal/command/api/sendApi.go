package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/internal/structures"
	"net/http"
	"strconv"
)

func Status(text string, i structures.CoinlistAccs) {
	reqBody, err := json.Marshal(map[string]string{
		"text":           text,
		"cid":            strconv.FormatInt(i.CID, 10),
		"need_send_file": "false",
		"type_message":   "default_message",
	})

	if err != nil {
		print(err)
	}

	resp, err := http.Post("http://investments-go-bot:3000/api/v1/status-captcha",
		"application/x-www-form-urlencoded", bytes.NewBuffer(reqBody))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))

}

func ChangeVerificationStatus(i structures.CoinlistAccs) {
	reqBody, err := json.Marshal(map[string]string{
		"id":     strconv.FormatInt(int64(i.Id), 10),
		"status": "yes",
	})
	if err != nil {
		print(err)
	}

	resp, err := http.Post("http://investments-api-ms-nginx/api/v1/coinlist/change/status",
		"application/x-www-form-urlencoded", bytes.NewBuffer(reqBody))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))

}
func FukVerificateStatus(i structures.CoinlistAccs) {
	reqBody, err := json.Marshal(map[string]string{
		"id":     strconv.FormatInt(int64(i.Id), 10),
		"status": "Later",
		"cid":    strconv.FormatInt(i.CID, 10),
	})
	if err != nil {
		print(err)
	}

	resp, err := http.Post("http://investments-api-ms-nginx/api/v1/coinlist/send/fukVerification",
		"application/x-www-form-urlencoded", bytes.NewBuffer(reqBody))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))

}
