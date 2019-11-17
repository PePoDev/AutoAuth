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
	AutoAuthData
	Login     AutoAuthLogin
	Logout    AutoAuthLogout
	Heartbeat AutoAuthHeartbeat
}

// AutoAuthData contains the data to use in main instance
type AutoAuthData struct {
	Name      string
	Encrypted bool

	Save []string

	IsRunning bool
	Try       int
}

// AutoAuthLogin contain login preset
type AutoAuthLogin struct {
	AutoAuthRequestSetting
}

// AutoAuthLogout contain logout preset
type AutoAuthLogout struct {
	AutoAuthRequestSetting
}

// AutoAuthHeartbeat contain heartbeat preset
type AutoAuthHeartbeat struct {
	AutoAuthRequestSetting
	Interval time.Duration
	Timeout  time.Duration
	Retry    int
}

// AutoAuthRequestSetting is struct to contain data for send request
type AutoAuthRequestSetting struct {
	Endpoint string
	Method   string
	Header   []string
	Body     []string
}
