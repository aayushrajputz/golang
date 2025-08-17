package main

import "fmt"

func main() {
	studentGrades := make(map[string]int)

	studentGrades["aayush"] = 100
	studentGrades["paras"]=56
	studentGrades["sahil"]=78
	studentGrades["pranav"]=89
	studentGrades["pranavi"]=90

	fmt.Println("makrs of parnav:",studentGrades["pranav"])
	studentGrades["pranav"]=98
	fmt.Println("new marks of pranav:",studentGrades["pranav"])

	delete(studentGrades,"sahil")
	fmt.Println(studentGrades["sahil"])

	grades, exists := studentGrades["sahil"]
	fmt.Println("grades of sahil:",grades)
	fmt.Println("shail exists:",exists)
	fmt.Println("length of map:",len(studentGrades))

	for index, value := range studentGrades{
		fmt.Printf("key is %s and value is %d\n",index,value)
	}
	person := map[string]int{
		"alice": 25,
		"bob": 30,
		"charlie": 35,
	}
	for index, value := range person{
		fmt.Printf("===========key is %s and value is %d\n",index,value)
	}
}