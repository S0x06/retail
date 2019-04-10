package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type App struct{
	RunMode string
	Addr  string
	Name string
	Url string
	JwtSecret string
}

type Mysql struct{
	Name string
	Addr string
	UserName string
	passWord string
}

type Redis struct{
	Addr string
	UserName string
	passWord string
}

type Consul struct{
	Addr string
}

type Zikpin struct{
	Addr string
}

type MQ struct{
	Addr string
}

type Config struct {
	Name string
	App *App
	Mysq *Mysql
	Redis *Redis
	Consul *Consul
	Zikpin *Zikpin
		MQ *MQ
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.New(); err != nil {
		return err
	}

	// 初始化日志包
	c.NewLog()

	// 监控配置文件变化并热加载程序
	c.watch()

	return nil
}

func (c *Config) New() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")  // 设置配置文件格式为YAML
	viper.AutomaticEnv()         // 读取匹配的环境变量
	viper.SetEnvPrefix("SERVER") // 读取环境变量的前缀为SERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

func (c *Config) NewLog() {
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


func NewApp *App{

	app := &App{
		RunMode:  viper.GetString("app.runmode"),
		Addr :  viper.GetString("app.addr"),
		Name:  viper.GetString("app.name"),
		Url : viper.GetString("app.url"),
		JwtSecret:  viper.GetString("app.jwt_secret")
	}

	return app
}


func NewMysql() *Mysql{
	mysql := &Mysql{
		UserName: viper.GetString("mysql.username"),
		PassWord: viper.GetString("mysql.password"),
		Addr: viper.GetString("mysql.addr"),
		Name: viper.GetString("mysql.name")
	}
	
	return mysql
}

func NewRedis() *Redis{

	redis := &Redis{
			Addr : viper.GetString("redis.addr"),
			UserName :viper.GetString("redis.username"),
			passWord :viper.GetString("redis.password"),
	} 
	return redis
}


func NewConsul() *Consul{
	consul := &Consul{
			Addr : viper.GetString("consul.addr"),
			UserName :viper.GetString("consul.username"),
			passWord :viper.GetString("consul.password"),
	} 
	return consul
}

func NewZikpin() *Zikpin{
	zikpin := &Zikpin{
			Addr : viper.GetString("zikpin.addr"),
			UserName :viper.GetString("zikpin.username"),
			passWord :viper.GetString("zikpin.password"),
	} 
	return zikpin
}

func (c *Config) NewMQ() *MQ{
	mq := &MQ{
			Addr : viper.GetString("mq.addr"),
			UserName :viper.GetString("mq.username"),
			passWord :viper.GetString("mq.password"),
	} 
	return mq
}

// 监控配置文件变化并热加载程序
func (c *Config) watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}
