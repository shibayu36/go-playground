package main

type LogCounter struct {
	Logs []Log
}

func (lc LogCounter) CountError() int {
	return 0
}
