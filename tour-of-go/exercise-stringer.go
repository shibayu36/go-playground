package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (i IPAddr) String() string {
	strs := make([]string, len(i))

	for i, val := range i {
		strs[i] = fmt.Sprintf("%v", val)
	}

	return strings.Join(strs, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
