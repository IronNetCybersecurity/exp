package exp

import (
	"net"
    "testing"
    "os"
    _ "log"
    "sync"
)

type Expression struct {
	Id             int64
	AnalyticType   string
	ExpirationDate float64
	Expression     Exp
}

var ipMap = Map{
	"src_ip": "192.168.1.61",	
}

var RulesExp = make([]*Expression, 0, 1)
var IpExp = make([]*Expression, 0, 1)

var IpList = [][]byte{}
var Ipinput = []IpParams{}
var input = []Map{}

var wgString sync.WaitGroup
var wgByte sync.WaitGroup


func dummyGet(s string) string{
	return s
}

func TestMain(m *testing.M) {
	id := 1
	for i := 1; i < 100000; i++ {
		nextRule := ContainsIp("src_ip", "10.10.10.0/24")
		if nextRule == nil {
			// log.Println("fail")
		}
		exObj := &Expression{
			Expression: nextRule,
			Id: int64(id),
		}
		RulesExp = append(RulesExp, exObj)

	nextIp := bitwiseContains("src_ip", "10.10.10.0/24")
	exObj = &Expression{
		Expression: nextIp,
		Id: int64(id),
	}
	IpExp = append(IpExp, exObj)

		id++


		ip1 := []byte{10,10,10,10}
		IpList = append(IpList, ip1)
	}

	nextRule := ContainsIp("src_ip", "100.10.10.0/24")
	exObj := &Expression{
		Expression: nextRule,
		Id: int64(id),
	}
	RulesExp = append(RulesExp, exObj)


	nextIp := bitwiseContains("src_ip", "100.10.10.0/24")
	exObj = &Expression{
		Expression: nextIp,
		Id: int64(id),
	}
	IpExp = append(IpExp, exObj)


		ip2 := []byte{100,10,10,10}
		IpList = append(IpList, ip2)

	laps := 40000
	for i := 0; i < laps; i++ {
		input = append(input, map[string]string{
			"src_ip": "100.10.10.10",
		})
		bIp := net.ParseIP("100.10.10.10")
		Ipinput = append(Ipinput, IpParams{
			Ip: bIp,
		})
	}

	os.Exit(m.Run())
}

func TestContainsIp(t *testing.T) {
	for key, value := range map[string]string{
		"src_ip": "192.168.1.0/24",
	} {
		if !ContainsIp(key, value).Eval(ipMap) {
			t.Errorf("Match(%q, %q) should evaluate to true", key, value)
		}
	}
}

func BenchmarkContainsIp_bit(b *testing.B) {
	// laps := 40000
	bits := [][]byte{}
	// ip1 := []byte{10,10,10,10}
	ip2 := []byte{100,10,10,10}
	mask := []byte{255,255,255,0}

	for i :=0; i<400; i++ {
		bits = append(bits, ip2)
	}
		// bits = append(bits, ip2)
	for _, r := range bits {
		for _, ip := range IpList {
			for j := 0; j < 4; j++ {
				if (r[j]&mask[j] != ip[j]&mask[j]) {
					 // log.Println("break")
					break
				}
				if j == 3 {
					 // log.Println("match", j, r, ip)
				}
			}
		}
	}
}

func BenchmarkContainsIp_aaaparse(b *testing.B) {
	// listNum := 99999

	// input = append(input, &map[string]string{
	// 		"SrcIp": "100.10.10.10",
	// })

	// log.Println("DEBUG len of rulesexp", len(RulesExp))
	// log.Println("start")
	for _, m := range input{
		// for _, r := range RulesExp {
		// 	if r.Expression.Eval(m) {
		// 		// log.Println("Match", i)
		// 		break
		// 	}
		// }
		wgString.Add(1)
		go evalString(m)
	}
	wgString.Wait()
	// log.Println("end")
}

func BenchmarkContainsIp_aaafast(b *testing.B) {
	// listNum := 99999

	// input = append(input, &map[string]string{
	// 		"SrcIp": "100.10.10.10",
	// })

	// log.Println("DEBUG len of rulesexp", len(RulesExp))
	// log.Println("start")
	for _, m := range Ipinput{
		// for _, r := range IpExp {
		// 	if r.Expression.Eval(m) {
		// 		// log.Println("Match", i)
		// 		break
		// 	}
		// }
		wgByte.Add(1)
		go evalByte(m)
	}
	wgByte.Wait()
	// log.Println("end")
}


func evalString(m Map) {
	defer wgString.Done()
	for _, r := range RulesExp {
		if r.Expression.Eval(m) {
			// log.Println("Match", i)
			break
		}
	}
}

func evalByte(m IpParams) {
	defer wgByte.Done()
	for _, r := range IpExp {
		if r.Expression.Eval(m) {
			// log.Println("Match", i)
			break
		}
	}
}
