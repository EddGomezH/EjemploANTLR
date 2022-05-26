package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/emivnajera/parser"
)

type TreeSyntax struct {
	*parser.BaseanalizadorListener
}

func NewTreeSyntax() *TreeSyntax {
	return new(TreeSyntax)
}

func interpretar(in string) {
	b := []byte("")
	err := ioutil.WriteFile("consola.txt", b, 0644)
	if err != nil {
		log.Fatal(err)
	}

	is, _ := antlr.NewFileStream("entrada.txt")

	// crear analizador lexico
	lexer := parser.NewanalizadorLexer(is)

	/*for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf(" %s ( %q ) \n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}*/

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// crear analizador sintactico
	p := parser.NewanalizadorParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(NewTreeSyntax(), p.Start())

	bytesLeidos, err := ioutil.ReadFile("consola.txt")
	if err != nil {
		log.Fatal(err)
	}

	contenido := string(bytesLeidos)
	fmt.Println(contenido)

}

func main() {
	file, err := os.Open("entrada.txt")

	if err != nil {
		log.Fatal(err)
	}
	// se invoca hasta el final
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	text := ""

	for fileScanner.Scan() {
		text += fileScanner.Text()
	}

	// accion de interpretar
	interpretar(text)
}
