package main

//Number ...
type Number struct {
	DID   string `json:"did"`
	ISOCC string `json:"isocc"`
}

//Response ...
type Response struct {
	MSG  string   `json:"msg"`
	CODE int      `json:"status"`
	DATA []Number `json:"data"`
}
