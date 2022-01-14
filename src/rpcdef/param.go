package rpcdef

const (
	TCP         = "tcp"
	TestPort    = "9973"
	LocalHost   = "localhost"
	DefaultAddr = LocalHost + ":" + TestPort
)

func GetAddr(ip, port string) string {
	return ip + ":" + port
}
