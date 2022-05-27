// Generated from /home/emivnajera/go/src/PersonalProjects/EjemploANTLR/analizador.g4 by ANTLR 4.8

import "github.com/emivnajera/Abstract"
import "github.com/emivnajera/Expresiones"
import "github.com/emivnajera/Instrucciones"
import "github.com/emivnajera/TS"
import "strings"

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
		RPRINT=1, PUNTOCOMA=2, PARA=3, PARC=4, ENTERO=5, FLOTANTE=6, RTRUE=7, 
		RFALSE=8, CHAR=9, STRING=10, MAS=11, MENOS=12, MUL=13, DIV=14, MOD=15, 
		MAYOR=16, MENOR=17, MAYORIGUAL=18, MENORIGUAL=19, IGUALIGUAL=20, DISTINTO=21, 
		OR=22, AND=23, NOT=24, COMMENT=25, WS=26;
	public static final int
		RULE_start = 0, RULE_instruccion = 1, RULE_expresion = 2, RULE_imprimir = 3, 
		RULE_finins = 4;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "instruccion", "expresion", "imprimir", "finins"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'print'", "';'", "'('", "')'", null, null, "'true'", "'false'", 
			null, null, "'+'", "'-'", "'*'", "'/'", "'%'", "'>'", "'<'", "'>='", 
			"'<='", "'=='", "'!='", "'||'", "'&&'", "'!'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "RPRINT", "PUNTOCOMA", "PARA", "PARC", "ENTERO", "FLOTANTE", "RTRUE", 
			"RFALSE", "CHAR", "STRING", "MAS", "MENOS", "MUL", "DIV", "MOD", "MAYOR", 
			"MENOR", "MAYORIGUAL", "MENORIGUAL", "IGUALIGUAL", "DISTINTO", "OR", 
			"AND", "NOT", "COMMENT", "WS"
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
			instrucciones := [] Abstract.Instruccion{};TSGlobal:=TS.TablaSimbolos{}
			setState(16);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==RPRINT) {
				{
				{
				setState(11);
				((StartContext)_localctx).instruccion = instruccion();
				instrucciones = append(instrucciones,((StartContext)_localctx).instruccion.nodo)
				}
				}
				setState(18);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}

			for _, n := range instrucciones {
			        n.Interpretar(&TSGlobal)
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
		public ImprimirContext imprimir() {
			return getRuleContext(ImprimirContext.class,0);
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
			enterOuterAlt(_localctx, 1);
			{
			setState(21);
			((InstruccionContext)_localctx).imprimir = imprimir();
			_localctx.nodo = ((InstruccionContext)_localctx).imprimir.nodo
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
		int _startState = 4;
		enterRecursionRule(_localctx, 4, RULE_expresion, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(45);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENTERO:
				{
				setState(25);
				((ExpresionContext)_localctx).ENTERO = match(ENTERO);
				_localctx.nodo = Expresiones.NewPrimitivo((((ExpresionContext)_localctx).ENTERO!=null?((ExpresionContext)_localctx).ENTERO.getText():null), TS.ENTERO, (((ExpresionContext)_localctx).ENTERO!=null?((ExpresionContext)_localctx).ENTERO.getLine():0), (((ExpresionContext)_localctx).ENTERO!=null?((ExpresionContext)_localctx).ENTERO.getCharPositionInLine():0))
				}
				break;
			case FLOTANTE:
				{
				setState(27);
				((ExpresionContext)_localctx).FLOTANTE = match(FLOTANTE);
				_localctx.nodo = Expresiones.NewPrimitivo((((ExpresionContext)_localctx).FLOTANTE!=null?((ExpresionContext)_localctx).FLOTANTE.getText():null),TS.FLOAT, (((ExpresionContext)_localctx).FLOTANTE!=null?((ExpresionContext)_localctx).FLOTANTE.getLine():0), (((ExpresionContext)_localctx).FLOTANTE!=null?((ExpresionContext)_localctx).FLOTANTE.getCharPositionInLine():0))
				}
				break;
			case RTRUE:
				{
				setState(29);
				((ExpresionContext)_localctx).RTRUE = match(RTRUE);
				_localctx.nodo = Expresiones.NewPrimitivo((((ExpresionContext)_localctx).RTRUE!=null?((ExpresionContext)_localctx).RTRUE.getText():null), TS.BOOLEAN,(((ExpresionContext)_localctx).RTRUE!=null?((ExpresionContext)_localctx).RTRUE.getLine():0), (((ExpresionContext)_localctx).RTRUE!=null?((ExpresionContext)_localctx).RTRUE.getCharPositionInLine():0))
				}
				break;
			case RFALSE:
				{
				setState(31);
				((ExpresionContext)_localctx).RFALSE = match(RFALSE);
				_localctx.nodo = Expresiones.NewPrimitivo((((ExpresionContext)_localctx).RFALSE!=null?((ExpresionContext)_localctx).RFALSE.getText():null), TS.BOOLEAN, (((ExpresionContext)_localctx).RFALSE!=null?((ExpresionContext)_localctx).RFALSE.getLine():0), (((ExpresionContext)_localctx).RFALSE!=null?((ExpresionContext)_localctx).RFALSE.getCharPositionInLine():0))
				}
				break;
			case CHAR:
				{
				setState(33);
				((ExpresionContext)_localctx).CHAR = match(CHAR);
				_localctx.nodo = Expresiones.NewPrimitivo(strings.Replace((((ExpresionContext)_localctx).CHAR!=null?((ExpresionContext)_localctx).CHAR.getText():null),"'","",2), TS.CARACTER, (((ExpresionContext)_localctx).CHAR!=null?((ExpresionContext)_localctx).CHAR.getLine():0), (((ExpresionContext)_localctx).CHAR!=null?((ExpresionContext)_localctx).CHAR.getCharPositionInLine():0))
				}
				break;
			case STRING:
				{
				setState(35);
				((ExpresionContext)_localctx).STRING = match(STRING);
				_localctx.nodo = Expresiones.NewPrimitivo(strings.Replace((((ExpresionContext)_localctx).STRING!=null?((ExpresionContext)_localctx).STRING.getText():null),"\"","",-1),TS.CADENA, (((ExpresionContext)_localctx).STRING!=null?((ExpresionContext)_localctx).STRING.getLine():0), (((ExpresionContext)_localctx).STRING!=null?((ExpresionContext)_localctx).STRING.getCharPositionInLine():0))
				}
				break;
			case MENOS:
				{
				setState(37);
				((ExpresionContext)_localctx).MENOS = match(MENOS);
				setState(38);
				((ExpresionContext)_localctx).expresion = expresion(15);
				_localctx.nodo = Expresiones.NewAritmetica(TS.MENOS, ((ExpresionContext)_localctx).expresion.nodo, nil, (((ExpresionContext)_localctx).MENOS!=null?((ExpresionContext)_localctx).MENOS.getLine():0), (((ExpresionContext)_localctx).MENOS!=null?((ExpresionContext)_localctx).MENOS.getCharPositionInLine():0))
				}
				break;
			case NOT:
				{
				setState(41);
				((ExpresionContext)_localctx).NOT = match(NOT);
				setState(42);
				((ExpresionContext)_localctx).expresion = expresion(3);
				_localctx.nodo = Expresiones.NewLogica(TS.NOT, ((ExpresionContext)_localctx).expresion.nodo, ((ExpresionContext)_localctx).expresion.nodo, (((ExpresionContext)_localctx).NOT!=null?((ExpresionContext)_localctx).NOT.getLine():0), (((ExpresionContext)_localctx).NOT!=null?((ExpresionContext)_localctx).NOT.getCharPositionInLine():0))
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(114);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(112);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(47);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(48);
						((ExpresionContext)_localctx).MUL = match(MUL);
						setState(49);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(15);
						_localctx.nodo = Expresiones.NewAritmetica(TS.POR, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MUL!=null?((ExpresionContext)_localctx).MUL.getLine():0), (((ExpresionContext)_localctx).MUL!=null?((ExpresionContext)_localctx).MUL.getCharPositionInLine():0))
						}
						break;
					case 2:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(52);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(53);
						((ExpresionContext)_localctx).DIV = match(DIV);
						setState(54);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(14);
						_localctx.nodo = Expresiones.NewAritmetica(TS.DIV, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).DIV!=null?((ExpresionContext)_localctx).DIV.getLine():0), (((ExpresionContext)_localctx).DIV!=null?((ExpresionContext)_localctx).DIV.getCharPositionInLine():0))
						}
						break;
					case 3:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(57);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(58);
						((ExpresionContext)_localctx).MOD = match(MOD);
						setState(59);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(13);
						_localctx.nodo = Expresiones.NewAritmetica(TS.MOD, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MOD!=null?((ExpresionContext)_localctx).MOD.getLine():0), (((ExpresionContext)_localctx).MOD!=null?((ExpresionContext)_localctx).MOD.getCharPositionInLine():0))
						}
						break;
					case 4:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(62);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(63);
						((ExpresionContext)_localctx).MAS = match(MAS);
						setState(64);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(12);
						_localctx.nodo = Expresiones.NewAritmetica(TS.MAS, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MAS!=null?((ExpresionContext)_localctx).MAS.getLine():0), (((ExpresionContext)_localctx).MAS!=null?((ExpresionContext)_localctx).MAS.getCharPositionInLine():0))
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(67);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(68);
						((ExpresionContext)_localctx).MENOS = match(MENOS);
						setState(69);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(11);
						_localctx.nodo = Expresiones.NewAritmetica(TS.MENOS, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MENOS!=null?((ExpresionContext)_localctx).MENOS.getLine():0), (((ExpresionContext)_localctx).MENOS!=null?((ExpresionContext)_localctx).MENOS.getCharPositionInLine():0))
						}
						break;
					case 6:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(72);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(73);
						((ExpresionContext)_localctx).DISTINTO = match(DISTINTO);
						setState(74);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(10);
						_localctx.nodo = Expresiones.NewRelacional(TS.DIFERENTE, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).DISTINTO!=null?((ExpresionContext)_localctx).DISTINTO.getLine():0), (((ExpresionContext)_localctx).DISTINTO!=null?((ExpresionContext)_localctx).DISTINTO.getCharPositionInLine():0))
						}
						break;
					case 7:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(77);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(78);
						((ExpresionContext)_localctx).MAYOR = match(MAYOR);
						setState(79);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(9);
						_localctx.nodo = Expresiones.NewRelacional(TS.MAYORQUE, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MAYOR!=null?((ExpresionContext)_localctx).MAYOR.getLine():0), (((ExpresionContext)_localctx).MAYOR!=null?((ExpresionContext)_localctx).MAYOR.getCharPositionInLine():0))
						}
						break;
					case 8:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(82);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(83);
						((ExpresionContext)_localctx).IGUALIGUAL = match(IGUALIGUAL);
						setState(84);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(8);
						_localctx.nodo = Expresiones.NewRelacional(TS.IGUALIGUAL, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).IGUALIGUAL!=null?((ExpresionContext)_localctx).IGUALIGUAL.getLine():0), (((ExpresionContext)_localctx).IGUALIGUAL!=null?((ExpresionContext)_localctx).IGUALIGUAL.getCharPositionInLine():0))
						}
						break;
					case 9:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(87);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(88);
						((ExpresionContext)_localctx).MENOR = match(MENOR);
						setState(89);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(7);
						_localctx.nodo = Expresiones.NewRelacional(TS.MENORQUE, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MENOR!=null?((ExpresionContext)_localctx).MENOR.getLine():0), (((ExpresionContext)_localctx).MENOR!=null?((ExpresionContext)_localctx).MENOR.getCharPositionInLine():0))
						}
						break;
					case 10:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(92);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(93);
						((ExpresionContext)_localctx).MAYORIGUAL = match(MAYORIGUAL);
						setState(94);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(6);
						_localctx.nodo = Expresiones.NewRelacional(TS.MAYORIGUAL, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MAYORIGUAL!=null?((ExpresionContext)_localctx).MAYORIGUAL.getLine():0), (((ExpresionContext)_localctx).MAYORIGUAL!=null?((ExpresionContext)_localctx).MAYORIGUAL.getCharPositionInLine():0))
						}
						break;
					case 11:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(97);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(98);
						((ExpresionContext)_localctx).MENORIGUAL = match(MENORIGUAL);
						setState(99);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(5);
						_localctx.nodo = Expresiones.NewRelacional(TS.MENORIGUAL, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MENORIGUAL!=null?((ExpresionContext)_localctx).MENORIGUAL.getLine():0), (((ExpresionContext)_localctx).MENORIGUAL!=null?((ExpresionContext)_localctx).MENORIGUAL.getCharPositionInLine():0))
						}
						break;
					case 12:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(102);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(103);
						((ExpresionContext)_localctx).OR = match(OR);
						setState(104);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(3);
						_localctx.nodo = Expresiones.NewLogica(TS.OR, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).OR!=null?((ExpresionContext)_localctx).OR.getLine():0), (((ExpresionContext)_localctx).OR!=null?((ExpresionContext)_localctx).OR.getCharPositionInLine():0))
						}
						break;
					case 13:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(107);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(108);
						((ExpresionContext)_localctx).AND = match(AND);
						setState(109);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(2);
						_localctx.nodo = Expresiones.NewLogica(TS.AND, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).AND!=null?((ExpresionContext)_localctx).AND.getLine():0), (((ExpresionContext)_localctx).AND!=null?((ExpresionContext)_localctx).AND.getCharPositionInLine():0))
						}
						break;
					}
					} 
				}
				setState(116);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
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
		enterRule(_localctx, 6, RULE_imprimir);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(117);
			((ImprimirContext)_localctx).RPRINT = match(RPRINT);
			setState(118);
			match(PARA);
			setState(119);
			((ImprimirContext)_localctx).expresion = expresion(0);
			setState(120);
			match(PARC);
			setState(121);
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
		enterRule(_localctx, 8, RULE_finins);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(124);
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
		case 2:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 14);
		case 1:
			return precpred(_ctx, 13);
		case 2:
			return precpred(_ctx, 12);
		case 3:
			return precpred(_ctx, 11);
		case 4:
			return precpred(_ctx, 10);
		case 5:
			return precpred(_ctx, 9);
		case 6:
			return precpred(_ctx, 8);
		case 7:
			return precpred(_ctx, 7);
		case 8:
			return precpred(_ctx, 6);
		case 9:
			return precpred(_ctx, 5);
		case 10:
			return precpred(_ctx, 4);
		case 11:
			return precpred(_ctx, 2);
		case 12:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\34\u0081\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\3\2\3\2\3\2\3\2\7\2\21\n\2\f\2\16\2\24"+
		"\13\2\3\2\3\2\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4\60\n\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\7\4s\n\4\f\4\16\4v\13\4\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\6\3\6\3\6\2\3\6\7\2\4\6\b\n\2\2\2\u0090\2\f\3\2\2\2\4"+
		"\27\3\2\2\2\6/\3\2\2\2\bw\3\2\2\2\n~\3\2\2\2\f\22\b\2\1\2\r\16\5\4\3\2"+
		"\16\17\b\2\1\2\17\21\3\2\2\2\20\r\3\2\2\2\21\24\3\2\2\2\22\20\3\2\2\2"+
		"\22\23\3\2\2\2\23\25\3\2\2\2\24\22\3\2\2\2\25\26\b\2\1\2\26\3\3\2\2\2"+
		"\27\30\5\b\5\2\30\31\b\3\1\2\31\5\3\2\2\2\32\33\b\4\1\2\33\34\7\7\2\2"+
		"\34\60\b\4\1\2\35\36\7\b\2\2\36\60\b\4\1\2\37 \7\t\2\2 \60\b\4\1\2!\""+
		"\7\n\2\2\"\60\b\4\1\2#$\7\13\2\2$\60\b\4\1\2%&\7\f\2\2&\60\b\4\1\2\'("+
		"\7\16\2\2()\5\6\4\21)*\b\4\1\2*\60\3\2\2\2+,\7\32\2\2,-\5\6\4\5-.\b\4"+
		"\1\2.\60\3\2\2\2/\32\3\2\2\2/\35\3\2\2\2/\37\3\2\2\2/!\3\2\2\2/#\3\2\2"+
		"\2/%\3\2\2\2/\'\3\2\2\2/+\3\2\2\2\60t\3\2\2\2\61\62\f\20\2\2\62\63\7\17"+
		"\2\2\63\64\5\6\4\21\64\65\b\4\1\2\65s\3\2\2\2\66\67\f\17\2\2\678\7\20"+
		"\2\289\5\6\4\209:\b\4\1\2:s\3\2\2\2;<\f\16\2\2<=\7\21\2\2=>\5\6\4\17>"+
		"?\b\4\1\2?s\3\2\2\2@A\f\r\2\2AB\7\r\2\2BC\5\6\4\16CD\b\4\1\2Ds\3\2\2\2"+
		"EF\f\f\2\2FG\7\16\2\2GH\5\6\4\rHI\b\4\1\2Is\3\2\2\2JK\f\13\2\2KL\7\27"+
		"\2\2LM\5\6\4\fMN\b\4\1\2Ns\3\2\2\2OP\f\n\2\2PQ\7\22\2\2QR\5\6\4\13RS\b"+
		"\4\1\2Ss\3\2\2\2TU\f\t\2\2UV\7\26\2\2VW\5\6\4\nWX\b\4\1\2Xs\3\2\2\2YZ"+
		"\f\b\2\2Z[\7\23\2\2[\\\5\6\4\t\\]\b\4\1\2]s\3\2\2\2^_\f\7\2\2_`\7\24\2"+
		"\2`a\5\6\4\bab\b\4\1\2bs\3\2\2\2cd\f\6\2\2de\7\25\2\2ef\5\6\4\7fg\b\4"+
		"\1\2gs\3\2\2\2hi\f\4\2\2ij\7\30\2\2jk\5\6\4\5kl\b\4\1\2ls\3\2\2\2mn\f"+
		"\3\2\2no\7\31\2\2op\5\6\4\4pq\b\4\1\2qs\3\2\2\2r\61\3\2\2\2r\66\3\2\2"+
		"\2r;\3\2\2\2r@\3\2\2\2rE\3\2\2\2rJ\3\2\2\2rO\3\2\2\2rT\3\2\2\2rY\3\2\2"+
		"\2r^\3\2\2\2rc\3\2\2\2rh\3\2\2\2rm\3\2\2\2sv\3\2\2\2tr\3\2\2\2tu\3\2\2"+
		"\2u\7\3\2\2\2vt\3\2\2\2wx\7\3\2\2xy\7\5\2\2yz\5\6\4\2z{\7\6\2\2{|\5\n"+
		"\6\2|}\b\5\1\2}\t\3\2\2\2~\177\7\4\2\2\177\13\3\2\2\2\6\22/rt";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}