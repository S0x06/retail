package config

import (
	"strings"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type Base struct{
	RunMode string
	Addr  string
	Name string
	Url string
	JwtSecret string
}

type Tls struct {
	Addr string
	Cert string
	Key string
}

type Mysql struct {
	Name string
	Addr string
	UserName string
	PassWord string
}

type Config struct {
	Path string
	Tls *Tls
	Base *Base
	Mysql *Mysql
}

//初始化配置
var Cfg = &Config{}

func init(){
	if err := Cfg.NewConfig(); err != nil {
		panic(err)
	}

	fmt.Println("All config load success")

}
	
// 配置
func (cfg *Config) NewConfig() error {

	if cfg.Path != "" {
		viper.SetConfigFile(cfg.Path) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
		cfg.Path = "conf/config"
	}
	viper.SetConfigType("yaml")  // 设置配置文件格式为YAML
	viper.AutomaticEnv()         // 读取匹配的环境变量
	viper.SetEnvPrefix("SERVER") // 读取环境变量的前缀为SERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	// 应用配置
	cfg.baseConfig()

	// TLS配置
	cfg.tlsConfig()

	//数据库配置
    cfg.mysqlConfig()

	// 初始化日志包
	cfg.logConfig()

	// 监控配置文件变化并热加载程序
	cfg.watch()

	return nil
}


//基础配置加载
func (cfg *Config)baseConfig() *Base{
	base := &Base{
		RunMode:  viper.GetString("base.runmode"),
		Addr :  viper.GetString("base.addr"),
		Name:  viper.GetString("base.name"),
		Url : viper.GetString("base.url"),
		JwtSecret:  viper.GetString("base.jwt_secret"),
	}

	cfg.Base = base
	return base
}


//TLS配置加载
func (cfg *Config)tlsConfig() *Tls{
	tls := &Tls{
		Addr :  viper.GetString("tls.addr"),
		Cert : viper.GetString("tls.cert"),
		Key : viper.GetString("tls.key"),
	}

	cfg.Tls = tls
	return tls
}


//mysql配置加载
func (cfg *Config)mysqlConfig() *Mysql{
	mysql := &Mysql{
		UserName: viper.GetString("mysql.username"),
		PassWord: viper.GetString("mysql.password"),
		Addr: viper.GetString("mysql.addr"),
		Name: viper.GetString("mysql.name"),
	}
	
	cfg.Mysql = mysql
	return mysql
}

//日记配置加载
func (cfg *Config) logConfig() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)
}


// 监控配置文件变化并热加载程序
func (cfg *Config) watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}
