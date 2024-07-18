package config

import (
	"gorm.io/gorm/logger"
	"strings"
)

type ServerConfig struct {
	Path string `json:"path"`
	Port int    `json:"port"`
}
type CORSWhitelist struct {
	AllowOrigin      string `mapstructure:"allow-origin"`
	AllowMethods     string `mapstructure:"allow-methods"`
	AllowHeaders     string `mapstructure:"allow-headers"`
	ExposeHeaders    string `mapstructure:"expose-headers"`
	AllowCredentials bool   `mapstructure:"allow-credentials"`
}

type CORSConfig struct {
	Mode      string          `mapstructure:"mode"`
	Whitelist []CORSWhitelist `mapstructure:"whitelist"`
}

type ExcelConfig struct {
	Dir string `mapstructure:"dir"`
}

type JWTConfig struct {
	SigningKey  string `mapstructure:"signing-key"`
	ExpiresTime string `mapstructure:"expires-time"`
	BufferTime  string `mapstructure:"buffer-time"`
	Issuer      string `mapstructure:"issuer"`
}

type LocalConfig struct {
	Path      string `mapstructure:"path"`
	StorePath string `mapstructure:"store-path"`
}

type MySQLConfig struct {
	Prefix       string `mapstructure:"prefix"`
	Port         string `mapstructure:"port"`
	Config       string `mapstructure:"config"`
	DBName       string `mapstructure:"db-name"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Path         string `mapstructure:"path"`
	Engine       string `mapstructure:"engine"`
	LogMode      string `mapstructure:"log-mode"`
	MaxIdleConns int    `mapstructure:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns"`
	Singular     bool   `mapstructure:"singular"`
	LogZap       bool   `mapstructure:"log-zap"`
}

type ZapConfig struct {
	Level         string `mapstructure:"level"`
	Prefix        string `mapstructure:"prefix"`
	Format        string `mapstructure:"format"`
	Director      string `mapstructure:"director"`
	EncodeLevel   string `mapstructure:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key"`
	ShowLine      bool   `mapstructure:"show-line"`
	LogInConsole  bool   `mapstructure:"log-in-console"`
	RetentionDay  int    `mapstructure:"retention-day"`
}

type CoreConfig struct {
	Path      string `mapstructure:"path"`
	Port      string `mapstructure:"port"`
	AccessKey string `mapstructure:"access-key"`
}

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	CORS   CORSConfig   `mapstructure:"cors"`
	Excel  ExcelConfig  `mapstructure:"excel"`
	JWT    JWTConfig    `mapstructure:"jwt"`
	Local  LocalConfig  `mapstructure:"local"`
	MySQL  MySQLConfig  `mapstructure:"mysql"`
	Zap    ZapConfig    `mapstructure:"zap"`
	Core   CoreConfig   `mapstructure:"core"`
}

func (c MySQLConfig) LogLevel() logger.LogLevel {
	switch strings.ToLower(c.LogMode) {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}

type SpecializedDB struct {
	Type        string `mapstructure:"type" json:"type" yaml:"type"`
	AliasName   string `mapstructure:"alias-name" json:"alias-name" yaml:"alias-name"`
	MySQLConfig `yaml:",inline" mapstructure:",squash"`
	Disable     bool `mapstructure:"disable" json:"disable" yaml:"disable"`
}
