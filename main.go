/*
@Time : 2021/3/20 10:39
@Author : Tux
@Description :
*/

package main

import (
	"log"

	"chapter01/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd execute err: %v", err)
	}
}
