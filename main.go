package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CrashArt/quiz/util"
)

func main() {
	var choise int
	fmt.Println("1.Start Game\n" +
		"2.Get rating\n" +
		"3.Exit")
	fmt.Println("Make your choice:")
	_, err := fmt.Fscan(os.Stdin, &choise)
	if err != nil {
		log.Println(err)
	}
	getChoice(choise)

}
func getChoice(choise int) {

	switch choise {

	case 1:
		fmt.Println("Game Start")
		util.Start()
	case 2:
		fmt.Println("Game rating")
		util.GetRarting()
		//getutil
		

	case 3:
		fmt.Println("Exit")
	default:
		fmt.Println("хуйню не выбирай")
	}
}
