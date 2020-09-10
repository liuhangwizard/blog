package config

//server
type VBConfigServer struct {
	Mode        string `toml:"mode"`
	Port        string `toml:"port"`
	Domain      string `toml:"domain"`
	Gzip        string `toml:"gzip"`
	SessionTime int    `toml:"sessiontime"`
	CookiesTime int    `toml:"cookiestime"`
}

//log
type VBConfigLog struct {
	Name       string `toml:"name"`
	Path       string `toml:"path"`
	MaxSize    int    `toml:"maxsize"`
	MaxBackups int    `toml:"maxbackups"`
	MaxAge     int    `toml:"maxage"`
	Compress   bool   `toml:"compress"`
}

//database
//database mysql
type VBConfigDBMySQL struct {
	Type         string `toml:"type"`
	Mode         string `toml:"mode"`
	Host         string `toml:"host"`
	Port         string `toml:"port"`
	Name         string `toml:"name"`
	Username     string `toml:"username"`
	Password     string `toml:"password"`
	Args         string `toml:"args"`
	Maxidleconns int    `toml:"maxidle"`
	Maxopenconns int    `toml:"maxopen"`
}

//database redis
type VBConfigDBRedis struct {
	Type       string `toml:"type"`
	Mode       string `toml:"mode"`
	Host       string `toml:"host"`
	Port       string `toml:"port"`
	Password   string `toml:"password"`
	PoolSize   int    `toml:"poolsize"`
	PoolActive int    `toml:"poolactive"`
	TimeOut    int    `toml:"timeout"`
	DB         int    `toml:"db"`
}

//oss
type VBConfigDBOSS struct {
	EndPoint        string `toml:"endpoint"`
	EndPointHost    int    `toml:"endpointhost"`
	Expire          int64  `toml:"expire"`
	Bucket          string `toml:"bucket"`
	Dirctory        string `toml:"dirctory"`
	AccessKeyId     string `toml:"accesskeyid"`
	AccessKeySecret string `toml:"accesskeysecret"`
}
