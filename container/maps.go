package main

import "fmt"

func main() {
	m := map[string]string{
		"Name": "Lee Sin",
		"Skills": "Yi Cool",
		"From": "Ya luo lan",
		"Language": "EN",
	}

	m1 := make(map[string] int ) //m1 == empty map

	var m2 map[string] int //m2 = nil

	fmt.Println(m,m1,m2)

	fmt.Println("Traversing Map:")  //map是无序的
	for k,v := range m{
		fmt.Println(k," : ",v)
	}
	fmt.Println("Traversing the key of Map:")
	for k := range m  {
		fmt.Println(k)
	}

	fmt.Println("Getting values from Map:")
	if name ,ok:= m["Name"];ok{
		fmt.Println(name,ok)
	}else {
		fmt.Println("Key does not exist!")
	}


	if skills,ok := m["top"];ok {  //key does not exist ,get the init value of the map
		fmt.Println(skills)
	}else {
		fmt.Println("Key does not exist!")
	}

	fmt.Println("Deleting values!")
	if language,ok := m["Language"];ok {
		fmt.Println(language,ok)
	}else {
		fmt.Println("Key does not exist!")
	}

	delete(m,"Language")
	if language,ok := m["Language"];ok {
		fmt.Println(language,ok)
	}else {
		fmt.Println("Key does not exist!")
	}
}
