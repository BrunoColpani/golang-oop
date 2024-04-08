package main

import "fmt"

type CurrentAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	brunoCurrentAccount := CurrentAccount{"Bruno Colpani", 589, 123456, 259.82}
	gabrielCurrentAccount := CurrentAccount{"Gabriel Colpani", 589, 123789, 3578.82}

	fmt.Println(brunoCurrentAccount, gabrielCurrentAccount)
}
