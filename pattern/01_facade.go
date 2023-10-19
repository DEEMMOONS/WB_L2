package pattern

import "fmt"

type Wallet struct{}

func (w *Wallet) Balance(id int) {
  fmt.Printf("Get balance data from user with id = %d\n", id)
}

type AddUser struct{}

func (au *AddUser) Add(id int) {
  fmt.Printf("Add user with id = ", id)
}

type Notification struct{}

func (n *Notification) SendWalletnotification(id int) {
  fmt.Printf("Send notification to user with id = ", id)
}

//Паттерн "Фасад" (Facade) - паттерн проектирования, который предоставляет унифицированный объект для доступа к набору компонентов в подсистеме.
//Он скрывает сложность и детали взаимодействия между клиентом и подсистемой, предоставляя простой объект, с которым клиент может взаимодействовать.
//Минусами являются необходимость в создании дополнительных абстракций и переусложнение поддержки при большом количестве подсистем.
