package kitex_conf

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
	Rpc        rpc        `yaml:"rpc"`         //服务ip配置
	Polaris    polaris    `yaml:"polaris"`     //北极星注册中心配置
	Jaeger     jaeger     `yaml:"jaeger"`      //链路配置
	Transport  transport  `yaml:"transport"`   //多路复用配置
	Limit      limit      `yaml:"limit"`       //限流器
	StatsLevel statsLevel `yaml:"stats_level"` //埋点策略&埋点粒度
}

// Service 服务名称配置
type Service struct {
	NameSpace  string `yaml:"namespace"`                     //服务空间名称
	ServerName string `yaml:"server_name" yaml:"serverName"` //服务名称
	ClientName string `yaml:"client_name" yaml:"clientName"` //客户端名称
	Version    string `yaml:"version"`                       //版本信息
}

// 服务地址端口配置
type rpc struct {
	Enable  bool   `yaml:"enable" `                 //是否启用rpc自定义配置
	Address string `yaml:"address"`                 //地址
	Network string `yaml:"net_work" yaml:"netWork"` //连接方式 (tcp udp)
}

// 注册中心配置
type polaris struct {
	Enable bool `yaml:"enable"` //是否启用注册中心，默认开启
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

//限流器配置
type limit struct {
	Enable         bool `yaml:"enable"`                                //是否启用多路复用
	MaxConnections int  `yaml:"max_connections" yaml:"maxConnections"` // 最大连接数
	MaxQPS         int  `yaml:"max_qps" yaml:"maxQps"`                 //最大qps
}

// **********************************公共对象*******************************

type statsLevel struct {
	LevelDisabled bool `yaml:"level_disabled" yaml:"levelDisabled"`
	LevelBase     bool `yaml:"level_base" yaml:"levelBase"`
	LevelDetailed bool `yaml:"level_detailed" yaml:"levelDetailed"`
}

// **********************************客户端对象******************************
//客户端配置
type Client struct {
	TimeoutControl timeOutControl `yaml:"timeout_control" yaml:"timeoutControl"` //超时控制
	ConnectionType connectionType `yaml:"connection_type" yaml:"connectionType"` // 连接类型
	FailureRetry   failureRetry   `yaml:"failure_retry" yaml:"failureRetry"`     //请求重试
	LoadBalancer   loadBalancer   `yaml:"load_balancer" yaml:"loadBalancer"`     //负载均衡
	CBSuite        cbsuite        `yaml:"cbsuite"`                               //熔断器
	StatsLevel     statsLevel     `yaml:"stats_level" yaml:"statsLevel"`         //埋点策略&埋点粒度
}

//超时控制
type timeOutControl struct {
	RpcTimeout     rpcTimeout     `yaml:"rpc_timeout" yaml:"rpcTimeout"`
	ConnectTimeOut connectTimeOut `yaml:"connect_time_out" yaml:"connectTimeOut"`
}

//连接类型（长链接 短链接）
type connectionType struct {
	ShortConnection shortConnection `yaml:"short_connection" yaml:"shortConnection"` //短链接
	LongConnection  longConnection  `yaml:"long_connection" yaml:"longConnection"`   //长链接
	ClientTransport clientTransport `yaml:"transport"`                               //客户端多路复用

}

//rpc超时控制
type rpcTimeout struct {
	Enable  bool   `yaml:"enable"`                  //是否启用rpc超时
	Timeout string `yaml:"time_out" yaml:"timeOut"` //超时时间 （默认 1s 单位："ns", "us" (or "µs"), "ms", "s", "m", "h"）
}

//connect超时控制
type connectTimeOut struct {
	Enable  bool   `yaml:"enable"`                  //是否启用rpc超时
	TimeOut string `yaml:"time_out" yaml:"timeOut"` //连接超时 （默认：50ms）
}

//短链接
type shortConnection struct {
	Enable bool `yaml:"enable"` //是否启用短链接
}

//长链接
type longConnection struct {
	Enable            bool   `yaml:"enable"`                                        //是否启用长链接
	MaxIdlePerAddress int    `yaml:"max_idle_per_address" yaml:"maxIdlePerAddress"` //最大空闲地址
	MinIdlePerAddress int    `yaml:"min_idle_per_address" yaml:"minIdlePerAddress"` //最小空闲地址
	MaxIdleGlobal     int    `yaml:"max_idle_global" yaml:"maxIdleGlobal"`          //最大空闲数
	MaxIdleTimeOut    string `yaml:"max_idle_time_out" yaml:"maxIdleTimeOut"`       //最大空闲超时
}

// 客户端多路复用
type clientTransport struct {
	Enable        bool `yaml:"enable"`                              //是否启用多路复用
	MuxConnection int  `yaml:"mux_connection" yaml:"muxConnection"` //连接数
}

//重试机制
type failureRetry struct {
	Enable        bool `yaml:"enable"`                               //是否启用请求重试机制
	MaxRetryTimes int  `yaml:"max_retry_times" yaml:"maxRetryTimes"` //重试次数
}

//负载均衡
type loadBalancer struct {
	Enable bool `yaml:"enable"` //是否启用负载均衡
}

//熔断器
type cbsuite struct {
	Enable bool `yaml:"enable"` //是否启用熔断器
}
