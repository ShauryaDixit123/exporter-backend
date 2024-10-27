package logging

import (
	mapset "github.com/deckarep/golang-set/v2"
)

var logLevelMap = map[LogLevel]mapset.Set[LogLevel]{
	INFO:     mapset.NewSet(INFO),
	WARN:     mapset.NewSet(INFO, WARN),
	ERROR:    mapset.NewSet(INFO, WARN, ERROR),
	CRITICAL: mapset.NewSet(INFO, WARN, ERROR, CRITICAL),
}

func CheckLoggingEligibility(
	logLevel LogLevel,
	appLogLevel LogLevel,
) bool {

	levels, found := logLevelMap[logLevel]

	if !found {
		return false
	} else {
		return levels.Contains(appLogLevel)
	}

}
