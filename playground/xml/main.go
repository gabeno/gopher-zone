package main

import (
	"encoding/xml"
	"fmt"
)

// field tags contain directives for encoder and decoder
type Plant struct {
	XMLName xml.Name `xml:"plant"`   // name of the xml element representing this struct
	Id      int      `xml:"id,attr"` // id is an attributed rather than a nested element
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	//
	// marshall
	//

	coffee := Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Kenya", "Ethiopia"}

	// emit xml
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// add a generic xml header to the output
	fmt.Println(xml.Header + string(out))

	//
	// unmarshall
	//

	// data structure to use to unmarshall
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	// nesting
	tomato := Plant{Id: 83, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{&coffee, &tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
