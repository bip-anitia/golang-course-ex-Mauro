package main

import (
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

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type BySalary []Person

func (s BySalary) Len() int           { return len(s) }
func (s BySalary) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s BySalary) Less(i, j int) bool { return s[i].Salary > s[j].Salary } // desc

type ByName []Person

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return n[i].Name < n[j].Name }

func main() {
	people := []Person{
		{"Alice", 28, 75000, "New York", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"Bob", 35, 95000, "Boston", time.Date(2018, 6, 15, 0, 0, 0, 0, time.UTC)},
		{"Charlie", 28, 85000, "New York", time.Date(2019, 3, 10, 0, 0, 0, 0, time.UTC)},
		{"Diana", 42, 110000, "Seattle", time.Date(2015, 9, 1, 0, 0, 0, 0, time.UTC)},
	}

	sort.Sort(ByAge(people))               // età crescente
	sort.Sort(sort.Reverse(ByAge(people))) // età decrescente
	sort.Sort(BySalary(people))            // salario decrescente
	sort.Sort(ByName(people))              // nome alfabetico
}
