# GetIPAndMACAddr
获取已启用网卡的IP和MAC地址

网络上大致都是如下做法：

	addrs, err := net.InterfaceAddrs()
		if err != nil {
			return false
		}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				gInnerIP = ipnet.IP.String()
				return true
			}
		}
	}

这种实现方式忽略了一个网卡可用性的问题，导致获取出来的IP可能不一定是想要的。需要通过判断net.FlagUp标志进行确认，排除掉无用的网卡。优化后的实现方式见main.go
 
### 这里贴一个文档里net包的flag方法

	type Flags uint

	const (
		FlagUp           Flags = 1 << iota // 接口在活动状态
		FlagBroadcast                      // 接口支持广播
		FlagLoopback                       // 接口是环回的
		FlagPointToPoint                   // 接口是点对点的
		FlagMulticast                      // 接口支持组播
	)
