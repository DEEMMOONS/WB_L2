package pattern
import "fmt"

 
type Button struct {
    isOn bool
}

func (l *Button) on() {
    l.isOn = true
    fmt.Println("Button On")
}

func (l *Button) off() {
    l.isOn = false
    fmt.Println("Button off")
}

 
type Command interface {
    execute()
}

 
type ButtonOnCommand struct {
    Button *Button
}

func (c *ButtonOnCommand) execute() {
    c.Button.on()
}

type ButtonOffCommand struct {
    Button *Button
}

func (c *ButtonOffCommand) execute() {
    c.Button.off()
}

type RemoteControl struct {
    command Command
}

func (rc *RemoteControl) pressButton() {
    rc.command.execute()
}


//Паттерн Command относится к поведенческим паттернам уровня объекта.
//Отделение отправителя от получателя: Команда позволяет разделить объект, который инициирует операцию (отправитель), от объекта, который выполняет операцию (получатель). Это может быть полезно, например, при реализации меню, кнопок, планирования операций и т. д.
//Паттерн позволяет легко добавить логгирование операций.
//Но при этом  может привести к увеличению количества классов: Внедрение команд может привести к увеличению числа классов в системе, что может усложнить код и его понимание. 
