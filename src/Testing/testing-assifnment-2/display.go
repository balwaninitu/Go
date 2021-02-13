package main

const (
	title = "Dental Appointment System"

	display = `  
1.Make appointment
2.List of Doctor and their availability
3.Search for available Doctor
4.Edit appointment
Select your choices:`
)

// type appointment struct{
// 	patientName string
// 	dateOfAppointment int
// 	timeOfAppointment int
// 	doctorName string
// 	next *appointment
// }

// type doctorList struct {
// 	name string
// 	head *appointment
// }

// func createAppointment (n string) *doctorList {
// 	return &doctorList{
// 		name: n,
// 	}
// }

// func (d *doctorList) addAppointmentDetails(pn string,dn string, dt int, t int)error {
// 	appt := &appointment{
// 		patientName: pn,
// 		doctorName: dn,
// 		dateOfAppointment: dt,
// 		timeOfAppointment: t,

// 	}
// 	if d.head == nil{
// 		d.head = appt
// 	}else {

// 	}

// }

// type appointment struct {
// 	next       *appointment
// 	doctorName string
// 	date       int
// 	month int
// 	time       int
// }

// type doctorList struct {
// 	head          *appointment

// }

// func

// func (d *doctorList) addDetails(n string, dt int, t int) error {
// 	fmt.Printf("Adding details %s %d %d\n", n, dt, t)
// 	dr := &appointment{
// 		doctorName: n,
// 		date:       dt,
// 		time:       t,
// 	}
// 	if d.head == nil {
// 		d.head = dr
// 	} else {
// 		currentNode := d.head
// 		for currentNode.next != nil {
// 			currentNode = currentNode.next
// 		}
// 		currentNode.next = dr
// 	}
// 	return nil
// }

// func main() {

// 	// docter1 := doctorList{"Sunny Je", "Mon-Fri", "10-6"}
// 	// docter2 := doctorList{"Joyce Tan", "Mon-Fri", "10-6"}
// 	// docter3 := doctorList{"Anthony Lim", "Mon-Fri", "10-6"}

// 	// //fmt.Printf("Doctor:%s Avaialble Days:%s,Available Time:%s\n",doctorName,doctorList.availableDay,doctorList.availableTime)

// // myAppointment := createAppointment("myAppointment")
// // fmt.Println("Created my appointment")

// }
