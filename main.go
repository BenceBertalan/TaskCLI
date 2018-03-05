package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var apiendpoint string

//Task : A representation of a Server-side task object
type Task struct {
	ID              int    `json:"ID"`
	Name            string `json:"Name"`
	Command         string `json:"Command"`
	Status          string `json:"Status"`
	Output          string `json:"Output"`
	CreatedDateTime string `json:"Created_DateTime"`
	LastRunDateTime string `json:"Last_Run_DateTime"`
}

func main() {
	apiptr := flag.String("apiendpoint", "", "URL of the API endpoint")
	commandptr := flag.String("command", "", "Task command. Can be either get, getbyid, or add")
	id := flag.String("id", "", "Task ID. Only used with the getbyid command")
	newtaskname := flag.String("taskname", "", "Name for the New Task to add. Only used when add command specified")
	newtaskcommand := flag.String("taskcommand", "", "Command to be executed by the New Task to add. Only used when add command specified")
	flag.Parse()

	if *apiptr == "" || *commandptr == "" {
		flag.Usage()
		os.Exit(1)
	}
	apiendpoint = *apiptr
	switch *commandptr {
	case "get":
		GetAllTasks()
	case "getbyid":
		idint, err := strconv.Atoi(*id)
		//fmt.Println("Getting task by id" + *taskid)
		if err != nil {
			fmt.Println("Invalid TaskID specified")
		} else {
			GetTaskByID(idint)
		}

	case "add":
		if *newtaskcommand == "" || *newtaskname == "" {
			fmt.Println("No Task command or name specified")
		} else {
			NewTask(*newtaskname, *newtaskcommand)
		}
	}

	//GetAllTasks()
	//GetTaskByID(25)
	//NewTask("Teszt", "ls -l")
}

//NewTask : Adds new task using the API post
func NewTask(Name string, Command string) {

	var t Task
	t.Command = Command
	t.Name = Name

	url := apiendpoint

	jsonStr, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
	}

	if resp.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	} else {
		fmt.Println("Error while adding the new task:" + err.Error())
	}
}

//GetAllTasks : Gets all task from the API endpoint
func GetAllTasks() {
	var alltasks []Task

	url := apiendpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	if err := json.NewDecoder(resp.Body).Decode(&alltasks); err != nil {
		log.Println(err)
	}

	for _, t := range alltasks {
		bytes, _ := json.Marshal(t)
		fmt.Println(string(bytes))
	}

}

//GetTaskByID a Task BY ID
func GetTaskByID(id int) {

	t := Task{}
	t.ID = -1

	idstr := strconv.Itoa(id)

	url := apiendpoint + "/" + idstr

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		log.Println(err)
	}
	if t.ID != -1 {
		b, _ := json.Marshal(t)
		fmt.Println(string(b))
	} else {
		fmt.Println("No Task found with the given ID")
	}
}
