package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, HasDMARC, dmarcRecord \n")

	for scanner.Scan(){
		checkDomain(scanner.Text())

	}

	if err := scanner.Err(); err != nil{
		log.Fatal(err, "could not read from input")
	}
}

func checkDomain(domain string){
	
}