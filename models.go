package main

//Number ...
type Number struct {
	DID   string `json:"did"`
	ISOCC string `json:"isocc"`
}

//Response ...
type Response struct {
	msg    string   `json:"msg"`
	status string   `json:"status"`
	data   []Number `json:"data"`
}
