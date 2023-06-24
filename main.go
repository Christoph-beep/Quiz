package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// todo

// 2. Creating file and directory for answering the answers
// done
// 3. writing answers to file with JSON

// 4. compare userinput with solution -> therefore use toLowerCase() for the userAnswers

// 5. Website stuff (handlers + forms for user input + get the user input in the backend  with formValue)
// done

func main() {
	fmt.Println("Starting Server on port 8080")
	//mux.HandleFunc("/quiz", quiz)
	//mux.HandleFunc("/", quiz)

	createAnswerDirectory()
	createAnswerFile()
	answerFileWrite("blau")
	answerFileWrite("hund")

	http.HandleFunc("/", quiz)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":8080", nil)

}

// Part 1

func process(w http.ResponseWriter, r *http.Request) {

	// When is the open call required ?
	//answers, err := os.Open("answers.txt")
	//defer answers.Close()

	answersContent, err := os.ReadFile("answerDirectory/answers.txt")
	answersContentString := string(answersContent)
	//if err != nil {
	//	fmt.Println("An error occured in process")
	//}
	// answers from user
	Question0 := r.FormValue("0")
	Question1 := r.FormValue("1")

	fmt.Println(Question0)
	fmt.Println(Question1)

	// check if answers were right
	rightAnswersCounter := 0

	fmt.Println(answersContentString + "inhalt file")
	if strings.Contains(answersContentString, Question0) {
		rightAnswersCounter++
		fmt.Println("right answer 1")
	} else {
		fmt.Println("wrong 1")
	}
	if strings.Contains(answersContentString, Question1) {
		rightAnswersCounter++
		fmt.Println("right answer 2")
	} else {
		fmt.Println("wrong 2")
	}

	processedData, err := template.ParseFiles("process.html")
	if err != nil {
		fmt.Println(err)
	}

	processedData.Execute(w, rightAnswersCounter)

}

func quiz(w http.ResponseWriter, r *http.Request) {

	// global variables

	quizpage, err := template.ParseFiles("quiz.html")
	if err != nil {
		fmt.Println(err)
	}
	quizpage.Execute(w, nil)

}

// create Folder for the file with answers
// if true, everything is right
func createAnswerDirectory() bool {
	_, err := os.Stat("answerDirectory")
	if err != nil {
		fmt.Println("Directory does not exist so far createAnswerDirectory")
		// directory is created
		os.MkdirAll("answerDirectory", os.ModePerm)
		fmt.Println("directory has been created")
		if err != nil {
			fmt.Println(err)
			return false
		}

	}
	fmt.Println("Directory already exists")
	return true

}

// Creating File with answers
// if true, everything is right
func createAnswerFile() bool {
	_, err := os.Stat("answerDirectory/answers.txt")
	if err != nil {
		fmt.Println("File does not exist so far")
		// file is created
		answerDirectory, err := os.Create("answerDirectory/answers.txt")
		fmt.Println("file has been created successfully")
		if err != nil {
			fmt.Println(err)
			return false
		}
		defer answerDirectory.Close()

	}
	fmt.Println("everything is working properly (-:)")
	return true

}

// Writing solutions to file -> can only be done by admins, needs to be implimented
func answerFileWrite(solutions string) {
	// reading part
	// check, if data has already been written to the file, therefore the file needs to be opened first

	file, err := os.ReadFile("answerDirectory/answers.txt")
	fileContentString := string(file)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// file information is stored here

	if fileContentString == solutions {
		fmt.Println("the following answers do already exist, therefore they do not need to be eddited anymore")
		return
	} else {
		// answers do not exist so far and can therefore be added

		err := os.WriteFile("answerDirectory/answers.txt", file, 0644)

		if err != nil {
			fmt.Println("error occured in answerFileWrite")
			fmt.Println(err)
			return
		}
		fmt.Println("data has been successfuly written into file")
	}

}