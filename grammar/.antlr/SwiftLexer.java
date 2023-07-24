// Generated from /Users/daniel/Desktop/USAC SEGUNDO SEMESTRE 2023/COMPILADORES 2 - 781/LABORATORIO/PROYECTOS/OLC2_202112145/grammar/SwiftLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SwiftLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		WHITESPACE=1, COMMENT=2, BLOCK_COMMENT=3, Kw_LET=4, Kw_VAR=5, Kw_FUNC=6, 
		Kw_STRUCT=7, Kw_IF=8, Kw_ELSE=9, Kw_SWITCH=10, Kw_CASE=11, Kw_DEFAULT=12, 
		Kw_FOR=13, Kw_WHILE=14, Kw_BREAK=15, Kw_CONTINUE=16, Kw_RETURN=17, Kw_DO=18, 
		Kw_INT=19, Kw_DOUBLE=20, Kw_BOOL=21, Kw_STRING=22, INT=23, DOUBLE=24, 
		BOOL=25, STRING=26, CHAR=27, ID=28, Op_ARROW=29, Op_MUL=30, Op_DIV=31, 
		Op_PLUS=32, Op_MINUS=33, Op_MOD=34, Op_ASSIGN=35, Op_MUL_ASSIGN=36, Op_DIV_ASSIGN=37, 
		Op_PLUS_ASSIGN=38, Op_MINUS_ASSIGN=39, Op_EQ=40, Op_NEQ=41, Op_LT=42, 
		Op_GT=43, Op_LE=44, Op_GE=45, Op_AND=46, Op_OR=47, Op_NOT=48, LPAREN=49, 
		RPAREN=50, LBRACE=51, RBRACE=52, LBRACKET=53, RBRACKET=54, COMMA=55, SEMICOLON=56, 
		COLON=57, DOT=58;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"WHITESPACE", "COMMENT", "BLOCK_COMMENT", "Kw_LET", "Kw_VAR", "Kw_FUNC", 
			"Kw_STRUCT", "Kw_IF", "Kw_ELSE", "Kw_SWITCH", "Kw_CASE", "Kw_DEFAULT", 
			"Kw_FOR", "Kw_WHILE", "Kw_BREAK", "Kw_CONTINUE", "Kw_RETURN", "Kw_DO", 
			"Kw_INT", "Kw_DOUBLE", "Kw_BOOL", "Kw_STRING", "INT", "DOUBLE", "BOOL", 
			"STRING", "CHAR", "ID", "Op_ARROW", "Op_MUL", "Op_DIV", "Op_PLUS", "Op_MINUS", 
			"Op_MOD", "Op_ASSIGN", "Op_MUL_ASSIGN", "Op_DIV_ASSIGN", "Op_PLUS_ASSIGN", 
			"Op_MINUS_ASSIGN", "Op_EQ", "Op_NEQ", "Op_LT", "Op_GT", "Op_LE", "Op_GE", 
			"Op_AND", "Op_OR", "Op_NOT", "LPAREN", "RPAREN", "LBRACE", "RBRACE", 
			"LBRACKET", "RBRACKET", "COMMA", "SEMICOLON", "COLON", "DOT"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, "'let'", "'var'", "'func'", "'struct'", "'if'", 
			"'else'", "'switch'", "'case'", "'default'", "'for'", "'while'", "'break'", 
			"'continue'", "'return'", "'do'", "'Int'", "'Double'", "'Bool'", "'String'", 
			null, null, null, null, null, null, "'->'", "'*'", "'/'", "'+'", "'-'", 
			"'%'", "'='", "'*='", "'/='", "'+='", "'-='", "'=='", "'!='", "'<'", 
			"'>'", "'<='", "'>='", "'&&'", "'||'", "'!'", "'('", "')'", "'{'", "'}'", 
			"'['", "']'", "','", "';'", "':'", "'.'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "WHITESPACE", "COMMENT", "BLOCK_COMMENT", "Kw_LET", "Kw_VAR", "Kw_FUNC", 
			"Kw_STRUCT", "Kw_IF", "Kw_ELSE", "Kw_SWITCH", "Kw_CASE", "Kw_DEFAULT", 
			"Kw_FOR", "Kw_WHILE", "Kw_BREAK", "Kw_CONTINUE", "Kw_RETURN", "Kw_DO", 
			"Kw_INT", "Kw_DOUBLE", "Kw_BOOL", "Kw_STRING", "INT", "DOUBLE", "BOOL", 
			"STRING", "CHAR", "ID", "Op_ARROW", "Op_MUL", "Op_DIV", "Op_PLUS", "Op_MINUS", 
			"Op_MOD", "Op_ASSIGN", "Op_MUL_ASSIGN", "Op_DIV_ASSIGN", "Op_PLUS_ASSIGN", 
			"Op_MINUS_ASSIGN", "Op_EQ", "Op_NEQ", "Op_LT", "Op_GT", "Op_LE", "Op_GE", 
			"Op_AND", "Op_OR", "Op_NOT", "LPAREN", "RPAREN", "LBRACE", "RBRACE", 
			"LBRACKET", "RBRACKET", "COMMA", "SEMICOLON", "COLON", "DOT"
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


	public SwiftLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "SwiftLexer.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2<\u0181\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\3\2\6\2y\n\2"+
		"\r\2\16\2z\3\2\3\2\3\3\3\3\3\3\3\3\7\3\u0083\n\3\f\3\16\3\u0086\13\3\3"+
		"\3\3\3\3\4\3\4\3\4\3\4\7\4\u008e\n\4\f\4\16\4\u0091\13\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\16"+
		"\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\30\6\30\u0103\n\30\r\30\16\30\u0104\3\31\7\31\u0108\n\31\f\31"+
		"\16\31\u010b\13\31\3\31\3\31\6\31\u010f\n\31\r\31\16\31\u0110\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\5\32\u011c\n\32\3\33\3\33\3\33\3\33"+
		"\7\33\u0122\n\33\f\33\16\33\u0125\13\33\3\33\3\33\3\34\3\34\3\34\3\34"+
		"\7\34\u012d\n\34\f\34\16\34\u0130\13\34\3\34\3\34\3\35\3\35\7\35\u0136"+
		"\n\35\f\35\16\35\u0139\13\35\3\36\3\36\3\36\3\37\3\37\3 \3 \3!\3!\3\""+
		"\3\"\3#\3#\3$\3$\3%\3%\3%\3&\3&\3&\3\'\3\'\3\'\3(\3(\3(\3)\3)\3)\3*\3"+
		"*\3*\3+\3+\3,\3,\3-\3-\3-\3.\3.\3.\3/\3/\3/\3\60\3\60\3\60\3\61\3\61\3"+
		"\62\3\62\3\63\3\63\3\64\3\64\3\65\3\65\3\66\3\66\3\67\3\67\38\38\39\3"+
		"9\3:\3:\3;\3;\3\u008f\2<\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f"+
		"\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63"+
		"\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62"+
		"c\63e\64g\65i\66k\67m8o9q:s;u<\3\2\n\6\2\13\f\17\17\"\"^^\4\2\f\f\17\17"+
		"\3\2\62;\6\2\f\f\17\17$$^^\5\2\f\f\17\17^^\6\2\f\f\17\17))^^\5\2C\\aa"+
		"c|\6\2\62;C\\aac|\2\u018c\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2"+
		"\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2"+
		"\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3"+
		"\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2"+
		"\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67"+
		"\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2"+
		"\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2"+
		"\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]"+
		"\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2"+
		"\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2"+
		"\3x\3\2\2\2\5~\3\2\2\2\7\u0089\3\2\2\2\t\u0097\3\2\2\2\13\u009b\3\2\2"+
		"\2\r\u009f\3\2\2\2\17\u00a4\3\2\2\2\21\u00ab\3\2\2\2\23\u00ae\3\2\2\2"+
		"\25\u00b3\3\2\2\2\27\u00ba\3\2\2\2\31\u00bf\3\2\2\2\33\u00c7\3\2\2\2\35"+
		"\u00cb\3\2\2\2\37\u00d1\3\2\2\2!\u00d7\3\2\2\2#\u00e0\3\2\2\2%\u00e7\3"+
		"\2\2\2\'\u00ea\3\2\2\2)\u00ee\3\2\2\2+\u00f5\3\2\2\2-\u00fa\3\2\2\2/\u0102"+
		"\3\2\2\2\61\u0109\3\2\2\2\63\u011b\3\2\2\2\65\u011d\3\2\2\2\67\u0128\3"+
		"\2\2\29\u0133\3\2\2\2;\u013a\3\2\2\2=\u013d\3\2\2\2?\u013f\3\2\2\2A\u0141"+
		"\3\2\2\2C\u0143\3\2\2\2E\u0145\3\2\2\2G\u0147\3\2\2\2I\u0149\3\2\2\2K"+
		"\u014c\3\2\2\2M\u014f\3\2\2\2O\u0152\3\2\2\2Q\u0155\3\2\2\2S\u0158\3\2"+
		"\2\2U\u015b\3\2\2\2W\u015d\3\2\2\2Y\u015f\3\2\2\2[\u0162\3\2\2\2]\u0165"+
		"\3\2\2\2_\u0168\3\2\2\2a\u016b\3\2\2\2c\u016d\3\2\2\2e\u016f\3\2\2\2g"+
		"\u0171\3\2\2\2i\u0173\3\2\2\2k\u0175\3\2\2\2m\u0177\3\2\2\2o\u0179\3\2"+
		"\2\2q\u017b\3\2\2\2s\u017d\3\2\2\2u\u017f\3\2\2\2wy\t\2\2\2xw\3\2\2\2"+
		"yz\3\2\2\2zx\3\2\2\2z{\3\2\2\2{|\3\2\2\2|}\b\2\2\2}\4\3\2\2\2~\177\7\61"+
		"\2\2\177\u0080\7\61\2\2\u0080\u0084\3\2\2\2\u0081\u0083\n\3\2\2\u0082"+
		"\u0081\3\2\2\2\u0083\u0086\3\2\2\2\u0084\u0082\3\2\2\2\u0084\u0085\3\2"+
		"\2\2\u0085\u0087\3\2\2\2\u0086\u0084\3\2\2\2\u0087\u0088\b\3\2\2\u0088"+
		"\6\3\2\2\2\u0089\u008a\7\61\2\2\u008a\u008b\7,\2\2\u008b\u008f\3\2\2\2"+
		"\u008c\u008e\13\2\2\2\u008d\u008c\3\2\2\2\u008e\u0091\3\2\2\2\u008f\u0090"+
		"\3\2\2\2\u008f\u008d\3\2\2\2\u0090\u0092\3\2\2\2\u0091\u008f\3\2\2\2\u0092"+
		"\u0093\7,\2\2\u0093\u0094\7\61\2\2\u0094\u0095\3\2\2\2\u0095\u0096\b\4"+
		"\2\2\u0096\b\3\2\2\2\u0097\u0098\7n\2\2\u0098\u0099\7g\2\2\u0099\u009a"+
		"\7v\2\2\u009a\n\3\2\2\2\u009b\u009c\7x\2\2\u009c\u009d\7c\2\2\u009d\u009e"+
		"\7t\2\2\u009e\f\3\2\2\2\u009f\u00a0\7h\2\2\u00a0\u00a1\7w\2\2\u00a1\u00a2"+
		"\7p\2\2\u00a2\u00a3\7e\2\2\u00a3\16\3\2\2\2\u00a4\u00a5\7u\2\2\u00a5\u00a6"+
		"\7v\2\2\u00a6\u00a7\7t\2\2\u00a7\u00a8\7w\2\2\u00a8\u00a9\7e\2\2\u00a9"+
		"\u00aa\7v\2\2\u00aa\20\3\2\2\2\u00ab\u00ac\7k\2\2\u00ac\u00ad\7h\2\2\u00ad"+
		"\22\3\2\2\2\u00ae\u00af\7g\2\2\u00af\u00b0\7n\2\2\u00b0\u00b1\7u\2\2\u00b1"+
		"\u00b2\7g\2\2\u00b2\24\3\2\2\2\u00b3\u00b4\7u\2\2\u00b4\u00b5\7y\2\2\u00b5"+
		"\u00b6\7k\2\2\u00b6\u00b7\7v\2\2\u00b7\u00b8\7e\2\2\u00b8\u00b9\7j\2\2"+
		"\u00b9\26\3\2\2\2\u00ba\u00bb\7e\2\2\u00bb\u00bc\7c\2\2\u00bc\u00bd\7"+
		"u\2\2\u00bd\u00be\7g\2\2\u00be\30\3\2\2\2\u00bf\u00c0\7f\2\2\u00c0\u00c1"+
		"\7g\2\2\u00c1\u00c2\7h\2\2\u00c2\u00c3\7c\2\2\u00c3\u00c4\7w\2\2\u00c4"+
		"\u00c5\7n\2\2\u00c5\u00c6\7v\2\2\u00c6\32\3\2\2\2\u00c7\u00c8\7h\2\2\u00c8"+
		"\u00c9\7q\2\2\u00c9\u00ca\7t\2\2\u00ca\34\3\2\2\2\u00cb\u00cc\7y\2\2\u00cc"+
		"\u00cd\7j\2\2\u00cd\u00ce\7k\2\2\u00ce\u00cf\7n\2\2\u00cf\u00d0\7g\2\2"+
		"\u00d0\36\3\2\2\2\u00d1\u00d2\7d\2\2\u00d2\u00d3\7t\2\2\u00d3\u00d4\7"+
		"g\2\2\u00d4\u00d5\7c\2\2\u00d5\u00d6\7m\2\2\u00d6 \3\2\2\2\u00d7\u00d8"+
		"\7e\2\2\u00d8\u00d9\7q\2\2\u00d9\u00da\7p\2\2\u00da\u00db\7v\2\2\u00db"+
		"\u00dc\7k\2\2\u00dc\u00dd\7p\2\2\u00dd\u00de\7w\2\2\u00de\u00df\7g\2\2"+
		"\u00df\"\3\2\2\2\u00e0\u00e1\7t\2\2\u00e1\u00e2\7g\2\2\u00e2\u00e3\7v"+
		"\2\2\u00e3\u00e4\7w\2\2\u00e4\u00e5\7t\2\2\u00e5\u00e6\7p\2\2\u00e6$\3"+
		"\2\2\2\u00e7\u00e8\7f\2\2\u00e8\u00e9\7q\2\2\u00e9&\3\2\2\2\u00ea\u00eb"+
		"\7K\2\2\u00eb\u00ec\7p\2\2\u00ec\u00ed\7v\2\2\u00ed(\3\2\2\2\u00ee\u00ef"+
		"\7F\2\2\u00ef\u00f0\7q\2\2\u00f0\u00f1\7w\2\2\u00f1\u00f2\7d\2\2\u00f2"+
		"\u00f3\7n\2\2\u00f3\u00f4\7g\2\2\u00f4*\3\2\2\2\u00f5\u00f6\7D\2\2\u00f6"+
		"\u00f7\7q\2\2\u00f7\u00f8\7q\2\2\u00f8\u00f9\7n\2\2\u00f9,\3\2\2\2\u00fa"+
		"\u00fb\7U\2\2\u00fb\u00fc\7v\2\2\u00fc\u00fd\7t\2\2\u00fd\u00fe\7k\2\2"+
		"\u00fe\u00ff\7p\2\2\u00ff\u0100\7i\2\2\u0100.\3\2\2\2\u0101\u0103\t\4"+
		"\2\2\u0102\u0101\3\2\2\2\u0103\u0104\3\2\2\2\u0104\u0102\3\2\2\2\u0104"+
		"\u0105\3\2\2\2\u0105\60\3\2\2\2\u0106\u0108\t\4\2\2\u0107\u0106\3\2\2"+
		"\2\u0108\u010b\3\2\2\2\u0109\u0107\3\2\2\2\u0109\u010a\3\2\2\2\u010a\u010c"+
		"\3\2\2\2\u010b\u0109\3\2\2\2\u010c\u010e\7\60\2\2\u010d\u010f\t\4\2\2"+
		"\u010e\u010d\3\2\2\2\u010f\u0110\3\2\2\2\u0110\u010e\3\2\2\2\u0110\u0111"+
		"\3\2\2\2\u0111\62\3\2\2\2\u0112\u0113\7v\2\2\u0113\u0114\7t\2\2\u0114"+
		"\u0115\7w\2\2\u0115\u011c\7g\2\2\u0116\u0117\7h\2\2\u0117\u0118\7c\2\2"+
		"\u0118\u0119\7n\2\2\u0119\u011a\7u\2\2\u011a\u011c\7g\2\2\u011b\u0112"+
		"\3\2\2\2\u011b\u0116\3\2\2\2\u011c\64\3\2\2\2\u011d\u0123\7$\2\2\u011e"+
		"\u0122\n\5\2\2\u011f\u0120\7^\2\2\u0120\u0122\t\6\2\2\u0121\u011e\3\2"+
		"\2\2\u0121\u011f\3\2\2\2\u0122\u0125\3\2\2\2\u0123\u0121\3\2\2\2\u0123"+
		"\u0124\3\2\2\2\u0124\u0126\3\2\2\2\u0125\u0123\3\2\2\2\u0126\u0127\7$"+
		"\2\2\u0127\66\3\2\2\2\u0128\u012e\7)\2\2\u0129\u012d\n\7\2\2\u012a\u012b"+
		"\7^\2\2\u012b\u012d\t\6\2\2\u012c\u0129\3\2\2\2\u012c\u012a\3\2\2\2\u012d"+
		"\u0130\3\2\2\2\u012e\u012c\3\2\2\2\u012e\u012f\3\2\2\2\u012f\u0131\3\2"+
		"\2\2\u0130\u012e\3\2\2\2\u0131\u0132\7)\2\2\u01328\3\2\2\2\u0133\u0137"+
		"\t\b\2\2\u0134\u0136\t\t\2\2\u0135\u0134\3\2\2\2\u0136\u0139\3\2\2\2\u0137"+
		"\u0135\3\2\2\2\u0137\u0138\3\2\2\2\u0138:\3\2\2\2\u0139\u0137\3\2\2\2"+
		"\u013a\u013b\7/\2\2\u013b\u013c\7@\2\2\u013c<\3\2\2\2\u013d\u013e\7,\2"+
		"\2\u013e>\3\2\2\2\u013f\u0140\7\61\2\2\u0140@\3\2\2\2\u0141\u0142\7-\2"+
		"\2\u0142B\3\2\2\2\u0143\u0144\7/\2\2\u0144D\3\2\2\2\u0145\u0146\7\'\2"+
		"\2\u0146F\3\2\2\2\u0147\u0148\7?\2\2\u0148H\3\2\2\2\u0149\u014a\7,\2\2"+
		"\u014a\u014b\7?\2\2\u014bJ\3\2\2\2\u014c\u014d\7\61\2\2\u014d\u014e\7"+
		"?\2\2\u014eL\3\2\2\2\u014f\u0150\7-\2\2\u0150\u0151\7?\2\2\u0151N\3\2"+
		"\2\2\u0152\u0153\7/\2\2\u0153\u0154\7?\2\2\u0154P\3\2\2\2\u0155\u0156"+
		"\7?\2\2\u0156\u0157\7?\2\2\u0157R\3\2\2\2\u0158\u0159\7#\2\2\u0159\u015a"+
		"\7?\2\2\u015aT\3\2\2\2\u015b\u015c\7>\2\2\u015cV\3\2\2\2\u015d\u015e\7"+
		"@\2\2\u015eX\3\2\2\2\u015f\u0160\7>\2\2\u0160\u0161\7?\2\2\u0161Z\3\2"+
		"\2\2\u0162\u0163\7@\2\2\u0163\u0164\7?\2\2\u0164\\\3\2\2\2\u0165\u0166"+
		"\7(\2\2\u0166\u0167\7(\2\2\u0167^\3\2\2\2\u0168\u0169\7~\2\2\u0169\u016a"+
		"\7~\2\2\u016a`\3\2\2\2\u016b\u016c\7#\2\2\u016cb\3\2\2\2\u016d\u016e\7"+
		"*\2\2\u016ed\3\2\2\2\u016f\u0170\7+\2\2\u0170f\3\2\2\2\u0171\u0172\7}"+
		"\2\2\u0172h\3\2\2\2\u0173\u0174\7\177\2\2\u0174j\3\2\2\2\u0175\u0176\7"+
		"]\2\2\u0176l\3\2\2\2\u0177\u0178\7_\2\2\u0178n\3\2\2\2\u0179\u017a\7."+
		"\2\2\u017ap\3\2\2\2\u017b\u017c\7=\2\2\u017cr\3\2\2\2\u017d\u017e\7<\2"+
		"\2\u017et\3\2\2\2\u017f\u0180\7\60\2\2\u0180v\3\2\2\2\17\2z\u0084\u008f"+
		"\u0104\u0109\u0110\u011b\u0121\u0123\u012c\u012e\u0137\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}