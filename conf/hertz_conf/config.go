package hertz_conf

// Logger 日志配置
type Logger struct {
	Enable     bool   `mapstructure:"enable" `     //是否启用自定义日志配置
	Filename   string `mapstructure:"file_name"`   //路径
	MaxSize    int    `mapstructure:"max_size"`    //日志的最大大小（M）
	MaxBackups int    `mapstructure:"max_backups"` //日志的最大保存数量
	MaxAge     int    `mapstructure:"max_age"`     //日志文件存储最大天数
	Compress   bool   `mapstructure:"compress"`    //是否执行压缩
	LocalTime  bool   `mapstructure:"local_time"`  //是否使用格式化时间辍
}

// Server 服务端配置
type Server struct {
	Http      http      `mapstructure:"http"`      //服务ip配置
	Polaris   polaris   `mapstructure:"polaris"`   //北极星注册中心配置
	Auth      auth      `mapstructure:"auth"`      // auth 身份认证配置
	Cors      cors      `mapstructure:"cors"`      //cors 配置
	Recovery  recovery  `mapstructure:"recovery"`  // recovery 配置
	Gzip      gzip      `mapstructure:"gzip"`      // 压缩配置
	I18n      i18n      `mapstructure:"i18n"`      // 国际化配置
	Swag      swag      `mapstructure:"swag"`      // swag文档
	Jaeger    jaeger    `mapstructure:"jaeger"`    //链路配置
	Transport transport `mapstructure:"transport"` //多路复用配置
}

// Service 服务名称配置
type Service struct {
	NameSpace  string `mapstructure:"namespace"`   //服务空间名称
	ServerName string `mapstructure:"server_name"` //服务名称
	ClientName string `mapstructure:"client_name"` //客户端名称
	Version    string `mapstructure:"version"`     //版本信息
}

// 服务地址端口配置
type http struct {
	Enable       bool   `mapstructure:"enable" `        //是否启用http自定义配置
	Address      string `mapstructure:"address"`        //地址
	ExitWaitTime int    `mapstructure:"exit_wait_time"` // 配置
}

// 注册中心配置
type polaris struct {
	Enable  bool   `mapstructure:"enable"` //是否启用注册中心，默认开启
	Network string `mapstructure:"network"`
	Address string `mapstructure:"address"`
}

//链路追踪配置
type jaeger struct {
	Enable   bool   `mapstructure:"enable"`   //是否启用链路追踪
	Endpoint string `mapstructure:"endpoint"` //地址
}

//多路复用配置
type transport struct {
	Enable bool `mapstructure:"enable"` //是否启用多路复用
}

// auth 配置
type auth struct {
	Enable bool   `mapstructure:"enable"` //是否启用auth配置
	AK     string `mapstructure:"ak"`
	SK     string `mapstructure:"sk"`
}

// cors配置
type cors struct {
	Enable bool `mapstructure:"enable"` //是否启用cors配置
}

// recovery 配置
type recovery struct {
	Enable bool `mapstructure:"enable"` //是否启用 recovery 配置
}

// gzip 压缩配置
type gzip struct {
	Enable             bool     `mapstructure:"enable"`              //是否启用 gzip 配置
	BestCompression    bool     `mapstructure:"best_compression"`    //提供最佳的文件压缩率
	BestSpeed          bool     `mapstructure:"best_speed"`          //提供了最佳的压缩速度
	DefaultCompression bool     `mapstructure:"default_compression"` //默认压缩率
	NoCompression      bool     `mapstructure:"no_compression"`      //不进行压缩
	Excluded           excluded `mapstructure:"excluded"`            // 设置不需要压缩的方式
}

type excluded struct {
	Enable              bool                `mapstructure:"enable"` //是否启用 excluded 配置
	ExcludedExtensions  excludedExtensions  `mapstructure:"excluded_extensions"`
	ExcludedPaths       excludedPaths       `mapstructure:"excluded_paths"`
	ExcludedPathRegexes excludedPathRegexes `mapstructure:"excluded_path_regexes"`
}

// 设置不需要 gzip 压缩的文件后缀
type excludedExtensions struct {
	Enable     bool   `mapstructure:"enable"`     //是否启用 excluded 配置
	Extensions string `mapstructure:"extensions"` // 文件后缀 数组用,连接 eg:".pdf", ".mp4"
}

// 设置不需要进行 gzip 压缩的文件路径
type excludedPaths struct {
	Enable bool   `mapstructure:"enable"` //是否启用 excludedPaths 配置
	Paths  string `mapstructure:"paths"`  //文件路径 数组用,连接 eg:/api/
}

// 设置自定义的正则表达式来过滤掉不需要 gzip 压缩的文件
type excludedPathRegexes struct {
	Enable  bool   `mapstructure:"enable"`  //是否启用 excludedPathRegexes 配置
	Regexes string `mapstructure:"regexes"` //文件路径 数组用,连接  eg: /api.*
}

// 设置国际化
type i18n struct {
	Enable bool `mapstructure:"enable"` //是否启用 i18n 配置
}

// 设置 swag
type swag struct {
	Enable bool `mapstructure:"enable"` //是否启用 swag 配置
}

// **********************************公共对象*******************************
type statsLevel struct {
	LevelDisabled bool `mapstructure:"level_disabled"`
	LevelBase     bool `mapstructure:"level_base"`
	LevelDetailed bool `mapstructure:"level_detailed"`
}

// Client **********************************客户端对象******************************
//客户端配置
type Client struct {
	TimeoutControl timeOutControl `mapstructure:"timeout_control"` //超时控制
	ConnectionType connectionType `mapstructure:"connection_type"` // 连接类型
	FailureRetry   failureRetry   `mapstructure:"failure_retry"`   //请求重试
	LoadBalancer   loadBalancer   `mapstructure:"load_balancer"`   //负载均衡
	CBSuite        cbsuite        `mapstructure:"cbsuite"`         //熔断器
	StatsLevel     statsLevel     `mapstructure:"stats_level"`     //埋点策略&埋点粒度
}

//超时控制
type timeOutControl struct {
	RpcTimeout     rpcTimeout     `mapstructure:"rpc_timeout"`
	ConnectTimeOut connectTimeOut `mapstructure:"connect_time_out"`
}

//连接类型（长链接 短链接）
type connectionType struct {
	ShortConnection shortConnection `mapstructure:"short_connection"` //短链接
	LongConnection  longConnection  `mapstructure:"long_connection"`  //长链接
	ClientTransport clientTransport `mapstructure:"transport"`        //客户端多路复用

}

//rpc超时控制
type rpcTimeout struct {
	Enable  bool   `mapstructure:"enable"`   //是否启用rpc超时
	Timeout string `mapstructure:"time_out"` //超时时间 （默认 1s 单位："ns", "us" (or "µs"), "ms", "s", "m", "h"）
}

//connect超时控制
type connectTimeOut struct {
	Enable  bool   `mapstructure:"enable"`   //是否启用rpc超时
	TimeOut string `mapstructure:"time_out"` //连接超时 （默认：50ms）
}

//短链接
type shortConnection struct {
	Enable bool `mapstructure:"enable"` //是否启用短链接
}

//长链接
type longConnection struct {
	Enable            bool   `mapstructure:"enable"`               //是否启用长链接
	MaxIdlePerAddress int    `mapstructure:"max_idle_per_address"` //最大空闲地址
	MinIdlePerAddress int    `mapstructure:"min_idle_per_address"` //最小空闲地址
	MaxIdleGlobal     int    `mapstructure:"max_idle_global"`      //最大空闲数
	MaxIdleTimeOut    string `mapstructure:"max_idle_time_out"`    //最大空闲超时
}

// 客户端多路复用
type clientTransport struct {
	Enable        bool `mapstructure:"enable"`         //是否启用多路复用
	MuxConnection int  `mapstructure:"mux_connection"` //连接数
}

//重试机制
type failureRetry struct {
	Enable        bool `mapstructure:"enable"`          //是否启用请求重试机制
	MaxRetryTimes int  `mapstructure:"max_retry_times"` //重试次数
}

//负载均衡
type loadBalancer struct {
	Enable bool `mapstructure:"enable"` //是否启用负载均衡
}

//熔断器
type cbsuite struct {
	Enable bool `mapstructure:"enable"` //是否启用熔断器
}
