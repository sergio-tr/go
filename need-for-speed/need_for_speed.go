package speed

const InitialBattery int = 100
const InitialDistance int = 0

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	if speed < 0 {
		panic("Speed cannot be negative")
	}
	if batteryDrain < 0 {
		panic("Battery drain cannot be negative")
	}

	return Car{
		battery:      InitialBattery,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     InitialDistance,
	}
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	if distance < 0 {
		panic("Distance cannot be negative")
	}

	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}

	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	numberDrives := car.battery / car.batteryDrain

	if car.battery%car.batteryDrain != 0 {
		numberDrives++
	}

	return numberDrives*car.speed >= track.distance
}
