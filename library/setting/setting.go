package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type AppConfig struct {
	AppMode     string `ini:"APP_MODE"`
	AppUrl      string `ini:"APP_URL"`
	AppLocal    string `ini:"APP_LOCAL"`
	PageSize    uint   `ini:"PAGE_SIZE"`
	JwtSecret   string `ini:"JWT_SECRET"`
	RuntimePath string `ini:"RUNTIME_PATH"`
}

type FileConfig struct {
	ImagePath string `ini:"IMAGE_PATH"`
	// ImageMaxSize  int      `ini:"APP_URL"`
	ImageAllowExt []string `ini:"IMAGE_ALLOW_EXT"`
	ExportPath    string   `ini:"EXPORT_PATH"`
	QrCodePath    string   `ini:"QRCODE_PATH"`
	// FontSavePath   string  `ini:"APP_URL"`
}

type LogConfig struct {
	SavePath   string `ini:"LOG_PATH"`
	Prefix     string `ini:"LOG_PREFIX"`
	FileExt    string `ini:"LOG_EXT"`
	TimeFormat string `ini:"TIME_FORMAT"`
}

type DatabaseConfig struct {
	DbType      string `ini:"DB_TYPE"`
	Port        string `ini:"DB_PORT"`
	Host        string `ini:"DB_HOST"`
	User        string `ini:"DB_USER"`
	Password    string `ini:"DB_PASSWORD"`
	DbName      string `ini:"DB_NAME"`
	TablePrefix string `ini:"TABLE_PREFIX"`
}

type RedisConfig struct {
	Host        string        `ini:"REDIS_HOST"`
	Port        string        `ini:"REDIS_PORT"`
	Password    string        `ini:"REDIS_PASSWORD"`
	MaxIdle     int           `ini:"MAX_IDLE"`
	MaxActive   int           `ini:"MAX_ACTIVE"`
	IdleTimeout time.Duration `ini:"IDLE_TIMEOUT"`
}

var AppCfg = &AppConfig{}
var FileCfg = &FileConfig{}
var LogCfg = &LogConfig{}
var DatabaseCfg = &DatabaseConfig{}
var RedisCfg = &RedisConfig{}
// 中文分词
// var Seg gse.Segmenter

// Setup initialize the configuration instance
func Initialized() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to load 'conf/app.ini': %v", err)
	}
	// 加载默认字典
	// Seg.LoadDict()
	// initialize the app setting
	cfg.Section("app").MapTo(AppCfg)
	cfg.Section("file").MapTo(FileCfg)
	cfg.Section("log").MapTo(LogCfg)
	cfg.Section("database").MapTo(DatabaseCfg)
	cfg.Section("redis").MapTo(RedisCfg)
}

// func testSeg() {
// 	// 分词文本
// 	tb := []byte("山达尔星联邦共和国联邦政府")
//
// 	// 处理分词结果
// 	// 支持普通模式和搜索模式两种分词，见代码中 ToString 函数的注释。
// 	// 搜索模式主要用于给搜索引擎提供尽可能多的关键字
// 	fmt.Println("输出分词结果, 类型为字符串, 使用搜索模式: ", setting.Seg.String(tb, true))
// 	fmt.Println("输出分词结果, 类型为 slice: ", setting.Seg.Slice(tb))
//
// 	segments := setting.Seg.Segment(tb)
// 	// 处理分词结果
// 	fmt.Println(gse.ToString(segments))
//
// 	text1 := []byte("上海地标建筑, 东方明珠电视台塔上海中心大厦")
// 	segments1 := setting.Seg.Segment([]byte(text1))
// 	fmt.Println(gse.ToString(segments1, true))
// }
