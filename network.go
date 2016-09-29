package exp

import "net"

// Conatins IP

type expContainsIp struct {
	key, str string
	cidrnet *net.IPNet
}

func (e expContainsIp) Eval(p Params) bool {
	testIp := net.ParseIP(p.Get(e.key))
	// net.ParseIP(p.Get(e.key))
	// return false
	return e.cidrnet.Contains(testIp)
}


func (e expContainsIp) String() string {
	return sprintf("[%s∋%s]", e.key, e.str)
}

// Contains is an expression that evaluates to true if substr falls within the cidr range
// given example:
//
// 192.168.1.0/24 will match all IPs that fall between
// 192.168.1.1 and 	192.168.1.254
//
// 192.168.1.0/32 will only match 192.168.1.0
func ContainsIp(key, substr string) Exp {
	_, cidrnet, err := net.ParseCIDR(substr)
	if err != nil {
		return nil
	}
	return expContainsIp{key, substr, cidrnet}
}


type expBitwiseContains struct {
	key string
	cidrnet *net.IPNet
}

// type IpParams interface {
// 	Get(string) string
// 	GetIp() net.IP
// }

type IpParams struct {
	Ip []byte
}

func (p IpParams) Get(s string) string {
	return s
}

func (e expBitwiseContains) Eval(p Params) bool {
	// testIp := net.ParseIP(p.Get(e.key))
	// return ((byte(e.cidrnet.IP) & byte(e.cidrnet.Mask)) == (byte(testIp.IP) & byte(e.cidrnet.Mask)))
	//return true


	// ip := IpParams(p)
	// return e.cidrnet.Contains(ip.GetIp())

	mask := e.cidrnet.Mask
	ip := e.cidrnet.IP

	switch p.(type) {
	case IpParams:
		testIp := p.(IpParams).Ip
		for j := 0; j < 4; j++ {
			if (ip[j]&mask[j] != testIp[j]&mask[j]) {
				break
			}
			return true
		}
	}

	// switch p.(type) {
	// case IpParams:
	// 	return e.cidrnet.Contains(p.(IpParams).Ip)
	// }
	return false
}


func (e expBitwiseContains) String() string {
	return sprintf("[%s∋%s]", e.key, e.cidrnet)
}

func bitwiseContains(key, substr string) Exp {
	_, cidrnet, err := net.ParseCIDR(substr)
	if err != nil {
		return nil
	}
	return expBitwiseContains{key, cidrnet}
}
