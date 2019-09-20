# GetIPAndMACAddr
获取已启用网卡的IP和MAC地址

网络上百度了下大致都是如下做法：

··addrs, err := net.InterfaceAddrs()
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
 ··
