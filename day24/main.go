package main

import (
	"day24/converter"
	"day24/mathutil"
	"day24/strutil"
	"day24/userutil"
	"day24/validatorr"
	"fmt"
)

func main() {
	fmt.Println(mathutil.Add(3, 5))
	fmt.Println(mathutil.Multiply(4, 6))
	fmt.Println(mathutil.Max(10, 20))
	fmt.Println(strutil.Reverse("Hello"))
	fmt.Println(strutil.ToUpper("hello"))
	fmt.Println(validatorr.IsEmpty(""))
	fmt.Println(validatorr.IsEmpty("!!!"))
	fmt.Println(validatorr.IsEmail("lpdemoi"))
	fmt.Println(validatorr.IsEmail("apskoj@odim."))
	sum, err := converter.RubToDol(2500, 90)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}
	sum2, err2 := converter.DolToRub(-500, 90)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(sum2)
	}
	user, err := userutil.NewUser("Ivan", 25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user, "succesful")
	}
	fmt.Println(userutil.FormatUser(user))
}
