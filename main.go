package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var reader *bufio.Reader

	reader = bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input := readLine(reader)
		x := eval(input)
		fmt.Print("$it = ")
		jprint(x)
	}
}

func readLine(r *bufio.Reader) interface{} {
	line, err := r.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			fmt.Println("goodbye :)")
			os.Exit(0)
		}
		panic(err)
	}
	line = strings.TrimSuffix(line, "\n")
	var js interface{}
	err = json.Unmarshal([]byte(line), &js)
	if err != nil {
		fmt.Println(err)
	}
	return js
}

func eval(source interface{}) interface{} {
	switch source.(type) {
	case string:
		return "unknown str LOL!!!@)!))@!"
	case []interface{}:
		return apply(source.([]interface{}))
	default:
		return source
	}
}

func apply(list []interface{}) interface{} {
	i := len(list) - 1
	for i >= 0 {
		if i > 1 {
			list[i-2] = funcallDyad(list[i-1], list[i-2], list[i])
			i -= 2
		} else if i > 0 {
			return funcallMonad(list[i-1], list[i])
		} else {
			return list[i]
		}
	}
	panic("apply fell off")
}

func funcallDyad(op interface{}, lhs interface{}, rhs interface{}) interface{} {
	la, lIsArray := lhs.([]interface{})
	ra, rIsArray := rhs.([]interface{})
	if !lIsArray && !rIsArray {
		return callOp(op, lhs, rhs)
	}
	if !lIsArray && rIsArray {
		var answers = make([]interface{}, len(ra))
		for i := range ra {
			answers[i] = callOp(op, lhs, ra[i])
		}
		return answers
	}
	if lIsArray && rIsArray {
		var answers = make([]interface{}, len(ra))
		for i := range ra {
			answers[i] = callOp(op, la[i], ra[i])
		}
		return answers
	}
	panic("funcallDyad fell off")
}

func funcallMonad(op interface{}, rhs interface{}) interface{} {
	return "foo"
}

func callOp(op interface{}, l, r interface{}) interface{} {
	fname := op.(string)
	switch fname {
	case "+":
		return l.(float64) + r.(float64)
	case "-":
		return l.(float64) - r.(float64)
	case "/":
		return l.(float64) / r.(float64)
	case "*":
		return l.(float64) * r.(float64)
	}
	panic("callOp fell off")
}

func jprint(x ...interface{}) (int, error) {
	for i := range x {
		js, err := json.Marshal(x[i])
		if err != nil {
			panic(err)
		}
		x[i] = string(js)
	}

	return fmt.Println(x...)
}
