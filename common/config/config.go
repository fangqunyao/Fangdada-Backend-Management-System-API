// Package config 文件配置

package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// 总配文件
type config struct {
	Server        server        `yaml:"server"`
	Db            db            `yaml:"db"`
	Redis         redis         `yaml:"redis"`
	Log           log           `yaml:"log"`
	ImageSettings imageSettings `yaml:"imageSettings"`
}

// 项目端口配置
type server struct {
	Address string `yaml:"address"`
	Model   string `yaml:"model"`
}

// 数据库配置
type db struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

// redis配置
type redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

// log日志
type log struct {
	Path  string `yaml:"path"`
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
}

// imageSettings图片上传配置
type imageSettings struct {
	UploadDir string `yaml:"uploadDir"`
	ImageHost string `yaml:"imageHost"`
}

var Config *config

// 配置初始化
func init() {
	// 使用 os.ReadFile
	yamlFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %v", err))
	}

	// 绑定值
	if err := yaml.Unmarshal(yamlFile, &Config); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %v", err))
	}
}
