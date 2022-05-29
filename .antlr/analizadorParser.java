// Generated from /home/emivnajera/go/src/PersonalProjects/EjemploANTLR/analizador.g4 by ANTLR 4.9.2

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
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

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
		RULE_parametro = 4, RULE_parametros = 5, RULE_funcion = 6, RULE_parametroll = 7, 
		RULE_parametrolls = 8, RULE_llamada = 9, RULE_expresion = 10, RULE_imprimir = 11, 
		RULE_finins = 12;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "instruccion", "declaracion", "asignacion", "parametro", "parametros", 
			"funcion", "parametroll", "parametrolls", "llamada", "expresion", "imprimir", 
			"finins"
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
			setState(32);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RPRINT) | (1L << RVAR) | (1L << RFUNC) | (1L << ID))) != 0)) {
				{
				{
				setState(27);
				((StartContext)_localctx).instruccion = instruccion();
				instrucciones = append(instrucciones,((StartContext)_localctx).instruccion.nodo)
				}
				}
				setState(34);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}

			for _, n := range instrucciones {
			   if(reflect.TypeOf(n).Name() == "Funcion"){
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
		public LlamadaContext llamada;
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
		public LlamadaContext llamada() {
			return getRuleContext(LlamadaContext.class,0);
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
			setState(52);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(37);
				((InstruccionContext)_localctx).imprimir = imprimir();
				_localctx.nodo = ((InstruccionContext)_localctx).imprimir.nodo
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(40);
				((InstruccionContext)_localctx).declaracion = declaracion();
				_localctx.nodo = ((InstruccionContext)_localctx).declaracion.nodo
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(43);
				((InstruccionContext)_localctx).asignacion = asignacion();
				_localctx.nodo = ((InstruccionContext)_localctx).asignacion.nodo
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(46);
				((InstruccionContext)_localctx).funcion = funcion();
				_localctx.nodo = ((InstruccionContext)_localctx).funcion.nodo
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(49);
				((InstruccionContext)_localctx).llamada = llamada();
				_localctx.nodo = ((InstruccionContext)_localctx).llamada.nodo
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
			setState(54);
			((DeclaracionContext)_localctx).RVAR = match(RVAR);
			setState(55);
			((DeclaracionContext)_localctx).ID = match(ID);
			setState(56);
			match(IGUAL);
			setState(57);
			((DeclaracionContext)_localctx).expresion = expresion(0);
			setState(58);
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
			setState(61);
			((AsignacionContext)_localctx).ID = match(ID);
			setState(62);
			match(IGUAL);
			setState(63);
			((AsignacionContext)_localctx).expresion = expresion(0);
			setState(64);
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
			setState(67);
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
			setState(70);
			((ParametrosContext)_localctx).parametro = parametro();
			_localctx.lista = append(_localctx.lista, ((ParametrosContext)_localctx).parametro.id)
			setState(78);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(72);
				match(COMA);
				setState(73);
				((ParametrosContext)_localctx).parametro = parametro();
				_localctx.lista = append(_localctx.lista, ((ParametrosContext)_localctx).parametro.id)
				}
				}
				setState(80);
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
			setState(115);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(81);
				((FuncionContext)_localctx).RFUNC = match(RFUNC);
				instrucciones := [] Abstract.Instruccion{}; parametros := []string{}
				setState(83);
				((FuncionContext)_localctx).ID = match(ID);
				setState(84);
				match(PARA);
				setState(85);
				match(PARC);
				setState(86);
				match(LLAVEA);
				setState(92);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RPRINT) | (1L << RVAR) | (1L << RFUNC) | (1L << ID))) != 0)) {
					{
					{
					setState(87);
					((FuncionContext)_localctx).instruccion = instruccion();
					instrucciones = append(instrucciones,((FuncionContext)_localctx).instruccion.nodo)
					}
					}
					setState(94);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(95);
				match(LLAVEC);
				_localctx.nodo = Instrucciones.NewFuncion((((FuncionContext)_localctx).ID!=null?((FuncionContext)_localctx).ID.getText():null),instrucciones, parametros, (((FuncionContext)_localctx).RFUNC!=null?((FuncionContext)_localctx).RFUNC.getLine():0), (((FuncionContext)_localctx).RFUNC!=null?((FuncionContext)_localctx).RFUNC.getCharPositionInLine():0))
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(97);
				((FuncionContext)_localctx).RFUNC = match(RFUNC);
				instrucciones := [] Abstract.Instruccion{};
				setState(99);
				((FuncionContext)_localctx).ID = match(ID);
				setState(100);
				match(PARA);
				setState(101);
				((FuncionContext)_localctx).parametros = parametros();
				setState(102);
				match(PARC);
				setState(103);
				match(LLAVEA);
				setState(109);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RPRINT) | (1L << RVAR) | (1L << RFUNC) | (1L << ID))) != 0)) {
					{
					{
					setState(104);
					((FuncionContext)_localctx).instruccion = instruccion();
					instrucciones = append(instrucciones,((FuncionContext)_localctx).instruccion.nodo)
					}
					}
					setState(111);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(112);
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

	public static class ParametrollContext extends ParserRuleContext {
		public Abstract.Instruccion nodo;
		public ExpresionContext expresion;
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ParametrollContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametroll; }
	}

	public final ParametrollContext parametroll() throws RecognitionException {
		ParametrollContext _localctx = new ParametrollContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_parametroll);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(117);
			((ParametrollContext)_localctx).expresion = expresion(0);
			_localctx.nodo = ((ParametrollContext)_localctx).expresion.nodo
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

	public static class ParametrollsContext extends ParserRuleContext {
		public []Abstract.Instruccion lista;
		public ParametrollContext parametroll;
		public List<ParametrollContext> parametroll() {
			return getRuleContexts(ParametrollContext.class);
		}
		public ParametrollContext parametroll(int i) {
			return getRuleContext(ParametrollContext.class,i);
		}
		public List<TerminalNode> COMA() { return getTokens(analizadorParser.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(analizadorParser.COMA, i);
		}
		public ParametrollsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametrolls; }
	}

	public final ParametrollsContext parametrolls() throws RecognitionException {
		ParametrollsContext _localctx = new ParametrollsContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_parametrolls);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(120);
			((ParametrollsContext)_localctx).parametroll = parametroll();
			_localctx.lista = append(_localctx.lista, ((ParametrollsContext)_localctx).parametroll.nodo)
			setState(128);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(122);
				match(COMA);
				setState(123);
				((ParametrollsContext)_localctx).parametroll = parametroll();
				_localctx.lista = append(_localctx.lista, ((ParametrollsContext)_localctx).parametroll.nodo)
				}
				}
				setState(130);
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

	public static class LlamadaContext extends ParserRuleContext {
		public Abstract.Instruccion nodo;
		public Token ID;
		public ParametrollsContext parametrolls;
		public TerminalNode ID() { return getToken(analizadorParser.ID, 0); }
		public TerminalNode PARA() { return getToken(analizadorParser.PARA, 0); }
		public TerminalNode PARC() { return getToken(analizadorParser.PARC, 0); }
		public FininsContext finins() {
			return getRuleContext(FininsContext.class,0);
		}
		public ParametrollsContext parametrolls() {
			return getRuleContext(ParametrollsContext.class,0);
		}
		public LlamadaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_llamada; }
	}

	public final LlamadaContext llamada() throws RecognitionException {
		LlamadaContext _localctx = new LlamadaContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_llamada);
		try {
			setState(145);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(131);
				((LlamadaContext)_localctx).ID = match(ID);
				parametros := []Abstract.Instruccion{}
				setState(133);
				match(PARA);
				setState(134);
				match(PARC);
				setState(135);
				finins();
				_localctx.nodo = Instrucciones.NewLlamada((((LlamadaContext)_localctx).ID!=null?((LlamadaContext)_localctx).ID.getText():null), parametros, (((LlamadaContext)_localctx).ID!=null?((LlamadaContext)_localctx).ID.getLine():0), (((LlamadaContext)_localctx).ID!=null?((LlamadaContext)_localctx).ID.getCharPositionInLine():0))
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(138);
				((LlamadaContext)_localctx).ID = match(ID);
				setState(139);
				match(PARA);
				setState(140);
				((LlamadaContext)_localctx).parametrolls = parametrolls();
				setState(141);
				match(PARC);
				setState(142);
				finins();
				_localctx.nodo = Instrucciones.NewLlamada((((LlamadaContext)_localctx).ID!=null?((LlamadaContext)_localctx).ID.getText():null), ((LlamadaContext)_localctx).parametrolls.lista, (((LlamadaContext)_localctx).ID!=null?((LlamadaContext)_localctx).ID.getLine():0), (((LlamadaContext)_localctx).ID!=null?((LlamadaContext)_localctx).ID.getCharPositionInLine():0))
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
		int _startState = 20;
		enterRecursionRule(_localctx, 20, RULE_expresion, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(170);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENTERO:
				{
				setState(148);
				((ExpresionContext)_localctx).ENTERO = match(ENTERO);
				_localctx.nodo = Expresiones.NewPrimitivo((((ExpresionContext)_localctx).ENTERO!=null?((ExpresionContext)_localctx).ENTERO.getText():null), TS.ENTERO, (((ExpresionContext)_localctx).ENTERO!=null?((ExpresionContext)_localctx).ENTERO.getLine():0), (((ExpresionContext)_localctx).ENTERO!=null?((ExpresionContext)_localctx).ENTERO.getCharPositionInLine():0))
				}
				break;
			case FLOTANTE:
				{
				setState(150);
				((ExpresionContext)_localctx).FLOTANTE = match(FLOTANTE);
				_localctx.nodo = Expresiones.NewPrimitivo((((ExpresionContext)_localctx).FLOTANTE!=null?((ExpresionContext)_localctx).FLOTANTE.getText():null),TS.FLOAT, (((ExpresionContext)_localctx).FLOTANTE!=null?((ExpresionContext)_localctx).FLOTANTE.getLine():0), (((ExpresionContext)_localctx).FLOTANTE!=null?((ExpresionContext)_localctx).FLOTANTE.getCharPositionInLine():0))
				}
				break;
			case RTRUE:
				{
				setState(152);
				((ExpresionContext)_localctx).RTRUE = match(RTRUE);
				_localctx.nodo = Expresiones.NewPrimitivo((((ExpresionContext)_localctx).RTRUE!=null?((ExpresionContext)_localctx).RTRUE.getText():null), TS.BOOLEAN,(((ExpresionContext)_localctx).RTRUE!=null?((ExpresionContext)_localctx).RTRUE.getLine():0), (((ExpresionContext)_localctx).RTRUE!=null?((ExpresionContext)_localctx).RTRUE.getCharPositionInLine():0))
				}
				break;
			case RFALSE:
				{
				setState(154);
				((ExpresionContext)_localctx).RFALSE = match(RFALSE);
				_localctx.nodo = Expresiones.NewPrimitivo((((ExpresionContext)_localctx).RFALSE!=null?((ExpresionContext)_localctx).RFALSE.getText():null), TS.BOOLEAN, (((ExpresionContext)_localctx).RFALSE!=null?((ExpresionContext)_localctx).RFALSE.getLine():0), (((ExpresionContext)_localctx).RFALSE!=null?((ExpresionContext)_localctx).RFALSE.getCharPositionInLine():0))
				}
				break;
			case CHAR:
				{
				setState(156);
				((ExpresionContext)_localctx).CHAR = match(CHAR);
				_localctx.nodo = Expresiones.NewPrimitivo(strings.Replace((((ExpresionContext)_localctx).CHAR!=null?((ExpresionContext)_localctx).CHAR.getText():null),"'","",2), TS.CARACTER, (((ExpresionContext)_localctx).CHAR!=null?((ExpresionContext)_localctx).CHAR.getLine():0), (((ExpresionContext)_localctx).CHAR!=null?((ExpresionContext)_localctx).CHAR.getCharPositionInLine():0))
				}
				break;
			case STRING:
				{
				setState(158);
				((ExpresionContext)_localctx).STRING = match(STRING);
				_localctx.nodo = Expresiones.NewPrimitivo(strings.Replace((((ExpresionContext)_localctx).STRING!=null?((ExpresionContext)_localctx).STRING.getText():null),"\"","",-1),TS.CADENA, (((ExpresionContext)_localctx).STRING!=null?((ExpresionContext)_localctx).STRING.getLine():0), (((ExpresionContext)_localctx).STRING!=null?((ExpresionContext)_localctx).STRING.getCharPositionInLine():0))
				}
				break;
			case MENOS:
				{
				setState(160);
				((ExpresionContext)_localctx).MENOS = match(MENOS);
				setState(161);
				((ExpresionContext)_localctx).expresion = expresion(16);
				_localctx.nodo = Expresiones.NewAritmetica(TS.MENOS, ((ExpresionContext)_localctx).expresion.nodo, nil, (((ExpresionContext)_localctx).MENOS!=null?((ExpresionContext)_localctx).MENOS.getLine():0), (((ExpresionContext)_localctx).MENOS!=null?((ExpresionContext)_localctx).MENOS.getCharPositionInLine():0))
				}
				break;
			case NOT:
				{
				setState(164);
				((ExpresionContext)_localctx).NOT = match(NOT);
				setState(165);
				((ExpresionContext)_localctx).expresion = expresion(4);
				_localctx.nodo = Expresiones.NewLogica(TS.NOT, ((ExpresionContext)_localctx).expresion.nodo, ((ExpresionContext)_localctx).expresion.nodo, (((ExpresionContext)_localctx).NOT!=null?((ExpresionContext)_localctx).NOT.getLine():0), (((ExpresionContext)_localctx).NOT!=null?((ExpresionContext)_localctx).NOT.getCharPositionInLine():0))
				}
				break;
			case ID:
				{
				setState(168);
				((ExpresionContext)_localctx).ID = match(ID);
				_localctx.nodo = Expresiones.NewIdentificador((((ExpresionContext)_localctx).ID!=null?((ExpresionContext)_localctx).ID.getText():null), (((ExpresionContext)_localctx).ID!=null?((ExpresionContext)_localctx).ID.getLine():0), (((ExpresionContext)_localctx).ID!=null?((ExpresionContext)_localctx).ID.getCharPositionInLine():0))
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(239);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(237);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(172);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(173);
						((ExpresionContext)_localctx).MUL = match(MUL);
						setState(174);
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
						setState(177);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(178);
						((ExpresionContext)_localctx).DIV = match(DIV);
						setState(179);
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
						setState(182);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(183);
						((ExpresionContext)_localctx).MOD = match(MOD);
						setState(184);
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
						setState(187);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(188);
						((ExpresionContext)_localctx).MAS = match(MAS);
						setState(189);
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
						setState(192);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(193);
						((ExpresionContext)_localctx).MENOS = match(MENOS);
						setState(194);
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
						setState(197);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(198);
						((ExpresionContext)_localctx).DISTINTO = match(DISTINTO);
						setState(199);
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
						setState(202);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(203);
						((ExpresionContext)_localctx).MAYOR = match(MAYOR);
						setState(204);
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
						setState(207);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(208);
						((ExpresionContext)_localctx).IGUALIGUAL = match(IGUALIGUAL);
						setState(209);
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
						setState(212);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(213);
						((ExpresionContext)_localctx).MENOR = match(MENOR);
						setState(214);
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
						setState(217);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(218);
						((ExpresionContext)_localctx).MAYORIGUAL = match(MAYORIGUAL);
						setState(219);
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
						setState(222);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(223);
						((ExpresionContext)_localctx).MENORIGUAL = match(MENORIGUAL);
						setState(224);
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
						setState(227);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(228);
						((ExpresionContext)_localctx).OR = match(OR);
						setState(229);
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
						setState(232);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(233);
						((ExpresionContext)_localctx).AND = match(AND);
						setState(234);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(3);
						_localctx.nodo = Expresiones.NewLogica(TS.AND, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).AND!=null?((ExpresionContext)_localctx).AND.getLine():0), (((ExpresionContext)_localctx).AND!=null?((ExpresionContext)_localctx).AND.getCharPositionInLine():0))
						}
						break;
					}
					} 
				}
				setState(241);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
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
		enterRule(_localctx, 22, RULE_imprimir);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(242);
			((ImprimirContext)_localctx).RPRINT = match(RPRINT);
			setState(243);
			match(PARA);
			setState(244);
			((ImprimirContext)_localctx).expresion = expresion(0);
			setState(245);
			match(PARC);
			setState(246);
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
		enterRule(_localctx, 24, RULE_finins);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(249);
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
		case 10:
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3#\u00fe\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\3\2\3\2\3\2\3\2\7\2!\n\2\f\2\16\2$\13\2"+
		"\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\5"+
		"\3\67\n\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6"+
		"\3\6\3\7\3\7\3\7\3\7\3\7\3\7\7\7O\n\7\f\7\16\7R\13\7\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\7\b]\n\b\f\b\16\b`\13\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\7\bn\n\b\f\b\16\bq\13\b\3\b\3\b\3\b\5\bv\n\b\3\t"+
		"\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\7\n\u0081\n\n\f\n\16\n\u0084\13\n\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13"+
		"\u0094\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f\u00ad\n\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\7\f\u00f0\n\f\f\f\16\f\u00f3\13\f\3\r\3\r"+
		"\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\2\3\26\17\2\4\6\b\n\f\16\20\22\24"+
		"\26\30\32\2\2\2\u0110\2\34\3\2\2\2\4\66\3\2\2\2\68\3\2\2\2\b?\3\2\2\2"+
		"\nE\3\2\2\2\fH\3\2\2\2\16u\3\2\2\2\20w\3\2\2\2\22z\3\2\2\2\24\u0093\3"+
		"\2\2\2\26\u00ac\3\2\2\2\30\u00f4\3\2\2\2\32\u00fb\3\2\2\2\34\"\b\2\1\2"+
		"\35\36\5\4\3\2\36\37\b\2\1\2\37!\3\2\2\2 \35\3\2\2\2!$\3\2\2\2\" \3\2"+
		"\2\2\"#\3\2\2\2#%\3\2\2\2$\"\3\2\2\2%&\b\2\1\2&\3\3\2\2\2\'(\5\30\r\2"+
		"()\b\3\1\2)\67\3\2\2\2*+\5\6\4\2+,\b\3\1\2,\67\3\2\2\2-.\5\b\5\2./\b\3"+
		"\1\2/\67\3\2\2\2\60\61\5\16\b\2\61\62\b\3\1\2\62\67\3\2\2\2\63\64\5\24"+
		"\13\2\64\65\b\3\1\2\65\67\3\2\2\2\66\'\3\2\2\2\66*\3\2\2\2\66-\3\2\2\2"+
		"\66\60\3\2\2\2\66\63\3\2\2\2\67\5\3\2\2\289\7\4\2\29:\7\35\2\2:;\7\36"+
		"\2\2;<\5\26\f\2<=\5\32\16\2=>\b\4\1\2>\7\3\2\2\2?@\7\35\2\2@A\7\36\2\2"+
		"AB\5\26\f\2BC\5\32\16\2CD\b\5\1\2D\t\3\2\2\2EF\7\35\2\2FG\b\6\1\2G\13"+
		"\3\2\2\2HI\5\n\6\2IP\b\7\1\2JK\7!\2\2KL\5\n\6\2LM\b\7\1\2MO\3\2\2\2NJ"+
		"\3\2\2\2OR\3\2\2\2PN\3\2\2\2PQ\3\2\2\2Q\r\3\2\2\2RP\3\2\2\2ST\7\34\2\2"+
		"TU\b\b\1\2UV\7\35\2\2VW\7\6\2\2WX\7\7\2\2X^\7\37\2\2YZ\5\4\3\2Z[\b\b\1"+
		"\2[]\3\2\2\2\\Y\3\2\2\2]`\3\2\2\2^\\\3\2\2\2^_\3\2\2\2_a\3\2\2\2`^\3\2"+
		"\2\2ab\7 \2\2bv\b\b\1\2cd\7\34\2\2de\b\b\1\2ef\7\35\2\2fg\7\6\2\2gh\5"+
		"\f\7\2hi\7\7\2\2io\7\37\2\2jk\5\4\3\2kl\b\b\1\2ln\3\2\2\2mj\3\2\2\2nq"+
		"\3\2\2\2om\3\2\2\2op\3\2\2\2pr\3\2\2\2qo\3\2\2\2rs\7 \2\2st\b\b\1\2tv"+
		"\3\2\2\2uS\3\2\2\2uc\3\2\2\2v\17\3\2\2\2wx\5\26\f\2xy\b\t\1\2y\21\3\2"+
		"\2\2z{\5\20\t\2{\u0082\b\n\1\2|}\7!\2\2}~\5\20\t\2~\177\b\n\1\2\177\u0081"+
		"\3\2\2\2\u0080|\3\2\2\2\u0081\u0084\3\2\2\2\u0082\u0080\3\2\2\2\u0082"+
		"\u0083\3\2\2\2\u0083\23\3\2\2\2\u0084\u0082\3\2\2\2\u0085\u0086\7\35\2"+
		"\2\u0086\u0087\b\13\1\2\u0087\u0088\7\6\2\2\u0088\u0089\7\7\2\2\u0089"+
		"\u008a\5\32\16\2\u008a\u008b\b\13\1\2\u008b\u0094\3\2\2\2\u008c\u008d"+
		"\7\35\2\2\u008d\u008e\7\6\2\2\u008e\u008f\5\22\n\2\u008f\u0090\7\7\2\2"+
		"\u0090\u0091\5\32\16\2\u0091\u0092\b\13\1\2\u0092\u0094\3\2\2\2\u0093"+
		"\u0085\3\2\2\2\u0093\u008c\3\2\2\2\u0094\25\3\2\2\2\u0095\u0096\b\f\1"+
		"\2\u0096\u0097\7\b\2\2\u0097\u00ad\b\f\1\2\u0098\u0099\7\t\2\2\u0099\u00ad"+
		"\b\f\1\2\u009a\u009b\7\n\2\2\u009b\u00ad\b\f\1\2\u009c\u009d\7\13\2\2"+
		"\u009d\u00ad\b\f\1\2\u009e\u009f\7\f\2\2\u009f\u00ad\b\f\1\2\u00a0\u00a1"+
		"\7\r\2\2\u00a1\u00ad\b\f\1\2\u00a2\u00a3\7\17\2\2\u00a3\u00a4\5\26\f\22"+
		"\u00a4\u00a5\b\f\1\2\u00a5\u00ad\3\2\2\2\u00a6\u00a7\7\33\2\2\u00a7\u00a8"+
		"\5\26\f\6\u00a8\u00a9\b\f\1\2\u00a9\u00ad\3\2\2\2\u00aa\u00ab\7\35\2\2"+
		"\u00ab\u00ad\b\f\1\2\u00ac\u0095\3\2\2\2\u00ac\u0098\3\2\2\2\u00ac\u009a"+
		"\3\2\2\2\u00ac\u009c\3\2\2\2\u00ac\u009e\3\2\2\2\u00ac\u00a0\3\2\2\2\u00ac"+
		"\u00a2\3\2\2\2\u00ac\u00a6\3\2\2\2\u00ac\u00aa\3\2\2\2\u00ad\u00f1\3\2"+
		"\2\2\u00ae\u00af\f\21\2\2\u00af\u00b0\7\20\2\2\u00b0\u00b1\5\26\f\22\u00b1"+
		"\u00b2\b\f\1\2\u00b2\u00f0\3\2\2\2\u00b3\u00b4\f\20\2\2\u00b4\u00b5\7"+
		"\21\2\2\u00b5\u00b6\5\26\f\21\u00b6\u00b7\b\f\1\2\u00b7\u00f0\3\2\2\2"+
		"\u00b8\u00b9\f\17\2\2\u00b9\u00ba\7\22\2\2\u00ba\u00bb\5\26\f\20\u00bb"+
		"\u00bc\b\f\1\2\u00bc\u00f0\3\2\2\2\u00bd\u00be\f\16\2\2\u00be\u00bf\7"+
		"\16\2\2\u00bf\u00c0\5\26\f\17\u00c0\u00c1\b\f\1\2\u00c1\u00f0\3\2\2\2"+
		"\u00c2\u00c3\f\r\2\2\u00c3\u00c4\7\17\2\2\u00c4\u00c5\5\26\f\16\u00c5"+
		"\u00c6\b\f\1\2\u00c6\u00f0\3\2\2\2\u00c7\u00c8\f\f\2\2\u00c8\u00c9\7\30"+
		"\2\2\u00c9\u00ca\5\26\f\r\u00ca\u00cb\b\f\1\2\u00cb\u00f0\3\2\2\2\u00cc"+
		"\u00cd\f\13\2\2\u00cd\u00ce\7\23\2\2\u00ce\u00cf\5\26\f\f\u00cf\u00d0"+
		"\b\f\1\2\u00d0\u00f0\3\2\2\2\u00d1\u00d2\f\n\2\2\u00d2\u00d3\7\27\2\2"+
		"\u00d3\u00d4\5\26\f\13\u00d4\u00d5\b\f\1\2\u00d5\u00f0\3\2\2\2\u00d6\u00d7"+
		"\f\t\2\2\u00d7\u00d8\7\24\2\2\u00d8\u00d9\5\26\f\n\u00d9\u00da\b\f\1\2"+
		"\u00da\u00f0\3\2\2\2\u00db\u00dc\f\b\2\2\u00dc\u00dd\7\25\2\2\u00dd\u00de"+
		"\5\26\f\t\u00de\u00df\b\f\1\2\u00df\u00f0\3\2\2\2\u00e0\u00e1\f\7\2\2"+
		"\u00e1\u00e2\7\26\2\2\u00e2\u00e3\5\26\f\b\u00e3\u00e4\b\f\1\2\u00e4\u00f0"+
		"\3\2\2\2\u00e5\u00e6\f\5\2\2\u00e6\u00e7\7\31\2\2\u00e7\u00e8\5\26\f\6"+
		"\u00e8\u00e9\b\f\1\2\u00e9\u00f0\3\2\2\2\u00ea\u00eb\f\4\2\2\u00eb\u00ec"+
		"\7\32\2\2\u00ec\u00ed\5\26\f\5\u00ed\u00ee\b\f\1\2\u00ee\u00f0\3\2\2\2"+
		"\u00ef\u00ae\3\2\2\2\u00ef\u00b3\3\2\2\2\u00ef\u00b8\3\2\2\2\u00ef\u00bd"+
		"\3\2\2\2\u00ef\u00c2\3\2\2\2\u00ef\u00c7\3\2\2\2\u00ef\u00cc\3\2\2\2\u00ef"+
		"\u00d1\3\2\2\2\u00ef\u00d6\3\2\2\2\u00ef\u00db\3\2\2\2\u00ef\u00e0\3\2"+
		"\2\2\u00ef\u00e5\3\2\2\2\u00ef\u00ea\3\2\2\2\u00f0\u00f3\3\2\2\2\u00f1"+
		"\u00ef\3\2\2\2\u00f1\u00f2\3\2\2\2\u00f2\27\3\2\2\2\u00f3\u00f1\3\2\2"+
		"\2\u00f4\u00f5\7\3\2\2\u00f5\u00f6\7\6\2\2\u00f6\u00f7\5\26\f\2\u00f7"+
		"\u00f8\7\7\2\2\u00f8\u00f9\5\32\16\2\u00f9\u00fa\b\r\1\2\u00fa\31\3\2"+
		"\2\2\u00fb\u00fc\7\5\2\2\u00fc\33\3\2\2\2\r\"\66P^ou\u0082\u0093\u00ac"+
		"\u00ef\u00f1";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}