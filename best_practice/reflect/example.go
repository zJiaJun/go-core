package reflect

import (
	"log"
	"reflect"
)

type person struct {
	Name string
}

func SetVal() {
	i := 10
	ivp := reflect.ValueOf(&i)
	ivp.Elem().SetInt(20)
	log.Println("SetVal i=", i)
}

func SetStructFieldValue() {
	p := &person{Name: "John"}
	ivp := reflect.ValueOf(p)
	elem := ivp.Elem()
	name := elem.FieldByName("Name")
	log.Println(name.CanSet())
	name.SetString("leon")
	log.Println("SetStructFieldValue name=", p.Name)
}

func GetKind() {
	p := person{Name: "John"}
	log.Println("GetKind p=", reflect.ValueOf(p).Kind())
	log.Println("GetKind &p=", reflect.ValueOf(&p).Kind())
}
