package config

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/burntsushi/toml"
)

// Config is toml
type Config struct {
	Path        string
	LogFile     string
	ServiceName string
	WebService  struct {
		Port    int
		Version string
	}
	Route struct {
		PAGE   string
		HELP   string
		UPLOAD string
	}
	ServiceIP struct {
		Prefix   string
		PRIVACY  string
		SIF2JSON string
		CSV2JSON string
	}
}

var (
	mux sync.Mutex
)

// newCfg :
func newCfg(configs ...string) *Config {
	defer func() {
		mux.Unlock()
	}()
	mux.Lock()
	for _, f := range configs {
		if _, e := os.Stat(f); e == nil {
			return (&Config{Path: f}).set()
		}
	}
	return nil
}

// set is
func (cfg *Config) set() *Config {
	f := cfg.Path /* make a copy of original for restoring */
	if _, e := toml.DecodeFile(f, cfg); e == nil {
		// modify some to save
		cfg.Path = f
		if abs, e := filepath.Abs(f); e == nil {
			cfg.Path = abs
		}
		if ver, e := gitver(); e == nil && ver != "" { /* successfully got git ver */
			cfg.WebService.Version = ver
		}
		// save
		cfg.save()

		return cfgRepl(cfg, map[string]interface{}{
			"[DATE]": time.Now().Format("2006-01-02"),
			"[v]":    cfg.WebService.Version,
		}).(*Config)
	}
	return nil
}

func (cfg *Config) save() {
	if f, e := os.OpenFile(cfg.Path, os.O_WRONLY|os.O_TRUNC, os.ModePerm); e == nil {
		defer f.Close()
		toml.NewEncoder(f).Encode(cfg)
	}
}

// InitEnvVarFromTOML : initialize the global variables
func InitEnvVarFromTOML(key string, configs ...string) bool {
	Cfg := newCfg(append(configs, "./config.toml", "./config/config.toml")...)
	if Cfg == nil {
		return false
	}
	struct2Env(key, Cfg)
	return true
}
