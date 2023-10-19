package pattern

type IObject interface {
  GetName() string
  SetName(name string)
}

type Object struct {
  Name string
}

func (p Object) GetName() string {
  return "Object " + p.Name
}

func (p Object) SetName(name string) {
  p.Name = name
}

type One struct {
  Object
}

func NewOne () IObject {
  return &One {
    Object: Object {
      Name : "one"
    }
  }
}

type Two struct {
  Object
}

func NewTwo () IObject {
  return &Two {
    Object: Object {
      Name : "two"
    }
  }
}

func getObject(objType string) IObject {
  switch objType {
    case "one":
      return NewOne()
    case "two":
      return NewTwo()
  }
}

/* пример вывода 

func main() {
  one := getObject("one")
  two := getObject("two")

  fmt.Println("Product:", one.GetName())
  fmt.Println("Product:", two.GetName())
}

*/


//	Паттерн "фабричный метод" (Factory Method) является порождающим паттерном проектирования, который предоставляет интерфейс для создания объектов, но оставляет решение о конкретных классах создаваемых объектов на подклассах. Этот паттерн позволяет делегировать создание объектов подклассам, что способствует более гибкой архитектуре программы.
