package main

import (
	"machine"
	"strconv"
)

var serial = machine.UART0

func main() {
	ini()
	println(getCoreFreq())
}

func ini() {
	println("----------------------------------------")
}

func getCoreFreq() string {
	return strconv.Itoa(int(machine.CPUFrequency())/1000000) + "MHz"
}

func print(str string) {
	serial.Write([]byte(str + "\r"))
}

func println(str string) {
	serial.Write([]byte(str + "\n" + "\r"))
}
