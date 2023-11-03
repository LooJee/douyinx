package config

type Config struct {
	ClientKey    string
	ClientSecret string
	DouYinHost   string
	CachePrefix  string
}

const (
	DefaultDouYinHost = "https://open.douyin.com"
	CachePrefix       = "douyinx"
)

type Option func(*Config)

func NewConfig(clientKey, clientSecret string, opts ...Option) *Config {
	conf := &Config{
		ClientKey:    clientKey,
		ClientSecret: clientSecret,
		DouYinHost:   DefaultDouYinHost,
		CachePrefix:  CachePrefix,
	}

	for _, opt := range opts {
		opt(conf)
	}

	return conf
}

func WithDouYinHost(host string) Option {
	return func(conf *Config) {
		conf.DouYinHost = host
	}
}

func WithCachePrefix(prefix string) Option {
	return func(conf *Config) {
		conf.CachePrefix = prefix + ":" + CachePrefix
	}
}
