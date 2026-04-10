// Package config 文件配置

package config

import (
	"fmt"
	"os"
	"path/filepath" // 1. 引入 path/filepath 包

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
	// --- 修改开始 ---

	// 1. 获取当前可执行文件（main）的绝对路径
	exePath, err := os.Executable()
	if err != nil {
		panic(fmt.Errorf("failed to get executable path: %v", err))
	}

	// 2. 获取可执行文件所在的目录（例如：/www/wwwroot/admin-go）
	exeDir := filepath.Dir(exePath)

	// 3. 拼接出 config.yaml 的绝对路径
	configPath := filepath.Join(exeDir, "config.yaml")

	// 4. 使用绝对路径读取文件
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Errorf("failed to read config file at %s: %v", configPath, err))
	}

	// --- 修改结束 ---

	// 绑定值
	if err := yaml.Unmarshal(yamlFile, &Config); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %v", err))
	}
}