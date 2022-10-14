package config

// UserSrvConfig user_srv的配置信息
type UserSrvConfig struct {
	// 使用consul服务发现，不再需要host和port
	//Host string `mapstructure:"host" json:"host"`
	//Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

// JWTConfig JWT的签名
type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

// AliSmsConfig 阿里短信服务配置信息
type AliSmsConfig struct {
	ApiKey       string `mapstructure:"key" json:"key"`
	ApiSecrect   string `mapstructure:"secrect" json:"secrect"`
	TemplateCode string `mapstructure:"template-code" json:"template-code"`
	SignName     string `mapstructure:"sign-name" json:"sign-name"`
	RegionId     string `mapstructure:"region-id" json:"region-id"`
	Width        int    `json:"width"`
}

// ConsulConfig consul配置信息
type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

// RedisConfig redis配置信息
type RedisConfig struct {
	Host   string `mapstructure:"host" json:"host"`
	Port   int    `mapstructure:"port" json:"port"`
	Expire int    `mapstructure:"expire" json:"expire"`
}

// ServerConfig 总的服务配置信息 从nacos中读取
type ServerConfig struct {
	Name        string        `mapstructure:"name" json:"name"`
	Host        string        `mapstructure:"host" json:"host"`
	Tags        []string      `mapstructure:"tags" json:"tags"`
	Port        int           `mapstructure:"port" json:"port"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	JWTInfo     JWTConfig     `mapstructure:"jwt" json:"jwt"`
	AliSmsInfo  AliSmsConfig  `mapstructure:"sms" json:"sms"`
	RedisInfo   RedisConfig   `mapstructure:"redis" json:"redis"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul" json:"consul"`
}

// NacosConfig nacos配置信息
type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}
