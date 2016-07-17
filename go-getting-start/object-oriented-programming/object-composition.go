package main 

func main() {
	mp := enhancedMessagePrinter{messagePrinter{"foo"}}
	mp.printMessage()
}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage(){
	println(mp.message)
}

type enhancedMessagePrinter struct {
	messagePrinter
}