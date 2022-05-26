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
		COMMENT=16, WS=17;
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
			null, null, "'+'", "'-'", "'*'", "'/'", "'%'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "RPRINT", "PUNTOCOMA", "PARA", "PARC", "ENTERO", "FLOTANTE", "RTRUE", 
			"RFALSE", "CHAR", "STRING", "MAS", "MENOS", "MUL", "DIV", "MOD", "COMMENT", 
			"WS"
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
		public Token MUL;
		public ExpresionContext opd;
		public Token DIV;
		public Token MOD;
		public Token MAS;
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
		public TerminalNode MUL() { return getToken(analizadorParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(analizadorParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(analizadorParser.MOD, 0); }
		public TerminalNode MAS() { return getToken(analizadorParser.MAS, 0); }
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
			setState(41);
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
				((ExpresionContext)_localctx).expresion = expresion(6);
				_localctx.nodo = Expresiones.NewAritmetica(TS.MENOS, ((ExpresionContext)_localctx).expresion.nodo, nil, (((ExpresionContext)_localctx).MENOS!=null?((ExpresionContext)_localctx).MENOS.getLine():0), (((ExpresionContext)_localctx).MENOS!=null?((ExpresionContext)_localctx).MENOS.getCharPositionInLine():0))
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(70);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(68);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(43);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(44);
						((ExpresionContext)_localctx).MUL = match(MUL);
						setState(45);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(6);
						_localctx.nodo = Expresiones.NewAritmetica(TS.POR, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MUL!=null?((ExpresionContext)_localctx).MUL.getLine():0), (((ExpresionContext)_localctx).MUL!=null?((ExpresionContext)_localctx).MUL.getCharPositionInLine():0))
						}
						break;
					case 2:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(48);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(49);
						((ExpresionContext)_localctx).DIV = match(DIV);
						setState(50);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(5);
						_localctx.nodo = Expresiones.NewAritmetica(TS.DIV, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).DIV!=null?((ExpresionContext)_localctx).DIV.getLine():0), (((ExpresionContext)_localctx).DIV!=null?((ExpresionContext)_localctx).DIV.getCharPositionInLine():0))
						}
						break;
					case 3:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(53);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(54);
						((ExpresionContext)_localctx).MOD = match(MOD);
						setState(55);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(4);
						_localctx.nodo = Expresiones.NewAritmetica(TS.MOD, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MOD!=null?((ExpresionContext)_localctx).MOD.getLine():0), (((ExpresionContext)_localctx).MOD!=null?((ExpresionContext)_localctx).MOD.getCharPositionInLine():0))
						}
						break;
					case 4:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(58);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(59);
						((ExpresionContext)_localctx).MAS = match(MAS);
						setState(60);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(3);
						_localctx.nodo = Expresiones.NewAritmetica(TS.MAS, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MAS!=null?((ExpresionContext)_localctx).MAS.getLine():0), (((ExpresionContext)_localctx).MAS!=null?((ExpresionContext)_localctx).MAS.getCharPositionInLine():0))
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.opi = _prevctx;
						_localctx.opi = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(63);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(64);
						((ExpresionContext)_localctx).MENOS = match(MENOS);
						setState(65);
						((ExpresionContext)_localctx).opd = ((ExpresionContext)_localctx).expresion = expresion(2);
						_localctx.nodo = Expresiones.NewAritmetica(TS.MENOS, ((ExpresionContext)_localctx).opi.nodo, ((ExpresionContext)_localctx).opd.nodo, (((ExpresionContext)_localctx).MENOS!=null?((ExpresionContext)_localctx).MENOS.getLine():0), (((ExpresionContext)_localctx).MENOS!=null?((ExpresionContext)_localctx).MENOS.getCharPositionInLine():0))
						}
						break;
					}
					} 
				}
				setState(72);
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
			setState(73);
			((ImprimirContext)_localctx).RPRINT = match(RPRINT);
			setState(74);
			match(PARA);
			setState(75);
			((ImprimirContext)_localctx).expresion = expresion(0);
			setState(76);
			match(PARC);
			setState(77);
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
			setState(80);
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
			return precpred(_ctx, 5);
		case 1:
			return precpred(_ctx, 4);
		case 2:
			return precpred(_ctx, 3);
		case 3:
			return precpred(_ctx, 2);
		case 4:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\23U\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\3\2\3\2\3\2\3\2\7\2\21\n\2\f\2\16\2\24\13\2"+
		"\3\2\3\2\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\5\4,\n\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\7\4G\n\4\f"+
		"\4\16\4J\13\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\2\3\6\7\2\4\6\b"+
		"\n\2\2\2[\2\f\3\2\2\2\4\27\3\2\2\2\6+\3\2\2\2\bK\3\2\2\2\nR\3\2\2\2\f"+
		"\22\b\2\1\2\r\16\5\4\3\2\16\17\b\2\1\2\17\21\3\2\2\2\20\r\3\2\2\2\21\24"+
		"\3\2\2\2\22\20\3\2\2\2\22\23\3\2\2\2\23\25\3\2\2\2\24\22\3\2\2\2\25\26"+
		"\b\2\1\2\26\3\3\2\2\2\27\30\5\b\5\2\30\31\b\3\1\2\31\5\3\2\2\2\32\33\b"+
		"\4\1\2\33\34\7\7\2\2\34,\b\4\1\2\35\36\7\b\2\2\36,\b\4\1\2\37 \7\t\2\2"+
		" ,\b\4\1\2!\"\7\n\2\2\",\b\4\1\2#$\7\13\2\2$,\b\4\1\2%&\7\f\2\2&,\b\4"+
		"\1\2\'(\7\16\2\2()\5\6\4\b)*\b\4\1\2*,\3\2\2\2+\32\3\2\2\2+\35\3\2\2\2"+
		"+\37\3\2\2\2+!\3\2\2\2+#\3\2\2\2+%\3\2\2\2+\'\3\2\2\2,H\3\2\2\2-.\f\7"+
		"\2\2./\7\17\2\2/\60\5\6\4\b\60\61\b\4\1\2\61G\3\2\2\2\62\63\f\6\2\2\63"+
		"\64\7\20\2\2\64\65\5\6\4\7\65\66\b\4\1\2\66G\3\2\2\2\678\f\5\2\289\7\21"+
		"\2\29:\5\6\4\6:;\b\4\1\2;G\3\2\2\2<=\f\4\2\2=>\7\r\2\2>?\5\6\4\5?@\b\4"+
		"\1\2@G\3\2\2\2AB\f\3\2\2BC\7\16\2\2CD\5\6\4\4DE\b\4\1\2EG\3\2\2\2F-\3"+
		"\2\2\2F\62\3\2\2\2F\67\3\2\2\2F<\3\2\2\2FA\3\2\2\2GJ\3\2\2\2HF\3\2\2\2"+
		"HI\3\2\2\2I\7\3\2\2\2JH\3\2\2\2KL\7\3\2\2LM\7\5\2\2MN\5\6\4\2NO\7\6\2"+
		"\2OP\5\n\6\2PQ\b\5\1\2Q\t\3\2\2\2RS\7\4\2\2S\13\3\2\2\2\6\22+FH";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}