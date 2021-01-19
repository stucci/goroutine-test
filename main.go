package main

import (
	"fmt"
	"time"
)

func main() {
	// try()
	// useSharedValue()
	// useChannel()
	// channelIsFirstClassObject()
	directionalChannel()
}

func try() {
	defer fmt.Println("main done")
	go func() {
		defer fmt.Println("goroutine1 done")
		time.Sleep(3 * time.Second)
	}()
	go func() {
		defer fmt.Println("goroutine2 done")
		time.Sleep(1 * time.Second)
	}()
	time.Sleep(5 * time.Second)
}

func useSharedValue() {
	var v int
	go func() {
		time.Sleep(2 * time.Second)
		v = 100
	}()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(v)
	}()
	time.Sleep(2 * time.Second)
}

func useChannel() {
	ch := make(chan int)
	go func() {
		ch <- 100
	}()
	go func() {
		v := <-ch
		fmt.Println(v)
	}()
	time.Sleep(2 * time.Second)
}

func makeCh() chan int {
	return make(chan int)
}

func recvCh(recv chan int) int {
	return <-recv
}

func channelIsFirstClassObject() {
	ch := makeCh()
	go func() { ch <- 100 }()
	fmt.Println(recvCh(ch))
}

func recvOnlyCh(recv <-chan int) int {
	return <-recv
}

func directionalChannel() {
	ch := makeCh()
	go func(ch chan<- int) { ch <- 100 }(ch)
	fmt.Println(recvOnlyCh(ch))
}
