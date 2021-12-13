package models

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"output,omitempty"`
}

type ExecuteRequest struct {
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

type DeleteRequest struct {
	Filename string `json:"filename"`
}

var CmdList = []string{"ls", "rm", "wc", "cat", "tr", "sort", "uniq", "tail", "head", "cut"}

var CmdTemplate = map[string]string{
	"wc":         "wc -w {{.folder}}* | tail -1 | cut -d ' ' -f1",
	"freq-words": "cat {{.folder}}* | tr -s ' ' '\\n' | sort | uniq -c | sort -n | {{.order}} -n {{.count}}",
}
