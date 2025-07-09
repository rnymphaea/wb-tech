package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Human - базовая структура для описания человека
type Human struct {
	FirstName  string
	LastName   string
	Age        int
	Profession string
}

// Introduce - метод Human для вывода инфмормации о человеке
func (h Human) Introduce() {
	fmt.Printf("Hello! My first name is %s, my last name is %s. I'm %d years old. My profession is %s.\n",
		h.FirstName, h.LastName, h.Age, h.Profession)
}

// Action - структура, встраивающая Human и добавляющая информацию о действии
type Action struct {
	Human
	CurrentTask string
	Duration    time.Duration
}

// Start - метод Action для начала работы над задачей
func (a Action) Start() {
	// демонстрация того, что Action может обращаться к полям встариваемой структуры Human
	fmt.Printf("Employee %s %s started working on %s.\n", a.FirstName, a.LastName, a.CurrentTask)

	// имитация работы
	time.Sleep(a.Duration)
}

// Finish - метод Action для завершения работы над задачей
func (a Action) Finish() {
	fmt.Printf("Employee %s %s finished working on %s.\n", a.FirstName, a.LastName, a.CurrentTask)
	fmt.Printf("%s by %s %s took %s.\n", a.CurrentTask, a.FirstName, a.LastName, a.Duration)
}

func main() {
	// создаём экземпляр структуры Human
	human := Human{
		FirstName:  "Ivan",
		LastName:   "Ivanov",
		Age:        30,
		Profession: "Go developer",
	}

	// создаём первое действие (Api optimization)
	apiOptimization := Action{
		Human:       human,
		CurrentTask: "Api optimization",
		Duration:    time.Duration(rand.Intn(5)+1) * time.Second,
	}
	
	// демонстрация того, что Action может обращаться к методам встраиваемой структуры Human
	apiOptimization.Introduce()

	// начало и завершение выполнения действия
	apiOptimization.Start()
	apiOptimization.Finish()
	
	// создаём второе действие (Method modification)
	methodModification := Action{
		Human:       human,
		CurrentTask: "Method modification",
		Duration:    time.Duration(rand.Intn(5)+1) * time.Second,
	}

	methodModification.Start()
	methodModification.Finish()
}
