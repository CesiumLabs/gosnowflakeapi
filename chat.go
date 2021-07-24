package main

import (
	"fmt"
	"net/url"
)

// The object for the [Client.Chat].
type ChatOptions struct {
	// The name of the bot to be set as.
	Name string
	// The gender of the bot to be set as.
	Gender string
	// The user id which is always default to 1.
	User uint32
	// The message the user sent.
	Message string
}

func (self ChatOptions) toQuery() string {
	return fmt.Sprintf("?message=%s&name=%s&gender=%s&user=%d", url.QueryEscape(self.Message), url.QueryEscape(self.Name), self.Gender, self.User)
}

// The chat response structure of the api.
type ChatResponse struct {
	// The message data sent by the chat api.
	Message string `json:"message"`
}

// Get the ai chat responses through options.
func (self *Client) Chat(options ChatOptions) (ChatResponse, error) {
	var response ChatResponse
	err := self.Fetch("GET", "/api/chatbot"+options.toQuery(), &response)
	return response, err
}

// Get the ai chat responses through all of the required options.
func (self *Client) ChatWith(message string, name string, gender string, user uint32) (ChatResponse, error) {
	var response ChatResponse
	err := self.Fetch("GET", fmt.Sprintf("/api/chatbot?message=%s&name=%s&gender=%s&user=%d", url.QueryEscape(message), url.QueryEscape(name), gender, user), &response)
	return response, err
}
