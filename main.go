package main
import ("fmt" 
		"strconv"
		"math")

func main(){
	var input int
	var output int
	fmt.Scan(&input)
	output = swapDigits(input)
	fmt.Println(output)
}


func swapDigits(num int) int{
	boundary := int(math.Pow(2,31))
	if num > boundary-1 || num <= -boundary {
		return 0
	} else {
		var numStr string
		var swapStr string
		var is_negative bool

		numStr = strconv.Itoa(int(num))
		length := len(numStr)
		if num < 0 {
			is_negative = true
			numStr = numStr[1:]
			length--
		}
		
		for i := length - 1; i >= 0; i-- {
			swapStr = swapStr + string(numStr[i])
        }

		swapNum, _ := strconv.Atoi(swapStr)
		if is_negative {
			swapNum = -swapNum
		}
		if swapNum > boundary-1 || swapNum <= -boundary {
			return 0
		}

		return swapNum
	}
}