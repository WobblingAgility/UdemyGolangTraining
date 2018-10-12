package main

import "fmt"

type vehicle struct {
	name  string
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	f150 := truck{
		vehicle: vehicle{
			name:  `F150`,
			doors: 2,
			color: `navy`,
		},
		fourWheel: true,
	}

	mazda3 := sedan{
		vehicle: vehicle{
			name:  `Mazda 3`,
			doors: 5,
			color: `silver`,
		},
		luxury: false,
	}

	fmt.Printf("The lovely %v is a %T with %d doors and comes in %s.\nBut does it have All wheel drive? %v\n", f150.name, f150, f150.doors, f150.color, f150.fourWheel)
	fmt.Printf("The lovely %v is a %T with %d doors and comes in %s.\nBut does it have All luxury options? %v\n", mazda3.name, mazda3, mazda3.doors, mazda3.color, mazda3.luxury)
}
