package main

import "fmt"

type Room struct {
	Name      string
	Area      float64
	Height    float64
	Furniture []string
}

type Apartment struct {
	Area   float64
	Height float64
	Rooms  []Room
}

func ForRooms() []Room {
	bedroom := Room{
		Name:      "Спальня",
		Area:      15.0,
		Height:    3.0,
		Furniture: []string{"шкаф-купе", "комод"},
	}

	livingRoom := Room{
		Name:      "Гостиная",
		Area:      20.0,
		Height:    3.0,
		Furniture: []string{"диван", "телевизор"},
	}

	kitchen := Room{
		Name:      "Кухня",
		Area:      10.0,
		Height:    3.0,
		Furniture: []string{"кухонный гарнитур", "микроволновая печь", "газовая плита"},
	}

	apartment := Apartment{
		Area:   46.0,
		Height: 3.0,
		Rooms:  []Room{bedroom, livingRoom, kitchen},
	}
	describeApartment(apartment)
	return []Room{bedroom, livingRoom, kitchen}
}

func describeApartment(apartment Apartment) {
	fmt.Println("Описание квартиры:")
	fmt.Printf("- Площадь: %.1f квадратных метров\n", apartment.Area)
	fmt.Printf("- Высота потолков: %.1f метров\n", apartment.Height)
	fmt.Printf("- Количество комнат: %d\n", len(apartment.Rooms))

	fmt.Println("\nМебель в комнатах:")
	for _, room := range apartment.Rooms {
		fmt.Printf("- %s:\n", room.Name)
		for _, furniture := range room.Furniture {
			fmt.Printf("  - %s\n", furniture)
		}
	}
}
