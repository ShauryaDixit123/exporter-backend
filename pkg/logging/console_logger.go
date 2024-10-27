package logging

import (
	"github.com/sanity-io/litter"
)

type consoleLogger struct {
	appName     string
	appLogLevel LogLevel
}

func NewConsoleLogger(appName string, appLogLevel LogLevel) *consoleLogger {
	return &consoleLogger{
		appName:     appName,
		appLogLevel: appLogLevel,
	}
}

func (l *consoleLogger) Info(
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

	litter.Dump(log)
}

func (l *consoleLogger) Warn(
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

	litter.Dump(log)
}

func (l *consoleLogger) Error(
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

	litter.Dump(log)
}

func (l *consoleLogger) Critical(
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

	litter.Dump(log)
}
