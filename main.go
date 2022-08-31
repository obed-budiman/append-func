package main

//create a system to handle insert, delete, take a look to array (append like function)
func main() {

	newAppend := NewAppend()
	newAppend.Insert(1)
	newAppend.CheckArray()

	newAppend.Insert(10)
	newAppend.Insert(20)
	newAppend.CheckArray()

	newAppend.RemoveLast()
	newAppend.CheckArray()

	newAppend.Insert(30)
	newAppend.Insert(40)
	newAppend.CheckArray()

	newAppend.Insert(50)
	newAppend.CheckArray()
}
