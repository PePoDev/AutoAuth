package auth

import (
	"time"

	"github.com/pepodev/xlog"
	"github.com/spf13/viper"
)

// LoadPresetFromPath will load preset from path and return it
func LoadPresetFromPath(dir string, fileName string) AutoAuthPreset {
	viper.SetConfigFile(fileName)
	viper.AddConfigPath(dir)
	if err := viper.ReadInConfig(); err != nil {
		xlog.Fatalf("fatal error config file: %s \n", err)
	}

	basePreset := BaseAutoAuthPreset{}
	if err := viper.Unmarshal(&basePreset); err != nil {
		xlog.Fatalf("%v", err)
	}

	return basePreset.AutoAuth
}

// BaseAutoAuthPreset is used to map to preset file like yml or json file
type BaseAutoAuthPreset struct {
	AutoAuth AutoAuthPreset `mapstructure:"autoauth"`
}

// AutoAuthPreset is base struct contain all configuration of preset file
type AutoAuthPreset struct {
	Encrypted bool
	Login     AutoAuthLogin
	Logout    AutoAuthLogout
	Heartbeat AutoAuthHeartbeat
	Save      []string
}

// AutoAuthLogin contain login preset
type AutoAuthLogin struct {
	Endpoint string
	Method   string
	Encoding string
	Header   []string
	Body     []string
}

// AutoAuthLogout contain logout preset
type AutoAuthLogout struct {
	Endpoint string
	Method   string
	Encoding string
	Header   []string
	Body     []string
}

// AutoAuthHeartbeat contain heartbeat preset
type AutoAuthHeartbeat struct {
	Endpoint string
	Method   string
	Encoding string
	Header   []string
	Body     []string
	Interval time.Duration
	Timeout  time.Duration
	Retry    int
}