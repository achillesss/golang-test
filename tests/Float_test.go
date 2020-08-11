package golangtest

import (
	"fmt"
	"math"
	"testing"
)

func printFloat(f float64, coinCode string) string {
	var vm int
	switch coinCode {
	case "BTC":
		vm = 4
		var _, mf = math.Modf(f)
		var bf = mf * 1e4
		var i, _ = math.Modf(bf)
		var ii = int(i)
		if ii == 0 {
			vm = 0
		} else {
			for ii%10 == 0 {
				vm--
				ii /= 10
			}
		}

	default:
		vm = 2
		var ex int
		var i, frac = math.Modf(f)
		switch frac {
		case 0:
			vm = 0

		default:
			var isSmall bool
			isSmall = i < 10

			// 计算小数点后第一个有效数字前的0的个数
			i, frac = math.Modf(frac)
			for i == 0 {
				ex++
				frac *= 10
				i, frac = math.Modf(frac)
			}

			vm = ex

			switch isSmall {
			case true:
				vm += 3
			default:
				vm += 1
			}

			// 去掉最后面的 0
			i, _ = math.Modf(f * math.Pow10(vm))
			if i != 0 {
				var ii = int(i)
				for ii%10 == 0 {
					vm--
					ii /= 10
				}
			}
		}
	}

	return fmt.Sprintf("vm: %[1]v, %.[1]*[2]f", vm, f)
}

func TestFloat(t *testing.T) {
	println(printFloat(1.00, ""))
	println(printFloat(10.00, ""))
	println(printFloat(12.10, ""))
	println(printFloat(12.00, ""))
	println(printFloat(12.000100123, ""))
	println(printFloat(12.000110123, ""))
	println(printFloat(0.000010123444555213982198, ""))
	println(printFloat(0.000, ""))
	println(printFloat(0.00011, ""))
	println(printFloat(0.00010, ""))
	println(printFloat(3.000, ""))
	println(printFloat(9751.500218575980000000, ""))

	//	println(printFloat(12.000100123, "BTC"))
	//	println(printFloat(12.000110123, "BTC"))
	//	println(printFloat(0.000110123, "BTC"))
	//	println(printFloat(0.000, "BTC"))
	//	println(printFloat(0.0001, "BTC"))
	//	println(printFloat(3.000, "BTC"))
}
