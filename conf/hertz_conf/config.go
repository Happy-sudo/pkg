package hertz_conf

// Logger 日志配置
type Logger struct {
	Enable     bool   `yaml:"enable" `     //是否启用自定义日志配置
	Filename   string `yaml:"file_name"`   //路径
	MaxSize    int    `yaml:"max_size"`    //日志的最大大小（M）
	MaxBackups int    `yaml:"max_backups"` //日志的最大保存数量
	MaxAge     int    `yaml:"max_age"`     //日志文件存储最大天数
	Compress   bool   `yaml:"compress"`    //是否执行压缩
	LocalTime  bool   `yaml:"local_time"`  //是否使用格式化时间辍
}

// Server 服务端配置
type Server struct {
	Http      http      `yaml:"http"`      //服务ip配置
	Polaris   polaris   `yaml:"polaris"`   //北极星注册中心配置
	Auth      auth      `yaml:"auth"`      // auth 身份认证配置
	Cors      cors      `yaml:"cors"`      //cors 配置
	Recovery  recovery  `yaml:"recovery"`  // recovery 配置
	Gzip      gzip      `yaml:"gzip"`      // 压缩配置
	I18n      i18n      `yaml:"i18n"`      // 国际化配置
	Swag      swag      `yaml:"swag"`      // swag文档
	Jaeger    jaeger    `yaml:"jaeger"`    //链路配置
	Transport transport `yaml:"transport"` //多路复用配置
}

// Service 服务名称配置
type Service struct {
	NameSpace  string `yaml:"namespace"`   //服务空间名称
	ServerName string `yaml:"server_name"` //服务名称
	ClientName string `yaml:"client_name"` //客户端名称
	Version    string `yaml:"version"`     //版本信息
}

// 服务地址端口配置
type http struct {
	Enable       bool   `yaml:"enable" `        //是否启用http自定义配置
	Address      string `yaml:"address"`        //地址
	ExitWaitTime int    `yaml:"exit_wait_time"` // 配置
}

// 注册中心配置
type polaris struct {
	Enable  bool   `yaml:"enable"` //是否启用注册中心，默认开启
	Network string `yaml:"network"`
	Address string `yaml:"address"`
}

//链路追踪配置
type jaeger struct {
	Enable   bool   `yaml:"enable"`   //是否启用链路追踪
	Endpoint string `yaml:"endpoint"` //地址
}

//多路复用配置
type transport struct {
	Enable bool `yaml:"enable"` //是否启用多路复用
}

// auth 配置
type auth struct {
	Enable bool   `yaml:"enable"` //是否启用auth配置
	AK     string `yaml:"ak"`
	SK     string `yaml:"sk"`
}

// cors配置
type cors struct {
	Enable bool `yaml:"enable"` //是否启用cors配置
}

// recovery 配置
type recovery struct {
	Enable bool `yaml:"enable"` //是否启用 recovery 配置
}

// gzip 压缩配置
type gzip struct {
	Enable             bool     `yaml:"enable"`              //是否启用 gzip 配置
	BestCompression    bool     `yaml:"best_compression"`    //提供最佳的文件压缩率
	BestSpeed          bool     `yaml:"best_speed"`          //提供了最佳的压缩速度
	DefaultCompression bool     `yaml:"default_compression"` //默认压缩率
	NoCompression      bool     `yaml:"no_compression"`      //不进行压缩
	Excluded           excluded `yaml:"excluded"`            // 设置不需要压缩的方式
}

type excluded struct {
	Enable              bool                `yaml:"enable"` //是否启用 excluded 配置
	ExcludedExtensions  excludedExtensions  `yaml:"excluded_extensions"`
	ExcludedPaths       excludedPaths       `yaml:"excluded_paths"`
	ExcludedPathRegexes excludedPathRegexes `yaml:"excluded_path_regexes"`
}

// 设置不需要 gzip 压缩的文件后缀
type excludedExtensions struct {
	Enable     bool   `yaml:"enable"`     //是否启用 excluded 配置
	Extensions string `yaml:"extensions"` // 文件后缀 数组用,连接 eg:".pdf", ".mp4"
}

// 设置不需要进行 gzip 压缩的文件路径
type excludedPaths struct {
	Enable bool   `yaml:"enable"` //是否启用 excludedPaths 配置
	Paths  string `yaml:"paths"`  //文件路径 数组用,连接 eg:/api/
}

// 设置自定义的正则表达式来过滤掉不需要 gzip 压缩的文件
type excludedPathRegexes struct {
	Enable  bool   `yaml:"enable"`  //是否启用 excludedPathRegexes 配置
	Regexes string `yaml:"regexes"` //文件路径 数组用,连接  eg: /api.*
}

// 设置国际化
type i18n struct {
	Enable bool `yaml:"enable"` //是否启用 i18n 配置
}

// 设置 swag
type swag struct {
	Enable bool `yaml:"enable"` //是否启用 swag 配置
}

// **********************************公共对象*******************************
type statsLevel struct {
	LevelDisabled bool `yaml:"level_disabled"`
	LevelBase     bool `yaml:"level_base"`
	LevelDetailed bool `yaml:"level_detailed"`
}

// Client **********************************客户端对象******************************
//客户端配置
type Client struct {
	TimeoutControl timeOutControl `yaml:"timeout_control"` //超时控制
	ConnectionType connectionType `yaml:"connection_type"` // 连接类型
	FailureRetry   failureRetry   `yaml:"failure_retry"`   //请求重试
	LoadBalancer   loadBalancer   `yaml:"load_balancer"`   //负载均衡
	CBSuite        cbsuite        `yaml:"cbsuite"`         //熔断器
	StatsLevel     statsLevel     `yaml:"stats_level"`     //埋点策略&埋点粒度
}

//超时控制
type timeOutControl struct {
	RpcTimeout     rpcTimeout     `yaml:"rpc_timeout"`
	ConnectTimeOut connectTimeOut `yaml:"connect_time_out"`
}

//连接类型（长链接 短链接）
type connectionType struct {
	ShortConnection shortConnection `yaml:"short_connection"` //短链接
	LongConnection  longConnection  `yaml:"long_connection"`  //长链接
	ClientTransport clientTransport `yaml:"transport"`        //客户端多路复用

}

//rpc超时控制
type rpcTimeout struct {
	Enable  bool   `yaml:"enable"`   //是否启用rpc超时
	Timeout string `yaml:"time_out"` //超时时间 （默认 1s 单位："ns", "us" (or "µs"), "ms", "s", "m", "h"）
}

//connect超时控制
type connectTimeOut struct {
	Enable  bool   `yaml:"enable"`   //是否启用rpc超时
	TimeOut string `yaml:"time_out"` //连接超时 （默认：50ms）
}

//短链接
type shortConnection struct {
	Enable bool `yaml:"enable"` //是否启用短链接
}

//长链接
type longConnection struct {
	Enable            bool   `yaml:"enable"`               //是否启用长链接
	MaxIdlePerAddress int    `yaml:"max_idle_per_address"` //最大空闲地址
	MinIdlePerAddress int    `yaml:"min_idle_per_address"` //最小空闲地址
	MaxIdleGlobal     int    `yaml:"max_idle_global"`      //最大空闲数
	MaxIdleTimeOut    string `yaml:"max_idle_time_out"`    //最大空闲超时
}

// 客户端多路复用
type clientTransport struct {
	Enable        bool `yaml:"enable"`         //是否启用多路复用
	MuxConnection int  `yaml:"mux_connection"` //连接数
}

//重试机制
type failureRetry struct {
	Enable        bool `yaml:"enable"`          //是否启用请求重试机制
	MaxRetryTimes int  `yaml:"max_retry_times"` //重试次数
}

//负载均衡
type loadBalancer struct {
	Enable bool `yaml:"enable"` //是否启用负载均衡
}

//熔断器
type cbsuite struct {
	Enable bool `yaml:"enable"` //是否启用熔断器
}
