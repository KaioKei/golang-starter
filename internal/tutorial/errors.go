package tutorial

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// https://go.dev/doc/tutorial/handle-errors
// https://go.dev/blog/error-handling-and-go
// https://www.digitalocean.com/community/tutorials/creating-custom-errors-in-go

// Usually, you can implement a default method to handle errors.
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// raiseSimpleError raise an error based on a simple condition
// if true, raise an error and return an empty string
// if false, do not raise and error and return a simple string
func raiseSimpleError(mybool bool) (string, error) {
	if mybool {
		return "", errors.New("you raised a simple error")
	}
	return "you did not raised a simple error", nil
}

type MyError struct {
	Msg         string
	MoreDetails string
}

// Error method is sufficient to say that MyError implements the Error interface
// Remember, interfaces in go are implicit. A Structure implements an interface when :
//   - the type is used in methods arguments
//   - the methods match all the interface functions
//
// The error interface requires only an Error method.
// Specific error implementations might have additional methods.
// For instance, the net package returns errors of type error, following the usual convention, but some
// of the error implementations have additional methods defined by the net called 'Timeout()' and
// 'Temporary()'. Be always careful about the interface you really implement !!
//
// But there, since the error interface has only one method called Error, all Error methods implements in
// fact the error interface. So the structure in argument has to be considered as a structure for a
// custom error.
// Finally, we can return multiple elements to add more context for this custom error
func (e *MyError) Error() (string, string) {
	// you can set default values for MyError struct and return them
	// to let default empty values (typed defaults), do not edit them here
	// you can still modify them in implementations for context
	e.Msg = "default message"
	e.MoreDetails = "default other details"
	return e.Msg, e.MoreDetails
}

// Here is a simplified example of the built-in error interface
// You have to implement it for custom errors
//   type error interface {
//	   Error() string
//   }

func raiseCustomError(mybool bool) (string, *MyError) {
	if mybool {
		// keeping the default message of the error, but we modify the default value of 'moreDetails' to
		// add the context about the boolean condition that raised this error.
		return "", &MyError{MoreDetails: fmt.Sprintf("because condition was: %t", mybool)}
	}
	return "you did not raised an error", nil
}

// WrappedError is a Wrapped error
// This type is used when you need to raise an error from another known error type
type WrappedError struct {
	Context string
	Err     error
}

func (e *WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", e.Context, e.Err)
}

func raiseWrappedError(mybool bool) (string, *WrappedError) {
	if mybool {
		// will raise a file not found error
		_, err := os.ReadFile("/dev/null")
		if err != nil {
			return "", &WrappedError{Context: "Original error: ", Err: err}
		}

	}
	return "you did not raised an error", nil
}

func Errors() {
	// handle simple errors
	// this error is simply handled with an 'if' statement
	// in this example, no error will be raised, but we can see that if we are turning the condition to
	// 'true', the code will end with a Fatal() message
	raiseCondition := false
	notAnError, err := raiseSimpleError(raiseCondition)
	if err != nil {
		// use log.Fatal to end the code with a log
		// here, it never happens
		log.Fatal("Will never happen")
		// Fatalf let you print more details
		//log.Fatalf("Error raised because condition was: %t", raiseCondition)
	}
	log.Printf("%s", notAnError)

	// raise a custom error
	emptyBecauseError, customError := raiseCustomError(true)
	if customError != nil {
		log.Println(err)
	} else {
		log.Printf("Never printed: %s", emptyBecauseError)
	}

	// raise wrapped error
	emptyBecauseError, wrappedError := raiseWrappedError(true)
	// generic and simple code to handle errors
	// this one will end the execution with Fatal() method
	log.Println(wrappedError)
}
