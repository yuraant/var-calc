package calc

import (
	"fmt"
	"strconv"
)

//Memory is a structur where we store all variables while app runs
type Memory struct {
	Variables []VariableDetails
}

//VariableDetails is a struct for every variable and its properties
type VariableDetails struct {
	Name       string
	Value      int
	Expression []string
	Printed    bool
}

//Resp type that func CalcVars returns
type Resp struct {
	Print     bool
	Variables map[string]int
}

//Store create storage in memory for variables
var Store = new(Memory)
var emptyStrSlice []string

//VarsCalc returns data for printing
func VarsCalc(variable string, args []string) Resp {
	var response Resp
	response.Variables = make(map[string]int)

	switch {
	case Store.CheckVar(variable) == false && noVarsInAgs(args) == true:
		value := sumIntSlice(convertStrToIntSlice(args))
		response.Variables[variable] = value
		response.Print = true
		Store.AddVar(variable, emptyStrSlice, value, true)
	case Store.CheckVar(variable) == false && noVarsInAgs(args) == false:
		Store.AddVar(variable, args, 0, false)
		vars := getVarFromAgs(args)
		for _, v := range vars {
			if Store.CheckVar(v) == true {
				value, err := Store.GetValue(v)
				if err == nil {
					Store.UpdateExpression(variable, v, value)
					expression, err := Store.GetExpression(variable)
					if err != nil {
						fmt.Println(err)
					}
					if noVarsInAgs(expression) == true {
						value := sumIntSlice(convertStrToIntSlice(expression))
						Store.SetValue(variable, value)
						response.Variables[variable] = value
						response.Print = true
					}
				}

			} else {
				Store.AddVar(v, emptyStrSlice, 0, false)
			}
		}
	case Store.CheckVar(variable) == true && Store.CheckIfValueEmpty(variable) == true && noVarsInAgs(args) == true:
		value := sumIntSlice(convertStrToIntSlice(args))
		response.Variables[variable] = value
		response.Print = true
		Store.SetValue(variable, value)
	case Store.CheckVar(variable) == true && Store.CheckIfValueEmpty(variable) == true && noVarsInAgs(args) == false:
		if Store.CheckIfExprIsEmpty(variable) == true {
			Store.SetExpression(variable, args)
		}
	case Store.CheckVar(variable) == true && Store.CheckIfValueEmpty(variable) == false:
		fmt.Println("variable already exists and cannot be overridden")
	}

	varsForResponse := checkIfDetermined(Store.GetAllNotPtintedVars())
	if varsForResponse.Print == true {
		for k, v := range varsForResponse.Variables {
			response.Variables[k] = v
		}
		response.Print = true
	}

	//fmt.Println(Store.Variables)
	return response
}

//
// Methods for struct Memory
//

//AddVar adds variable with ditails to memeory
func (m *Memory) AddVar(name string, exp []string, value int, print bool) {
	var varDetails VariableDetails
	varDetails.Expression = exp
	varDetails.Value = value
	varDetails.Name = name
	varDetails.Printed = print
	m.Variables = append(m.Variables, varDetails)
}

//CheckVar if variable exist in memory struct
func (m *Memory) CheckVar(v string) bool {
	for _, val := range m.Variables {
		if val.Name == v {
			return true
		}
	}
	return false
}

//GetAllNotPtintedVars rom memory
func (m *Memory) GetAllNotPtintedVars() []string {
	var vars []string
	for _, val := range m.Variables {
		if val.Printed == false {
			vars = append(vars, val.Name)
		}
	}
	return vars
}

//GetValue of variable from memory struct
func (m *Memory) GetValue(v string) (int, error) {
	for _, val := range m.Variables {
		if val.Name == v && val.Printed == true {
			return val.Value, nil
		}
	}
	return 0, fmt.Errorf("Variable %s has not determined yet", v)
}

//CheckIfValueEmpty method for decision
func (m *Memory) CheckIfValueEmpty(v string) bool {
	for _, val := range m.Variables {
		if val.Name == v && val.Printed == false && val.Value == 0 {
			return true
		}
	}
	return false
}

//SetValue determined value
func (m *Memory) SetValue(name string, value int) {
	for key, val := range m.Variables {
		if val.Name == name {
			m.Variables[key].Value = value
			m.Variables[key].Printed = true
		}
	}
}

//CheckIfExprIsEmpty method for decision
func (m *Memory) CheckIfExprIsEmpty(v string) bool {
	for _, val := range m.Variables {
		if val.Name == v && len(val.Expression) > 0 {
			return false
		}
	}
	return true
}

//GetExpression gest property Expression form variable in memory
func (m *Memory) GetExpression(v string) ([]string, error) {
	for _, val := range m.Variables {
		if val.Name == v {
			return val.Expression, nil
		}
	}
	return emptyStrSlice, fmt.Errorf("Variable %s does not exist", v)
}

//SetExpression sets property Expression for variable in memory
func (m *Memory) SetExpression(v string, args []string) {
	for key, val := range m.Variables {
		if val.Name == v {
			m.Variables[key].Expression = args
		}
	}

}

//UpdateExpression in memory for variable
func (m *Memory) UpdateExpression(name, argName string, argValue int) {
	for _, v := range m.Variables {
		if v.Name == name {
			for key, val := range v.Expression {
				if val == argName {
					v.Expression[key] = strconv.Itoa(argValue)
				}
			}
		}
	}
}

//
// Auxiliary functions
//

func noVarsInAgs(args []string) bool {
	for _, v := range args {
		if _, err := strconv.Atoi(v); err != nil {
			return false
		}
	}
	return true
}

func convertStrToIntSlice(args []string) []int {
	var intSlice = []int{}
	for _, i := range args {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, j)
	}
	return intSlice
}

func sumIntSlice(sl []int) int {
	sum := 0
	for i := range sl {
		sum += sl[i]
	}
	return sum
}

func getVarFromAgs(args []string) []string {
	var vars []string
	for _, v := range args {
		if _, err := strconv.Atoi(v); err != nil {
			vars = append(vars, v)
		}
	}
	return vars
}

func checkIfDetermined(vars []string) Resp {
	var response Resp
	var updated bool
	response.Variables = make(map[string]int)
	if len(vars) > 0 {
		for _, varName := range vars {
			expr, err := Store.GetExpression(varName)
			if err != nil {
				fmt.Println(err)
			}
			if len(expr) > 0 {
				varsInExpr := getVarFromAgs(expr)
				for _, varNameInExpresion := range varsInExpr {
					if Store.CheckVar(varNameInExpresion) == true {
						varValue, err := Store.GetValue(varNameInExpresion)
						if err == nil {
							Store.UpdateExpression(varName, varNameInExpresion, varValue)
							updatedExpr, err := Store.GetExpression(varName)
							if err != nil {
								fmt.Println(err)
							}
							if noVarsInAgs(updatedExpr) == true {
								value := sumIntSlice(convertStrToIntSlice(updatedExpr))
								Store.SetValue(varName, value)
								response.Variables[varName] = value
								response.Print = true

							}
							updated = true
						}
					}
				}
			}
		}
	}
	if updated == true {
		result := checkIfDetermined(Store.GetAllNotPtintedVars())
		if result.Print == true {
			for k, v := range result.Variables {
				response.Variables[k] = v
			}
		}
		updated = false
	}
	return response
}
