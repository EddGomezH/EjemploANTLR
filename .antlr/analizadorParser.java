// Generated from /home/emivnajera/go/src/PersonalProjects/EjemploANTLR/analizador.g4 by ANTLR 4.8

import "github.com/emivnajera/Abstract"
import "github.com/emivnajera/Expresiones"
import "github.com/emivnajera/Instrucciones"
import "github.com/emivnajera/TS"
import "strings"
import "reflect"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class analizadorParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		RPRINT=1, RVAR=2, PUNTOCOMA=3, PARA=4, PARC=5, ENTERO=6, FLOTANTE=7, RTRUE=8, 
		RFALSE=9, CHAR=10, STRING=11, MAS=12, MENOS=13, MUL=14, DIV=15, MOD=16, 
		MAYOR=17, MENOR=18, MAYORIGUAL=19, MENORIGUAL=20, IGUALIGUAL=21, DISTINTO=22, 
		OR=23, AND=24, NOT=25, RFUNC=26, ID=27, IGUAL=28, LLAVEA=29, LLAVEC=30, 
		COMA=31, COMMENT=32, WS=33;
	public static final int
		RULE_start = 0, RULE_instruccion = 1, RULE_declaracion = 2, RULE_asignacion = 3, 
		RULE_parametro = 4, RULE_parametros = 5, RULE_funcion = 6, RULE_expresion = 7, 
		RULE_imprimir = 8, RULE_finins = 9;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "instruccion", "declaracion", "asignacion", "parametro", "parametros", 
			"funcion", "expresion", "imprimir", "finins"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'print'", "'var'", "';'", "'('", "')'", null, null, "'true'", 
			"'false'", null, null, "'+'", "'-'", "'*'", "'/'", "'%'", "'>'", "'<'", 
			"'>='", "'<='", "'=='", "'!='", "'||'", "'&&'", "'!'", "'func'", null, 
			"'='", "'{'", "'}'", "','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "RPRINT", "RVAR", "PUNTOCOMA", "PARA", "PARC", "ENTERO", "FLOTANTE", 
			"RTRUE", "RFALSE", "CHAR", "STRING", "MAS", "MENOS", "MUL", "DIV", "MOD", 
			"MAYOR", "MENOR", "MAYORIGUAL", "MENORIGUAL", "IGUALIGUAL", "DISTINTO", 
			"OR", "AND", "NOT", "RFUNC", "ID", "IGUAL", "LLAVEA", "LLAVEC", "COMA", 
			"COMMENT", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "analizador.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }



	public analizadorParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public InstruccionContext instruccion;
		public List<InstruccionContext> instruccion() {
			return getRuleContexts(InstruccionContext.class);
		}
		public InstruccionContext instruccion(int i) {
			return getRuleContext(InstruccionContext.class,i);
		}
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			instrucciones := [] Abstract.Instruccion{};TSGlobal:=TS.TablaSimbolos{}; var Funciones []interface{};
			setState(26);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RPRINT) | (1L << RVAR) | (1L << RFUNC) | (1L << ID))) != 0)) {
				{
				{
				setState(21);
				((StartContext)_localctx).instruccion = instruccion();
				instrucciones = append(instrucciones,((StartContext)_localctx).instruccion.nodo)
				}
				}
				setState(28);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}

			for _, n := range instrucciones {
			   if(reflect.TypeOf(n).Name() == "Funcion"){
			        fmt.Println("Creacion de Funcion")
			        Funciones = append(Funciones, n.(Instrucciones.Funcion))
			    }else{
			        n.Interpretar(&TSGlobal, &Funciones)
			}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstruccionContext extends ParserRuleContext {
		public Abstract.Instruccion nodo;
		public ImprimirContext imprimir;
		public DeclaracionContext declaracion;
		public AsignacionContext asignacion;
		public FuncionContext funcion;
		public ImprimirContext imprimir() {
			return getRuleContext(ImprimirContext.class,0);
		}
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public AsignacionContext asignacion() {
			return getRuleContext(AsignacionContext.class,0);
		}
		public FuncionContext funcion() {
			return getRuleContext(FuncionContext.class,0);
		}
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_instruccion);
		try {
			setState(43);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RPRINT:
				enterOuterAlt(_localctx, 1);
				{
				setState(31);
				((InstruccionContext)_localctx).imprimir = imprimir();
				_localctx.nodo = ((InstruccionContext)_localctx).imprimir.nodo
				}
				break;
			case RVAR:
				enterOuterAlt(_localctx, 2);
				{
				setState(34);
				((InstruccionContext)_localctx).declaracion = declaracion();
				_localctx.nodo = ((InstruccionContext)_localctx).declaracion.nodo
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(37);
				((InstruccionContext)_localctx).asignacion = asignacion();
				_localctx.nodo = ((InstruccionContext)_localctx).asignacion.nodo
				}
				break;
			case RFUNC:
				enterOuterAlt(_localctx, 4);
				{
				setState(40);
				((InstruccionContext)_localctx).funcion = funcion();
				_localctx.nodo = ((InstruccionContext)_localctx).funcion.nodo
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DeclaracionContext extends ParserRuleContext {
		public Abstract.Instruccion nodo;
		public Token RVAR;
		public Token ID;
		public ExpresionContext expresion;
		public TerminalNode RVAR() { return getToken(analizadorParser.RVAR, 0); }
		public TerminalNode ID() { return getToken(analizadorParser.ID, 0); }
		public TerminalNode IGUAL() { return getToken(analizadorParser.IGUAL, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public FininsContext finins() {
			return getRuleContext(FininsContext.class,0);
		}
		public DeclaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion; }
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(45);
			((DeclaracionContext)_localctx).RVAR = match(RVAR);
			setState(46);
			((DeclaracionContext)_localctx).ID = match(ID);
			setState(47);
			match(IGUAL);
			setState(48);
			((DeclaracionContext)_localctx).expresion = expresion(0);
			setState(49);
			finins();
			_localctx.nodo = Instrucciones.NewDeclaracion((((DeclaracionContext)_localctx).ID!=null?((DeclaracionContext)_localctx).ID.getText():null), ((DeclaracionContext)_localctx).expresion.nodo, (((DeclaracionContext)_localctx).RVAR!=null?((DeclaracionContext)_localctx).RVAR.getLine():0), (((DeclaracionContext)_localctx).RVAR!=null?((DeclaracionContext)_localctx).RVAR.getCharPositionInLine():0))
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AsignacionContext extends ParserRuleContext {
		public Abstract.Instruccion nodo;
		public Token ID;
		public ExpresionContext expresion;
		public TerminalNode ID() { return getToken(analizadorParser.ID, 0); }
		public TerminalNode IGUAL() { return getToken(analizadorParser.IGUAL, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public FininsContext finins() {
			return getRuleContext(FininsContext.class,0);
		}
		public AsignacionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacion; }
	}

	public final AsignacionContext asignacion() throws RecognitionException {
		AsignacionContext _localctx = new AsignacionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_asignacion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(52);
			((AsignacionContext)_localctx).ID = match(ID);
			setState(53);
			match(IGUAL);
			setState(54);
			((AsignacionContext)_localctx).expresion = expresion(0);
			setState(55);
			finins();
			_localctx.nodo = Instrucciones.NewAsignacion((((AsignacionContext)_localctx).ID!=null?((AsignacionContext)_localctx).ID.getText():null), ((AsignacionContext)_localctx).expresion.nodo, (((AsignacionContext)_localctx).ID!=null?((AsignacionContext)_localctx).ID.getLine():0), (((AsignacionContext)_localctx).ID!=null?((AsignacionContext)_localctx).ID.getCharPositionInLine():0))
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParametroContext extends ParserRuleContext {
		public string id;
		public Token ID;
		public TerminalNode ID() { return getToken(analizadorParser.ID, 0); }
		public ParametroContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametro; }
	}

	public final ParametroContext parametro() throws RecognitionException {
		ParametroContext _localctx = new ParametroContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_parametro);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(58);
			((ParametroContext)_localctx).ID = match(ID);
			_localctx.id= (((ParametroContext)_localctx).ID!=null?((ParametroContext)_localctx).ID.getText():null)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParametrosContext extends ParserRuleContext {
		public []string lista;
		public ParametroContext parametro;
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public List<TerminalNode> COMA() { return getTokens(analizadorParser.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(analizadorParser.COMA, i);
		}
		public ParametrosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametros; }
	}

	public final ParametrosContext parametros() throws RecognitionException {
		ParametrosContext _localctx = new ParametrosContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_parametros);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(61);
			((ParametrosContext)_localctx).parametro = parametro();
			_localctx.lista = append(_localctx.lista, ((ParametrosContext)_localctx).parametro.id)
			setState(69);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(63);
				match(COMA);
				setState(64);
				((ParametrosContext)_localctx).parametro = parametro();
				_localctx.lista = append(_localctx.lista, ((ParametrosContext)_localctx).parametro.id)
				}
				}
				setState(71);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FuncionContext extends ParserRuleContext {
		public Abstract.Instruccion nodo;
		public Token RFUNC;
		public Token ID;
		public InstruccionContext instruccion;
		public ParametrosContext parametros;
		public TerminalNode RFUNC() { return getToken(analizadorParser.RFUNC, 0); }
		public TerminalNode ID() { return getToken(analizadorParser.ID, 0); }
		public TerminalNode PARA() { return getToken(analizadorParser.PARA, 0); }
		public TerminalNode PARC() { return getToken(analizadorParser.PARC, 0); }
		public TerminalNode LLAVEA() { return getToken(analizadorParser.LLAVEA, 0); }
		public TerminalNode LLAVEC() { return getToken(analizadorParser.LLAVEC, 0); }
		public List<InstruccionContext> instruccion() {
			return getRuleContexts(InstruccionContext.class);
		}
		public InstruccionContext instruccion(int i) {
			return getRuleContext(InstruccionContext.class,i);
		}
		public ParametrosContext parametros() {
			return getRuleContext(ParametrosContext.class,0);
		}
		public FuncionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcion; }
	}

	public final FuncionContext funcion() throws RecognitionException {
		FuncionContext _localctx = new FuncionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_funcion);
		int _la;
		try {
			setState(106);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(72);
				((FuncionContext)_localctx).RFUNC = match(RFUNC);
				instrucciones := [] Abstract.Instruccion{}; parametros := []string{}
				setState(74);
				((FuncionContext)_localctx).ID = match(ID);
				setState(75);
				match(PARA);
				setState(76);
				match(PARC);
				setState(77);
				match(LLAVEA);
				setState(83);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RPRINT) | (1L << RVAR) | (1L << RFUNC) | (1L << ID))) != 0)) {
					{
					{
					setState(78);
					((FuncionContext)_localctx).instruccion = instruccion();
					instrucciones = append(instrucciones,((FuncionContext)_localctx).instruccion.nodo)
					}
					}
					setState(85);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(86);
				match(LLAVEC);
				_localctx.nodo = Instrucciones.NewFuncion((((FuncionContext)_localctx).ID!=null?((FuncionContext)_localctx).ID.getText():null),instrucciones, parametros, (((FuncionContext)_localctx).RFUNC!=null?((FuncionContext)_localctx).RFUNC.getLine():0), (((FuncionContext)_localctx).RFUNC!=null?((FuncionContext)_localctx).RFUNC.getCharPositionInLine():0))
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(88);
				((FuncionContext)_localctx).RFUNC = match(RFUNC);
				instrucciones := [] Abstract.Instruccion{};
				setState(90);
				((FuncionContext)_localctx).ID = match(ID);
				setState(91);
				match(PARA);
				setState(92);
				((FuncionContext)_localctx).parametros = parametros();
				setState(93);
				match(PARC);
				setState(94);
				match(LLAVEA);
				setState(100);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RPRINT) | (1L << RVAR) | (1L << RFUNC) | (1L << ID))) != 0)) {
					{
					{
					setState(95);
					((FuncionContext)_localctx).instruccion = instruccion();
					instrucciones = append(instrucciones,((FuncionContext)_localctx).instruccion.nodo)
					}
					}
					setState(102);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(103);
				match(LLAVEC);
				_localctx.nodo = Instrucciones.NewFuncion((((FuncionContext)_localctx).ID!=null?((FuncionContext)_localctx).ID.getText():null),instrucciones, ((FuncionContext)_localctx).parametros.lista, (((FuncionContext)_localctx).RFUNC!=null?((FuncionContext)_localctx).RFUNC.getLine():0), (((FuncionContext)_localctx).RFUNC!=null?((FuncionContext)_localctx).RFUNC.getCharPositionInLine():0))
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpresionContext extends ParserRuleContext {
		public Abstract.Instruccion nodo;
		public ExpresionContext opi;
		public Token ENTERO;
		public Token FLOTANTE;
		public Token RTRUE;
		public Token RFALSE;
		public Token CHAR;
		public Token STRING;
		public Token MENOS;
		public ExpresionContext expresion;
		public Token NOT;
		public Token ID;
		public Token MUL;
		public ExpresionContext opd;
		public Token DIV;
		public Token MOD;
		public Token MAS;
		public Token DISTINTO;
		public Token MAYOR;
		public Token IGUALIGUAL;
		public Token MENOR;
		public Token MAYORIGUAL;
		public Token MENORIGUAL;
		public Token OR;
		public Token AND;
		public TerminalNode ENTERO() { return getToken(analizadorParser.ENTERO, 0); }
		public TerminalNode FLOTANTE() { return getToken(analizadorParser.FLOTANTE, 0); }
		public TerminalNode RTRUE() { return getToken(analizadorParser.RTRUE, 0); }
		public TerminalNode RFALSE() { return getToken(analizadorParser.RFALSE, 0); }
		public TerminalNode CHAR() { return getToken(analizadorParser.CHAR, 0); }
		public TerminalNode STRING() { return getToken(analizadorParser.STRING, 0); }
		public TerminalNode MENOS() { return getToken(analizadorParser.MENOS, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode NOT() { return getToken(analizadorParser.NOT, 0); }
		public TerminalNode ID() { return getToken(analizadorParser.ID, 0); }
		public TerminalNode MUL() { return getToken(analizadorParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(analizadorParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(analizadorParser.MOD, 0); }
		public TerminalNode MAS() { return getToken(analizadorParser.MAS, 0); }
		public TerminalNode DISTINTO() { return getToken(analizadorParser.DISTINTO, 0); }
		public TerminalNode MAYOR() { return getToken(analizadorParser.MAYOR, 0); }
		public TerminalNode IGUALIGUAL() { return getToken(analizadorParser.IGUALIGUAL, 0); }
		public TerminalNode MENOR() { return getToken(analizadorParser.MENOR, 0); }
		public TerminalNode MAYORIGUAL() { return getToken(analizadorParser.MAYORIGUAL, 0); }
		public TerminalNode MENORIGUAL() { return getToken(analizadorParser.MENORIGUAL, 0); }
		public TerminalNode OR() { return getToken(analizadorParser.OR, 0); }
		public TerminalNode AND() { return getToken(analizadorParser.AND, 0); }
		public ExpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresion; }
	}

	public final ExpresionContext expresion() throws RecognitionException {
		return expresion(0);
	}

	private ExpresionContext expresion(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpresionContext _localctx = new ExpresionContext(_ctx, _parentState);
		ExpresionContext _prevctx = _localctx;
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_expresion, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(131);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENTERO:
				{
				setState(109);
				((ExpresionContext)_localctx).ENTERO = match(ENTERO);
				_localctx.nodo = Expresiones.NewPrimitivo((((ExpresionContext)_localctx).ENTERO!=null?((ExpresionContext)_localctx).ENTERO.getText():null), TS.ENTERO, (((ExpresionContext)_localctx).ENTERO!=null?((ExpresionContext)_localctx).ENTERO.getLine():0), (((ExpresionContext)_localctx).ENTERO!=null?((ExpresionContext)_localctx).ENTERO.getCharPositionInLine():0))
				}
				break;
			case FLOTANTE:
				{
				setState(111);
				((ExpresionContext)_localctx).FLOTANTE = match(FLOTANTE);
				_localctx.nodo = Expresiones.NewPrimitivo((((ExpresionContext)_localctx).FLOTANTE!=null?((ExpresionContext)_localctx).FLOTANTE.getText():null),TS.FLOAT, (((ExpresionContext)_localctx).FLOTANTE!=null?((ExpresionContext)_localctx).FLOTANTE.getLine():0), (((ExpresionContext)_localctx).FLOTANTE!=null?((ExpresionContext)_localctx).FLOTANTE.getCharPositionInLine():0))
				}
				break;
			case RTRUE:
				{
				setState(113);
				((ExpresionContext)_localctx).RTRUE = match(RTRUE);
				_localctx.nodo = Expresiones.NewPrimitivo((((ExpresionContext)_localctx).RTRUE!=null?((ExpresionContext)_localctx).RTRUE.getText():null), TS.BOOLEAN,(((ExpresionContext)_localctx).RTRUE!=null?((ExpresionContext)_localctx).RTRUE.getLine():0), (((ExpresionContext)_localctx).RTRUE!=null?((ExpresionContext)_localctx).RTRUE.getCharPositionInLine():0))
				}
				break;
			case RFALSE:
				{
				setState(115);
				((ExpresionContext)_localctx).RFALSE = match(RFALSE);
				_localctx.nodo = Expresiones.NewPrimitivo((((ExpresionContext)_localctx).RFALSE!=null?((ExpresionContext)_localctx).RFALSE.getText():null), TS.BOOLEAN, (((ExpresionContext)_localctx).RFALSE!=null?((ExpresionContext)_localctx).RFALSE.getLine():0), (((ExpresionContext)_localctx).RFALSE!=null?((ExpresionContext)_localctx).RFALSE.getCharPositionInLine():0))
				}
				break;
			case CHAR:
				{
				setState(117);
				((ExpresionContext)_localctx).CHAR = match(CHAR);
				_localctx.nodo = Expresiones.NewPrimitivo(strings.Replace((((ExpresionContext)_localctx).CHAR!=null?((ExpresionContext)_localctx).CHAR.getText():null),"'","",2), TS.CARACTER, (((ExpresionContext)_localctx).CHAR!=null?((ExpresionContext)_localctx).CHAR.getLine():0), (((ExpresionContext)_localctx).CHAR!=null?((ExpresionContext)_localctx).CHAR.getCharPositionInLine():0))
				}
				break;
			case STRING:
				{
				setState(119);
				((ExpresionContext)_localctx).STRING = match(STRING);
				_localctx.nodo = Expresiones.NewPrimitivo(strings.Replace((((ExpresionContext)_localctx).STRING!=null?((ExpresionContext)_localctx).STRING.getText():null),"\"","",-1),TS.CADENA, (((ExpresionContext)_localctx).STRING!=null?((ExpresionContext)_localctx).STRING.getLine():0), (((ExpresionContext)_localctx).STRING!=null?((ExpresionContext)_localctx).STRING.getCharPositionInLine():0))
				}
				break;
			case MENOS:
				{
				setState(121);
				((ExpresionContext)_localctx).MENOS = match(MENOS);
				setState(122);
				((ExpresionContext)_localctx).expresion = expresion(16);
				_localctx.nodo = Expresiones.NewAritmetica(TS.MENOS, ((ExpresionContext)_localctx).expresion.nodo, nil, (((ExpresionContext)_localctx).MENOS!=null?((ExpresionContext)_localctx).MENOS.getLine():0), (((ExpresionContext)_localctx).MENOS!=null?((ExpresionContext)_localctx).MENOS.getCharPositionInLine():0))
				}
				break;
			case NOT:
				{
				setState(125);
				((ExpresionContext)_localctx).NOT = match(NOT);
				setState(126);
				((ExpresionContext)_localctx).expresion = expresion(4);
				_localctx.nodo = Expresiones.NewLogica(TS.NOT, ((ExpresionContext)_localctx).expresion.nodo, ((ExpresionContext)_localctx).expresion.nodo, (((ExpresionContext)_localctx).NOT!=null?((ExpresionContext)_localctx).NOT.getLine():0), (((ExpresionContext)_localctx).NOT!=null?((ExpresionContext)_localctx).NOT.getCharPositionInLine():0))
				}
				break;
			case ID:
				{
				setState(129);
				((ExpresionContext)_localctx).ID = match(ID);
				_localctx.nodo = Expresiones.NewIdentificador((((ExpresionContext)_localctx).ID!=null?((ExpresionContext)_localctx).ID.getText():null), (((ExpresionContext)_localctx).ID!=null?((ExpresionContext)_localctx).ID.getLine():0), (((ExpresionContext)_localctx).ID!=null?((ExpresionContext)_localctx).ID.getCharPositionInLine():0))
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(200);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(198);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(133);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(134);
						((ExpresionContext)_localctx).MUL = match(MUL);
						setState(135);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(16);
						_localctx.nodo = Expresiones.NewAritmetica(TS.POR, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MUL!=null?((ExpresionContext)_localctx).MUL.getLine():0), (((ExpresionContext)_localctx).MUL!=null?((ExpresionContext)_localctx).MUL.getCharPositionInLine():0))
						}
						break;
					case 2:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(138);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(139);
						((ExpresionContext)_localctx).DIV = match(DIV);
						setState(140);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(15);
						_localctx.nodo = Expresiones.NewAritmetica(TS.DIV, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).DIV!=null?((ExpresionContext)_localctx).DIV.getLine():0), (((ExpresionContext)_localctx).DIV!=null?((ExpresionContext)_localctx).DIV.getCharPositionInLine():0))
						}
						break;
					case 3:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(143);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(144);
						((ExpresionContext)_localctx).MOD = match(MOD);
						setState(145);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(14);
						_localctx.nodo = Expresiones.NewAritmetica(TS.MOD, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MOD!=null?((ExpresionContext)_localctx).MOD.getLine():0), (((ExpresionContext)_localctx).MOD!=null?((ExpresionContext)_localctx).MOD.getCharPositionInLine():0))
						}
						break;
					case 4:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(148);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(149);
						((ExpresionContext)_localctx).MAS = match(MAS);
						setState(150);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(13);
						_localctx.nodo = Expresiones.NewAritmetica(TS.MAS, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MAS!=null?((ExpresionContext)_localctx).MAS.getLine():0), (((ExpresionContext)_localctx).MAS!=null?((ExpresionContext)_localctx).MAS.getCharPositionInLine():0))
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(153);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(154);
						((ExpresionContext)_localctx).MENOS = match(MENOS);
						setState(155);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(12);
						_localctx.nodo = Expresiones.NewAritmetica(TS.MENOS, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MENOS!=null?((ExpresionContext)_localctx).MENOS.getLine():0), (((ExpresionContext)_localctx).MENOS!=null?((ExpresionContext)_localctx).MENOS.getCharPositionInLine():0))
						}
						break;
					case 6:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(158);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(159);
						((ExpresionContext)_localctx).DISTINTO = match(DISTINTO);
						setState(160);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(11);
						_localctx.nodo = Expresiones.NewRelacional(TS.DIFERENTE, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).DISTINTO!=null?((ExpresionContext)_localctx).DISTINTO.getLine():0), (((ExpresionContext)_localctx).DISTINTO!=null?((ExpresionContext)_localctx).DISTINTO.getCharPositionInLine():0))
						}
						break;
					case 7:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(163);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(164);
						((ExpresionContext)_localctx).MAYOR = match(MAYOR);
						setState(165);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(10);
						_localctx.nodo = Expresiones.NewRelacional(TS.MAYORQUE, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MAYOR!=null?((ExpresionContext)_localctx).MAYOR.getLine():0), (((ExpresionContext)_localctx).MAYOR!=null?((ExpresionContext)_localctx).MAYOR.getCharPositionInLine():0))
						}
						break;
					case 8:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(168);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(169);
						((ExpresionContext)_localctx).IGUALIGUAL = match(IGUALIGUAL);
						setState(170);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(9);
						_localctx.nodo = Expresiones.NewRelacional(TS.IGUALIGUAL, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).IGUALIGUAL!=null?((ExpresionContext)_localctx).IGUALIGUAL.getLine():0), (((ExpresionContext)_localctx).IGUALIGUAL!=null?((ExpresionContext)_localctx).IGUALIGUAL.getCharPositionInLine():0))
						}
						break;
					case 9:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(173);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(174);
						((ExpresionContext)_localctx).MENOR = match(MENOR);
						setState(175);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(8);
						_localctx.nodo = Expresiones.NewRelacional(TS.MENORQUE, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MENOR!=null?((ExpresionContext)_localctx).MENOR.getLine():0), (((ExpresionContext)_localctx).MENOR!=null?((ExpresionContext)_localctx).MENOR.getCharPositionInLine():0))
						}
						break;
					case 10:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(178);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(179);
						((ExpresionContext)_localctx).MAYORIGUAL = match(MAYORIGUAL);
						setState(180);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(7);
						_localctx.nodo = Expresiones.NewRelacional(TS.MAYORIGUAL, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MAYORIGUAL!=null?((ExpresionContext)_localctx).MAYORIGUAL.getLine():0), (((ExpresionContext)_localctx).MAYORIGUAL!=null?((ExpresionContext)_localctx).MAYORIGUAL.getCharPositionInLine():0))
						}
						break;
					case 11:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(183);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(184);
						((ExpresionContext)_localctx).MENORIGUAL = match(MENORIGUAL);
						setState(185);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(6);
						_localctx.nodo = Expresiones.NewRelacional(TS.MENORIGUAL, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MENORIGUAL!=null?((ExpresionContext)_localctx).MENORIGUAL.getLine():0), (((ExpresionContext)_localctx).MENORIGUAL!=null?((ExpresionContext)_localctx).MENORIGUAL.getCharPositionInLine():0))
						}
						break;
					case 12:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(188);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(189);
						((ExpresionContext)_localctx).OR = match(OR);
						setState(190);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(4);
						_localctx.nodo = Expresiones.NewLogica(TS.OR, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).OR!=null?((ExpresionContext)_localctx).OR.getLine():0), (((ExpresionContext)_localctx).OR!=null?((ExpresionContext)_localctx).OR.getCharPositionInLine():0))
						}
						break;
					case 13:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(193);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(194);
						((ExpresionContext)_localctx).AND = match(AND);
						setState(195);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(3);
						_localctx.nodo = Expresiones.NewLogica(TS.AND, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).AND!=null?((ExpresionContext)_localctx).AND.getLine():0), (((ExpresionContext)_localctx).AND!=null?((ExpresionContext)_localctx).AND.getCharPositionInLine():0))
						}
						break;
					}
					} 
				}
				setState(202);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ImprimirContext extends ParserRuleContext {
		public Abstract.Instruccion nodo;
		public Token RPRINT;
		public ExpresionContext expresion;
		public TerminalNode RPRINT() { return getToken(analizadorParser.RPRINT, 0); }
		public TerminalNode PARA() { return getToken(analizadorParser.PARA, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode PARC() { return getToken(analizadorParser.PARC, 0); }
		public FininsContext finins() {
			return getRuleContext(FininsContext.class,0);
		}
		public ImprimirContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_imprimir; }
	}

	public final ImprimirContext imprimir() throws RecognitionException {
		ImprimirContext _localctx = new ImprimirContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_imprimir);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(203);
			((ImprimirContext)_localctx).RPRINT = match(RPRINT);
			setState(204);
			match(PARA);
			setState(205);
			((ImprimirContext)_localctx).expresion = expresion(0);
			setState(206);
			match(PARC);
			setState(207);
			finins();
			_localctx.nodo = Instrucciones.NewImprimir(((ImprimirContext)_localctx).expresion.nodo,(((ImprimirContext)_localctx).RPRINT!=null?((ImprimirContext)_localctx).RPRINT.getLine():0),(((ImprimirContext)_localctx).RPRINT!=null?((ImprimirContext)_localctx).RPRINT.getCharPositionInLine():0))
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FininsContext extends ParserRuleContext {
		public TerminalNode PUNTOCOMA() { return getToken(analizadorParser.PUNTOCOMA, 0); }
		public FininsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_finins; }
	}

	public final FininsContext finins() throws RecognitionException {
		FininsContext _localctx = new FininsContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_finins);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(210);
			match(PUNTOCOMA);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 7:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 15);
		case 1:
			return precpred(_ctx, 14);
		case 2:
			return precpred(_ctx, 13);
		case 3:
			return precpred(_ctx, 12);
		case 4:
			return precpred(_ctx, 11);
		case 5:
			return precpred(_ctx, 10);
		case 6:
			return precpred(_ctx, 9);
		case 7:
			return precpred(_ctx, 8);
		case 8:
			return precpred(_ctx, 7);
		case 9:
			return precpred(_ctx, 6);
		case 10:
			return precpred(_ctx, 5);
		case 11:
			return precpred(_ctx, 3);
		case 12:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3#\u00d7\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\3\2\3\2\3\2\3\2\7\2\33\n\2\f\2\16\2\36\13\2\3\2\3\2\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\5\3.\n\3\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\7\7F\n"+
		"\7\f\7\16\7I\13\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\7\bT\n\b\f\b\16"+
		"\bW\13\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\7\be\n\b\f\b"+
		"\16\bh\13\b\3\b\3\b\3\b\5\bm\n\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u0086\n\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\7\t\u00c9\n\t\f\t\16"+
		"\t\u00cc\13\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\2\3\20\f\2\4"+
		"\6\b\n\f\16\20\22\24\2\2\2\u00e9\2\26\3\2\2\2\4-\3\2\2\2\6/\3\2\2\2\b"+
		"\66\3\2\2\2\n<\3\2\2\2\f?\3\2\2\2\16l\3\2\2\2\20\u0085\3\2\2\2\22\u00cd"+
		"\3\2\2\2\24\u00d4\3\2\2\2\26\34\b\2\1\2\27\30\5\4\3\2\30\31\b\2\1\2\31"+
		"\33\3\2\2\2\32\27\3\2\2\2\33\36\3\2\2\2\34\32\3\2\2\2\34\35\3\2\2\2\35"+
		"\37\3\2\2\2\36\34\3\2\2\2\37 \b\2\1\2 \3\3\2\2\2!\"\5\22\n\2\"#\b\3\1"+
		"\2#.\3\2\2\2$%\5\6\4\2%&\b\3\1\2&.\3\2\2\2\'(\5\b\5\2()\b\3\1\2).\3\2"+
		"\2\2*+\5\16\b\2+,\b\3\1\2,.\3\2\2\2-!\3\2\2\2-$\3\2\2\2-\'\3\2\2\2-*\3"+
		"\2\2\2.\5\3\2\2\2/\60\7\4\2\2\60\61\7\35\2\2\61\62\7\36\2\2\62\63\5\20"+
		"\t\2\63\64\5\24\13\2\64\65\b\4\1\2\65\7\3\2\2\2\66\67\7\35\2\2\678\7\36"+
		"\2\289\5\20\t\29:\5\24\13\2:;\b\5\1\2;\t\3\2\2\2<=\7\35\2\2=>\b\6\1\2"+
		">\13\3\2\2\2?@\5\n\6\2@G\b\7\1\2AB\7!\2\2BC\5\n\6\2CD\b\7\1\2DF\3\2\2"+
		"\2EA\3\2\2\2FI\3\2\2\2GE\3\2\2\2GH\3\2\2\2H\r\3\2\2\2IG\3\2\2\2JK\7\34"+
		"\2\2KL\b\b\1\2LM\7\35\2\2MN\7\6\2\2NO\7\7\2\2OU\7\37\2\2PQ\5\4\3\2QR\b"+
		"\b\1\2RT\3\2\2\2SP\3\2\2\2TW\3\2\2\2US\3\2\2\2UV\3\2\2\2VX\3\2\2\2WU\3"+
		"\2\2\2XY\7 \2\2Ym\b\b\1\2Z[\7\34\2\2[\\\b\b\1\2\\]\7\35\2\2]^\7\6\2\2"+
		"^_\5\f\7\2_`\7\7\2\2`f\7\37\2\2ab\5\4\3\2bc\b\b\1\2ce\3\2\2\2da\3\2\2"+
		"\2eh\3\2\2\2fd\3\2\2\2fg\3\2\2\2gi\3\2\2\2hf\3\2\2\2ij\7 \2\2jk\b\b\1"+
		"\2km\3\2\2\2lJ\3\2\2\2lZ\3\2\2\2m\17\3\2\2\2no\b\t\1\2op\7\b\2\2p\u0086"+
		"\b\t\1\2qr\7\t\2\2r\u0086\b\t\1\2st\7\n\2\2t\u0086\b\t\1\2uv\7\13\2\2"+
		"v\u0086\b\t\1\2wx\7\f\2\2x\u0086\b\t\1\2yz\7\r\2\2z\u0086\b\t\1\2{|\7"+
		"\17\2\2|}\5\20\t\22}~\b\t\1\2~\u0086\3\2\2\2\177\u0080\7\33\2\2\u0080"+
		"\u0081\5\20\t\6\u0081\u0082\b\t\1\2\u0082\u0086\3\2\2\2\u0083\u0084\7"+
		"\35\2\2\u0084\u0086\b\t\1\2\u0085n\3\2\2\2\u0085q\3\2\2\2\u0085s\3\2\2"+
		"\2\u0085u\3\2\2\2\u0085w\3\2\2\2\u0085y\3\2\2\2\u0085{\3\2\2\2\u0085\177"+
		"\3\2\2\2\u0085\u0083\3\2\2\2\u0086\u00ca\3\2\2\2\u0087\u0088\f\21\2\2"+
		"\u0088\u0089\7\20\2\2\u0089\u008a\5\20\t\22\u008a\u008b\b\t\1\2\u008b"+
		"\u00c9\3\2\2\2\u008c\u008d\f\20\2\2\u008d\u008e\7\21\2\2\u008e\u008f\5"+
		"\20\t\21\u008f\u0090\b\t\1\2\u0090\u00c9\3\2\2\2\u0091\u0092\f\17\2\2"+
		"\u0092\u0093\7\22\2\2\u0093\u0094\5\20\t\20\u0094\u0095\b\t\1\2\u0095"+
		"\u00c9\3\2\2\2\u0096\u0097\f\16\2\2\u0097\u0098\7\16\2\2\u0098\u0099\5"+
		"\20\t\17\u0099\u009a\b\t\1\2\u009a\u00c9\3\2\2\2\u009b\u009c\f\r\2\2\u009c"+
		"\u009d\7\17\2\2\u009d\u009e\5\20\t\16\u009e\u009f\b\t\1\2\u009f\u00c9"+
		"\3\2\2\2\u00a0\u00a1\f\f\2\2\u00a1\u00a2\7\30\2\2\u00a2\u00a3\5\20\t\r"+
		"\u00a3\u00a4\b\t\1\2\u00a4\u00c9\3\2\2\2\u00a5\u00a6\f\13\2\2\u00a6\u00a7"+
		"\7\23\2\2\u00a7\u00a8\5\20\t\f\u00a8\u00a9\b\t\1\2\u00a9\u00c9\3\2\2\2"+
		"\u00aa\u00ab\f\n\2\2\u00ab\u00ac\7\27\2\2\u00ac\u00ad\5\20\t\13\u00ad"+
		"\u00ae\b\t\1\2\u00ae\u00c9\3\2\2\2\u00af\u00b0\f\t\2\2\u00b0\u00b1\7\24"+
		"\2\2\u00b1\u00b2\5\20\t\n\u00b2\u00b3\b\t\1\2\u00b3\u00c9\3\2\2\2\u00b4"+
		"\u00b5\f\b\2\2\u00b5\u00b6\7\25\2\2\u00b6\u00b7\5\20\t\t\u00b7\u00b8\b"+
		"\t\1\2\u00b8\u00c9\3\2\2\2\u00b9\u00ba\f\7\2\2\u00ba\u00bb\7\26\2\2\u00bb"+
		"\u00bc\5\20\t\b\u00bc\u00bd\b\t\1\2\u00bd\u00c9\3\2\2\2\u00be\u00bf\f"+
		"\5\2\2\u00bf\u00c0\7\31\2\2\u00c0\u00c1\5\20\t\6\u00c1\u00c2\b\t\1\2\u00c2"+
		"\u00c9\3\2\2\2\u00c3\u00c4\f\4\2\2\u00c4\u00c5\7\32\2\2\u00c5\u00c6\5"+
		"\20\t\5\u00c6\u00c7\b\t\1\2\u00c7\u00c9\3\2\2\2\u00c8\u0087\3\2\2\2\u00c8"+
		"\u008c\3\2\2\2\u00c8\u0091\3\2\2\2\u00c8\u0096\3\2\2\2\u00c8\u009b\3\2"+
		"\2\2\u00c8\u00a0\3\2\2\2\u00c8\u00a5\3\2\2\2\u00c8\u00aa\3\2\2\2\u00c8"+
		"\u00af\3\2\2\2\u00c8\u00b4\3\2\2\2\u00c8\u00b9\3\2\2\2\u00c8\u00be\3\2"+
		"\2\2\u00c8\u00c3\3\2\2\2\u00c9\u00cc\3\2\2\2\u00ca\u00c8\3\2\2\2\u00ca"+
		"\u00cb\3\2\2\2\u00cb\21\3\2\2\2\u00cc\u00ca\3\2\2\2\u00cd\u00ce\7\3\2"+
		"\2\u00ce\u00cf\7\6\2\2\u00cf\u00d0\5\20\t\2\u00d0\u00d1\7\7\2\2\u00d1"+
		"\u00d2\5\24\13\2\u00d2\u00d3\b\n\1\2\u00d3\23\3\2\2\2\u00d4\u00d5\7\5"+
		"\2\2\u00d5\25\3\2\2\2\13\34-GUfl\u0085\u00c8\u00ca";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}