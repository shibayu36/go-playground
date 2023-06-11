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

func (lc LogCounter) GroupByUser() map[string][]Log {
	group := make(map[string][]Log)
	for _, log := range lc.Logs {
		user := log.User
		if user == "" {
			user = "guest"
		}
		group[user] = append(group[user], log)
	}
	return group
}
