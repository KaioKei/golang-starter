package tutorial

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Goroutine is a thread managed by the Go runtime.
// For a given function 'f(x, y, z)', the evaluation of f, x, y, and z happens in the current goroutine
// and the execution of f happens in the new goroutine.
// However, Goroutines run in the same address space, so access to shared memory must be synchronized.
//
// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values into another goroutine.
// By default, channels 'send' and 'receive' until the other side is ready. This allows goroutines to
// synchronize without explicit locks or condition variables.
// Channels are a typed conduit.
// Create a new channel with make(chan type), so channels are typed by the values they convey.
// You can 'send' and 'receive' values with the channel operator '<-' (the data flows in the direction of
// the arrow.). Send a value into a channel using the 'channel <-' syntax while the '<-channel' syntax
// receives a value from the channel.
// When using channels as function parameters, you should specify if a channel is meant to only send or
// receive values. This specificity increases the type-safety of the program.
// Like maps and slices, channels must be created before use. Check the examples below.
// A sender can close a channel to indicate that no more values will be sent. Note that Channels aren't
// like files; you don't usually need to close them. Closing is only necessary when the receiver must be
// told there are no more values coming, such as to terminate a range loop.
// BEST PRACTICE : Only the sender should close a channel, never the receiver. Sending on a closed
// channel will cause a panic. Also, a channel should be owned by a single entity.

func Goroutines() {

	// start a thread for the next function.
	// the function, and so the result, is executed in a new thread (goroutine).
	// A thread happens in the same address space so there is no guaranteed sync.
	// The following calls won't be synchronized so multiple "hello" can happen without "world".
	go say("hello")
	say("world")

	// Channels are the pipes that connect concurrent goroutines. You can send values into channels from
	// one goroutine and receive those values into another goroutine.
	// Channels are not meant to be used outside goroutines.
	stringChannel := make(chan string)
	// ping will send data in the channel
	go ping(stringChannel)
	// pong will receive the data in the channel
	go pong(stringChannel)

	// A channel cannot receive 2 values if less than 2 values have been sent in the channel.
	// The following lines will raise an error if there are uncommented :
	//errorChan := make(chan string)
	//errorChan <- "hello"
	//msg := <-errorChan
	//msg = <-errorChan

	// Channels are 'blocking' by default.
	// It means that you cannot send 2 data in a channel if you do not consume the first one before.
	// Note : check the 'synchronize goroutines' example with the 'done' channel for an example of
	// unbuffered channels
	blockingChannel := make(chan string)
	// If we execute the following line, the execution will raise a deadlock error.
	// Indeed, blocking channels block receiving data if the sender is not ready :
	//log.Println(<-blockingChannel)
	// Otherwise, we can start sending data in the channel and receive it further :
	go func() { blockingChannel <- "blocking code" }()
	log.Println(<-blockingChannel)
	close(blockingChannel)
	// With blocking Channels, you are sure that the following code will be executed after the
	// channel is consumed (data was received)
	log.Println("code blocked finished")

	// Use 'buffered channels' to avoid blocking the code.
	// By default, (like previously), channels are unbuffered, meaning that they will only accept sends
	// (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
	// It means that because a channel is buffered, we can send values into the channel without
	// a corresponding concurrent receive and we can start receiving data before any data was sent.
	// Note : unbuffered channel is a common pattern to prevent goroutine leaks in case the channel is
	// never read.
	bufferedChannel := make(chan string, 1)
	go func() { bufferedChannel <- "not blocking code" }()
	log.Println(<-bufferedChannel)
	close(bufferedChannel)
	log.Println("this string maybe printed after further code")

	// WARNING ! the sending order is not always respected for channel receivers.
	// Indeed, since goroutines are sending results, the second goroutines has chances to finish before
	// the first goroutine. So the order of the result is unpredictable.
	// If you want to keep track of order, choose different channels or learn about waitGroup in the
	// further examples.
	noOrderChannel := make(chan int)
	go func() { noOrderChannel <- 1 }()
	go func() { noOrderChannel <- 2 }()
	go func() { noOrderChannel <- 3 }()
	int1, int2, int3 := <-noOrderChannel, <-noOrderChannel, <-noOrderChannel
	close(noOrderChannel)
	log.Println("Unordered results:", int1, int2, int3)

	// Loop range over channels.
	// range iterates over each element as it’s received from queue. Because we closed the channel above,
	// the iteration terminates after receiving the 2 elements.
	c1 := make(chan string, 3)
	c1 <- "one"
	c1 <- "two"
	c1 <- "three"
	close(c1)
	for elem := range c1 {
		log.Println(elem)
	}

	// Goroutines and channels are useful to RECEIVE data (listens for values).
	// In this example, 'listenWorker' is an arbitrary function that uses a channels and waits for values
	// We call this function in a goroutine : it starts working before we send data to work with.
	// Then we send data to the worker through a channel that we provide as an argument.
	// You can notice that the channel is closed from the sender, here the main function, while the
	// worker receives the data.
	jobs := make(chan int, 3)
	log.Println("starts listening for values")
	go listenWorker(jobs)
	log.Println("starts sending values")
	for j := 1; j <= 3; j++ {
		jobs <- j
		log.Println("sent job", j)
	}
	close(jobs)
	log.Println("sent all values")
	log.Println("stop listening for values")

	// Goroutines and channels are also useful to SEND data.
	// In this example, 'sendWorker' sends an 'echo' string in a channel using a goroutine.
	// Then the method 'staticReceiver' reads the channel until the sender closes it.
	echoc := make(chan string)
	log.Println("start sending echo")
	go sendWorker(echoc, "echo")
	// The receiver is in a function in the same thread
	log.Println("start receiving echo")
	staticReceiver(echoc)
	log.Println("stop receiving echo")

	// We can use channels to synchronize execution across goroutines.
	// Here’s an example of using a blocking receive to wait for a goroutine to finish.
	// The done channel will be used to notify a goroutine that this function’s work is done.
	// If you removed the <- done line from this program, the program would exit before the worker even
	// started.
	echoc = make(chan string, 2)
	done := make(chan bool)
	log.Println("starts synced goroutine")
	go sendWorkerSync(echoc, done, "echo sync")

	log.Println("Waiting for the synced goroutine to finish ...")
	// here the code is blocked until a value is sent in the 'done' channel
	// this is due to the unbuffered channel properties
	<-done
	log.Println("Synced goroutine has finished, receive data from channel :")
	simpleStaticReceiver(echoc)

	// We can use a 'wait group' to wait for multiple goroutines to finish
	// Note: if a WaitGroup is explicitly passed into functions, it should be done by pointers.
	// The first example uses the same function to start multiple goroutines.
	// The defer statement defers the execution of a function until the surrounding function returns.
	var wg sync.WaitGroup
	echoc = make(chan string, 2)
	log.Println("starts multiple goroutines with the same function")
	for i := 1; i <= 3; i++ {
		// We add a new instance in the wait group for each goroutine
		wg.Add(1)
		id := i
		go func() {
			// The defer statement tells the wait group that the goroutine have finished when the
			// 'say' function returns
			defer wg.Done()
			say(fmt.Sprintf("I am the goroutine number %d", id))
		}()
	}
	log.Println("waiting for multiple goroutines to finish ...")
	// wait for all the goroutines in the wait group to be 'done'
	wg.Wait()
	log.Println("all the goroutines have finished")

	// This second example is more intuitive since we know how many goroutine we start, with different
	// functions.
	var wg2 sync.WaitGroup
	wg2.Add(2)
	pingc := make(chan string, 2)
	pongc := make(chan string, 2)
	log.Println("starts multiple goroutines with different functions")
	// Note: give a wait group as a pointer to functions
	// the functions defer 'wg.Done()' in their implementation
	go sendWorkerWait(pingc, "ping", &wg2)
	go sendWorkerWait(pongc, "pong", &wg2)
	log.Println("waiting for multiple goroutines to finish ...")
	wg2.Wait()
	log.Println("all the goroutines have finished, receive data from channels :")
	simpleStaticReceiver(pingc)
	simpleStaticReceiver(pongc)

	// Asynchronous goroutines are useful to parallelize tasks.
	// But it is sometimes hard to correlate their results.
	// In this example, 'sendWokerPing' sends a 'ping' string in a channel using a goroutine.
	// 'sendWokerPong' sends a 'pong' string in a channel using another goroutine.
	// Then the method 'staticReceiver' reads the channels until the sender closes the channels.
	pingc = make(chan string, 2)
	pongc = make(chan string, 2)
	log.Println("starts async ping and pong")
	go sendWorker(pingc, "ping")
	go sendWorker(pongc, "pong")
	asyncPingPongReceiver(pingc, pongc)

	// Timeouts are an elegant way to avoid leaks for asynchronous goroutines
	// Timeouts are important for programs that connect to external resources or that otherwise need to
	// bound execution times.
	// In the following example, res := <-c1 awaits the result and <-time.After awaits a value to be sent
	// after the timeout of 1s. Since 'select' proceeds with the first receive that’s ready, we’ll take
	// the timeout case if the operation takes more than the allowed 1s. If we allow a longer timeout of
	// 3s, then the receive from c2 will succeed and we’ll print the result.
	ct1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ct1 <- "result 1"
	}()
	// select statement to receive the data
	// change the timeout value
	timeoutValue := 1 // seconds
	select {
	case res := <-ct1:
		log.Println("never printed due to the timeout", res)
	case <-time.After(time.Duration(timeoutValue) * time.Second):
		// the timeout will happen because our goroutine waits for 2 seconds
		log.Println("timeout 1")
	}

}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		log.Println(s)
	}
}

// ping function only accept channels to send data (injecting data into the channel)
// this behavior is defined by passing 'c chan<- type' in arguments
func ping(c chan<- string) {
	c <- "ping"
	// best practice : the sender close the channel when no more values are sent into the channel
	close(c)
}

// pong function only accept channels to receive data
// this behavior is defined by passing 'c <-chan type' in arguments
func pong(c <-chan string) {
	// for a receiver, you can check if data was sent to a channel with '_, ok'
	// 'ok' is a boolean equals to true when a channel contains at least one data, and false if no data
	// was sent.
	message, ok := <-c
	if ok {
		log.Println("received", message)
		log.Println("pong")
	}
}

func listenWorker(jobs <-chan int) {
	for {
		j, more := <-jobs
		if more {
			log.Println("received job", j)
		} else {
			log.Println("received all jobs")
			return
		}
	}
}

func sendWorker(c chan<- string, data string) {
	for i := 0; i < 2; i++ {
		log.Println("sender:", data)
		c <- data
	}
	close(c)
	log.Println("stop sending", data)
}

func sendWorkerSync(c chan<- string, done chan<- bool, data string) {
	// same as sendWorker func.
	sendWorker(c, data)
	// Send a value to notify that we’re done.
	done <- true
	close(done)
}

func sendWorkerWait(c chan<- string, data string, wg *sync.WaitGroup) {
	defer wg.Done()
	sendWorker(c, data)
}

// This function is useful because you can add as many channels as you want by copy-pasting the 'case'
// statement of the first channel for a new one, then by adding the channel in the if statement to close
// the for loop.
func staticReceiver(c chan string) {
	for {
		// 'select' properties :
		// If one or more of the communications can proceed, a single one that can proceed is chosen via
		// a uniform pseudo-random selection. Otherwise, if there is a 'default' case, that case is
		// chosen.
		// If there is no 'default' case, the "select" statement blocks until at least one of the
		// communications can proceed.
		// So, without a default case, the code will block until some data is available in either of the
		// channels. It implicitly waits for the other goroutines to wake up and write to their channel.
		select {
		// 'ok' is a boolean value equals true when more values was sent to the channel
		// by checking 'ok' value, we are sure to read all the data in a channel before breaking the loop
		case data, ok := <-c:
			if data != "" {
				log.Println("receiver:", data)
			}
			// no more data in the channel
			if !ok {
				c = nil
			}
		}
		// breaking the loop
		if c == nil {
			return
		}
	}
}

func asyncPingPongReceiver(pingc chan string, pongc chan string) {
	for {
		select {
		case data, ok := <-pingc:
			if data != "" {
				log.Println("receiver:", data)
			}
			if !ok {
				pingc = nil
			}
		case data, ok := <-pongc:
			if data != "" {
				log.Println("receiver:", data)
			}
			if !ok {
				pongc = nil
			}
		}
		if pingc == nil && pongc == nil {
			return
		}
	}
}

func simpleStaticReceiver(c chan string) {
	for elem := range c {
		log.Println(elem)
	}
}
