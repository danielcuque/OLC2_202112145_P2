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
		Kw_IF=7, Kw_ELSE=8, Kw_SWITCH=9, Kw_CASE=10, Kw_DEFAULT=11, Kw_FOR=12, 
		Kw_WHILE=13, Kw_BREAK=14, Kw_CONTINUE=15, Kw_RETURN=16, Kw_DO=17, Kw_INT=18, 
		Kw_DOUBLE=19, Kw_BOOL=20, Kw_STRING=21, INT=22, DOUBLE=23, BOOL=24, STRING=25, 
		CHAR=26, ID=27, Op_PLUS=28, Op_MINUS=29, Op_MUL=30, Op_DIV=31, Op_MOD=32, 
		Op_ASSIGN=33, Op_EQ=34, Op_NEQ=35, Op_LT=36, Op_GT=37, Op_LE=38, Op_GE=39, 
		Op_AND=40, Op_OR=41, Op_NOT=42, Op_INC=43, Op_DEC=44, LPAREN=45, RPAREN=46, 
		LBRACE=47, RBRACE=48, LBRACKET=49, RBRACKET=50, COMMA=51, SEMICOLON=52, 
		COLON=53, DOT=54;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"WHITESPACE", "COMMENT", "BLOCK_COMMENT", "Kw_LET", "Kw_VAR", "Kw_FUNC", 
			"Kw_IF", "Kw_ELSE", "Kw_SWITCH", "Kw_CASE", "Kw_DEFAULT", "Kw_FOR", "Kw_WHILE", 
			"Kw_BREAK", "Kw_CONTINUE", "Kw_RETURN", "Kw_DO", "Kw_INT", "Kw_DOUBLE", 
			"Kw_BOOL", "Kw_STRING", "INT", "DOUBLE", "BOOL", "STRING", "CHAR", "ID", 
			"Op_PLUS", "Op_MINUS", "Op_MUL", "Op_DIV", "Op_MOD", "Op_ASSIGN", "Op_EQ", 
			"Op_NEQ", "Op_LT", "Op_GT", "Op_LE", "Op_GE", "Op_AND", "Op_OR", "Op_NOT", 
			"Op_INC", "Op_DEC", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACKET", 
			"RBRACKET", "COMMA", "SEMICOLON", "COLON", "DOT"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, "'let'", "'var'", "'func'", "'if'", "'else'", 
			"'switch'", "'case'", "'default'", "'for'", "'while'", "'break'", "'continue'", 
			"'return'", "'do'", "'Int'", "'Double'", "'Bool'", "'String'", null, 
			null, null, null, null, null, "'+'", "'-'", "'*'", "'/'", "'%'", "'='", 
			"'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'&&'", "'||'", "'!'", 
			"'++'", "'--'", "'('", "')'", "'{'", "'}'", "'['", "']'", "','", "';'", 
			"':'", "'.'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "WHITESPACE", "COMMENT", "BLOCK_COMMENT", "Kw_LET", "Kw_VAR", "Kw_FUNC", 
			"Kw_IF", "Kw_ELSE", "Kw_SWITCH", "Kw_CASE", "Kw_DEFAULT", "Kw_FOR", "Kw_WHILE", 
			"Kw_BREAK", "Kw_CONTINUE", "Kw_RETURN", "Kw_DO", "Kw_INT", "Kw_DOUBLE", 
			"Kw_BOOL", "Kw_STRING", "INT", "DOUBLE", "BOOL", "STRING", "CHAR", "ID", 
			"Op_PLUS", "Op_MINUS", "Op_MUL", "Op_DIV", "Op_MOD", "Op_ASSIGN", "Op_EQ", 
			"Op_NEQ", "Op_LT", "Op_GT", "Op_LE", "Op_GE", "Op_AND", "Op_OR", "Op_NOT", 
			"Op_INC", "Op_DEC", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACKET", 
			"RBRACKET", "COMMA", "SEMICOLON", "COLON", "DOT"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\28\u0169\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\3\2\6\2q\n\2\r\2\16\2r\3\2\3\2\3\3\3"+
		"\3\3\3\3\3\7\3{\n\3\f\3\16\3~\13\3\3\3\3\3\3\4\3\4\3\4\3\4\7\4\u0086\n"+
		"\4\f\4\16\4\u0089\13\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6"+
		"\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3"+
		"\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3"+
		"\21\3\21\3\21\3\21\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3"+
		"\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3"+
		"\26\3\26\3\27\6\27\u00f4\n\27\r\27\16\27\u00f5\3\30\7\30\u00f9\n\30\f"+
		"\30\16\30\u00fc\13\30\3\30\3\30\6\30\u0100\n\30\r\30\16\30\u0101\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\5\31\u010d\n\31\3\32\3\32\3\32"+
		"\3\32\7\32\u0113\n\32\f\32\16\32\u0116\13\32\3\32\3\32\3\33\3\33\3\33"+
		"\3\33\7\33\u011e\n\33\f\33\16\33\u0121\13\33\3\33\3\33\3\34\3\34\7\34"+
		"\u0127\n\34\f\34\16\34\u012a\13\34\3\35\3\35\3\36\3\36\3\37\3\37\3 \3"+
		" \3!\3!\3\"\3\"\3#\3#\3#\3$\3$\3$\3%\3%\3&\3&\3\'\3\'\3\'\3(\3(\3(\3)"+
		"\3)\3)\3*\3*\3*\3+\3+\3,\3,\3,\3-\3-\3-\3.\3.\3/\3/\3\60\3\60\3\61\3\61"+
		"\3\62\3\62\3\63\3\63\3\64\3\64\3\65\3\65\3\66\3\66\3\67\3\67\3\u0087\2"+
		"8\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20"+
		"\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37"+
		"= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8\3"+
		"\2\n\6\2\13\f\17\17\"\"^^\4\2\f\f\17\17\3\2\62;\6\2\f\f\17\17$$^^\5\2"+
		"\f\f\17\17^^\6\2\f\f\17\17))^^\5\2C\\aac|\6\2\62;C\\aac|\2\u0174\2\3\3"+
		"\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2"+
		"\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3"+
		"\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2"+
		"%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61"+
		"\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2"+
		"\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I"+
		"\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2"+
		"\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2"+
		"\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\3p"+
		"\3\2\2\2\5v\3\2\2\2\7\u0081\3\2\2\2\t\u008f\3\2\2\2\13\u0093\3\2\2\2\r"+
		"\u0097\3\2\2\2\17\u009c\3\2\2\2\21\u009f\3\2\2\2\23\u00a4\3\2\2\2\25\u00ab"+
		"\3\2\2\2\27\u00b0\3\2\2\2\31\u00b8\3\2\2\2\33\u00bc\3\2\2\2\35\u00c2\3"+
		"\2\2\2\37\u00c8\3\2\2\2!\u00d1\3\2\2\2#\u00d8\3\2\2\2%\u00db\3\2\2\2\'"+
		"\u00df\3\2\2\2)\u00e6\3\2\2\2+\u00eb\3\2\2\2-\u00f3\3\2\2\2/\u00fa\3\2"+
		"\2\2\61\u010c\3\2\2\2\63\u010e\3\2\2\2\65\u0119\3\2\2\2\67\u0124\3\2\2"+
		"\29\u012b\3\2\2\2;\u012d\3\2\2\2=\u012f\3\2\2\2?\u0131\3\2\2\2A\u0133"+
		"\3\2\2\2C\u0135\3\2\2\2E\u0137\3\2\2\2G\u013a\3\2\2\2I\u013d\3\2\2\2K"+
		"\u013f\3\2\2\2M\u0141\3\2\2\2O\u0144\3\2\2\2Q\u0147\3\2\2\2S\u014a\3\2"+
		"\2\2U\u014d\3\2\2\2W\u014f\3\2\2\2Y\u0152\3\2\2\2[\u0155\3\2\2\2]\u0157"+
		"\3\2\2\2_\u0159\3\2\2\2a\u015b\3\2\2\2c\u015d\3\2\2\2e\u015f\3\2\2\2g"+
		"\u0161\3\2\2\2i\u0163\3\2\2\2k\u0165\3\2\2\2m\u0167\3\2\2\2oq\t\2\2\2"+
		"po\3\2\2\2qr\3\2\2\2rp\3\2\2\2rs\3\2\2\2st\3\2\2\2tu\b\2\2\2u\4\3\2\2"+
		"\2vw\7\61\2\2wx\7\61\2\2x|\3\2\2\2y{\n\3\2\2zy\3\2\2\2{~\3\2\2\2|z\3\2"+
		"\2\2|}\3\2\2\2}\177\3\2\2\2~|\3\2\2\2\177\u0080\b\3\2\2\u0080\6\3\2\2"+
		"\2\u0081\u0082\7\61\2\2\u0082\u0083\7,\2\2\u0083\u0087\3\2\2\2\u0084\u0086"+
		"\13\2\2\2\u0085\u0084\3\2\2\2\u0086\u0089\3\2\2\2\u0087\u0088\3\2\2\2"+
		"\u0087\u0085\3\2\2\2\u0088\u008a\3\2\2\2\u0089\u0087\3\2\2\2\u008a\u008b"+
		"\7,\2\2\u008b\u008c\7\61\2\2\u008c\u008d\3\2\2\2\u008d\u008e\b\4\2\2\u008e"+
		"\b\3\2\2\2\u008f\u0090\7n\2\2\u0090\u0091\7g\2\2\u0091\u0092\7v\2\2\u0092"+
		"\n\3\2\2\2\u0093\u0094\7x\2\2\u0094\u0095\7c\2\2\u0095\u0096\7t\2\2\u0096"+
		"\f\3\2\2\2\u0097\u0098\7h\2\2\u0098\u0099\7w\2\2\u0099\u009a\7p\2\2\u009a"+
		"\u009b\7e\2\2\u009b\16\3\2\2\2\u009c\u009d\7k\2\2\u009d\u009e\7h\2\2\u009e"+
		"\20\3\2\2\2\u009f\u00a0\7g\2\2\u00a0\u00a1\7n\2\2\u00a1\u00a2\7u\2\2\u00a2"+
		"\u00a3\7g\2\2\u00a3\22\3\2\2\2\u00a4\u00a5\7u\2\2\u00a5\u00a6\7y\2\2\u00a6"+
		"\u00a7\7k\2\2\u00a7\u00a8\7v\2\2\u00a8\u00a9\7e\2\2\u00a9\u00aa\7j\2\2"+
		"\u00aa\24\3\2\2\2\u00ab\u00ac\7e\2\2\u00ac\u00ad\7c\2\2\u00ad\u00ae\7"+
		"u\2\2\u00ae\u00af\7g\2\2\u00af\26\3\2\2\2\u00b0\u00b1\7f\2\2\u00b1\u00b2"+
		"\7g\2\2\u00b2\u00b3\7h\2\2\u00b3\u00b4\7c\2\2\u00b4\u00b5\7w\2\2\u00b5"+
		"\u00b6\7n\2\2\u00b6\u00b7\7v\2\2\u00b7\30\3\2\2\2\u00b8\u00b9\7h\2\2\u00b9"+
		"\u00ba\7q\2\2\u00ba\u00bb\7t\2\2\u00bb\32\3\2\2\2\u00bc\u00bd\7y\2\2\u00bd"+
		"\u00be\7j\2\2\u00be\u00bf\7k\2\2\u00bf\u00c0\7n\2\2\u00c0\u00c1\7g\2\2"+
		"\u00c1\34\3\2\2\2\u00c2\u00c3\7d\2\2\u00c3\u00c4\7t\2\2\u00c4\u00c5\7"+
		"g\2\2\u00c5\u00c6\7c\2\2\u00c6\u00c7\7m\2\2\u00c7\36\3\2\2\2\u00c8\u00c9"+
		"\7e\2\2\u00c9\u00ca\7q\2\2\u00ca\u00cb\7p\2\2\u00cb\u00cc\7v\2\2\u00cc"+
		"\u00cd\7k\2\2\u00cd\u00ce\7p\2\2\u00ce\u00cf\7w\2\2\u00cf\u00d0\7g\2\2"+
		"\u00d0 \3\2\2\2\u00d1\u00d2\7t\2\2\u00d2\u00d3\7g\2\2\u00d3\u00d4\7v\2"+
		"\2\u00d4\u00d5\7w\2\2\u00d5\u00d6\7t\2\2\u00d6\u00d7\7p\2\2\u00d7\"\3"+
		"\2\2\2\u00d8\u00d9\7f\2\2\u00d9\u00da\7q\2\2\u00da$\3\2\2\2\u00db\u00dc"+
		"\7K\2\2\u00dc\u00dd\7p\2\2\u00dd\u00de\7v\2\2\u00de&\3\2\2\2\u00df\u00e0"+
		"\7F\2\2\u00e0\u00e1\7q\2\2\u00e1\u00e2\7w\2\2\u00e2\u00e3\7d\2\2\u00e3"+
		"\u00e4\7n\2\2\u00e4\u00e5\7g\2\2\u00e5(\3\2\2\2\u00e6\u00e7\7D\2\2\u00e7"+
		"\u00e8\7q\2\2\u00e8\u00e9\7q\2\2\u00e9\u00ea\7n\2\2\u00ea*\3\2\2\2\u00eb"+
		"\u00ec\7U\2\2\u00ec\u00ed\7v\2\2\u00ed\u00ee\7t\2\2\u00ee\u00ef\7k\2\2"+
		"\u00ef\u00f0\7p\2\2\u00f0\u00f1\7i\2\2\u00f1,\3\2\2\2\u00f2\u00f4\t\4"+
		"\2\2\u00f3\u00f2\3\2\2\2\u00f4\u00f5\3\2\2\2\u00f5\u00f3\3\2\2\2\u00f5"+
		"\u00f6\3\2\2\2\u00f6.\3\2\2\2\u00f7\u00f9\t\4\2\2\u00f8\u00f7\3\2\2\2"+
		"\u00f9\u00fc\3\2\2\2\u00fa\u00f8\3\2\2\2\u00fa\u00fb\3\2\2\2\u00fb\u00fd"+
		"\3\2\2\2\u00fc\u00fa\3\2\2\2\u00fd\u00ff\7\60\2\2\u00fe\u0100\t\4\2\2"+
		"\u00ff\u00fe\3\2\2\2\u0100\u0101\3\2\2\2\u0101\u00ff\3\2\2\2\u0101\u0102"+
		"\3\2\2\2\u0102\60\3\2\2\2\u0103\u0104\7v\2\2\u0104\u0105\7t\2\2\u0105"+
		"\u0106\7w\2\2\u0106\u010d\7g\2\2\u0107\u0108\7h\2\2\u0108\u0109\7c\2\2"+
		"\u0109\u010a\7n\2\2\u010a\u010b\7u\2\2\u010b\u010d\7g\2\2\u010c\u0103"+
		"\3\2\2\2\u010c\u0107\3\2\2\2\u010d\62\3\2\2\2\u010e\u0114\7$\2\2\u010f"+
		"\u0113\n\5\2\2\u0110\u0111\7^\2\2\u0111\u0113\t\6\2\2\u0112\u010f\3\2"+
		"\2\2\u0112\u0110\3\2\2\2\u0113\u0116\3\2\2\2\u0114\u0112\3\2\2\2\u0114"+
		"\u0115\3\2\2\2\u0115\u0117\3\2\2\2\u0116\u0114\3\2\2\2\u0117\u0118\7$"+
		"\2\2\u0118\64\3\2\2\2\u0119\u011f\7)\2\2\u011a\u011e\n\7\2\2\u011b\u011c"+
		"\7^\2\2\u011c\u011e\t\6\2\2\u011d\u011a\3\2\2\2\u011d\u011b\3\2\2\2\u011e"+
		"\u0121\3\2\2\2\u011f\u011d\3\2\2\2\u011f\u0120\3\2\2\2\u0120\u0122\3\2"+
		"\2\2\u0121\u011f\3\2\2\2\u0122\u0123\7)\2\2\u0123\66\3\2\2\2\u0124\u0128"+
		"\t\b\2\2\u0125\u0127\t\t\2\2\u0126\u0125\3\2\2\2\u0127\u012a\3\2\2\2\u0128"+
		"\u0126\3\2\2\2\u0128\u0129\3\2\2\2\u01298\3\2\2\2\u012a\u0128\3\2\2\2"+
		"\u012b\u012c\7-\2\2\u012c:\3\2\2\2\u012d\u012e\7/\2\2\u012e<\3\2\2\2\u012f"+
		"\u0130\7,\2\2\u0130>\3\2\2\2\u0131\u0132\7\61\2\2\u0132@\3\2\2\2\u0133"+
		"\u0134\7\'\2\2\u0134B\3\2\2\2\u0135\u0136\7?\2\2\u0136D\3\2\2\2\u0137"+
		"\u0138\7?\2\2\u0138\u0139\7?\2\2\u0139F\3\2\2\2\u013a\u013b\7#\2\2\u013b"+
		"\u013c\7?\2\2\u013cH\3\2\2\2\u013d\u013e\7>\2\2\u013eJ\3\2\2\2\u013f\u0140"+
		"\7@\2\2\u0140L\3\2\2\2\u0141\u0142\7>\2\2\u0142\u0143\7?\2\2\u0143N\3"+
		"\2\2\2\u0144\u0145\7@\2\2\u0145\u0146\7?\2\2\u0146P\3\2\2\2\u0147\u0148"+
		"\7(\2\2\u0148\u0149\7(\2\2\u0149R\3\2\2\2\u014a\u014b\7~\2\2\u014b\u014c"+
		"\7~\2\2\u014cT\3\2\2\2\u014d\u014e\7#\2\2\u014eV\3\2\2\2\u014f\u0150\7"+
		"-\2\2\u0150\u0151\7-\2\2\u0151X\3\2\2\2\u0152\u0153\7/\2\2\u0153\u0154"+
		"\7/\2\2\u0154Z\3\2\2\2\u0155\u0156\7*\2\2\u0156\\\3\2\2\2\u0157\u0158"+
		"\7+\2\2\u0158^\3\2\2\2\u0159\u015a\7}\2\2\u015a`\3\2\2\2\u015b\u015c\7"+
		"\177\2\2\u015cb\3\2\2\2\u015d\u015e\7]\2\2\u015ed\3\2\2\2\u015f\u0160"+
		"\7_\2\2\u0160f\3\2\2\2\u0161\u0162\7.\2\2\u0162h\3\2\2\2\u0163\u0164\7"+
		"=\2\2\u0164j\3\2\2\2\u0165\u0166\7<\2\2\u0166l\3\2\2\2\u0167\u0168\7\60"+
		"\2\2\u0168n\3\2\2\2\17\2r|\u0087\u00f5\u00fa\u0101\u010c\u0112\u0114\u011d"+
		"\u011f\u0128\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}