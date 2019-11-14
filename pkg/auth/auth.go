package auth

import (
	"time"

	"github.com/pepodev/autoauth/internal/message"
	"github.com/pepodev/autoauth/internal/utils"

	"github.com/pepodev/xlog"
)

// StartAutoLogin will start corutine to detect internet connection and send login request.
func (preset AutoAuthPreset) StartAutoLogin() {
	xlog.Infof("\n%s", message.GetWelcome())
	xlog.Info("AutoAuth Started")

	preset.IsRunning = true

	go func() {
		for preset.IsRunning {
			if !preset.IsHeatbeatAlive() {
				status := preset.RequestLogin()
				if !status {
					xlog.Info("Login fail")
					preset.Try++
				}
			}
			time.Sleep(time.Second * preset.Heartbeat.Interval)
		}
		xlog.Info("corutine has stopped by user")
	}()
}

// StopAutoLogin will stop AutoAuth
func (preset AutoAuthPreset) StopAutoLogin() {
	if !preset.IsRunning {
		xlog.Info("AutoAuth is not started yet")
		return
	}
	preset.IsRunning = false
}

// RequestLogin will create request to authentication service
func (preset *AutoAuthPreset) RequestLogin() bool {
	err := utils.Do(preset.Login.Endpoint,
		preset.Login.Method,
		preset.Login.Header,
		preset.Login.Body)

	if err != nil {
		xlog.Errorf("Login to %s is Error", preset.Login.Endpoint)
		return false
	}
	xlog.Infof("Login to %s is OK", preset.Login.Endpoint)
	return true
}

// RequestLogout send logout request
func (preset *AutoAuthPreset) RequestLogout() bool {
	err := utils.Do(preset.Logout.Endpoint,
		preset.Logout.Method,
		preset.Logout.Header,
		preset.Logout.Body)

	if err != nil {
		xlog.Errorf("Logout to %s is Error", preset.Logout.Endpoint)
		return false
	}
	xlog.Infof("Logout to %s is OK", preset.Logout.Endpoint)
	return true
}

// IsHeatbeatAlive send request to heartbeat endpoint and return status of request
func (preset *AutoAuthPreset) IsHeatbeatAlive() bool {
	err := utils.Do(preset.Heartbeat.Endpoint,
		preset.Heartbeat.Method,
		preset.Heartbeat.Header,
		preset.Heartbeat.Body)

	if err != nil {
		xlog.Errorf("Heartbeat to %s is Error", preset.Heartbeat.Endpoint)
		return false
	}
	xlog.Infof("Heartbeat to %s is OK", preset.Heartbeat.Endpoint)
	return true
}
