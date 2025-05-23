package main

import "fmt"

type validationError struct {
	Message string
}

type notFoundError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		// Karena error berupa interface, kita balikkan berupa pointer
		return &validationError{Message: "id is empty"}
	}

	if id != "nathan" {
		return &notFoundError{Message: "data not found"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)

	// ===== MENGECEK JENIS ERROR =====
	if err != nil {
		// ===== Kalo pake if-else =====
		// ok adalah boolean (true / false)
		//if validationError, ok := err.(*validationError); ok {
		//	fmt.Println(ok) // true
		//	fmt.Println("Validation Error =", validationError.Message)
		//} else if notFoundError, ok := err.(*notFoundError); ok {
		//	fmt.Println(ok) // true
		//	fmt.Println("NotFound Error =", notFoundError.Message)
		//} else {
		//	fmt.Println("Error =", err.Error())
		//}

		// ===== Kalo pake switch =====
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation Error:", finalError.Error())
		case *notFoundError:
			fmt.Println("Not found Error:", finalError.Error())
		default:
			fmt.Println("Unknown error:", finalError.Error())
		}
	} else {
		fmt.Println("Save Data Success")
	}
}
