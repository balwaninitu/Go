package main

import "fmt"

//Lauret is..
type Laureats struct {
	next  *Laureats
	name  string
	field string
	year  int
}

//Lauret is..
type LaureatsList struct {
	head *Laureats
	name string
}

func createLaureatsList(n string) *LaureatsList {
	return &LaureatsList{
		name: n,
	}
}

func (l_list *LaureatsList) addLaureat(n string, f string, y int) error {
	fmt.Printf("adding %s %s %d\n", n, f, y)

	l := &Laureats{
		name:  n,
		field: f,
		year:  y,
	}

	if l_list.head == nil {
		l_list.head = l
	} else {
		currentNode := l_list.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = l
	}
	return nil
}

func (l_list *LaureatsList) deleteLaureat(n string, f string, y int) error {
	fmt.Printf("deleteing %s %s %d\n", n, f, y)

	currentNode := l_list.head
	if currentNode == nil {
		fmt.Println("Empty list")
		return nil
	}

	// head
	if currentNode.name == n && currentNode.field == f && currentNode.year == y {
		if currentNode == l_list.head {
			l_list.head = currentNode.next
		}
		return nil
	}

	// others
	fmt.Printf("*currentNode %+v\n", *currentNode)
	for currentNode.next != nil {
		fmt.Printf("*currentNode.next: %+v\n", *currentNode.next)
		next := currentNode.next
		if next.name == n && next.field == f && next.year == y {
			fmt.Println("matching laureat found in next")
			currentNode.next = next.next
			break
		}
		currentNode = currentNode.next
	}
	return nil
}

func (l_list *LaureatsList) showAllLaureats() error {
	fmt.Printf("\nCurrent list:\n")
	fmt.Println("-------------")
	currentNode := l_list.head
	if currentNode == nil {
		fmt.Println("Empty list")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}
	fmt.Println("-------------")
	return nil
}

func main() {

	myLaureatsList := createLaureatsList("myLaureatsList")
	fmt.Println("Created LaureatsList")

	myLaureatsList.addLaureat("Marie Curie", "Physics", 1903)
	myLaureatsList.addLaureat("John Steinbeck", "Literature", 1962)
	myLaureatsList.addLaureat("Walther Nernst", "Chemistry", 1920)
	myLaureatsList.addLaureat("Kim Dae-jung", "Peace", 2000)
	myLaureatsList.showAllLaureats()

	myLaureatsList.deleteLaureat("Walther Nernst", "Chemistry", 1920)
	myLaureatsList.deleteLaureat("Marie Curie", "Physics", 1903)
	myLaureatsList.showAllLaureats()

	myLaureatsList.addLaureat("Robert Solow", "Economy", 1987)
	myLaureatsList.showAllLaureats()
}
