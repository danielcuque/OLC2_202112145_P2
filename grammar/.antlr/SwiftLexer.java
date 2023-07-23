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
		WHITESPACE=1, Kw_LET=2, Kw_VAR=3, Kw_FUNC=4, Kw_IF=5, Kw_ELSE=6, Kw_SWITCH=7, 
		Kw_CASE=8, Kw_DEFAULT=9, Kw_FOR=10, Kw_WHILE=11, Kw_BREAK=12, Kw_CONTINUE=13, 
		Kw_RETURN=14, Kw_DO=15, Kw_INT=16, Kw_FLOAT=17, Kw_DOUBLE=18, Kw_BOOL=19, 
		Kw_STRING=20, ID=21, INT=22, FLOAT=23, DOUBLE=24, BOOL=25, STRING=26, 
		CHAR=27, Op_PLUS=28, Op_MINUS=29, Op_MUL=30, Op_DIV=31, Op_MOD=32, Op_ASSIGN=33, 
		Op_EQ=34, Op_NEQ=35, Op_LT=36, Op_GT=37, Op_LE=38, Op_GE=39, Op_AND=40, 
		Op_OR=41, Op_NOT=42, Op_INC=43, Op_DEC=44, LPAREN=45, RPAREN=46, LBRACE=47, 
		RBRACE=48, LBRACKET=49, RBRACKET=50, COMMA=51, SEMICOLON=52, COLON=53, 
		DOT=54;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"WHITESPACE", "Kw_LET", "Kw_VAR", "Kw_FUNC", "Kw_IF", "Kw_ELSE", "Kw_SWITCH", 
			"Kw_CASE", "Kw_DEFAULT", "Kw_FOR", "Kw_WHILE", "Kw_BREAK", "Kw_CONTINUE", 
			"Kw_RETURN", "Kw_DO", "Kw_INT", "Kw_FLOAT", "Kw_DOUBLE", "Kw_BOOL", "Kw_STRING", 
			"ID", "INT", "FLOAT", "DOUBLE", "BOOL", "STRING", "CHAR", "Op_PLUS", 
			"Op_MINUS", "Op_MUL", "Op_DIV", "Op_MOD", "Op_ASSIGN", "Op_EQ", "Op_NEQ", 
			"Op_LT", "Op_GT", "Op_LE", "Op_GE", "Op_AND", "Op_OR", "Op_NOT", "Op_INC", 
			"Op_DEC", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET", 
			"COMMA", "SEMICOLON", "COLON", "DOT"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'let'", "'var'", "'func'", "'if'", "'else'", "'switch'", 
			"'case'", "'default'", "'for'", "'while'", "'break'", "'continue'", "'return'", 
			"'do'", "'Int'", "'Float'", "'Double'", "'Bool'", "'String'", null, null, 
			null, null, null, null, null, "'+'", "'-'", "'*'", "'/'", "'%'", "'='", 
			"'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'&&'", "'||'", "'!'", 
			"'++'", "'--'", "'('", "')'", "'{'", "'}'", "'['", "']'", "','", "';'", 
			"':'", "'.'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "WHITESPACE", "Kw_LET", "Kw_VAR", "Kw_FUNC", "Kw_IF", "Kw_ELSE", 
			"Kw_SWITCH", "Kw_CASE", "Kw_DEFAULT", "Kw_FOR", "Kw_WHILE", "Kw_BREAK", 
			"Kw_CONTINUE", "Kw_RETURN", "Kw_DO", "Kw_INT", "Kw_FLOAT", "Kw_DOUBLE", 
			"Kw_BOOL", "Kw_STRING", "ID", "INT", "FLOAT", "DOUBLE", "BOOL", "STRING", 
			"CHAR", "Op_PLUS", "Op_MINUS", "Op_MUL", "Op_DIV", "Op_MOD", "Op_ASSIGN", 
			"Op_EQ", "Op_NEQ", "Op_LT", "Op_GT", "Op_LE", "Op_GE", "Op_AND", "Op_OR", 
			"Op_NOT", "Op_INC", "Op_DEC", "LPAREN", "RPAREN", "LBRACE", "RBRACE", 
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\28\u0162\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\3\2\6\2q\n\2\r\2\16\2r\3\2\3\2\3\3\3"+
		"\3\3\3\3\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\7\3\7\3\7"+
		"\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r"+
		"\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24"+
		"\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\7\26\u00e2\n\26"+
		"\f\26\16\26\u00e5\13\26\3\27\6\27\u00e8\n\27\r\27\16\27\u00e9\3\30\7\30"+
		"\u00ed\n\30\f\30\16\30\u00f0\13\30\3\30\3\30\6\30\u00f4\n\30\r\30\16\30"+
		"\u00f5\3\31\7\31\u00f9\n\31\f\31\16\31\u00fc\13\31\3\31\3\31\6\31\u0100"+
		"\n\31\r\31\16\31\u0101\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\5"+
		"\32\u010d\n\32\3\33\3\33\3\33\3\33\7\33\u0113\n\33\f\33\16\33\u0116\13"+
		"\33\3\33\3\33\3\34\3\34\3\34\3\34\7\34\u011e\n\34\f\34\16\34\u0121\13"+
		"\34\3\34\3\34\3\35\3\35\3\36\3\36\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#"+
		"\3#\3$\3$\3$\3%\3%\3&\3&\3\'\3\'\3\'\3(\3(\3(\3)\3)\3)\3*\3*\3*\3+\3+"+
		"\3,\3,\3,\3-\3-\3-\3.\3.\3/\3/\3\60\3\60\3\61\3\61\3\62\3\62\3\63\3\63"+
		"\3\64\3\64\3\65\3\65\3\66\3\66\3\67\3\67\2\28\3\3\5\4\7\5\t\6\13\7\r\b"+
		"\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26"+
		"+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S"+
		"+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8\3\2\t\6\2\13\f\17\17\"\"^"+
		"^\5\2C\\aac|\6\2\62;C\\aac|\3\2\62;\6\2\f\f\17\17$$^^\5\2\f\f\17\17^^"+
		"\6\2\f\f\17\17))^^\2\u016d\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2"+
		"\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2"+
		"\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3"+
		"\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2"+
		"\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67"+
		"\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2"+
		"\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2"+
		"\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]"+
		"\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2"+
		"\2\2\2k\3\2\2\2\2m\3\2\2\2\3p\3\2\2\2\5v\3\2\2\2\7z\3\2\2\2\t~\3\2\2\2"+
		"\13\u0083\3\2\2\2\r\u0086\3\2\2\2\17\u008b\3\2\2\2\21\u0092\3\2\2\2\23"+
		"\u0097\3\2\2\2\25\u009f\3\2\2\2\27\u00a3\3\2\2\2\31\u00a9\3\2\2\2\33\u00af"+
		"\3\2\2\2\35\u00b8\3\2\2\2\37\u00bf\3\2\2\2!\u00c2\3\2\2\2#\u00c6\3\2\2"+
		"\2%\u00cc\3\2\2\2\'\u00d3\3\2\2\2)\u00d8\3\2\2\2+\u00df\3\2\2\2-\u00e7"+
		"\3\2\2\2/\u00ee\3\2\2\2\61\u00fa\3\2\2\2\63\u010c\3\2\2\2\65\u010e\3\2"+
		"\2\2\67\u0119\3\2\2\29\u0124\3\2\2\2;\u0126\3\2\2\2=\u0128\3\2\2\2?\u012a"+
		"\3\2\2\2A\u012c\3\2\2\2C\u012e\3\2\2\2E\u0130\3\2\2\2G\u0133\3\2\2\2I"+
		"\u0136\3\2\2\2K\u0138\3\2\2\2M\u013a\3\2\2\2O\u013d\3\2\2\2Q\u0140\3\2"+
		"\2\2S\u0143\3\2\2\2U\u0146\3\2\2\2W\u0148\3\2\2\2Y\u014b\3\2\2\2[\u014e"+
		"\3\2\2\2]\u0150\3\2\2\2_\u0152\3\2\2\2a\u0154\3\2\2\2c\u0156\3\2\2\2e"+
		"\u0158\3\2\2\2g\u015a\3\2\2\2i\u015c\3\2\2\2k\u015e\3\2\2\2m\u0160\3\2"+
		"\2\2oq\t\2\2\2po\3\2\2\2qr\3\2\2\2rp\3\2\2\2rs\3\2\2\2st\3\2\2\2tu\b\2"+
		"\2\2u\4\3\2\2\2vw\7n\2\2wx\7g\2\2xy\7v\2\2y\6\3\2\2\2z{\7x\2\2{|\7c\2"+
		"\2|}\7t\2\2}\b\3\2\2\2~\177\7h\2\2\177\u0080\7w\2\2\u0080\u0081\7p\2\2"+
		"\u0081\u0082\7e\2\2\u0082\n\3\2\2\2\u0083\u0084\7k\2\2\u0084\u0085\7h"+
		"\2\2\u0085\f\3\2\2\2\u0086\u0087\7g\2\2\u0087\u0088\7n\2\2\u0088\u0089"+
		"\7u\2\2\u0089\u008a\7g\2\2\u008a\16\3\2\2\2\u008b\u008c\7u\2\2\u008c\u008d"+
		"\7y\2\2\u008d\u008e\7k\2\2\u008e\u008f\7v\2\2\u008f\u0090\7e\2\2\u0090"+
		"\u0091\7j\2\2\u0091\20\3\2\2\2\u0092\u0093\7e\2\2\u0093\u0094\7c\2\2\u0094"+
		"\u0095\7u\2\2\u0095\u0096\7g\2\2\u0096\22\3\2\2\2\u0097\u0098\7f\2\2\u0098"+
		"\u0099\7g\2\2\u0099\u009a\7h\2\2\u009a\u009b\7c\2\2\u009b\u009c\7w\2\2"+
		"\u009c\u009d\7n\2\2\u009d\u009e\7v\2\2\u009e\24\3\2\2\2\u009f\u00a0\7"+
		"h\2\2\u00a0\u00a1\7q\2\2\u00a1\u00a2\7t\2\2\u00a2\26\3\2\2\2\u00a3\u00a4"+
		"\7y\2\2\u00a4\u00a5\7j\2\2\u00a5\u00a6\7k\2\2\u00a6\u00a7\7n\2\2\u00a7"+
		"\u00a8\7g\2\2\u00a8\30\3\2\2\2\u00a9\u00aa\7d\2\2\u00aa\u00ab\7t\2\2\u00ab"+
		"\u00ac\7g\2\2\u00ac\u00ad\7c\2\2\u00ad\u00ae\7m\2\2\u00ae\32\3\2\2\2\u00af"+
		"\u00b0\7e\2\2\u00b0\u00b1\7q\2\2\u00b1\u00b2\7p\2\2\u00b2\u00b3\7v\2\2"+
		"\u00b3\u00b4\7k\2\2\u00b4\u00b5\7p\2\2\u00b5\u00b6\7w\2\2\u00b6\u00b7"+
		"\7g\2\2\u00b7\34\3\2\2\2\u00b8\u00b9\7t\2\2\u00b9\u00ba\7g\2\2\u00ba\u00bb"+
		"\7v\2\2\u00bb\u00bc\7w\2\2\u00bc\u00bd\7t\2\2\u00bd\u00be\7p\2\2\u00be"+
		"\36\3\2\2\2\u00bf\u00c0\7f\2\2\u00c0\u00c1\7q\2\2\u00c1 \3\2\2\2\u00c2"+
		"\u00c3\7K\2\2\u00c3\u00c4\7p\2\2\u00c4\u00c5\7v\2\2\u00c5\"\3\2\2\2\u00c6"+
		"\u00c7\7H\2\2\u00c7\u00c8\7n\2\2\u00c8\u00c9\7q\2\2\u00c9\u00ca\7c\2\2"+
		"\u00ca\u00cb\7v\2\2\u00cb$\3\2\2\2\u00cc\u00cd\7F\2\2\u00cd\u00ce\7q\2"+
		"\2\u00ce\u00cf\7w\2\2\u00cf\u00d0\7d\2\2\u00d0\u00d1\7n\2\2\u00d1\u00d2"+
		"\7g\2\2\u00d2&\3\2\2\2\u00d3\u00d4\7D\2\2\u00d4\u00d5\7q\2\2\u00d5\u00d6"+
		"\7q\2\2\u00d6\u00d7\7n\2\2\u00d7(\3\2\2\2\u00d8\u00d9\7U\2\2\u00d9\u00da"+
		"\7v\2\2\u00da\u00db\7t\2\2\u00db\u00dc\7k\2\2\u00dc\u00dd\7p\2\2\u00dd"+
		"\u00de\7i\2\2\u00de*\3\2\2\2\u00df\u00e3\t\3\2\2\u00e0\u00e2\t\4\2\2\u00e1"+
		"\u00e0\3\2\2\2\u00e2\u00e5\3\2\2\2\u00e3\u00e1\3\2\2\2\u00e3\u00e4\3\2"+
		"\2\2\u00e4,\3\2\2\2\u00e5\u00e3\3\2\2\2\u00e6\u00e8\t\5\2\2\u00e7\u00e6"+
		"\3\2\2\2\u00e8\u00e9\3\2\2\2\u00e9\u00e7\3\2\2\2\u00e9\u00ea\3\2\2\2\u00ea"+
		".\3\2\2\2\u00eb\u00ed\t\5\2\2\u00ec\u00eb\3\2\2\2\u00ed\u00f0\3\2\2\2"+
		"\u00ee\u00ec\3\2\2\2\u00ee\u00ef\3\2\2\2\u00ef\u00f1\3\2\2\2\u00f0\u00ee"+
		"\3\2\2\2\u00f1\u00f3\7\60\2\2\u00f2\u00f4\t\5\2\2\u00f3\u00f2\3\2\2\2"+
		"\u00f4\u00f5\3\2\2\2\u00f5\u00f3\3\2\2\2\u00f5\u00f6\3\2\2\2\u00f6\60"+
		"\3\2\2\2\u00f7\u00f9\t\5\2\2\u00f8\u00f7\3\2\2\2\u00f9\u00fc\3\2\2\2\u00fa"+
		"\u00f8\3\2\2\2\u00fa\u00fb\3\2\2\2\u00fb\u00fd\3\2\2\2\u00fc\u00fa\3\2"+
		"\2\2\u00fd\u00ff\7\60\2\2\u00fe\u0100\t\5\2\2\u00ff\u00fe\3\2\2\2\u0100"+
		"\u0101\3\2\2\2\u0101\u00ff\3\2\2\2\u0101\u0102\3\2\2\2\u0102\62\3\2\2"+
		"\2\u0103\u0104\7v\2\2\u0104\u0105\7t\2\2\u0105\u0106\7w\2\2\u0106\u010d"+
		"\7g\2\2\u0107\u0108\7h\2\2\u0108\u0109\7c\2\2\u0109\u010a\7n\2\2\u010a"+
		"\u010b\7u\2\2\u010b\u010d\7g\2\2\u010c\u0103\3\2\2\2\u010c\u0107\3\2\2"+
		"\2\u010d\64\3\2\2\2\u010e\u0114\7$\2\2\u010f\u0113\n\6\2\2\u0110\u0111"+
		"\7^\2\2\u0111\u0113\t\7\2\2\u0112\u010f\3\2\2\2\u0112\u0110\3\2\2\2\u0113"+
		"\u0116\3\2\2\2\u0114\u0112\3\2\2\2\u0114\u0115\3\2\2\2\u0115\u0117\3\2"+
		"\2\2\u0116\u0114\3\2\2\2\u0117\u0118\7$\2\2\u0118\66\3\2\2\2\u0119\u011f"+
		"\7)\2\2\u011a\u011e\n\b\2\2\u011b\u011c\7^\2\2\u011c\u011e\t\7\2\2\u011d"+
		"\u011a\3\2\2\2\u011d\u011b\3\2\2\2\u011e\u0121\3\2\2\2\u011f\u011d\3\2"+
		"\2\2\u011f\u0120\3\2\2\2\u0120\u0122\3\2\2\2\u0121\u011f\3\2\2\2\u0122"+
		"\u0123\7)\2\2\u01238\3\2\2\2\u0124\u0125\7-\2\2\u0125:\3\2\2\2\u0126\u0127"+
		"\7/\2\2\u0127<\3\2\2\2\u0128\u0129\7,\2\2\u0129>\3\2\2\2\u012a\u012b\7"+
		"\61\2\2\u012b@\3\2\2\2\u012c\u012d\7\'\2\2\u012dB\3\2\2\2\u012e\u012f"+
		"\7?\2\2\u012fD\3\2\2\2\u0130\u0131\7?\2\2\u0131\u0132\7?\2\2\u0132F\3"+
		"\2\2\2\u0133\u0134\7#\2\2\u0134\u0135\7?\2\2\u0135H\3\2\2\2\u0136\u0137"+
		"\7>\2\2\u0137J\3\2\2\2\u0138\u0139\7@\2\2\u0139L\3\2\2\2\u013a\u013b\7"+
		">\2\2\u013b\u013c\7?\2\2\u013cN\3\2\2\2\u013d\u013e\7@\2\2\u013e\u013f"+
		"\7?\2\2\u013fP\3\2\2\2\u0140\u0141\7(\2\2\u0141\u0142\7(\2\2\u0142R\3"+
		"\2\2\2\u0143\u0144\7~\2\2\u0144\u0145\7~\2\2\u0145T\3\2\2\2\u0146\u0147"+
		"\7#\2\2\u0147V\3\2\2\2\u0148\u0149\7-\2\2\u0149\u014a\7-\2\2\u014aX\3"+
		"\2\2\2\u014b\u014c\7/\2\2\u014c\u014d\7/\2\2\u014dZ\3\2\2\2\u014e\u014f"+
		"\7*\2\2\u014f\\\3\2\2\2\u0150\u0151\7+\2\2\u0151^\3\2\2\2\u0152\u0153"+
		"\7}\2\2\u0153`\3\2\2\2\u0154\u0155\7\177\2\2\u0155b\3\2\2\2\u0156\u0157"+
		"\7]\2\2\u0157d\3\2\2\2\u0158\u0159\7_\2\2\u0159f\3\2\2\2\u015a\u015b\7"+
		".\2\2\u015bh\3\2\2\2\u015c\u015d\7=\2\2\u015dj\3\2\2\2\u015e\u015f\7<"+
		"\2\2\u015fl\3\2\2\2\u0160\u0161\7\60\2\2\u0161n\3\2\2\2\17\2r\u00e3\u00e9"+
		"\u00ee\u00f5\u00fa\u0101\u010c\u0112\u0114\u011d\u011f\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}