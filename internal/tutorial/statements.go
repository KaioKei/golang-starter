package tutorial

import (
	"log"
)

func Statements() {
	log.Printf("Golang statements")
	ifStatement()
	switchStatement()
	forStatement()
}

func ifStatement() {
	log.Printf("If statement")
	const ifCondition bool = true
	const elseIfCondition bool = true

	// basic statement
	if ifCondition {
		log.Printf("Printed")
	} else if elseIfCondition {
		log.Printf("Also printed")
	} else {
		log.Printf("Not printed")
	}

	// short statement
	// allows var instantiation & condition on the former var in one line
	// this kind of vars only exist during if statement
	if shortCondition := true; shortCondition {
		log.Printf("Printed")
	} else {
		log.Printf("Not printed")
	}
}

func switchStatement() {
	log.Printf("Switch statement")
	dayNumber := 6

	// basic statement
	switch dayNumber {
	case 1:
		log.Println("Monday")
	case 2:
		log.Println("Tuesday")
	case 3:
		log.Println("Wednesday")
	case 4:
		log.Println("Thursday")
	case 5:
		log.Println("Friday")
	case 6:
		log.Println("Saturday, yay !")
	case 7:
		log.Println("Sunday, yay !")
	default:
		log.Println("Does not exist")
	}

	// short statement
	// allows var instantiation & condition on the former var in one line
	// this kind of vars only exist during switch statement
	switch diceNumber := 6; diceNumber {
	case 1:
		log.Println("fail")
	case 2, 3, 4, 5:
		// allow multiple cases with same code
		log.Println("pass")
	case 6:
		log.Println("Success, yay !")
	default:
		log.Println("Does not exist on this dice")
	}

	// no statement
	// faster than if statements for static comparisons
	test := 1
	switch {
	case test == 1:
		log.Println("Printed")
	case test == 2:
		// allow multiple cases with same code
		log.Println("Not Printed")
	case test == 3:
		log.Println("Not Printed")
	default:
		log.Println("test is not in {1,2,3}")
	}
}

func forStatement() {
	log.Printf("For statement")

	//for initialization; condition; increment {
	//	// loop body
	//}

	//short statement
	for i := 0; i < 10; i++ {
		log.Printf("%d\n", i)
	}

	// statement with existing index
	i := 0
	for ; i < 10; i++ {
		log.Printf("%d\n", i)
	}

	// no increment
	i = 0
	for i < 10 {
		log.Printf("%d\n", i)
		break
	}

	//while loop
	condition := true
	j := 0
	for condition {
		j++
		if j == 10 {
			condition = false
		}
		log.Printf("%d\n", i)
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
		log.Printf("%d ", odd)
	}
}
