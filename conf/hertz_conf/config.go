package hertz_conf

// Logger 日志配置
type Logger struct {
	Enable     bool   `json:"enable" `                       //是否启用自定义日志配置
	Filename   string `json:"file_name" json:"fileName"`     //路径
	MaxSize    int    `json:"max_size" json:"maxSize"`       //日志的最大大小（M）
	MaxBackups int    `json:"max_backups" json:"maxBackups"` //日志的最大保存数量
	MaxAge     int    `json:"max_age" json:"maxAge"`         //日志文件存储最大天数
	Compress   bool   `json:"compress"`                      //是否执行压缩
	LocalTime  bool   `json:"local_time" json:"localTime"`   //是否使用格式化时间辍
}

// Server 服务端配置
type Server struct {
	Http      http      `json:"http"`      //服务ip配置
	Polaris   polaris   `json:"polaris"`   //北极星注册中心配置
	Auth      auth      `json:"auth"`      // auth 身份认证配置
	Cors      cors      `json:"cors"`      //cors 配置
	Recovery  recovery  `json:"recovery"`  // recovery 配置
	Gzip      gzip      `json:"gzip"`      // 压缩配置
	I18n      i18n      `json:"i18n"`      // 国际化配置
	Swag      swag      `json:"swag"`      // swag文档
	Jaeger    jaeger    `json:"jaeger"`    //链路配置
	Transport transport `json:"transport"` //多路复用配置
}

// Service 服务名称配置
type Service struct {
	NameSpace  string `json:"namespace"`                     //服务空间名称
	ServerName string `json:"server_name" json:"serverName"` //服务名称
	ClientName string `json:"client_name" json:"clientName"` //客户端名称
	Version    string `json:"version"`                       //版本信息
}

// 服务地址端口配置
type http struct {
	Enable       bool   `json:"enable" `        //是否启用http自定义配置
	Address      string `json:"address"`        //地址
	ExitWaitTime int    `json:"exit_wait_time"` // 配置
}

// 注册中心配置
type polaris struct {
	Enable  bool   `json:"enable"` //是否启用注册中心，默认开启
	Network string `json:"network"`
	Address string `json:"address"`
}

//链路追踪配置
type jaeger struct {
	Enable   bool   `json:"enable"`   //是否启用链路追踪
	Endpoint string `json:"endpoint"` //地址
}

//多路复用配置
type transport struct {
	Enable bool `json:"enable"` //是否启用多路复用
}

// auth 配置
type auth struct {
	Enable bool   `json:"enable"` //是否启用auth配置
	AK     string `json:"ak"`
	SK     string `json:"sk"`
}

// cors配置
type cors struct {
	Enable bool `json:"enable"` //是否启用cors配置
}

// recovery 配置
type recovery struct {
	Enable bool `json:"enable"` //是否启用 recovery 配置
}

// gzip 压缩配置
type gzip struct {
	Enable             bool     `json:"enable"`              //是否启用 gzip 配置
	BestCompression    bool     `json:"best_compression"`    //提供最佳的文件压缩率
	BestSpeed          bool     `json:"best_speed"`          //提供了最佳的压缩速度
	DefaultCompression bool     `json:"default_compression"` //默认压缩率
	NoCompression      bool     `json:"no_compression"`      //不进行压缩
	Excluded           excluded `json:"excluded"`            // 设置不需要压缩的方式
}

type excluded struct {
	Enable              bool                `json:"enable"` //是否启用 excluded 配置
	ExcludedExtensions  excludedExtensions  `json:"excluded_extensions"`
	ExcludedPaths       excludedPaths       `json:"excluded_paths"`
	ExcludedPathRegexes excludedPathRegexes `json:"excluded_path_regexes"`
}

// 设置不需要 gzip 压缩的文件后缀
type excludedExtensions struct {
	Enable     bool   `json:"enable"`     //是否启用 excluded 配置
	Extensions string `json:"extensions"` // 文件后缀 数组用,连接 eg:".pdf", ".mp4"
}

// 设置不需要进行 gzip 压缩的文件路径
type excludedPaths struct {
	Enable bool   `json:"enable"` //是否启用 excludedPaths 配置
	Paths  string `json:"paths"`  //文件路径 数组用,连接 eg:/api/
}

// 设置自定义的正则表达式来过滤掉不需要 gzip 压缩的文件
type excludedPathRegexes struct {
	Enable  bool   `json:"enable"`  //是否启用 excludedPathRegexes 配置
	Regexes string `json:"regexes"` //文件路径 数组用,连接  eg: /api.*
}

// 设置国际化
type i18n struct {
	Enable bool `json:"enable"` //是否启用 i18n 配置
}

// 设置 swag
type swag struct {
	Enable bool `json:"enable"` //是否启用 swag 配置
}

// **********************************公共对象*******************************
type statsLevel struct {
	LevelDisabled bool `json:"level_disabled"`
	LevelBase     bool `json:"level_base"`
	LevelDetailed bool `json:"level_detailed"`
}

// Client **********************************客户端对象******************************
//客户端配置
type Client struct {
	TimeoutControl timeOutControl `json:"timeout_control" json:"timeoutControl"` //超时控制
	ConnectionType connectionType `json:"connection_type" json:"connectionType"` // 连接类型
	FailureRetry   failureRetry   `json:"failure_retry" json:"failureRetry"`     //请求重试
	LoadBalancer   loadBalancer   `json:"load_balancer" json:"loadBalancer"`     //负载均衡
	CBSuite        cbsuite        `json:"cbsuite"`                               //熔断器
	StatsLevel     statsLevel     `json:"stats_level"`                           //埋点策略&埋点粒度
}

//超时控制
type timeOutControl struct {
	RpcTimeout     rpcTimeout     `json:"rpc_timeout" json:"rpcTimeout"`
	ConnectTimeOut connectTimeOut `json:"connect_time_out" json:"connectTimeOut"`
}

//连接类型（长链接 短链接）
type connectionType struct {
	ShortConnection shortConnection `json:"short_connection" json:"shortConnection"` //短链接
	LongConnection  longConnection  `json:"long_connection" json:"longConnection"`   //长链接
	ClientTransport clientTransport `json:"transport"`                               //客户端多路复用

}

//rpc超时控制
type rpcTimeout struct {
	Enable  bool   `json:"enable"`                  //是否启用rpc超时
	Timeout string `json:"time_out" json:"timeout"` //超时时间 （默认 1s 单位："ns", "us" (or "µs"), "ms", "s", "m", "h"）
}

//connect超时控制
type connectTimeOut struct {
	Enable  bool   `json:"enable"`                  //是否启用rpc超时
	TimeOut string `json:"time_out" json:"timeOut"` //连接超时 （默认：50ms）
}

//短链接
type shortConnection struct {
	Enable bool `json:"enable"` //是否启用短链接
}

//长链接
type longConnection struct {
	Enable            bool   `json:"enable"`                                        //是否启用长链接
	MaxIdlePerAddress int    `json:"max_idle_per_address" json:"maxIdlePerAddress"` //最大空闲地址
	MinIdlePerAddress int    `json:"min_idle_per_address" json:"minIdlePerAddress"` //最小空闲地址
	MaxIdleGlobal     int    `json:"max_idle_global" json:"maxIdleGlobal"`          //最大空闲数
	MaxIdleTimeOut    string `json:"max_idle_time_out" json:"maxIdleTimeOut"`       //最大空闲超时
}

// 客户端多路复用
type clientTransport struct {
	Enable        bool `json:"enable"`                              //是否启用多路复用
	MuxConnection int  `json:"mux_connection" json:"muxConnection"` //连接数
}

//重试机制
type failureRetry struct {
	Enable        bool `json:"enable"`                                 //是否启用请求重试机制
	MaxRetryTimes int  `json:"max_retry_times" json:"max_retry_times"` //重试次数
}

//负载均衡
type loadBalancer struct {
	Enable bool `json:"enable"` //是否启用负载均衡
}

//熔断器
type cbsuite struct {
	Enable bool `json:"enable"` //是否启用熔断器
}
