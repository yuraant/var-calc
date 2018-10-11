package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"var-calc/calc"
)

func main() {
	var calcResp calc.Resp
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("You can enter your expressions in format \"<variable-name> = <arg> [ + <arg> ... ]\"")
	for scanner.Scan() {
		variable, args := strToArg(scanner.Text())
		calcResp = calc.VarsCalc(variable, args)
		if calcResp.Print {
			for k, v := range calcResp.Variables {
				fmt.Printf("===> %s = %d\n", k, v)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

// strToArg split console string to variable and agruments
func strToArg(s string) (string, []string) {
	result := strings.SplitN(s, "=", 2)
	variable := strings.TrimSpace(result[0])
	args := strings.Split(result[1], "+")
	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}
	return variable, args
}
