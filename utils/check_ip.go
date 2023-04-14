package utils

import (
	"errors"
	"net"
	"regexp"
	"strings"
)

func IsIPv4(ipAddr string) (string, error) {
	ip := net.ParseIP(ipAddr)
	if ip != nil && strings.Contains(ipAddr, ".") {
		return ipAddr, nil
	}
	return ipAddr, errors.New("IP address not valid")
}

func IsPort(input string) (string, error) {
	pattern := "\\d+"
	res, _ := regexp.MatchString(pattern, input)
	if res {
		return input, nil
	}
	return input, errors.New("Port number not valid")
}
