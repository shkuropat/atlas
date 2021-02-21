// Copyright 2020 The Atlas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"runtime"
	"strings"

	hd "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	conf "github.com/spf13/viper"
)

// configFile defines what app's config file should be used. In case it is specified, bootstrapConfig is not used
var configFile string

// SetConfigFile sets configFile
func SetConfigFile(file string) {
	configFile = file
}

const (
	// What PathSet is to be specified
	root = "root"
	home = "home"

	// These must be the same as runtime.GOOS values, because they are used in context of runtime.GOOS
	windows = "windows"
	linux   = "linux"
)

// PathSet specifies set of paths to be used to find configuration.
// Ex.:
// 	pathSet["root"] = /etc, /etc/somedir
//	pathSet["home"] = home-relative dirs, ex.: .atlas-config, .atlas-client
type PathSet map[string][]string

// PathOptions specifies set of PathSet's for different platforms.
// Ex.:
// pathOptions["linux"] = PathSet with Linux paths
// pathOptions["windows"] = PathSet with Windows paths
type PathOptions map[string]PathSet

// BootstrapConfig specifies how to bootstrap application's configuration.
// It provides configuration options on:
//   1. where to look for config files
//   2. what config file type to look for
//   3. should ENV VARs be used as config options
//   ... etc
type BootstrapConfig struct {
	pathOptions PathOptions
	// envVarPrefix specifies prefix used to search for env variables to be used as config options
	envVarPrefix string
	// configFile specifies config filename without extension. Used in combination with configType. Ex: "config"
	configFile string
	// configType specifies config file extension. Used in combination with configFile. Ex.: "yaml"
	configType string
}

// NewBootstrapConfig
func NewBootstrapConfig() *BootstrapConfig {
	return &BootstrapConfig{
		pathOptions:  make(PathOptions),
		envVarPrefix: "",
		configFile:   "config",
		configType:   "yaml",
	}
}

// SetEnvVarPrefix
func (c *BootstrapConfig) SetEnvVarPrefix(prefix string) *BootstrapConfig {
	c.envVarPrefix = strings.Replace(strings.ToUpper(prefix), "-", "_", -1)
	return c
}

// SetConfigFile
func (c *BootstrapConfig) SetConfigFile(file string) *BootstrapConfig {
	c.configFile = file
	return c
}

// SetConfigType
func (c *BootstrapConfig) SetConfigType(_type string) *BootstrapConfig {
	c.configType = _type
	return c
}

// AddWindowsPaths
func (c *BootstrapConfig) AddWindowsPaths(rootPaths, homeRelativePaths []string) *BootstrapConfig {
	pathSet := make(PathSet)
	pathSet[root] = rootPaths
	pathSet[home] = homeRelativePaths

	c.pathOptions[windows] = pathSet
	return c
}

// AddLinuxPaths
func (c *BootstrapConfig) AddLinuxPaths(rootPaths, homeRelativePaths []string) *BootstrapConfig {
	pathSet := make(PathSet)
	pathSet[root] = rootPaths
	pathSet[home] = homeRelativePaths

	c.pathOptions[linux] = pathSet
	return c
}

// getRootPaths
func (c *BootstrapConfig) getRootPaths() []string {
	return c.getPaths(root)
}

// getHomePaths
func (c *BootstrapConfig) getHomePaths() []string {
	return c.getPaths(home)
}

// getPaths returns running-OS-specific paths from BootstrapConfig
func (c *BootstrapConfig) getPaths(what string) []string {
	if _, ok := c.pathOptions[runtime.GOOS]; ok {
		paths, _ := c.pathOptions[runtime.GOOS][what]
		return paths
	}
	return nil
}

// InitConfig initializes application config according to provided BootstrapConfig options and
// reads in found app's config file and ENV variables if set.
func InitConfig(bootstrapConfig *BootstrapConfig) {
	log.Info("InitConfig()")

	if bootstrapConfig == nil {
		// Provide some default bootstrap config
		bootstrapConfig = NewBootstrapConfig()
	}

	if configFile != "" {
		// Look for explicitly specified config file
		conf.SetConfigFile(configFile)
		log.Infof("InitConfig() - looking for explicitly specified config: %s", configFile)
	} else {
		// Config file is not explicitly specified, we need to find it
		// We need to search for config file in pre-defined set of paths, such as /etc/, /home/, etc

		// We'll look for config file in root-based list of dirs, such as /etc, /opt/etc ...
		for _, path := range bootstrapConfig.getRootPaths() {
			log.Infof("InitConfig() - add root path to look for config: %v", path)
			conf.AddConfigPath(path)
		}

		// We'll look for default config file in HOMEDIR-based list of dirs, such as $HOME/.atlas ...
		// Find HOMEDIR dir
		homedir, err := hd.Dir()
		if err != nil {
			log.Fatalf("InitConfig() - unable to find homedir %v", err)
		}
		// Build HOMEDIR-based path
		for _, path := range bootstrapConfig.getHomePaths() {
			homeRelativePath := homedir + "/" + path
			log.Infof("InitConfig() - add home relative path to look for config: %v : %v", path, homeRelativePath)
			conf.AddConfigPath(homeRelativePath)
		}

		// At last specify config file (without extension) and config type (extension)
		log.Infof("InitConfig() - add config file name to look for: %v", bootstrapConfig.configFile)
		log.Infof("InitConfig() - add config file type to look for: %v", bootstrapConfig.configType)
		conf.SetConfigName(bootstrapConfig.configFile)
		conf.SetConfigType(bootstrapConfig.configType) // REQUIRED if the config file does not have the extension in the name
	}

	if bootstrapConfig.envVarPrefix != "" {
		// As we have ENV prefix specified, setup ENV vars search
		log.Infof("InitConfig() - set env prefix: %v_", bootstrapConfig.envVarPrefix)
		// By default empty environment variables are considered unset and will fall back to the next configuration source.
		// To treat empty environment variables as set, use the AllowEmptyEnv method.
		conf.AllowEmptyEnv(false)
		// Check for an env var with a name matching the key upper-cased and prefixed with the EnvPrefix
		// Prefix has "_" added automatically, so no need to say 'ATLAS_'
		conf.SetEnvPrefix(bootstrapConfig.envVarPrefix)
		// SetEnvKeyReplacer allows you to use a strings.Replacer object to rewrite Env keys to an extent.
		// This is useful if you want to use - or something in your Get() calls, but want your environmental variables to use _ delimiters.
		conf.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
		// Check ENV variables for all keys set in config, default & flags
		conf.AutomaticEnv()
	}

	// Read configuration

	if err := conf.ReadInConfig(); err == nil {
		log.Infof("InitConfig() - config file used: %s", conf.ConfigFileUsed())
	} else if _, ok := err.(conf.ConfigFileNotFoundError); ok {
		// Config file not found
		log.Infof("InitConfig() - no config file found")
	} else {
		// Config file was found but another error was produced
		log.Errorf("InitConfig() - unable to read config file: %v", err)
	}
}
