package auth

import (
	"time"

	"github.com/pepodev/autoauth/internal/message"
	"github.com/pepodev/autoauth/internal/utils"
	"github.com/valyala/fasthttp"

	"github.com/pepodev/xlog"
)

// StartAutoLogin will start corutine to detect internet connection and send login request.
func (preset AutoAuthPreset) StartAutoLogin() {
	xlog.Infof("\n%s", message.GetWelcome())
	xlog.Info("AutoAuth Started")

	preset.IsRunning = true

	go func() {
		for preset.IsRunning {
			time.Sleep(time.Second * preset.Heartbeat.Interval)

			err := preset.IsHeatbeatAlive()
			if err == nil {
				continue
			}

			err = preset.RequestLogin()
			if err != nil {
				xlog.Errorf("Login fail with err: %v", err)
				preset.Try++
			}
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
func (preset *AutoAuthPreset) RequestLogin() error {
	resp, err := utils.Do(preset.Login.Endpoint,
		preset.Login.Method,
		preset.Login.Header,
		preset.Login.Body,
		preset.Login.Timeout)
	defer fasthttp.ReleaseResponse(resp)

	if err != nil {
		xlog.Errorf("Login to %s is Error", preset.Login.Endpoint)
		return err
	}
	xlog.Infof("Login to %s is OK", preset.Login.Endpoint)
	return nil
}

// RequestLogout send logout request
func (preset *AutoAuthPreset) RequestLogout() error {
	resp, err := utils.Do(preset.Logout.Endpoint,
		preset.Logout.Method,
		preset.Logout.Header,
		preset.Logout.Body,
		preset.Logout.Timeout)
	defer fasthttp.ReleaseResponse(resp)

	if err != nil {
		xlog.Errorf("Logout to %s is Error", preset.Logout.Endpoint)
		return err
	}
	xlog.Infof("Logout to %s is OK", preset.Logout.Endpoint)
	return nil
}

// IsHeatbeatAlive send request to heartbeat endpoint and return status of request
func (preset *AutoAuthPreset) IsHeatbeatAlive() error {
	resp, err := utils.Do(preset.Heartbeat.Endpoint,
		preset.Heartbeat.Method,
		preset.Heartbeat.Header,
		preset.Heartbeat.Body,
		preset.Heartbeat.Timeout)
	defer fasthttp.ReleaseResponse(resp)

	if err != nil || resp.StatusCode() == 302 {
		xlog.Errorf("Heartbeat to %s is Error", preset.Heartbeat.Endpoint)
		return err
	}
	xlog.Infof("Heartbeat to %s is OK", preset.Heartbeat.Endpoint)
	return nil
}
