/**
 * Created by chaolinding on 2018/3/16.
 */

package main

import (
	"test"
	"log"
)

func main() {
	pByte := test.Code()
	person := test.Decode(pByte)
	log.Println( person.Phones[0].Type)
}