package util

import "gorm.io/gorm"

type DBOptional struct {
	DriverName   string `yaml:"DriverName"`
	Timeout      string `yaml:"Timeout"`      // Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	ReadTimeout  string `yaml:"ReadTimeout"`  // Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	WriteTimeout string `yaml:"WriteTimeout"` // Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	User         string `yaml:"User"`
	Password     string `yaml:"Password"`
	DBName       string `yaml:"DBName"`
	DBCharset    string `yaml:"DBCharset"`
	DBHostname   string `yaml:"DBHostname"`
	DBPort       string `yaml:"DBPort"`
	MaxIdleConns int    `yaml:"MaxIdleConns"`
	MaxOpenConns int    `yaml:"MaxOpenConns"`
	ExtraDSNMap  map[string]string
	Config       gorm.Config
}
