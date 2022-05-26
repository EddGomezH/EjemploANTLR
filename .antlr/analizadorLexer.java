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
		RFALSE=8, CHAR=9, STRING=10, COMMENT=11, WS=12;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"RPRINT", "PUNTOCOMA", "PARA", "PARC", "ENTERO", "FLOTANTE", "RTRUE", 
			"RFALSE", "CHAR", "STRING", "COMMENT", "WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'print'", "';'", "'('", "')'", null, null, "'true'", "'false'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "RPRINT", "PUNTOCOMA", "PARA", "PARC", "ENTERO", "FLOTANTE", "RTRUE", 
			"RFALSE", "CHAR", "STRING", "COMMENT", "WS"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\16h\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3"+
		"\6\6\6)\n\6\r\6\16\6*\3\7\6\7.\n\7\r\7\16\7/\3\7\3\7\6\7\64\n\7\r\7\16"+
		"\7\65\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\5\n"+
		"G\n\n\3\n\3\n\3\13\3\13\3\13\3\13\7\13O\n\13\f\13\16\13R\13\13\3\13\3"+
		"\13\3\f\3\f\3\f\3\f\6\fZ\n\f\r\f\16\f[\5\f^\n\f\3\f\3\f\3\r\6\rc\n\r\r"+
		"\r\16\rd\3\r\3\r\2\2\16\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f"+
		"\27\r\31\16\3\2\6\3\2\62;\4\2C\\c|\4\2\f\f\17\17\5\2\13\f\17\17\"\"\2"+
		"r\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2"+
		"\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2"+
		"\2\31\3\2\2\2\3\33\3\2\2\2\5!\3\2\2\2\7#\3\2\2\2\t%\3\2\2\2\13(\3\2\2"+
		"\2\r-\3\2\2\2\17\67\3\2\2\2\21<\3\2\2\2\23B\3\2\2\2\25J\3\2\2\2\27U\3"+
		"\2\2\2\31b\3\2\2\2\33\34\7r\2\2\34\35\7t\2\2\35\36\7k\2\2\36\37\7p\2\2"+
		"\37 \7v\2\2 \4\3\2\2\2!\"\7=\2\2\"\6\3\2\2\2#$\7*\2\2$\b\3\2\2\2%&\7+"+
		"\2\2&\n\3\2\2\2\')\t\2\2\2(\'\3\2\2\2)*\3\2\2\2*(\3\2\2\2*+\3\2\2\2+\f"+
		"\3\2\2\2,.\t\2\2\2-,\3\2\2\2./\3\2\2\2/-\3\2\2\2/\60\3\2\2\2\60\61\3\2"+
		"\2\2\61\63\7\60\2\2\62\64\t\2\2\2\63\62\3\2\2\2\64\65\3\2\2\2\65\63\3"+
		"\2\2\2\65\66\3\2\2\2\66\16\3\2\2\2\678\7v\2\289\7t\2\29:\7w\2\2:;\7g\2"+
		"\2;\20\3\2\2\2<=\7h\2\2=>\7c\2\2>?\7n\2\2?@\7u\2\2@A\7g\2\2A\22\3\2\2"+
		"\2BF\7)\2\2CG\t\3\2\2DG\5\31\r\2EG\t\2\2\2FC\3\2\2\2FD\3\2\2\2FE\3\2\2"+
		"\2GH\3\2\2\2HI\7)\2\2I\24\3\2\2\2JP\7$\2\2KO\t\3\2\2LO\5\31\r\2MO\t\2"+
		"\2\2NK\3\2\2\2NL\3\2\2\2NM\3\2\2\2OR\3\2\2\2PN\3\2\2\2PQ\3\2\2\2QS\3\2"+
		"\2\2RP\3\2\2\2ST\7$\2\2T\26\3\2\2\2UV\7\61\2\2VW\7\61\2\2W]\3\2\2\2XZ"+
		"\n\4\2\2YX\3\2\2\2Z[\3\2\2\2[Y\3\2\2\2[\\\3\2\2\2\\^\3\2\2\2]Y\3\2\2\2"+
		"]^\3\2\2\2^_\3\2\2\2_`\b\f\2\2`\30\3\2\2\2ac\t\5\2\2ba\3\2\2\2cd\3\2\2"+
		"\2db\3\2\2\2de\3\2\2\2ef\3\2\2\2fg\b\r\2\2g\32\3\2\2\2\f\2*/\65FNP[]d"+
		"\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}