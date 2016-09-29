package exp

import (
    "testing"
    "os"
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
var input = []Map{}

var wgString sync.WaitGroup
const ruleCnt = 100000
const ipCount = 40000

func TestMain(m *testing.M) {
	for i := 1; i < ruleCnt; i++ {
		nextRule := ContainsIp("src_ip", "10.10.10.0/24")
		exObj := &Expression{
			Expression: nextRule,
		}
		RulesExp = append(RulesExp, exObj)
	}

	nextRule := ContainsIp("src_ip", "100.10.10.0/24")
	exObj := &Expression{
		Expression: nextRule,
	}
	RulesExp = append(RulesExp, exObj)

	for i := 0; i < ipCount; i++ {
		input = append(input, map[string]string{
			"src_ip": "100.10.10.10",
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

func BenchmarkContainsIp_100kIps(b *testing.B) {
	m := map[string]string{
		"src_ip": "100.10.10.10",
	}
	for i := 0; i < ipCount; i++ {
		wgString.Add(1)
		go evalString(m)
	}
	wgString.Wait()
}

func evalString(m Map) {
	defer wgString.Done()
	for _, r := range RulesExp {
		if r.Expression.Eval(m) {
			break
		}
	}
}
