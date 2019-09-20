package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"net"

)

func init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	termbox.SetCursor(0, 0)
	termbox.HideCursor()
}


//func getMacAddrs() (macAddrs []string) {
//	netInterfaces, err := net.Interfaces()
//	if err != nil {
//		fmt.Printf("fail to get net interfaces: %v", err)
//		return macAddrs
//	}
//
//	for _, netInterface := range netInterfaces {
//		macAddr := netInterface.HardwareAddr.String()
//		if len(macAddr) == 0 {
//			continue
//		}
//		macAddrs = append(macAddrs, macAddr)
//	}
//
//
//	return macAddrs
//}
//
//func getIPs() (name []string,ips []string,macAddrs []string) {
//	netInterfaces, err := net.Interfaces()
//	if err != nil {
//		fmt.Println("net.Interfaces failed, err:", err.Error())
//		return name,ips,macAddrs
//	}
//
//	for i := 0; i < len(netInterfaces); i++ {
//		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
//			addrs, _ := netInterfaces[i].Addrs()
//
//			macAddr := netInterfaces[i].HardwareAddr.String()
//			if len(macAddr) == 0 {
//				continue
//			}
//			macAddrs = append(macAddrs, macAddr)
//			if addrs != nil {
//				name = append(name,netInterfaces[i].Name)
//			}
//
//			for _, address := range addrs {
//				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
//					if ipnet.IP.To4() != nil {
//						ips = append(ips, ipnet.IP.String())
//					}
//				}
//			}
//		}
//	}
//	return name,ips,macAddrs
//}
type EthModel struct{
	Name 	string
	Ip 		string
	Mac 	string
}

func get() {
	eth := EthModel{}
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			ip, _ := netInterfaces[i].Addrs()

			for _, address := range ip {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						eth.Name = netInterfaces[i].Name
						eth.Mac = netInterfaces[i].HardwareAddr.String()
						eth.Ip = ipnet.IP.String()

						//mac := netInterfaces[i].HardwareAddr.String()
						//name := netInterfaces[i].Name
						//fmt.Printf("name:%s ip:%s mac:%s\n",name,ipnet.IP.String(),mac)
						fmt.Println(eth)
					}
				}
			}
		}
	}

}


func pause() {
	fmt.Println("请按任意键继续...")
Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			break Loop
		}
	}
}


func main() {
	//name,ip,mac := getIPs()
	//fmt.Printf("%q\n",name)
	//fmt.Printf("%q\n",ip)
	//fmt.Printf("%q\n",mac)
	get()

	//fmt.Printf("mac addrs: %q\n", getMacAddrs())
	pause()
}
