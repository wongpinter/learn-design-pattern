package builders

type BuildProccess interface {
	SetWheels() BuildProccess
	SetSeats() BuildProccess
	SetStructure() BuildProccess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProccess
}

func (m *ManufacturingDirector) Construct() {
	m.builder.SetWheels().SetSeats().SetStructure()
}

func (m *ManufacturingDirector) SetBuilder(b BuildProccess) {
	m.builder = b
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type CarBuilder struct {
	attribute VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProccess {
	c.attribute.Wheels = 4

	return c
}

func (c *CarBuilder) SetSeats() BuildProccess {
	c.attribute.Seats = 5

	return c
}

func (c *CarBuilder) SetStructure() BuildProccess {
	c.attribute.Structure = "Car"

	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.attribute
}
