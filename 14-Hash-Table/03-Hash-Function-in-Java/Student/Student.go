package Student
import (
	"hash/fnv"
	"strconv"
)
type Student struct{
	grade int
	cls int
	firstName string
	lastName string
}
func NewStudent(grade ,cls int,firstName,lastName string) *Student{
	return &Student{grade:grade,cls :cls,firstName:firstName,lastName:lastName}
}

func (this *Student)HashCode() int{
	B := 31
	hash := 0
	hash = hash * B + int(HashCode(this.grade))
	hash = hash * B + int(HashCode(this.cls))
	hash = hash * B + int(HashCode(this.firstName))
	hash = hash * B + int(HashCode(this.lastName))
	return hash
}
func (this *Student)Equals(another Student) bool{
	if this == &another{
		return true
	}
	if nil == &another {
		return false
	}
	return this.grade == another.grade &&
	this.cls == another.cls &&
	this.firstName == another.firstName &&
	this.lastName == another.lastName;
}

func HashCode(source interface{}) uint64 {
	switch source.(type) {
	case int:
		source = strconv.Itoa(source.(int))
	case float64:
		source = strconv.FormatFloat(source.(float64),'f',6,64)
	}
	h := fnv.New64a()
	h.Write([]byte(source.(string)))
	return h.Sum64()
}