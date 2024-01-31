package config

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// RR return result
	RR = ReturnResult{}
)

// Result result
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ReturnResult return result model
type ReturnResult struct {
	Success          Result `mapstructure:"success"`
	BadRequest       Result `mapstructure:"bad_request"`
	ConnectionError  Result `mapstructure:"connection_error"`
	DatabaseNotFound Result `mapstructure:"database_not_found"`
}

// Error error
func (rs Result) Error() string {
	return rs.Message
}

// NewMessageError new message error
func NewMessageError(code int, message string) *Result {
	if code == http.StatusOK {
		code = http.StatusInternalServerError
	}
	return &Result{
		Code:    code,
		Message: message,
	}
}

// InitReturnResult init return result
func InitReturnResult(configPath string) error {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName("result")
	v.SetConfigType("yml")

	if err := v.ReadInConfig(); err != nil {
		logrus.Error("read config file error:", err)
		return err
	}

	if err := v.Unmarshal(&RR); err != nil {
		logrus.Error("unmarshal config error:", err)
		return err
	}

	return nil
}
