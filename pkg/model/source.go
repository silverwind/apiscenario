package model

type Source int

//go:generate enumer -type=Source -json -linecomment  -output source_generated.go
const (
	ResponseStatus Source = iota //response_status
	ResponseTime                 //response_time
	ResponseJson                 //response_json
)
