package types

type NewTaskPayload struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	Due      string `json:"due"`
	Priority int16  `json:"priority,omitempty"`
}

type ApiConfigFile struct {
	Key string `json:"key"`
}
