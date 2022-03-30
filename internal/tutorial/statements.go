package tutorial

import (
	"fmt"
)

func statements() {
	fmt.Printf("Golang statements")
	ifStatement()
	switchStatement()
	forStatement()
}

func ifStatement() {
	fmt.Printf("If statement")
	const ifCondition bool = true
	const elseIfCondition bool = true

	// basic statement
	if ifCondition {
		fmt.Printf("Printed")
	} else if elseIfCondition {
		fmt.Printf("Also printed")
	} else {
		fmt.Printf("Not printed")
	}

	// short statement
	// allows var instantiation & condition on the former var in one line
	// this kind of vars only exist during if statement
	if shortCondition := true; shortCondition {
		fmt.Printf("Printed")
	} else {
		fmt.Printf("Not printed")
	}
}

func switchStatement() {
	fmt.Printf("Switch statement")
	dayNumber := 6

	// basic statement
	switch dayNumber {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday, yay !")
	case 7:
		fmt.Println("Sunday, yay !")
	default:
		fmt.Println("Does not exist")
	}

	// short statement
	// allows var instantiation & condition on the former var in one line
	// this kind of vars only exist during switch statement
	switch diceNumber := 6; diceNumber {
	case 1:
		fmt.Println("fail")
	case 2, 3, 4, 5:
		// allow multiple cases with same code
		fmt.Println("pass")
	case 6:
		fmt.Println("Success, yay !")
	default:
		fmt.Println("Does not exist on this dice")
	}

	// no statement
	// faster than if statements for static comparisons
	test := 1
	switch {
	case test == 1:
		fmt.Println("Printed")
	case test == 2:
		// allow multiple cases with same code
		fmt.Println("Not Printed")
	case test == 3:
		fmt.Println("Not Printed")
	default:
		fmt.Println("test is not in {1,2,3}")
	}
}

func forStatement() {
	fmt.Printf("For statement")

	//for initialization; condition; increment {
	//	// loop body
	//}

	// short statement
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	// statement with existing index
	i := 0
	for ; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	// no increment
	i = 0
	for i < 10 {
		fmt.Printf("%d\n", i)
	}

	// while loop
	condition := true
	j := 0
	for condition {
		j++
		if j == 10 {
			condition = false
		}
		fmt.Printf("%d\n", i)
	}

	// infinite
	for {
		// break it !!!
		break
	}

	// continue
	for odd := 1; odd <= 10; odd++ {
		// will avoid to print fair numbers
		if odd%2 == 0 {
			continue
		}
		fmt.Printf("%d ", odd)
	}
}
