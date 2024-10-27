package logging

type Logger interface {
	Info(
		statement string,
		description string,
		searchTerms map[string]any,
		references map[string]any,
	)

	Warn(
		statement string,
		description string,
		logError *error,
		searchTerms map[string]any,
		references map[string]any,
	)

	Error(
		statement string,
		description string,
		logError error,
		searchTerms map[string]any,
		references map[string]any,
	)

	Critical(
		statement string,
		description string,
		logError error,
		searchTerms map[string]any,
		references map[string]any,
	)
}
