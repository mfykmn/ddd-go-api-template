package config

import (
	"github.com/BurntSushi/toml"
)

// Config アプリケーションの設定値を持つ構造体
type Config struct {
	Server   Server `toml:"server"`
	Redis    Redis  `toml:"redis"`
	DBMaster DB     `toml:"dbm"`
	DBSlave  DB     `toml:"dbs"`
}

// Server　APIサーバーの設定を保持する構造体
type Server struct {
	Port string `toml:"port"`
}

// DB データベースの設定を保持する構造体
type DB struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
}

// Redis Redisの設定を保持する構造体
type Redis struct {
	Host                  string `toml:"host"`
	Port                  string `toml:"port"`
	PoolMaxIdle           int    `toml:"pool_max_idle"`
	PoolIdleTimeoutSecond int    `toml:"pool_idle_timeout_second"`
}

// New Configのコンストラクタ
func New(config *Config, configPath string) error {
	_, err := toml.DecodeFile(configPath, config)
	return err
}
