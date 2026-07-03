package main

func main() {
	day := "sun"

	// if day == "sat" {
	// 	println("Yes, Sleep More")
	// } else {
	// 	println("Work Day")
	// }

	// switch day { // tagged switch
	// case "sat":
	// 	println("Yes, Sleep more")
	// case "fri":
	// 	println("Go to Sleep")
	// case "sun":
	// 	println("go to work")
	// default:
	// 	println("another boring day")
	// }

	switch { // normal switch
	case day == "sat":
		println("Yes, Sleep more")
	case day == "fri":
		println("Go to Sleep")
	case day == "sun":
		println("go to work")
	default:
		println("another boring day")
	}
}
