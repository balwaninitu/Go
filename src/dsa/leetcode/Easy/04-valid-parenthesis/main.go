package main

import "fmt"

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	slice := []rune{}

	m := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		//check if open parenthesis
		if v == ')' || v == ']' || v == '}' {
			if len(slice) > 0 && m[v] == slice[len(slice)-1] {
				slice = slice[:len(slice)-1]
				continue
			}
		}
		slice = append(slice, v)
	}

	return len(slice) == 0
}

//Quite simple this one, we need a stack like data structure,
//then we just need to push if brackets are open and pop if a closed is encountered,
//then a check if the pope bracket is in order.
func isValid1(s string) bool {

	var stack []rune

	for _, brac := range s {
		n := len(stack) - 1

		if brac == '}' {
			if n < 0 {
				return false
			}
			current := stack[n]
			stack = stack[:n]
			if current != '{' {
				return false
			}
		} else if brac == ']' {
			if n < 0 {
				return false
			}
			current := stack[n]
			stack = stack[:n]
			if current != '[' {
				return false
			}
		} else if brac == ')' {
			if n < 0 {
				return false
			}
			current := stack[n]
			stack = stack[:n]
			if current != '(' {
				return false
			}
		} else {
			stack = append(stack, brac)
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}

func isValid2(s string) bool {
	//empty string is valid
	if s == "" {
		return true
	}
	//length of string should be even
	if len(s)%2 != 0 {
		return false
	}
	//create map of all 3 pairs
	//opening will be key and closing braces will be value
	m := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	//create an array
	stack := []rune{}

	//checking top value of stack with last value
	//compare in the give string
	for _, value := range s {
		//checking string value with map key
		if value == '(' || value == '[' || value == '{' {
			//if its opening parenthesis add to stack
			//we need to pass value correspondig to map hence append m[value]
			//add open parenthesis to array, it will be key in the map
			stack = append(stack, m[value])
		} else if len(stack) == 0 {
			return false
			//make sure if closing parenthesis match with opening
			//len(stack)-1 --> last element in array
		} else if stack[len(stack)-1] != value {
			return false
			//stack will have length untill last element
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	//return true only when all elements pop from array
	//there should be pair of opening & closing brackets not two open or two close
	return true && len(stack) == 0

}

func main() {

	ans := isValid2("{}")

	fmt.Println(ans)
}
