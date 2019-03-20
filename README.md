# Simulating-Microchip-AN526-8x8-Unsigned-Multiplication
This is used to simulate the 8x8 Unsigned Multiplication PIC16 assembly language routine found in the Application Note 526 Source Code of Microchip.

I am not sure if 'simulate' is the right word.

## Compiling the program
```
go build main.go
```

## Using the program
You can execute:
```
go run main.go 8bitNum1 8bitNum2
```
If you have built the program, you can execute(if you are on Windows):
```
./main.exe 8bitNum1 8bitNum2
```
or(if you are on a Linux-based distro:
```
./main 8bitNum1 8bitNum2
```

## Sample Execution
C is the Carry bit.

H is the Most Significant Byte of the product.

L is the Least Significant Byte of the product.

As you can see by the end on H and L, the product is displayed. It is separated into 2 because remember that it is PIC16.
```
go run main.go 123 24

     123 x       24 =             2952 (base 10)
01111011 x 00011000 = 0000101110001000 (base 2)

        C 8Bit Val
0
        0 00000000 H
        0 00000000 L
1
        0 00000000 H
        0 00000000 L
2
        0 00000000 H
        0 00000000 L
3
        1 00111101 H
        0 10000000 L
4
        0 01011100 H
        0 01000000 L
5
        0 00101110 H
        0 00100000 L
6
        0 00010111 H
        0 00010000 L
7
        1 00001011 H
        0 10001000 L

```

## The PIC16 assembly language program
I tried to do this by hand.

The program goes like this:
```
Same    equ     1

;
;****   Define a macro for adding & right shifting  **
;
mult    MACRO   bit             ; Begin macro
	btfsc   mulplr,bit
	addwf   H_byte,Same
	rrf     H_byte,Same
	rrf     L_byte,Same
	ENDM                    ; End of macro
;
; *****************************         Begin Multiplier Routine
mpy_F   clrf    H_byte
	clrf    L_byte
	movf    mulcnd,W        ; move the multiplicand to W reg.
	bcf     STATUS,C    ; Clear the carry bit in the status Reg.
	mult    0
	mult    1
	mult    2
	mult    3
	mult    4
	mult    5
	mult    6
	mult    7
;
	retlw   0
```

## Links
* http://ww1.microchip.com/downloads/en/AppNotes/00526e.pdf
* http://ww1.microchip.com/downloads/en/AppNotes/00526.zip
