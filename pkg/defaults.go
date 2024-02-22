package pkg

var (
	DefaultConfigPath = "/etc/avoid/config.yml"
	DefaultDNSKey     = "default"

	DefaultAvoidDNSPort    = 9000
	DefaultAvoidDNSPortENV = "AVOID_DNS_PORT"
	DefaultAvoidDNSAddr    = "0.0.0.0"
	DefaultAvoidDNSAddrENV = "AVOID_DNS_ADDR"

	DefaultEtcdHost    = "etcd"
	DefaultEtcdPort    = 2379
	DefaultEtcdHostENV = "ETCD_HOST"
	DefaultEtcdPortENV = "ETCD_PORT"
)
