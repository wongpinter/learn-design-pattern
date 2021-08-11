package builders

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car should be 4 and they were: %v", car.Wheels)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car should be 5 and they were: %v", car.Seats)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car should be 'car'")
	}
}
