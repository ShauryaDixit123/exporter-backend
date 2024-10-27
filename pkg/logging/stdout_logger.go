package logging

import (
	"encoding/json"
	"fmt"
)

type stdOutLogger struct {
	appName     string
	appLogLevel LogLevel
}

func NewStdOutLogger(appName string, appLogLevel LogLevel) *stdOutLogger {
	return &stdOutLogger{
		appName:     appName,
		appLogLevel: appLogLevel,
	}
}

func (l *stdOutLogger) Info(
	statement string,
	description string,
	searchTerms map[string]any,
	references map[string]any,
) {

	if !CheckLoggingEligibility(INFO, l.appLogLevel) {
		return
	}

	log := CreateLog(
		l.appName,
		INFO,
		statement,
		description,
		nil,
		searchTerms,
		references,
	)

	logJson, err := json.Marshal(log)

	if err != nil {
		fmt.Println("unable to marshal json from struct")
		return
	}

	fmt.Println(string(logJson))
}

func (l *stdOutLogger) Warn(
	statement string,
	description string,
	logError *error,
	searchTerms map[string]any,
	references map[string]any,
) {

	if !CheckLoggingEligibility(WARN, l.appLogLevel) {
		return
	}

	log := CreateLog(
		l.appName,
		WARN,
		statement,
		description,
		logError,
		searchTerms,
		references,
	)

	logJson, err := json.Marshal(log)

	if err != nil {
		fmt.Println("unable to marshal json from struct")
		return
	}

	fmt.Println(string(logJson))
}

func (l *stdOutLogger) Error(
	statement string,
	description string,
	logError error,
	searchTerms map[string]any,
	references map[string]any,
) {

	if !CheckLoggingEligibility(ERROR, l.appLogLevel) {
		return
	}

	log := CreateLog(
		l.appName,
		ERROR,
		statement,
		description,
		&logError,
		searchTerms,
		references,
	)

	logJson, err := json.Marshal(log)

	if err != nil {
		fmt.Println("unable to marshal json from struct")
		return
	}

	fmt.Println(string(logJson))
}

func (l *stdOutLogger) Critical(
	statement string,
	description string,
	logError error,
	searchTerms map[string]any,
	references map[string]any,
) {

	if !CheckLoggingEligibility(CRITICAL, l.appLogLevel) {
		return
	}

	log := CreateLog(
		l.appName,
		CRITICAL,
		statement,
		description,
		&logError,
		searchTerms,
		references,
	)

	logJson, err := json.Marshal(log)

	if err != nil {
		fmt.Println("unable to marshal json from struct")
		return
	}

	fmt.Println(string(logJson))
}
