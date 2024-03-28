package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name     string
	Gender   string
	Father   []*Person // to track Father's Name
	Children []*Person // track for Childrens
	Wives    []*Person // field to track wives
}

type FamilyTree struct {
	People map[string]*Person
}

func NewFamilyTree() *FamilyTree {
	return &FamilyTree{
		People: make(map[string]*Person),
	}
}

func (ft *FamilyTree) AddPerson(name, gender string) {
	ft.People[name] = &Person{
		Name:   name,
		Gender: gender,
	}
}

func (ft *FamilyTree) ConnectRelationship(name1, relation, name2 string) {
	person1 := ft.People[name1]
	person2 := ft.People[name2]

	if person1 == nil || person2 == nil {
		fmt.Println("Error: Person not found")
		return
	}

	switch relation {
	case "son": // Handle Son & Father's relationship
		person2.Children = append(person2.Children, person1)
		person1.Father = append(person1.Father, person2)
	case "daughter": // Handle daughter relationship
		person2.Children = append(person2.Children, person1)
		person1.Father = append(person1.Father, person2)
	case "wife": // Handle wife relationship
		person1.Wives = append(person1.Wives, person2)
		person2.Wives = append(person2.Wives, person1)
	default:
		fmt.Println("Error: Invalid relationship")
	}

}
func (ft *FamilyTree) FatherOf(name string) string {
	person := ft.People[name]
	if person == nil {
		return "Error: Person not found"
	}

	for _, child := range person.Father {
		if child.Gender == "male" {
			return child.Name
		}
	}
	return "Unknown" // If father not found
}

func (ft *FamilyTree) CountSons(name string) int {
	person := ft.People[name]
	if person == nil {
		fmt.Println("Error: Person not found")
		return 0
	}

	sons := 0
	for _, child := range person.Children {
		if child.Gender == "male" {
			sons++
		}
	}
	return sons
}

func (ft *FamilyTree) CountDaughters(name string) int {
	person := ft.People[name]
	if person == nil {
		fmt.Println("Error: Person not found")
		return 0
	}

	daughters := 0
	for _, child := range person.Children {
		if child.Gender == "female" {
			daughters++
		}
	}
	return daughters
}

func (ft *FamilyTree) CountWives(name string) int {
	person := ft.People[name]
	if person == nil {
		fmt.Println("Error: Person not found")
		return 0
	}

	wives := 0
	for _, partner := range person.Wives {
		if partner.Gender == "female" {
			wives++
		}
	}
	return wives
}

func main() {
	ft := NewFamilyTree()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("family-tree> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "add":
			if len(args) >= 4 {
				ft.AddPerson(args[1]+" "+args[2], args[3]) // Assuming gender provided as third argument
			}
		case "connect":
			if len(args) >= 8 && args[3] == "as" && args[5] == "of" {
				ft.ConnectRelationship(args[1]+" "+args[2], args[4], args[6]+" "+args[7])
			}
		case "count":
			if len(args) >= 5 && args[2] == "of" {
				switch args[1] {
				case "sons":
					fmt.Println(ft.CountSons(args[3] + " " + args[4]))
				case "daughters":
					fmt.Println(ft.CountDaughters(args[3] + " " + args[4]))
				case "wives":
					fmt.Println(ft.CountWives(args[3] + " " + args[4]))
				}
			}
		case "father":
			if len(args) >= 4 && args[1] == "of" {
				fmt.Println(ft.FatherOf(args[2] + " " + args[3]))
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid Argument, Please try again!")
		}
	}
}
