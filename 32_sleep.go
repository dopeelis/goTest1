package main

import "time"

func main() {
	sleep(5)
}

// Реализация функции sleep
func sleep(durationSec int) {
	<-time.After(time.Second * time.Duration(durationSec))
}
