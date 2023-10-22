package tutorial

import (
	"log"
	"time"
)

// CHECKS FOR THE GOLANG GOROUTINES AND CHANNELS TUTORIAL FIRST TO FULLY UNDERSTAND SOME OF THE FEATURES!

// Timers are for when you want to do something once in the future.
// Tickers are for when you want to do something repeatedly at regular intervals.

// Timers represent a single event in the future. You tell the timer how long you want to wait, and it
// provides a channel that will be notified at that time.
// If you just want to wait, you could have used time.Sleep.
// One reason a timer may be useful is that you can cancel the timer before it fires.
// Tickers use a similar mechanism to timers: a channel that is sent values.
// Tickers can be stopped like timers.
// Once a ticker is stopped it won’t receive any more values on its channel.
func Timers() {
	log.Println("Timers tutorial")

	// TIMERS
	// This timer will wait 2 seconds.
	// The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer fired.
	log.Println("Timing 2 seconds ...")
	timer1 := time.NewTimer(2 * time.Second)
	// wait for the channel, and so, wait for the timer
	<-timer1.C
	log.Println("... The 2 seconds ended !")

	// Cancel a timer before it fires
	log.Println("Starting timer2 of 1 second ...")
	timer2 := time.NewTimer(time.Second)
	// starting a goroutine
	go func() {
		// non-blocking channel
		<-timer2.C
		log.Println("Timer2 fired !")
	}()
	// the timer should stop before a second passed
	myStop := timer2.Stop()
	if myStop {
		log.Println("Timer2 stopped before it fired")
	}

	// Or simply wait
	log.Println("Waiting 2 seconds (not a timer) ...")
	time.Sleep(2 * time.Second)
	log.Println("... The 2 seconds ended !")

	// TICKERS
	// Here we’ll use the select builtin on the channel to await the values as they arrive every 500ms.
	// we will synchronize the channels with a boolean channel
	ticker := time.NewTicker(time.Second)
	done := make(chan bool) // done channel for sync
	// here we execute a goroutine
	// continue until the done channel receives a value
	// in this case we stop the ticker
	log.Println("Ticker started for 6 seconds")
	go func() {
		// infinite channel's select loop
		for {
			select {
			case <-done:
				// stopping loop
				return
			case t := <-ticker.C:
				log.Println("Tick at", t)
			}
		}
	}()
	// here we wait for 6 seconds
	// during that time, the goroutine will run, printing for values
	// then we stop the ticker and sending a value inside the 'done' channel
	// resulting in stopping the select loop of the ticker's channel
	time.Sleep(6 * time.Second)
	ticker.Stop()
	done <- true
	log.Println("Ticker stopped")

}
