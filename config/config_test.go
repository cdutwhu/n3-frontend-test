package config

import "testing"

func TestLoad(t *testing.T) {
	cfg := NewCfg("./config.toml")
	fPln(cfg.Path)
	fPln(cfg.LogFile)
	fPln(cfg.ServiceName)
	fPln(cfg.WebService)
}