# var-calc
This is small summing calculator for variables which can determine value of variable immediately if expression contains numbers or already determined variables.
In case expression contains some variables that cannot be determine right now, such variables will be stored in memory and will be printed as soon as they can be determined

## Build and Run
Application requires GO being installed on your machine Clone current repo. then run commands in console.
```
# cd [to repo folder] 
# go build
# ./var-calc
```
## Example of output
```
HOST:var-calc user$ ./var-calc
You can enter your expressions in format "<variable-name> = <arg> [ + <arg> ... ]"
a = 3
===> a = 3
b = 2 + 4
===> b = 6
c = a
===> c = 3
d = b + 7
===> d = 13
e = b + c + d
===> e = 22
f = g
g = 1 + 2 + 3
===> g = 6
===> f = 6
h = i + e + j
j = k + 1
i = e + 5
===> i = 27
k = a + b
===> k = 9
===> j = 10
===> h = 59
l = k + k
===> l = 18
````
## Use as library
Calc can be used as a library and be imported into your project 
````
https://github.com/yuraant/var-calc/calc
````
```Golang
import (
	
	"var-calc/calc"
  	OR
  	"github.com/yuraant/var-calc/calc"
)
```
You need to use func calc.VarsCalc
This function takes 2 parameters "VariableName" `string` and "Arguments" `[]string` 
Returns struct
```Golang
type Resp struct {
	Print     bool
	Variables map[string]int
}
```
`Print` is marker: shold we print variable or not
`Variables` is map of variables: Name `string`and value `ìnt`
