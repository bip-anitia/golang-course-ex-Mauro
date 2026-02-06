package main

import (
	"fmt"
	"sort"
	"time"
)

type Person struct {
	Name     string
	Age      int
	Salary   float64
	City     string
	JoinDate time.Time
}

type Movie struct {
	Title    string
	Year     int
	Rating   float64
	Duration int
	Genre    string
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool {
	if a[i].Age != a[j].Age {
		return a[i].Age < a[j].Age
	}
	return a[i].Name < a[j].Name
}

type BySalary []Person

func (s BySalary) Len() int           { return len(s) }
func (s BySalary) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s BySalary) Less(i, j int) bool {
	if s[i].Salary != s[j].Salary {
		return s[i].Salary > s[j].Salary
	}
	return s[i].Name < s[j].Name
}

type ByName []Person

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool {
	if n[i].Name != n[j].Name {
		return n[i].Name < n[j].Name
	}
	return n[i].Age < n[j].Age
}

func movieScore(movie Movie) float64 {
	if movie.Duration <= 0 {
		return 0
	}
	return movie.Rating * float64(movie.Year) / float64(movie.Duration)
}

func printPeople(title string, people []Person) {
	fmt.Println(title)
	for index, person := range people {
		fmt.Printf("  %d. %-10s (%d years, $%.0f, %s)\n", index+1, person.Name, person.Age, person.Salary, person.City)
	}
	fmt.Println()
}

func printMovies(title string, movies []Movie) {
	fmt.Println(title)
	for index, movie := range movies {
		fmt.Printf("  %d. %-18s (%d, %.1f, %dmin, %s)\n", index+1, movie.Title, movie.Year, movie.Rating, movie.Duration, movie.Genre)
	}
	fmt.Println()
}

func printMoviesWithScore(title string, movies []Movie) {
	fmt.Println(title)
	for index, movie := range movies {
		fmt.Printf("  %d. %-18s (Score: %.1f)\n", index+1, movie.Title, movieScore(movie))
	}
	fmt.Println()
}

func main() {
	people := []Person{
		{"Alice", 28, 75000, "New York", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"Bob", 35, 95000, "Boston", time.Date(2018, 6, 15, 0, 0, 0, 0, time.UTC)},
		{"Charlie", 28, 85000, "New York", time.Date(2019, 3, 10, 0, 0, 0, 0, time.UTC)},
		{"Diana", 42, 110000, "Seattle", time.Date(2015, 9, 1, 0, 0, 0, 0, time.UTC)},
	}

	movies := []Movie{
		{"The Godfather", 1972, 9.2, 175, "Crime"},
		{"The Dark Knight", 2008, 9.0, 152, "Action"},
		{"Inception", 2010, 8.8, 148, "Sci-Fi"},
		{"Interstellar", 2014, 8.6, 169, "Sci-Fi"},
	}

	printPeople("Original people:", people)

	peopleByName := append([]Person(nil), people...)
	sort.Sort(ByName(peopleByName))
	printPeople("Sorted by Name:", peopleByName)

	peopleByAge := append([]Person(nil), people...)
	sort.Sort(ByAge(peopleByAge))
	printPeople("Sorted by Age:", peopleByAge)

	peopleByAgeDesc := append([]Person(nil), people...)
	sort.Sort(sort.Reverse(ByAge(peopleByAgeDesc)))
	printPeople("Sorted by Age (descending):", peopleByAgeDesc)

	peopleBySalary := append([]Person(nil), people...)
	sort.Sort(BySalary(peopleBySalary))
	printPeople("Sorted by Salary (descending):", peopleBySalary)

	peopleByCityAge := append([]Person(nil), people...)
	sort.Slice(peopleByCityAge, func(i, j int) bool {
		if peopleByCityAge[i].City != peopleByCityAge[j].City {
			return peopleByCityAge[i].City < peopleByCityAge[j].City
		}
		return peopleByCityAge[i].Age < peopleByCityAge[j].Age
	})
	printPeople("Sorted by City, then Age:", peopleByCityAge)

	peopleByNameLength := append([]Person(nil), people...)
	sort.Slice(peopleByNameLength, func(i, j int) bool {
		firstNameLen := len(peopleByNameLength[i].Name)
		secondNameLen := len(peopleByNameLength[j].Name)
		if firstNameLen != secondNameLen {
			return firstNameLen < secondNameLen
		}
		return peopleByNameLength[i].Name < peopleByNameLength[j].Name
	})
	printPeople("Sorted by Name Length:", peopleByNameLength)

	printMovies("Original movies:", movies)

	moviesByYear := append([]Movie(nil), movies...)
	sort.Slice(moviesByYear, func(i, j int) bool {
		return moviesByYear[i].Year < moviesByYear[j].Year
	})
	printMovies("Movies sorted by Year:", moviesByYear)

	moviesByRating := append([]Movie(nil), movies...)
	sort.SliceStable(moviesByRating, func(i, j int) bool {
		return moviesByRating[i].Rating > moviesByRating[j].Rating
	})
	printMovies("Movies sorted by Rating:", moviesByRating)

	moviesByDuration := append([]Movie(nil), movies...)
	sort.Slice(moviesByDuration, func(i, j int) bool {
		return moviesByDuration[i].Duration < moviesByDuration[j].Duration
	})
	printMovies("Movies sorted by Duration:", moviesByDuration)

	moviesByGenreRating := append([]Movie(nil), movies...)
	sort.Slice(moviesByGenreRating, func(i, j int) bool {
		if moviesByGenreRating[i].Genre != moviesByGenreRating[j].Genre {
			return moviesByGenreRating[i].Genre < moviesByGenreRating[j].Genre
		}
		return moviesByGenreRating[i].Rating > moviesByGenreRating[j].Rating
	})
	printMovies("Movies sorted by Genre, then Rating:", moviesByGenreRating)

	moviesByScore := append([]Movie(nil), movies...)
	sort.Slice(moviesByScore, func(i, j int) bool {
		return movieScore(moviesByScore[i]) > movieScore(moviesByScore[j])
	})
	printMoviesWithScore("Movies sorted by Custom Score:", moviesByScore)
}
