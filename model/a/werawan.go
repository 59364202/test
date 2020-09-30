package a

import (
	"net"
	"time"
)

type check_current_ip struct {
	Current_ip string `json:"current_ip"` // example:`192.168.12.191`
}

func Werawan_Province() ([]*check_current_ip, error) {

	hostName := "api2.thaiwater.net"
	portNum := "9200"
	seconds := 5
	timeOut := time.Duration(seconds) * time.Second
	result := ""

	conn, err := net.DialTimeout("tcp", hostName+":"+portNum, timeOut)

	if err == nil {
		result = conn.RemoteAddr().String()
	}

	//	fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())
	//	fmt.Printf("Local Address : %s \n", conn.LocalAddr().String())

	rs := make([]*check_current_ip, 0)

	p := &check_current_ip{
		Current_ip: result,
	}

	rs = append(rs, p)

	return rs, err

}
