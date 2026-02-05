package main

import (
	"fmt"
)

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type BySalary []Person

func (s BySalary) Len() int      { return len(s) }
func (s BySalary) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s BySalary) Less(i, j int) bool {
	return s[i].Salary > s[j].Salary
}

type ByName []Person

func (n ByName) Len() int      { return len(n) }
func (n ByName) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool {
    return n[i].Name < n[j].Name
}

func main() {
	// TODO: Implementare custom sorting
	fmt.Println("Custom Sort with sort.Interface")
}
