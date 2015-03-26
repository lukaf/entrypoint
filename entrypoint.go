package main

import (
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func getIp() string {
	hostname, _ := os.Hostname()
	ips, _ := net.LookupHost(hostname)
	if len(ips) > 0 {
		return ips[0]
	}
	return "127.0.0.1"
}

func parseArgs(arguments []string) []string {
	var parsed []string

	ip := getIp()

	for _, value := range arguments {
		parsed = append(
			parsed,
			strings.Replace(value, "%ip%", ip, -1),
		)
	}

	return parsed
}

func main() {
	var command string
	var arguments []string = []string{""}
	var err error

	log.SetFlags(0)

	switch len(os.Args) {

	case 1:
		log.Println("No arguments, nothing to do.")
		return

	case 2:
		command, err = exec.LookPath(os.Args[1])
		if err != nil {
			log.Println(err)
			return
		}

	default:
		command, err = exec.LookPath(os.Args[1])
		if err != nil {
			log.Println(err)
			return
		}
		arguments = parseArgs(os.Args[1:])
	}

	syscall.Exec(command, arguments, os.Environ())
}
