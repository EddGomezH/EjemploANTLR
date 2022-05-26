// Generated from /home/emivnajera/go/src/PersonalProjects/EjemploANTLR/analizador.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class analizadorLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		RPRINT=1, PUNTOCOMA=2, PARA=3, PARC=4, ENTERO=5, FLOTANTE=6, RTRUE=7, 
		RFALSE=8, CHAR=9, STRING=10, MAS=11, MENOS=12, MUL=13, DIV=14, MOD=15, 
		COMMENT=16, WS=17;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"RPRINT", "PUNTOCOMA", "PARA", "PARC", "ENTERO", "FLOTANTE", "RTRUE", 
			"RFALSE", "CHAR", "STRING", "MAS", "MENOS", "MUL", "DIV", "MOD", "COMMENT", 
			"WS"
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


	public analizadorLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "analizador.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\23|\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\6\6\63\n\6\r\6\16"+
		"\6\64\3\7\6\78\n\7\r\7\16\79\3\7\3\7\6\7>\n\7\r\7\16\7?\3\b\3\b\3\b\3"+
		"\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\5\nQ\n\n\3\n\3\n\3\13\3"+
		"\13\3\13\3\13\7\13Y\n\13\f\13\16\13\\\13\13\3\13\3\13\3\f\3\f\3\r\3\r"+
		"\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\21\3\21\6\21n\n\21\r\21\16"+
		"\21o\5\21r\n\21\3\21\3\21\3\22\6\22w\n\22\r\22\16\22x\3\22\3\22\2\2\23"+
		"\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20"+
		"\37\21!\22#\23\3\2\6\3\2\62;\4\2C\\c|\4\2\f\f\17\17\5\2\13\f\17\17\"\""+
		"\2\u0086\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2"+
		"\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27"+
		"\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2"+
		"\2\2#\3\2\2\2\3%\3\2\2\2\5+\3\2\2\2\7-\3\2\2\2\t/\3\2\2\2\13\62\3\2\2"+
		"\2\r\67\3\2\2\2\17A\3\2\2\2\21F\3\2\2\2\23L\3\2\2\2\25T\3\2\2\2\27_\3"+
		"\2\2\2\31a\3\2\2\2\33c\3\2\2\2\35e\3\2\2\2\37g\3\2\2\2!i\3\2\2\2#v\3\2"+
		"\2\2%&\7r\2\2&\'\7t\2\2\'(\7k\2\2()\7p\2\2)*\7v\2\2*\4\3\2\2\2+,\7=\2"+
		"\2,\6\3\2\2\2-.\7*\2\2.\b\3\2\2\2/\60\7+\2\2\60\n\3\2\2\2\61\63\t\2\2"+
		"\2\62\61\3\2\2\2\63\64\3\2\2\2\64\62\3\2\2\2\64\65\3\2\2\2\65\f\3\2\2"+
		"\2\668\t\2\2\2\67\66\3\2\2\289\3\2\2\29\67\3\2\2\29:\3\2\2\2:;\3\2\2\2"+
		";=\7\60\2\2<>\t\2\2\2=<\3\2\2\2>?\3\2\2\2?=\3\2\2\2?@\3\2\2\2@\16\3\2"+
		"\2\2AB\7v\2\2BC\7t\2\2CD\7w\2\2DE\7g\2\2E\20\3\2\2\2FG\7h\2\2GH\7c\2\2"+
		"HI\7n\2\2IJ\7u\2\2JK\7g\2\2K\22\3\2\2\2LP\7)\2\2MQ\t\3\2\2NQ\5#\22\2O"+
		"Q\t\2\2\2PM\3\2\2\2PN\3\2\2\2PO\3\2\2\2QR\3\2\2\2RS\7)\2\2S\24\3\2\2\2"+
		"TZ\7$\2\2UY\t\3\2\2VY\5#\22\2WY\t\2\2\2XU\3\2\2\2XV\3\2\2\2XW\3\2\2\2"+
		"Y\\\3\2\2\2ZX\3\2\2\2Z[\3\2\2\2[]\3\2\2\2\\Z\3\2\2\2]^\7$\2\2^\26\3\2"+
		"\2\2_`\7-\2\2`\30\3\2\2\2ab\7/\2\2b\32\3\2\2\2cd\7,\2\2d\34\3\2\2\2ef"+
		"\7\61\2\2f\36\3\2\2\2gh\7\'\2\2h \3\2\2\2ij\7\61\2\2jk\7\61\2\2kq\3\2"+
		"\2\2ln\n\4\2\2ml\3\2\2\2no\3\2\2\2om\3\2\2\2op\3\2\2\2pr\3\2\2\2qm\3\2"+
		"\2\2qr\3\2\2\2rs\3\2\2\2st\b\21\2\2t\"\3\2\2\2uw\t\5\2\2vu\3\2\2\2wx\3"+
		"\2\2\2xv\3\2\2\2xy\3\2\2\2yz\3\2\2\2z{\b\22\2\2{$\3\2\2\2\f\2\649?PXZ"+
		"oqx\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}