package hex

import (
	"fmt"
	"strconv"
)

func parseInt(str string) (int64, error) {
	return strconv.ParseInt(str, 0, 64)
}

func printBinary(num int64) {
	printf("%b\n", num)
}

func printOctal(num int64) {
	printf("%o\n", num)

}

func printDecimal(num int64) {
	printf("%d\n", num)
}

func printHex(num int64) {
	printf("%x\n", num)
}

func printUpperHex(num int64) {
	printf("%X\n", num)
}

func printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
