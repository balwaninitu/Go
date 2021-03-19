package main

func main() {

	var (
		minecraft        = game{title: "Minecraft", price: 5}
		worldOfWarcraft  = game{title: "World of warcraft", price: 19}
		eliteDangerous   = game{title: "Elite Dangerous", price: 54}
		candleInTomb     = book{title: "Candle in the tomb", price: 20}
		barneyAndFriends = book{title: "Barney and Friends", price: 10}
		razerBT          = computerAccessories{title: "Razer BT earpiece", price: 159}
		razerKeyboard    = computerAccessories{title: "Razer Keyboard", price: 110}
		logitechMouse    = computerAccessories{title: "Logitech Mouse", price: 80}
	)

	var store list

	store = append(store, &minecraft, &worldOfWarcraft, &eliteDangerous, &candleInTomb, &barneyAndFriends, &razerBT, &razerKeyboard, &logitechMouse)

	store.print()

}
