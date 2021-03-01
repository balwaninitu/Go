package main

type bankbalance struct {
	currency string
	amount   float64
}

func (b *bankbalance) deposit(amt float64) float64 {
	//deposited := b.amount + amt
	return b.amount + amt
}

// func (b *bankbalance) withdraw(amt float64) float64 {
// 	withdraw := b.amount - amt
// 	return withdraw
// }

// func (b *bankbalance) display() {

// }

// func (b *bankbalance) processBankBalance() {
// 	fmt.Println("deosit money")
// 	fmt.Scanln(&b.deposit)
// }

func main() {

	// //	b := bankbalance{}
	// var money float64
	// fmt.Println("deposit money")
	// fmt.Scanln(&money)

	//deposited := b.deposit(5.4)

	//fmt.Println(deposited)

}
