package config

// Config for go-deepseek client.
//
//	ApiKey - go-deepseek API key.
//	TimeoutSeconds - http client timeout used by go-deepseek client.
//	DisableRequestValidation - disable request validation by go-deepseek client.
type Config struct {
	ApiKey                   string
	TimeoutSeconds           int
	DisableRequestValidation bool
}
