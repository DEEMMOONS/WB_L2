package main

import "fmt"

type Request struct {
    amount int
}

type Handler interface {
    SetNext(handler Handler)
    Handle(request Request)
}

type BaseHandler struct {
    next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
    h.next = handler
}

func (h *BaseHandler) Handle(request Request) {
    if h.next != nil {
        h.next.Handle(request)
    }
}

type BankHandler struct {
    BaseHandler
}

func (h *BankHandler) Handle(request Request) {
    if request.amount <= 100 {
        fmt.Println("Request handled by BankHandler")
    } else {
        h.BaseHandler.Handle(request)
    }
}

type ManagerHandler struct {
    BaseHandler
}

func (h *ManagerHandler) Handle(request Request) {
    if request.amount <= 1000 {
        fmt.Println("Request handled by ManagerHandler")
    } else {
        h.BaseHandler.Handle(request)
    }
}

type CEOHandler struct {
    BaseHandler
}

func (h *CEOHandler) Handle(request Request) {
    if request.amount <= 10000 {
        fmt.Println("Request handled by CEOHandler")
    } else {
        fmt.Println("Request can't be handled")
    }
}

/* bankHandler := &BankHandler{}
    managerHandler := &ManagerHandler{}
    ceoHandler := &CEOHandler{}

    bankHandler.SetNext(managerHandler)
    managerHandler.SetNext(ceoHandler)
*/


// Паттерн "цепочка вызовов" (или "цепочка методов") является структурным паттерном проектирования, который используется для создания последовательной цепочки операций, где каждая операция выполняется над объектом, и результат этой операции становится входными данными для следующей операции. Этот паттерн также известен как "Method Chaining". Применяется он для улучшения читаемости кода и упрощения работы с объектами, предоставляя возможность вызывать несколько методов объекта в одной строке.

// Паттерн делает код более читаемым, поскольку каждый вызов метода добавляет ясность к тому, что происходит.
// Но цепочка методов может стать слишком длинной и трудной для понимания, если она неадекватно используется или если включает слишком много методов и может быть проблема с ограничением типов
