package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/CrashArt/quiz/entity"
)

const (
	taskFile   = "json/task.json"
	ratingFile = "json/rating.json"
)

func Start() {
	//var name string
	calculate()
}

func calculate() {

	startTime := time.Now()
	questions, err := getTasks()
	//SHAFL q=shafl questions

	if err != nil {
		log.Println(err)
	}
	correct, incorrect := askQuestions(questions)
	endTime := time.Now()
	resultTime := endTime.Sub(startTime)
	fmt.Println(correct, incorrect, resultTime)
	gamer := &entity.Rating{
		Name:           "Test",
		CorrectAnswers: correct,
		Duration:       resultTime,
	}
	addToRating(gamer)
	//getRarting()

}

func getTasks() ([]*entity.Task, error) {
	var tasks []*entity.Task
	jsonFile, err := os.Open(taskFile)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	value, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(value, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func askQuestions(questions []*entity.Task) (int, int) {
	var correct, incorrect int
	var userAnswer string
	for _, task := range questions {
		fmt.Println(task.Question)
		fmt.Fscan(os.Stdin, &userAnswer)
		if userAnswer == task.Answer {
			correct++
		} else {
			incorrect++
		}

	}
	return correct, incorrect
}

func addToRating(user *entity.Rating) {
	var rating []*entity.Rating //get rating
	rating = append(rating, user)
	result, err := json.Marshal(rating)
	if err != nil {
		log.Println(err)
	}
	file, _ := os.Create(ratingFile)
	file.WriteString(string(result))
	defer file.Close()
	fmt.Println("TOP")
}

//get rating.
func GetRarting() {
	var rating []entity.Rating
	jsonFile, err := os.Open(ratingFile)
	if err != nil {
		fmt.Print(err)
	}
	defer jsonFile.Close()
	value, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(value, &rating)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(rating[0])
}

//func шафл ,
