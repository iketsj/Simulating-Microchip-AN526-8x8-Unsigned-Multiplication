package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {

	checkNumberOfArguments()
	multiplicand, multiplier := checkValueOfArguments()
	product := uint16(multiplicand) * uint16(multiplier)
	fmt.Println()
	fmt.Printf("%8d x %8d = %16d (base 10)\n", multiplicand, multiplier, product)
	fmt.Printf("%0.8b x %0.8b = %0.16b (base 2)\n", multiplicand, multiplier, product)
	fmt.Println()

	simulateMicrochip8x8MultiplicationAN256(multiplicand, multiplier)
	fmt.Println()
}

func checkNumberOfArguments() {
	if len(os.Args) != 3 {
		var osProgramFormat string
		var errorMessage string

		if runtime.GOOS == "windows" {
			osProgramFormat = fmt.Sprintf("./main.exe")
		} else if runtime.GOOS == "linux" {
			osProgramFormat = fmt.Sprintf("./main")
		}

		fmt.Println(osProgramFormat)
		errorMessage = fmt.Sprintf("Usage: %s 8bitNum1 8bitNum2", osProgramFormat)
		panic(errorMessage)
	}
}

func checkValueOfArguments() (uint8, uint8) {
	var err error
	var errorMessage string
	var numbers [2]uint64

	for i := 1; i < 3; i++ {
		numbers[i-1], err = strconv.ParseUint(os.Args[i], 10, 8)
		if err != nil {
			errorMessage = fmt.Sprintf("\"%s\" should be an 8bit number!\n%s", os.Args[i], err)
			panic(errorMessage)
		}
	}

	multiplicand := uint8(numbers[0])
	multiplier := uint8(numbers[1])
	return multiplicand, multiplier
}

func simulateMicrochip8x8MultiplicationAN256(multiplicand uint8, multiplier uint8) {
	var prodH uint8 = 0
	var prodL uint8 = 0
	var carry uint8 = 0
	accumulator := multiplicand
	fmt.Println("\tC 8Bit Val")
	for i := 0; i < 8; i++ {
		fmt.Printf("%d\n", i)
		if checkBit(multiplier, uint8(i)) == 1 {
			carry = uint8((uint16(accumulator) + uint16(prodH)) >> 8)
			prodH = uint8(uint16(accumulator) + uint16(prodH))
		}
		prodH, carry = rotateRightThroughCarry(prodH, carry)
		fmt.Printf("\t%d %0.8b H\n", carry, prodH)
		prodL, carry = rotateRightThroughCarry(prodL, carry)
		fmt.Printf("\t%d %0.8b L\n", carry, prodL)
	}
}

func rotateRightThroughCarry(value uint8, carry uint8) (uint8, uint8) {
	var newValue uint8
	if carry == 1 {
		newValue = (value >> 1) | (uint8(1) << 7)
	} else {
		newValue = (value >> 1) & ^(uint8(1) << 7)
	}
	return newValue, value & 1
}

func checkBit(multiplier uint8, bitPosition uint8) uint8 {
	return ((multiplier >> bitPosition) & 1)
}
