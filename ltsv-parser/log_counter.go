package main

type LogCounter struct {
	Logs []Log
}

func (lc LogCounter) CountError() int {
	var count int
	for _, log := range lc.Logs {
		if log.Status >= 500 {
			count++
		}
	}
	return count
}
