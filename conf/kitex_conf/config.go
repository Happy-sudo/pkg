package kitex_conf

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
	Rpc        rpc        `mapstructure:"rpc"`         //服务ip配置
	Polaris    polaris    `mapstructure:"polaris"`     //北极星注册中心配置
	Jaeger     jaeger     `mapstructure:"jaeger"`      //链路配置
	Transport  transport  `mapstructure:"transport"`   //多路复用配置
	Limit      limit      `mapstructure:"limit"`       //限流器
	StatsLevel statsLevel `mapstructure:"stats_level"` //埋点策略&埋点粒度
}

// Service 服务名称配置
type Service struct {
	NameSpace  string `mapstructure:"namespace"`                             //服务空间名称
	ServerName string `mapstructure:"server_name" mapstructure:"serverName"` //服务名称
	ClientName string `mapstructure:"client_name" mapstructure:"clientName"` //客户端名称
	Version    string `mapstructure:"version"`                               //版本信息
}

// 服务地址端口配置
type rpc struct {
	Enable  bool   `mapstructure:"enable" `                         //是否启用rpc自定义配置
	Address string `mapstructure:"address"`                         //地址
	Network string `mapstructure:"net_work" mapstructure:"netWork"` //连接方式 (tcp udp)
}

// 注册中心配置
type polaris struct {
	Enable bool `mapstructure:"enable"` //是否启用注册中心，默认开启
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

//限流器配置
type limit struct {
	Enable         bool `mapstructure:"enable"`                                        //是否启用多路复用
	MaxConnections int  `mapstructure:"max_connections" mapstructure:"maxConnections"` // 最大连接数
	MaxQPS         int  `mapstructure:"max_qps" mapstructure:"maxQps"`                 //最大qps
}

// **********************************公共对象*******************************

type statsLevel struct {
	LevelDisabled bool `mapstructure:"level_disabled" mapstructure:"levelDisabled"`
	LevelBase     bool `mapstructure:"level_base" mapstructure:"levelBase"`
	LevelDetailed bool `mapstructure:"level_detailed" mapstructure:"levelDetailed"`
}

// **********************************客户端对象******************************
//客户端配置
type Client struct {
	TimeoutControl timeOutControl `mapstructure:"timeout_control" mapstructure:"timeoutControl"` //超时控制
	ConnectionType connectionType `mapstructure:"connection_type" mapstructure:"connectionType"` // 连接类型
	FailureRetry   failureRetry   `mapstructure:"failure_retry" mapstructure:"failureRetry"`     //请求重试
	LoadBalancer   loadBalancer   `mapstructure:"load_balancer" mapstructure:"loadBalancer"`     //负载均衡
	CBSuite        cbsuite        `mapstructure:"cbsuite"`                                       //熔断器
	StatsLevel     statsLevel     `mapstructure:"stats_level" mapstructure:"statsLevel"`         //埋点策略&埋点粒度
}

//超时控制
type timeOutControl struct {
	RpcTimeout     rpcTimeout     `mapstructure:"rpc_timeout" mapstructure:"rpcTimeout"`
	ConnectTimeOut connectTimeOut `mapstructure:"connect_time_out" mapstructure:"connectTimeOut"`
}

//连接类型（长链接 短链接）
type connectionType struct {
	ShortConnection shortConnection `mapstructure:"short_connection" mapstructure:"shortConnection"` //短链接
	LongConnection  longConnection  `mapstructure:"long_connection" mapstructure:"longConnection"`   //长链接
	ClientTransport clientTransport `mapstructure:"transport"`                                       //客户端多路复用

}

//rpc超时控制
type rpcTimeout struct {
	Enable  bool   `mapstructure:"enable"`                          //是否启用rpc超时
	Timeout string `mapstructure:"time_out" mapstructure:"timeOut"` //超时时间 （默认 1s 单位："ns", "us" (or "µs"), "ms", "s", "m", "h"）
}

//connect超时控制
type connectTimeOut struct {
	Enable  bool   `mapstructure:"enable"`                          //是否启用rpc超时
	TimeOut string `mapstructure:"time_out" mapstructure:"timeOut"` //连接超时 （默认：50ms）
}

//短链接
type shortConnection struct {
	Enable bool `mapstructure:"enable"` //是否启用短链接
}

//长链接
type longConnection struct {
	Enable            bool   `mapstructure:"enable"`                                                //是否启用长链接
	MaxIdlePerAddress int    `mapstructure:"max_idle_per_address" mapstructure:"maxIdlePerAddress"` //最大空闲地址
	MinIdlePerAddress int    `mapstructure:"min_idle_per_address" mapstructure:"minIdlePerAddress"` //最小空闲地址
	MaxIdleGlobal     int    `mapstructure:"max_idle_global" mapstructure:"maxIdleGlobal"`          //最大空闲数
	MaxIdleTimeOut    string `mapstructure:"max_idle_time_out" mapstructure:"maxIdleTimeOut"`       //最大空闲超时
}

// 客户端多路复用
type clientTransport struct {
	Enable        bool `mapstructure:"enable"`                                      //是否启用多路复用
	MuxConnection int  `mapstructure:"mux_connection" mapstructure:"muxConnection"` //连接数
}

//重试机制
type failureRetry struct {
	Enable        bool `mapstructure:"enable"`                                       //是否启用请求重试机制
	MaxRetryTimes int  `mapstructure:"max_retry_times" mapstructure:"maxRetryTimes"` //重试次数
}

//负载均衡
type loadBalancer struct {
	Enable bool `mapstructure:"enable"` //是否启用负载均衡
}

//熔断器
type cbsuite struct {
	Enable bool `mapstructure:"enable"` //是否启用熔断器
}
