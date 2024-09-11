package main

import "fmt"

type Mouse struct {
	UUID  string
	Name  string
	Wired bool
}

// Guard to implement String() for pointer and regular values
var _ fmt.Stringer = Mouse{} // Uncomment func below

// func (m Mouse) String() string {
// 	return m.Name
// }

// Can use it but only for the pointers
// var _ fmt.Stringer = (*Mouse)(nil)
//
// func (m *Mouse) String() string {
// 	return m.Name
// }

func main() {
	m := Mouse{UUID: "1111-2222", Name: "Razer Pro", Wired: true}
	fmt.Println(m)
	fmt.Println(&m)
	fmt.Println((*Mouse)(nil))

}
