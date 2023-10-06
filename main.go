package main

import (
	"strings"
)

type Person struct {
	activeRoom string
	inventory  [2]string
	keys       bool
	notes      bool
	backpack   bool
}

var person Person

type Rooms struct {
	hello   string
	objects string
	qwest   string
	going   string
}

var (
	door  bool
	rooms = make(map[string]Rooms)
)

func LookAround() string {
	var answer string
	for name, element := range rooms {
		if person.activeRoom == name {
			answer += element.hello + element.objects + element.qwest + element.going
			if person.activeRoom == "кухня" {
				if person.backpack {
					answer = strings.Replace(answer, "собрать рюкзак и ", "", -1)
				}
			}
			if person.activeRoom == "комната" {
				if person.keys && person.notes && person.backpack {
					return "пустая комната. можно пройти - коридор"
				}
				if person.keys {
					answer = strings.Replace(answer, "ключи, ", "", -1)
				}
				if person.notes {
					answer = strings.Replace(answer, "конспекты, ", "", -1)
				}
				if person.backpack {
					answer = strings.Replace(answer, ", на стуле: рюкзак", "", -1)
				}
			}
			return answer
		}
	}
	return "not implemented"
}

func Go(command string) string {
	if command == "кухня" && person.activeRoom == "коридор" {
		person.activeRoom = "кухня"
		return "кухня, ничего интересного. можно пройти - коридор"
	} else if command == "коридор" && person.activeRoom != "коридор" {
		person.activeRoom = "коридор"
		return "ничего интересного. можно пройти - кухня, комната, улица"
	} else if command == "комната" && person.activeRoom == "коридор" {
		person.activeRoom = "комната"
		return "ты в своей комнате. можно пройти - коридор"
	} else if command == "улица" && person.activeRoom == "коридор" {
		if door {
			return "на улице весна. можно пройти - домой"
		}
		return "дверь закрыта"
	}
	return "нет пути в " + command
}

func Take(command string) string {
	if person.backpack {
		if command == "ключи" {
			if !person.keys {
				person.inventory[0] = command
				person.keys = true
				return "предмет добавлен в инвентарь: " + person.inventory[0]
			}
			return "нет такого"
		} else if command == "конспекты" && !person.notes {
			person.inventory[1] = command
			person.notes = true
			return "предмет добавлен в инвентарь: " + person.inventory[1]
		} else {
			return "нет такого"
		}
	}
	return "некуда класть"
}

func PutOn() string {
	person.backpack = true
	return "вы надели: рюкзак"
}

func Apply(command1, command2 string) string {
	for _, elem := range person.inventory {
		if command1 == elem {
			if command1 == "ключи" && person.keys {
				door = true
				if command2 == "дверь" {
					return "дверь открыта"
				} else {
					return "не к чему применить"
				}
			}
		} else {
			return "нет предмета в инвентаре - " + command1
		}
	}
	return "не к чему применить"
}

func main() {

}

func initGame() {
	person.activeRoom = "кухня"
	person.inventory[0] = ""
	person.inventory[1] = ""
	person.keys = false
	person.notes = false
	person.backpack = false
	rooms["кухня"] = Rooms{
		hello:   "ты находишься на кухне, ",
		objects: "на столе: чай, ",
		qwest:   "надо собрать рюкзак и идти в универ. ",
		going:   "можно пройти - коридор",
	}
	rooms["комната"] = Rooms{
		objects: "на столе: ключи, конспекты, на стуле: рюкзак. ",
		going:   "можно пройти - коридор",
	}
	door = false
}

func handleCommand(command string) string {
	commandDo := []string(strings.Split(command, " "))
	switch {
	case commandDo[0] == "осмотреться":
		return LookAround()
	case commandDo[0] == "идти":
		commandGo := commandDo[1]
		return Go(commandGo)
	case string(commandDo[0]) == "взять":
		commandGo := commandDo[1]
		return Take(commandGo)
	case string(commandDo[0]) == "надеть":
		return PutOn()
	case string(commandDo[0]) == "применить":
		commandGo1 := commandDo[1]
		commandGo2 := commandDo[2]
		return Apply(commandGo1, commandGo2)
	default:
		return "неизвестная команда"
	}
}
