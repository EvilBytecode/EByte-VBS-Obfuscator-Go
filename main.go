package main

import (
	"fmt"
	"os"

	"Ebyte-VBS-Obf/obfuscation"
	"Ebyte-VBS-Obf/utils"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("[+] Usage: go run obfuscator.go inFile.vbs outFile.vbs")
		os.Exit(1)
	}

	utils.SeedRandom()

	splitter := "*"

	numOfChars := utils.RandomInt(5, 60)
	pld := utils.RandomString(numOfChars)
	array := utils.RandomString(numOfChars)
	temp := utils.RandomString(numOfChars)
	x := utils.RandomString(numOfChars)

	subOne := utils.RandomString(numOfChars)
	subTwo := utils.RandomString(numOfChars)

	inputFile, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("[-] Error reading input file:", err)
		os.Exit(1)
	}

	outputFile, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("[-] Error creating output file:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	fmt.Fprintln(outputFile, obfuscation.RandCapitalization("Dim "+pld+", "+array+", "+temp))
	fmt.Fprintln(outputFile, obfuscation.RandCapitalization("Sub "+subOne))
	fmt.Fprintf(outputFile, "%s = \"%s\"\n", obfuscation.RandCapitalization(pld), obfuscation.Obfu(string(inputFile)))
	fmt.Fprintf(outputFile, "%s = Split(%s, chr(eval(%s)))\n", obfuscation.RandCapitalization(array), pld, obfuscation.Obfu(splitter))
	fmt.Fprintln(outputFile, obfuscation.RandCapitalization("for each "+x+" in "+array))
	fmt.Fprintln(outputFile, obfuscation.RandCapitalization(temp+" = "+temp+" & chr(eval("+x+"))"))
	fmt.Fprintln(outputFile, obfuscation.RandCapitalization("next"))
	fmt.Fprintln(outputFile, obfuscation.RandCapitalization(subTwo))
	fmt.Fprintln(outputFile, obfuscation.RandCapitalization("eval(execute("+temp+"))"))
	fmt.Fprintln(outputFile, obfuscation.RandCapitalization("End Sub"))
	fmt.Fprintln(outputFile, obfuscation.RandCapitalization(subOne))

	fmt.Println("[+] EByte Obf Done!")
}
