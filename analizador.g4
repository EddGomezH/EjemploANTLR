grammar analizador;


@parser::header{
import "github.com/emivnajera/Abstract"
import "github.com/emivnajera/Expresiones"
import "github.com/emivnajera/Instrucciones"
import "github.com/emivnajera/TS"
import "strings"
import "reflect"
}

@parser::members{
}

//Tokens
RPRINT: 'print';
RVAR:   'var';
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
MAYOR: '>';
MENOR: '<';
MAYORIGUAL:'>=';
MENORIGUAL:'<=';
IGUALIGUAL:'==';
DISTINTO:'!=';
OR:'||';
AND:'&&';
NOT: '!';
RFUNC:'func';
ID: ([A-Za-z]|'_')([A-Za-z]|[0-9]|'_')*;
IGUAL: '=';
LLAVEA:'{';
LLAVEC:'}';
COMA:',';

start: {instrucciones := [] Abstract.Instruccion{};TSGlobal:=TS.TablaSimbolos{}; var Funciones []interface{};}(instruccion{instrucciones = append(instrucciones,$instruccion.nodo)})*{
for _, n := range instrucciones {
   if(reflect.TypeOf(n).Name() == "Funcion"){
        Funciones = append(Funciones, n.(Instrucciones.Funcion))
    }else{
        n.Interpretar(&TSGlobal, &Funciones)
}
}
};

instruccion returns[Abstract.Instruccion nodo]
:imprimir
{$nodo = $imprimir.nodo}
|declaracion
{$nodo = $declaracion.nodo}
|asignacion
{$nodo = $asignacion.nodo}
|funcion
{$nodo = $funcion.nodo}
|llamada
{$nodo = $llamada.nodo}
;


declaracion returns[Abstract.Instruccion  nodo]
:RVAR ID IGUAL expresion finins
{$nodo = Instrucciones.NewDeclaracion($ID.text, $expresion.nodo, $RVAR.line, $RVAR.pos)}
;

asignacion returns[Abstract.Instruccion nodo]
:ID IGUAL expresion finins
{$nodo = Instrucciones.NewAsignacion($ID.text, $expresion.nodo, $ID.line, $ID.pos)}
;

parametro returns[string id]
:ID {$id= $ID.text}
;

parametros returns[[]string lista]:
parametro{$lista = append($lista, $parametro.id)} (COMA parametro{$lista = append($lista, $parametro.id)})*
;

funcion returns[Abstract.Instruccion nodo]
:RFUNC {instrucciones := [] Abstract.Instruccion{}; parametros := []string{}} ID PARA PARC LLAVEA (instruccion {instrucciones = append(instrucciones,$instruccion.nodo)})* LLAVEC
{$nodo = Instrucciones.NewFuncion($ID.text,instrucciones, parametros, $RFUNC.line, $RFUNC.pos)}
|RFUNC {instrucciones := [] Abstract.Instruccion{};} ID PARA parametros PARC LLAVEA (instruccion {instrucciones = append(instrucciones,$instruccion.nodo)})* LLAVEC
{$nodo = Instrucciones.NewFuncion($ID.text,instrucciones, $parametros.lista, $RFUNC.line, $RFUNC.pos)}
;

parametroll returns[Abstract.Instruccion  nodo]
:expresion {$nodo = $expresion.nodo}
;

parametrolls returns[[]Abstract.Instruccion  lista]
:parametroll{$lista = append($lista, $parametroll.nodo)} (COMA parametroll{$lista = append($lista, $parametroll.nodo)})*
;

llamada returns[Abstract.Instruccion  nodo]
: ID {parametros := []Abstract.Instruccion{}}PARA PARC finins
{$nodo = Instrucciones.NewLlamada($ID.text, parametros, $ID.line, $ID.pos)}
| ID PARA parametrolls PARC finins
{$nodo = Instrucciones.NewLlamada($ID.text, $parametrolls.lista, $ID.line, $ID.pos)}
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
|<assoc=left>opi=expresion DISTINTO opd=expresion
{$nodo = Expresiones.NewRelacional(TS.DIFERENTE, $opi.nodo, $opd.nodo, $DISTINTO.line, $DISTINTO.pos)}
|<assoc=left>opi=expresion MAYOR opd=expresion
{$nodo = Expresiones.NewRelacional(TS.MAYORQUE, $opi.nodo, $opd.nodo, $MAYOR.line, $MAYOR.pos)}
|<assoc=left>opi=expresion IGUALIGUAL opd=expresion
{$nodo = Expresiones.NewRelacional(TS.IGUALIGUAL, $opi.nodo, $opd.nodo, $IGUALIGUAL.line, $IGUALIGUAL.pos)}
|<assoc=left>opi=expresion MENOR opd=expresion
{$nodo = Expresiones.NewRelacional(TS.MENORQUE, $opi.nodo, $opd.nodo, $MENOR.line, $MENOR.pos)}
|<assoc=left>opi=expresion MAYORIGUAL opd=expresion
{$nodo = Expresiones.NewRelacional(TS.MAYORIGUAL, $opi.nodo, $opd.nodo, $MAYORIGUAL.line, $MAYORIGUAL.pos)}
|<assoc=left>opi=expresion MENORIGUAL opd=expresion
{$nodo = Expresiones.NewRelacional(TS.MENORIGUAL, $opi.nodo, $opd.nodo, $MENORIGUAL.line, $MENORIGUAL.pos)}
|<assoc=right>NOT expresion
{$nodo = Expresiones.NewLogica(TS.NOT, $expresion.nodo, $expresion.nodo, $NOT.line, $NOT.pos)}
|<assoc=left>opi=expresion OR opd=expresion
{$nodo = Expresiones.NewLogica(TS.OR, $opi.nodo, $opd.nodo, $OR.line, $OR.pos)}
|<assoc=left>opi=expresion AND opd=expresion
{$nodo = Expresiones.NewLogica(TS.AND, $opi.nodo, $opd.nodo, $AND.line, $AND.pos)}
|ID
{$nodo = Expresiones.NewIdentificador($ID.text, $ID.line, $ID.pos)}
;


imprimir returns[Abstract.Instruccion  nodo]: RPRINT PARA expresion PARC finins
{$nodo = Instrucciones.NewImprimir($expresion.nodo,$RPRINT.line,$RPRINT.pos)};

finins: PUNTOCOMA;

COMMENT: '//' ( ~[\r\n]+)?->skip;
WS : [ \t\r\n]+ -> skip ;


