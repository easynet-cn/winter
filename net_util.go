package winter

import (
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"net/http"
)

func ExternalIP() net.IP {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip
		}
	}
	return nil
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}

func LocalIP() string {
	ip := "127.0.0.1"

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Printf("获取本地IP发生异常 : %v\n", err)
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()

				break
			}
		}
	}

	return ip
}

func ReadFileFromUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if bytes, err := io.ReadAll(resp.Body); err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http status code is %d", resp.StatusCode)
	} else {
		return bytes, nil
	}
}

func ReadFileFromUrlAsBase64(url string) (string, error) {
	if bytes, err := ReadFileFromUrl(url); err != nil {
		return "", err
	} else {
		return base64.StdEncoding.EncodeToString(bytes), nil
	}
}
