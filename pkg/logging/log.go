package logging

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"time"
)

type LogLevel string

const (
	INFO     LogLevel = "CONSOLE"
	WARN     LogLevel = "WARN"
	ERROR    LogLevel = "ERROR"
	CRITICAL LogLevel = "CRITICAL"
)

type Log struct {
	AppName     string
	LogLevel    LogLevel
	Statment    string
	Description string
	SearchTerms string
	References  map[string]string
	IpAddress   string
	Hostname    string
	Timestamp   string
}

func CreateLog(
	appName string,
	LogLevel LogLevel,
	statment string,
	description string,
	logError *error,
	searchTerms map[string]any,
	references map[string]any,
) Log {

	searchTermsKeys := make([]string, 0, len(searchTerms))

	for k, v := range searchTerms {
		searchTermsKeys = append(searchTermsKeys, fmt.Sprintf("%s=%v", k, v))
	}

	terms := strings.Join(searchTermsKeys, "|")

	refs := map[string]string{}

	if logError != nil {
		references["errMessage"] = *logError
		b := make([]byte, 2048) // adjust buffer size to be larger than expected stack
		n := runtime.Stack(b, false)
		references["stackTrace"] = string(b[:n])
	}

	for k, v := range references {
		refs[k] = fmt.Sprintf("%v", v)
	}

	hostname, err := os.Hostname()

	if err != nil {
		fmt.Println("Unable to fetch Machine Hostname")
	}

	ip, err := net.LookupIP(hostname)

	if err != nil {
		fmt.Println("Unable to fetch IP address from Hostname: ", hostname)
	}

	return Log{
		AppName:     appName,
		LogLevel:    LogLevel,
		Statment:    statment,
		Description: description,
		SearchTerms: terms,
		References:  refs,
		IpAddress:   fmt.Sprintf("%v", ip),
		Hostname:    hostname,
		Timestamp:   time.Now().UTC().Format(time.RFC3339Nano),
	}

}
