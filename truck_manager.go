package main 

import(
	"error"
)
var ErrTruckNotFound = error.New("truck not found")

type FleetManager interface{
	AddTruck(id string,cargo int) error
	GetTruck(id string) (*Truck,error)
	RemoveTruck(id string )error
	UpdateTruckCargo(id string,cargo int) error
}

type Truck struct{
	ID string
	Cargo int
}
type TruckManager struct{
	trucks map[string]*Truck
}
func (m *TruckManager) AddTruck(id string,cargo int) error{
	m.trucks := &Truck{ID:id,Cargo:cargo}
	return nil

}
func (m *TruckManager) GetTruck(id string) (*Truck,error){
	truck,ok := m.trucks[id]
	if !ok{
		return nil,ErrTruckNotFound
	}
	return truck,nil
}
func (m *TruckManager) RemoveTruck(id string) error{
	delete(m.trucks,id)
	return nil
}
func (m *TruckManager) UpdateTruckCargo(id string,cargo int) error{
	m.trucks[id] = cargo
	return nil
}

func TruckCompany(truck FleetManager) error{
	if 
}

func NewTruckManager() TruckManager{
	return TruckManager{
		trucks: make(map[string]*truck),
	}
}
