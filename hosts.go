/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

type reqAction struct {
		action   string `json:"action"`
		hostname string `json:"hostname"`
}

func (self *Client) MuteHost(hostname string) ([]string, error) {
	var out reqAction
	err := self.doJsonRequest("POST", "/v1/search?q=hosts:"+search, nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
