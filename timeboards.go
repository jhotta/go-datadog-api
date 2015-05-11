/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

import (
	"fmt"
)

// Graph represents a graph that might exist on a Timeboard.
type TimeGraph struct {
	Title      string     `json:"title"`
	Events     []struct{} `json:"events"`
	Definition struct {
		Viz      string `json:"viz"`
		Requests []struct {
			Query   string `json:"q"`
			Stacked bool   `json:"stacked"`
		} `json:"requests"`
	} `json:"definition"`
}

// Timeboard represents a user created Timeboard. This is the full Timeboard
// struct when we load a Timeboard in detail.
type Timeboard struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Title       string  `json:"title"`
	Graphs      []TimeGraph `json:"graphs"`
}

// TimeboardLite represents a user created Timeboard. This is the mini
// struct when we load the summaries.
type TimeboardLite struct {
	Id          int    `json:"id,string"` // TODO: Remove ',string'.
	Resource    string `json:"resource"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

// reqGetTimeboards from /api/v1/dash
type reqGetTimeboards struct {
	Timeboards []TimeboardLite `json:"dashes"`
}

// reqGetTimeboard from /api/v1/dash/:Timeboard_id
type reqGetTimeboard struct {
	Resource  string    `json:"resource"`
	Url       string    `json:"url"`
	Timeboard Timeboard `json:"dash"`
}

// GetTimeboard returns a single Timeboard created on this account.
func (self *Client) GetTimeboard(id int) (*Timeboard, error) {
	var out reqGetTimeboard
	err := self.doJsonRequest("GET", fmt.Sprintf("/v1/dash/%d", id), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out.Timeboard, nil
}

// GetTimeboards returns a list of all Timeboards created on this account.
func (self *Client) GetTimeboards() ([]TimeboardLite, error) {
	var out reqGetTimeboards
	err := self.doJsonRequest("GET", "/v1/dash", nil, &out)
	if err != nil {
		return nil, err
	}
	return out.Timeboards, nil
}

// DeleteTimeboard deletes a Timeboard by the identifier.
func (self *Client) DeleteTimeboard(id int) error {
	return self.doJsonRequest("DELETE", fmt.Sprintf("/v1/dash/%d", id), nil, nil)
}

// CreateTimeboard creates a new Timeboard when given a Timeboard struct. Note
// that the Id, Resource, Url and similar elements are not used in creation.
func (self *Client) CreateTimeboard(dash *Timeboard) (*Timeboard, error) {
	var out reqGetTimeboard
	err := self.doJsonRequest("POST", "/v1/dash", dash, &out)
	if err != nil {
		return nil, err
	}
	return &out.Timeboard, nil
}

// UpdateTimeboard in essence takes a Timeboard struct and persists it back to
// the server. Use this if you've updated your local and need to push it back.
func (self *Client) UpdateTimeboard(dash *Timeboard) error {
	return self.doJsonRequest("PUT", fmt.Sprintf("/v1/dash/%d", dash.Id),
		dash, nil)
}
