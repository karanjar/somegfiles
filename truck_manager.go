package main 

import(
	"errors"
	"fmt"
)
var ErrTruckNotFound = errors.New("truck not found")

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
	m.trucks[id] = &Truck{ID:id,Cargo:cargo}
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
	truck,ok := m.trucks[id]
	if !ok{
		return ErrTruckNotFound
	}
	truck.Cargo = cargo
	return ErrTruckNotFound
}


func NewTruckManager() *TruckManager{
	return &TruckManager{
		trucks: make(map[string]*Truck),
	}
}
func main(){
	mgr := NewTruckManager()

	fmt.Println("Adding the truck")
	err := mgr.AddTruck("T101",100)
	if err != nil{
		fmt.Errorf("error adding truck: %w\n",err)
	}

	truck,err := mgr.GetTruck("T101")
	if err != nil{
		fmt.Errorf("error getting truck: %w",err)
	}
	fmt.Printf("get truck with truck ID:%s and cargo:%d\n",truck.ID,truck.Cargo)
	
	err = mgr.RemoveTruck("T101")
	if err != nil{
		fmt.Errorf("error removing truck: %w\n",err)
	}
	fmt.Println("truck removed")

	err = mgr.UpdateTruckCargo("T102",750)
	if err!=nil{
		fmt.Errorf("error update truck cargo:%w",err)

	}

	fmt.Printf("truck:%s,cargo:%d",truck.ID,truck.Cargo)

}

