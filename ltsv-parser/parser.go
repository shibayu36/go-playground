package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Parser struct {
	filename string
}

func (p Parser) Parse() ([]Log, error) {
	file, err := os.Open(p.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var logs []Log
	for scanner.Scan() {
		log, err := p.lineToLog(scanner.Text())
		if err != nil {
			// Ignore invalid line
			continue
		}
		logs = append(logs, *log)
	}

	return logs, nil
}

func (p Parser) lineToLog(line string) (*Log, error) {
	data := make(map[string]string)

	for _, field := range strings.Split(line, "\t") {
		kv := strings.SplitN(field, ":", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid field: %s", field)
		}
		if kv[1] == "-" {
			kv[1] = ""
		}
		data[kv[0]] = kv[1]
	}

	log := &Log{}
	var err error

	log.Host = data["host"]
	log.User = data["user"]
	log.Epoch, err = strconv.Atoi(data["epoch"])
	if err != nil {
		return nil, err
	}
	log.Req = data["req"]
	log.Status, err = strconv.Atoi(data["status"])
	if err != nil {
		return nil, err
	}
	log.Size, err = strconv.Atoi(data["size"])
	if err != nil {
		return nil, err
	}
	log.Referer = data["referer"]

	return log, nil
}
