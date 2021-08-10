package testdata

const constId = "ID"

var myId = "Id"

type StructWithId struct {
	Id string
}

type InterfaceWithId interface {
	GetById(myId string) (otherId string)
}

func GetId(id string) {}

func MapSomethingToId(something int) string {
	someId := "some Id"
	return myId + someId
}

func GetID() (MyId string) {
	return constId
}

func MyIdAndSomeOtherIdAndId() {}
