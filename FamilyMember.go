package main

import "fmt"

type FamilyMember struct {
	Name string
}

type Family struct {
	Members []FamilyMember
	Pets    []string
}

func ForFamilyMembers() {
	familyMembers := []FamilyMember{
		{Name: "Папа"},
		{Name: "Мама"},
		{Name: "Я"},
		{Name: "Сестра"},
	}

	pets := []string{"Йоркширский терьер"}

	family := Family{
		Members: familyMembers,
		Pets:    pets,
	}
	describeFamilymember(family)
}

func describeFamilymember(family Family) {
	fmt.Println("\nСемья:")
	for _, member := range family.Members {
		fmt.Println("- ", member.Name)
	}

	fmt.Println("\nДомашние питомцы:")
	for _, pet := range family.Pets {
		fmt.Println("- ", pet)
	}
}
