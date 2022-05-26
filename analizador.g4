grammar analizador;


@parser::header{
import "github.com/emivnajera/Abstract"
import "github.com/emivnajera/Expresiones"
import "github.com/emivnajera/Instrucciones"
import "github.com/emivnajera/TS"
import "strings"
}

@parser::members{
}

//Tokens
RPRINT: 'print';
PUNTOCOMA: ';';
PARA: '(';
PARC: ')';
ENTERO: [0-9]+;
FLOTANTE: [0-9]+'.'[0-9]+;
RTRUE: 'true';
RFALSE: 'false';
CHAR: '\''([A-Za-z]|WS|[0-9])'\'';
STRING: '"'([A-Za-z]|WS|[0-9])*'"';
MAS: '+';
MENOS: '-';
MUL: '*';
DIV: '/';
MOD: '%';

start: {instrucciones := [] Abstract.Instruccion{};TSGlobal:=TS.TablaSimbolos{}}(instruccion{instrucciones = append(instrucciones,$instruccion.nodo)})*{
for _, n := range instrucciones {
        n.Interpretar(&TSGlobal)
}
};


instruccion returns[Abstract.Instruccion nodo]
:imprimir
{$nodo = $imprimir.nodo}
;


expresion returns[Abstract.Instruccion  nodo]
:ENTERO
{$nodo = Expresiones.NewPrimitivo($ENTERO.text, TS.ENTERO, $ENTERO.line, $ENTERO.pos)}
|FLOTANTE
{$nodo = Expresiones.NewPrimitivo($FLOTANTE.text,TS.FLOAT, $FLOTANTE.line, $FLOTANTE.pos)}
|RTRUE
{$nodo = Expresiones.NewPrimitivo($RTRUE.text, TS.BOOLEAN,$RTRUE.line, $RTRUE.pos)}
|RFALSE
{$nodo = Expresiones.NewPrimitivo($RFALSE.text, TS.BOOLEAN, $RFALSE.line, $RFALSE.pos)}
|CHAR
{$nodo = Expresiones.NewPrimitivo(strings.Replace($CHAR.text,"'","",2), TS.CARACTER, $CHAR.line, $CHAR.pos)}
|STRING
{$nodo = Expresiones.NewPrimitivo(strings.Replace($STRING.text,"\"","",-1),TS.CADENA, $STRING.line, $STRING.pos)}
|<assoc=right>MENOS expresion
{$nodo = Expresiones.NewAritmetica(TS.MENOS, $expresion.nodo, nil, $MENOS.line, $MENOS.pos)}
|<assoc=left>opi=expresion MUL opd=expresion
{$nodo = Expresiones.NewAritmetica(TS.POR, $opi.nodo, $opd.nodo, $MUL.line, $MUL.pos)}
|<assoc=left>opi=expresion DIV opd=expresion
{$nodo = Expresiones.NewAritmetica(TS.DIV, $opi.nodo, $opd.nodo, $DIV.line, $DIV.pos)}
|<assoc=left>opi=expresion MOD opd=expresion
{$nodo = Expresiones.NewAritmetica(TS.MOD, $opi.nodo, $opd.nodo, $MOD.line, $MOD.pos)}
|<assoc=left>opi=expresion MAS opd=expresion
{$nodo = Expresiones.NewAritmetica(TS.MAS, $opi.nodo, $opd.nodo, $MAS.line, $MAS.pos)}
|<assoc=left>opi=expresion MENOS opd=expresion
{$nodo = Expresiones.NewAritmetica(TS.MENOS, $opi.nodo, $opd.nodo, $MENOS.line, $MENOS.pos)}
;


imprimir returns[Abstract.Instruccion  nodo]: RPRINT PARA expresion PARC finins
{$nodo = Instrucciones.NewImprimir($expresion.nodo,$RPRINT.line,$RPRINT.pos)};

finins: PUNTOCOMA;

COMMENT: '//' ( ~[\r\n]+)?->skip;
WS : [ \t\r\n]+ -> skip ;


