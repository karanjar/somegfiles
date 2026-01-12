package main

import(
	"fmt"
	"errors"
	"log"

)
var(
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound = errors.New("truck not found ")
)
type inTruck interface{
	loadTruck() error
	unloadTruck() error
}
type Truck struct{
	id string
	cargo int 
}

type electricTruck struct {
	id string
	cargo int
	battery float64

}
func (t *Truck) loadTruck() error{
	return nil
}
func (t *Truck) unloadTruck() error{
	return nil
}
func (e *electricTruck)loadTruck() error{
	return nil
}
func (e *electricTruck) unloadTruck() error{
	return nil
}

func processTruck(truck inTruck) error{
	

	
	if err := truck.loadTruck(); err != nil{
		return fmt.Errorf("error loadind cargo: %w",err)
	}

	if err := truck.unloadTruck(); err != nil{
		return fmt.Errorf("error unloading the cargo: %w",err)
	}

	return nil
}
func main(){

	truck := Truck{
		id: "truck-1",
	}	
		
	etruck := electricTruck{
		id:"e fr",
	}

	if err := processTruck(&truck); err != nil{
		log.Fatalf("Error processing the truck: %s",err)	
	}


	if err := processTruck(&etruck); err != nil{
		log.Fatalf("there was an error processing the electric truck: %s",err)
	}

	fmt.Printf("processed trucks %+v\n",truck)
	fmt.Printf("processed electric trucks %+v\n",etruck)






	
}
