package main

import (
	"fmt"
	"time"
)

type Sex string

const (
	Male   Sex = "Male"
	Female Sex = "Female"
)

type Person struct {
	Name         string
	BirthdayDate time.Time
}

type Patient struct {
	Person
	Sex       Sex
	Height    int
	Weight    float32
	Diagnosis string
}

func (p Person) Age() int {
	now := time.Now()
	age := now.Year() - p.BirthdayDate.Year()

	if now.Month() < p.BirthdayDate.Month() ||
		(now.Month() == p.BirthdayDate.Month() && now.Day() < p.BirthdayDate.Day()) {
		age--
	}

	return age
}

func AverageHeight(patients []Patient) float64 {
	if len(patients) == 0 {
		return 0
	}
	sum := 0
	for _, v := range patients {
		sum += v.Height
	}
	return float64(sum) / float64(len(patients))
}

func ShortestHeight(patients []Patient) int {
	if len(patients) == 0 {
		return 0
	}
	minHeight := patients[0].Height
	for _, v := range patients[1:] {
		if v.Height < minHeight {
			minHeight = v.Height
		}
	}
	return minHeight
}

func main() {
	patients := []Patient{
		{
			Person: Person{
				Name:         "Dmytro",
				BirthdayDate: time.Date(2005, time.November, 15, 0, 0, 0, 0, time.Local),
			},
			Sex:       Male,
			Height:    180,
			Weight:    75.5,
			Diagnosis: "Healthy",
		},
		{
			Person: Person{
				Name:         "Olena",
				BirthdayDate: time.Date(2004, time.May, 20, 0, 0, 0, 0, time.Local),
			},
			Sex:       Female,
			Height:    168,
			Weight:    58.0,
			Diagnosis: "Headache",
		},
	}

	fmt.Printf("%s — вік: %d\n", patients[0].Name, patients[0].Age())
	fmt.Printf("Середній зріст: %.1f см\n", AverageHeight(patients))
	fmt.Printf("Найменший зріст: %d см\n", ShortestHeight(patients))
}
