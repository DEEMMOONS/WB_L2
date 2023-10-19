package pattern

import (
	"fmt"
)

type Visitable interface {
	Accept(Visitor)
}

type IntegerElement struct {
	Value int
}

func (i *IntegerElement) Accept(visitor Visitor) {
	visitor.VisitInteger(i)
}

type StringElement struct {
	Value string
}

func (s *StringElement) Accept(visitor Visitor) {
	visitor.VisitString(s)
}

type Visitor interface {
	VisitInteger(*IntegerElement)
	VisitString(*StringElement)
}

type PrintVisitor struct{}

func (p *PrintVisitor) VisitInteger(element *IntegerElement) {
	fmt.Printf("integer: %d\n", element.Value)
}

func (p *PrintVisitor) VisitString(element *StringElement) {
	fmt.Printf("string : %s\n", element.Value)
}

 


//Паттерн "Посетитель" (Visitor) - это поведенческий паттерн проектирования, который позволяет добавлять новые операции к объектам без изменения их классов.
//Основным плюсом является разделение кода: Паттерн "Посетитель" помогает разделить код операций от структуры объектов, что упрощает поддержку и добавление новых операций.
//Но при этом может привести к нарушению  инкапсуляции: Посетитель может требовать открытия доступа к внутренним данным объектов, что может нарушить инкапсуляцию.
