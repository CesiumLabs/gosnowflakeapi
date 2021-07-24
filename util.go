package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Normal expected unusual responses
var UnusualResponses []int = []int{400, 401, 404, 429, 500}

func isUnusualResponse(code int) bool {
	for _, i := range UnusualResponses {
		if i == code {
			return true
		}
	}

	return false
}

// Simply fetches from the snowflake api and returns a json parsed object!
func (self *Client) Fetch(method string, url string, structure interface{}) error {
	client := &http.Client{}

	req, err := http.NewRequest(method, "https://api.snowflakedev.org"+url, nil)
	if err != nil {
		return errors.New("UnexpectedError: Failed making a request")
	}

	req.Header.Add("Authorization", self.Token)

	res, err := client.Do(req)
	if err != nil {
		return errors.New("UnexpectedError: Failed making a request")
	}

	defer res.Body.Close()

	data, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return errors.New("UnexpectedError: Failed reading request data")
	}

	fmt.Println(string(data))

	unusual := isUnusualResponse(res.StatusCode)
	if unusual {
		return errors.New("SnowflakeApiError: Snowflake api sent an unusual api response as " + string(data) + " with status code as " + strconv.Itoa(res.StatusCode) + "!")
	}

	marshallErr := json.Unmarshal(data, structure)
	if marshallErr != nil {
		return errors.New("UnexpectedError: Failed while marshalling json: " + marshallErr.Error())
	}

	return nil
}

// Simply fetches from the snowflake api and returns binary!
func (self *Client) FetchBinary(method string, url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, "https://api.snowflakedev.org"+url, nil)
	if err != nil {
		return []byte{}, errors.New("UnexpectedError: Failed making a request")
	}

	req.Header.Add("Authorization", self.Token)

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("UnexpectedError: Failed making a request")
	}

	defer res.Body.Close()

	data, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return []byte{}, errors.New("UnexpectedError: Failed reading request data")
	}

	unusual := isUnusualResponse(res.StatusCode)
	if unusual {
		return []byte{}, errors.New("SnowflakeApiError: Snowflake api sent an unusual api response as " + string(data) + " with status code as " + strconv.Itoa(res.StatusCode) + "!")
	}

	return data, nil
}
