package main 

import (
	"fmt"
	"sync"
)

func main(){
	var m sync.Mutex
	var a int //what is value of a
	done := make(chan struct{})
	m.Lock()

	//how is defer used (and why?)
	go func(){
		m.Lock() //if this is commented it is unlocked twice below, which causes an error
		fmt.Println(a) //lock used for a memory location, create struct using data and lock together. Mutex atands by itself
		m.Unlock() //who should be locked or unlocked (instead, you can put defer m.Unlock right underneath Lock, and it will unlok it at the end of the scope)
		close(done) //this closes channel
	}(

	go func (){ //no matter what way or order it runs, it should always run
		a = 1
		a.Unlock() //channel - one line of code to lock, one line to unlock
	}()

	//what does this do
	<-done //this waits forever because its empty, stopping it from exiting. Line 19
	//
}