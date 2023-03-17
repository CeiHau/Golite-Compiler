// Generated from /Users/wxh/Documents/UChicago/2023 Winter/Compiler/milestone3/grammars/GoliteParser.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class GoliteParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		NUMBER=1, PLUS=2, MINUS=3, ASSIGN=4, ASTERISK=5, FSLASH=6, PRINT=7, LET=8, 
		TYPE=9, STRUCT=10, INT_TYPE=11, BOOL_TYPE=12, VAR=13, FUNC=14, IF=15, 
		ELSE=16, FOR=17, RETURN=18, SCAN=19, DELETE=20, PRINTF=21, NEW=22, TRUE=23, 
		FALSE=24, NIL=25, LEFTCURL=26, RIGHTCURL=27, LEFTBRAC=28, RIGHTBRAC=29, 
		DOT=30, COMMA=31, EXCLAMATION=32, OR=33, AND=34, EQUALS=35, NOTEQUALS=36, 
		MORETHAN=37, LESSTHAN=38, GEQ=39, LEQ=40, ID=41, SEMICOLON=42, COMMENT=43, 
		STRING=44, WS=45;
	public static final int
		RULE_program = 0, RULE_types = 1, RULE_typeDeclaration = 2, RULE_fields = 3, 
		RULE_decl = 4, RULE_type = 5, RULE_declarations = 6, RULE_declaration = 7, 
		RULE_ids = 8, RULE_functions = 9, RULE_function = 10, RULE_parameters = 11, 
		RULE_returnType = 12, RULE_statements = 13, RULE_statement = 14, RULE_read = 15, 
		RULE_block = 16, RULE_delete = 17, RULE_assignment = 18, RULE_print = 19, 
		RULE_conditional = 20, RULE_loop = 21, RULE_return = 22, RULE_invocation = 23, 
		RULE_arguments = 24, RULE_lValue = 25, RULE_expression = 26, RULE_boolTerm = 27, 
		RULE_equalTerm = 28, RULE_opRelationTerm = 29, RULE_relationTerm = 30, 
		RULE_opSimpleTerm = 31, RULE_simpleTerm = 32, RULE_opTerm = 33, RULE_term = 34, 
		RULE_opUnaryTerm = 35, RULE_unaryTerm = 36, RULE_selectorTerm = 37, RULE_factor = 38;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "types", "typeDeclaration", "fields", "decl", "type", "declarations", 
			"declaration", "ids", "functions", "function", "parameters", "returnType", 
			"statements", "statement", "read", "block", "delete", "assignment", "print", 
			"conditional", "loop", "return", "invocation", "arguments", "lValue", 
			"expression", "boolTerm", "equalTerm", "opRelationTerm", "relationTerm", 
			"opSimpleTerm", "simpleTerm", "opTerm", "term", "opUnaryTerm", "unaryTerm", 
			"selectorTerm", "factor"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'+'", "'-'", "'='", "'*'", "'/'", "'print'", "'let'", "'type'", 
			"'struct'", "'int'", "'bool'", "'var'", "'func'", "'if'", "'else'", "'for'", 
			"'return'", "'scan'", "'delete'", "'printf'", "'new'", "'true'", "'false'", 
			"'nil'", "'{'", "'}'", "'('", "')'", "'.'", "','", "'!'", "'||'", "'&&'", 
			"'=='", "'!='", "'>'", "'<'", "'>='", "'<='", null, "';'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "NUMBER", "PLUS", "MINUS", "ASSIGN", "ASTERISK", "FSLASH", "PRINT", 
			"LET", "TYPE", "STRUCT", "INT_TYPE", "BOOL_TYPE", "VAR", "FUNC", "IF", 
			"ELSE", "FOR", "RETURN", "SCAN", "DELETE", "PRINTF", "NEW", "TRUE", "FALSE", 
			"NIL", "LEFTCURL", "RIGHTCURL", "LEFTBRAC", "RIGHTBRAC", "DOT", "COMMA", 
			"EXCLAMATION", "OR", "AND", "EQUALS", "NOTEQUALS", "MORETHAN", "LESSTHAN", 
			"GEQ", "LEQ", "ID", "SEMICOLON", "COMMENT", "STRING", "WS"
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
	public String getGrammarFileName() { return "GoliteParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GoliteParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgramContext extends ParserRuleContext {
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public DeclarationsContext declarations() {
			return getRuleContext(DeclarationsContext.class,0);
		}
		public FunctionsContext functions() {
			return getRuleContext(FunctionsContext.class,0);
		}
		public TerminalNode EOF() { return getToken(GoliteParser.EOF, 0); }
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(78);
			types();
			setState(79);
			declarations();
			setState(80);
			functions();
			setState(81);
			match(EOF);
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

	public static class TypesContext extends ParserRuleContext {
		public List<TypeDeclarationContext> typeDeclaration() {
			return getRuleContexts(TypeDeclarationContext.class);
		}
		public TypeDeclarationContext typeDeclaration(int i) {
			return getRuleContext(TypeDeclarationContext.class,i);
		}
		public TypesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_types; }
	}

	public final TypesContext types() throws RecognitionException {
		TypesContext _localctx = new TypesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_types);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(86);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TYPE) {
				{
				{
				setState(83);
				typeDeclaration();
				}
				}
				setState(88);
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

	public static class TypeDeclarationContext extends ParserRuleContext {
		public TerminalNode TYPE() { return getToken(GoliteParser.TYPE, 0); }
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public TerminalNode STRUCT() { return getToken(GoliteParser.STRUCT, 0); }
		public TerminalNode LEFTCURL() { return getToken(GoliteParser.LEFTCURL, 0); }
		public FieldsContext fields() {
			return getRuleContext(FieldsContext.class,0);
		}
		public TerminalNode RIGHTCURL() { return getToken(GoliteParser.RIGHTCURL, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public TypeDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeDeclaration; }
	}

	public final TypeDeclarationContext typeDeclaration() throws RecognitionException {
		TypeDeclarationContext _localctx = new TypeDeclarationContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_typeDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(89);
			match(TYPE);
			setState(90);
			match(ID);
			setState(91);
			match(STRUCT);
			setState(92);
			match(LEFTCURL);
			setState(93);
			fields();
			setState(94);
			match(RIGHTCURL);
			setState(95);
			match(SEMICOLON);
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

	public static class FieldsContext extends ParserRuleContext {
		public List<DeclContext> decl() {
			return getRuleContexts(DeclContext.class);
		}
		public DeclContext decl(int i) {
			return getRuleContext(DeclContext.class,i);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(GoliteParser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(GoliteParser.SEMICOLON, i);
		}
		public FieldsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fields; }
	}

	public final FieldsContext fields() throws RecognitionException {
		FieldsContext _localctx = new FieldsContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_fields);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(97);
			decl();
			setState(98);
			match(SEMICOLON);
			setState(104);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ID) {
				{
				{
				setState(99);
				decl();
				setState(100);
				match(SEMICOLON);
				}
				}
				setState(106);
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

	public static class DeclContext extends ParserRuleContext {
		public TypeContext ty;
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public DeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decl; }
	}

	public final DeclContext decl() throws RecognitionException {
		DeclContext _localctx = new DeclContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_decl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(107);
			match(ID);
			setState(108);
			((DeclContext)_localctx).ty = type();
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

	public static class TypeContext extends ParserRuleContext {
		public TerminalNode INT_TYPE() { return getToken(GoliteParser.INT_TYPE, 0); }
		public TerminalNode BOOL_TYPE() { return getToken(GoliteParser.BOOL_TYPE, 0); }
		public TerminalNode ASTERISK() { return getToken(GoliteParser.ASTERISK, 0); }
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_type);
		try {
			setState(114);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT_TYPE:
				enterOuterAlt(_localctx, 1);
				{
				setState(110);
				match(INT_TYPE);
				}
				break;
			case BOOL_TYPE:
				enterOuterAlt(_localctx, 2);
				{
				setState(111);
				match(BOOL_TYPE);
				}
				break;
			case ASTERISK:
				enterOuterAlt(_localctx, 3);
				{
				setState(112);
				match(ASTERISK);
				setState(113);
				match(ID);
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

	public static class DeclarationsContext extends ParserRuleContext {
		public List<DeclarationContext> declaration() {
			return getRuleContexts(DeclarationContext.class);
		}
		public DeclarationContext declaration(int i) {
			return getRuleContext(DeclarationContext.class,i);
		}
		public DeclarationsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarations; }
	}

	public final DeclarationsContext declarations() throws RecognitionException {
		DeclarationsContext _localctx = new DeclarationsContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_declarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(119);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(116);
				declaration();
				}
				}
				setState(121);
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

	public static class DeclarationContext extends ParserRuleContext {
		public TypeContext ty;
		public TerminalNode VAR() { return getToken(GoliteParser.VAR, 0); }
		public IdsContext ids() {
			return getRuleContext(IdsContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public DeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaration; }
	}

	public final DeclarationContext declaration() throws RecognitionException {
		DeclarationContext _localctx = new DeclarationContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(122);
			match(VAR);
			setState(123);
			ids();
			setState(124);
			((DeclarationContext)_localctx).ty = type();
			setState(125);
			match(SEMICOLON);
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

	public static class IdsContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(GoliteParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(GoliteParser.ID, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(GoliteParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(GoliteParser.COMMA, i);
		}
		public IdsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ids; }
	}

	public final IdsContext ids() throws RecognitionException {
		IdsContext _localctx = new IdsContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_ids);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(127);
			match(ID);
			setState(132);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(128);
				match(COMMA);
				setState(129);
				match(ID);
				}
				}
				setState(134);
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

	public static class FunctionsContext extends ParserRuleContext {
		public List<FunctionContext> function() {
			return getRuleContexts(FunctionContext.class);
		}
		public FunctionContext function(int i) {
			return getRuleContext(FunctionContext.class,i);
		}
		public FunctionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functions; }
	}

	public final FunctionsContext functions() throws RecognitionException {
		FunctionsContext _localctx = new FunctionsContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_functions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(138);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FUNC) {
				{
				{
				setState(135);
				function();
				}
				}
				setState(140);
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

	public static class FunctionContext extends ParserRuleContext {
		public TerminalNode FUNC() { return getToken(GoliteParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public ParametersContext parameters() {
			return getRuleContext(ParametersContext.class,0);
		}
		public TerminalNode LEFTCURL() { return getToken(GoliteParser.LEFTCURL, 0); }
		public DeclarationsContext declarations() {
			return getRuleContext(DeclarationsContext.class,0);
		}
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public TerminalNode RIGHTCURL() { return getToken(GoliteParser.RIGHTCURL, 0); }
		public ReturnTypeContext returnType() {
			return getRuleContext(ReturnTypeContext.class,0);
		}
		public FunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function; }
	}

	public final FunctionContext function() throws RecognitionException {
		FunctionContext _localctx = new FunctionContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_function);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(141);
			match(FUNC);
			setState(142);
			match(ID);
			setState(143);
			parameters();
			setState(145);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << ASTERISK) | (1L << INT_TYPE) | (1L << BOOL_TYPE))) != 0)) {
				{
				setState(144);
				returnType();
				}
			}

			setState(147);
			match(LEFTCURL);
			setState(148);
			declarations();
			setState(149);
			statements();
			setState(150);
			match(RIGHTCURL);
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

	public static class ParametersContext extends ParserRuleContext {
		public TerminalNode LEFTBRAC() { return getToken(GoliteParser.LEFTBRAC, 0); }
		public TerminalNode RIGHTBRAC() { return getToken(GoliteParser.RIGHTBRAC, 0); }
		public List<DeclContext> decl() {
			return getRuleContexts(DeclContext.class);
		}
		public DeclContext decl(int i) {
			return getRuleContext(DeclContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(GoliteParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(GoliteParser.COMMA, i);
		}
		public ParametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameters; }
	}

	public final ParametersContext parameters() throws RecognitionException {
		ParametersContext _localctx = new ParametersContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_parameters);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			match(LEFTBRAC);
			setState(161);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(153);
				decl();
				setState(158);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(154);
					match(COMMA);
					setState(155);
					decl();
					}
					}
					setState(160);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(163);
			match(RIGHTBRAC);
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

	public static class ReturnTypeContext extends ParserRuleContext {
		public TypeContext ty;
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public ReturnTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnType; }
	}

	public final ReturnTypeContext returnType() throws RecognitionException {
		ReturnTypeContext _localctx = new ReturnTypeContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_returnType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(165);
			((ReturnTypeContext)_localctx).ty = type();
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

	public static class StatementsContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public StatementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statements; }
	}

	public final StatementsContext statements() throws RecognitionException {
		StatementsContext _localctx = new StatementsContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_statements);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(170);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << IF) | (1L << FOR) | (1L << RETURN) | (1L << SCAN) | (1L << DELETE) | (1L << PRINTF) | (1L << LEFTCURL) | (1L << ID))) != 0)) {
				{
				{
				setState(167);
				statement();
				}
				}
				setState(172);
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

	public static class StatementContext extends ParserRuleContext {
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public AssignmentContext assignment() {
			return getRuleContext(AssignmentContext.class,0);
		}
		public PrintContext print() {
			return getRuleContext(PrintContext.class,0);
		}
		public DeleteContext delete() {
			return getRuleContext(DeleteContext.class,0);
		}
		public ConditionalContext conditional() {
			return getRuleContext(ConditionalContext.class,0);
		}
		public LoopContext loop() {
			return getRuleContext(LoopContext.class,0);
		}
		public ReturnContext return() {
			return getRuleContext(ReturnContext.class,0);
		}
		public ReadContext read() {
			return getRuleContext(ReadContext.class,0);
		}
		public InvocationContext invocation() {
			return getRuleContext(InvocationContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_statement);
		try {
			setState(182);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(173);
				block();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(174);
				assignment();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(175);
				print();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(176);
				delete();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(177);
				conditional();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(178);
				loop();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(179);
				return();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(180);
				read();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(181);
				invocation();
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

	public static class ReadContext extends ParserRuleContext {
		public TerminalNode SCAN() { return getToken(GoliteParser.SCAN, 0); }
		public LValueContext lValue() {
			return getRuleContext(LValueContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public ReadContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_read; }
	}

	public final ReadContext read() throws RecognitionException {
		ReadContext _localctx = new ReadContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_read);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(184);
			match(SCAN);
			setState(185);
			lValue();
			setState(186);
			match(SEMICOLON);
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

	public static class BlockContext extends ParserRuleContext {
		public TerminalNode LEFTCURL() { return getToken(GoliteParser.LEFTCURL, 0); }
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public TerminalNode RIGHTCURL() { return getToken(GoliteParser.RIGHTCURL, 0); }
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_block);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(188);
			match(LEFTCURL);
			setState(189);
			statements();
			setState(190);
			match(RIGHTCURL);
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

	public static class DeleteContext extends ParserRuleContext {
		public TerminalNode DELETE() { return getToken(GoliteParser.DELETE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public DeleteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_delete; }
	}

	public final DeleteContext delete() throws RecognitionException {
		DeleteContext _localctx = new DeleteContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_delete);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(192);
			match(DELETE);
			setState(193);
			expression();
			setState(194);
			match(SEMICOLON);
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

	public static class AssignmentContext extends ParserRuleContext {
		public LValueContext lValue() {
			return getRuleContext(LValueContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(GoliteParser.ASSIGN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(196);
			lValue();
			setState(197);
			match(ASSIGN);
			setState(198);
			expression();
			setState(199);
			match(SEMICOLON);
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

	public static class PrintContext extends ParserRuleContext {
		public TerminalNode PRINTF() { return getToken(GoliteParser.PRINTF, 0); }
		public TerminalNode LEFTBRAC() { return getToken(GoliteParser.LEFTBRAC, 0); }
		public TerminalNode STRING() { return getToken(GoliteParser.STRING, 0); }
		public TerminalNode RIGHTBRAC() { return getToken(GoliteParser.RIGHTBRAC, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public List<TerminalNode> COMMA() { return getTokens(GoliteParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(GoliteParser.COMMA, i);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public PrintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_print; }
	}

	public final PrintContext print() throws RecognitionException {
		PrintContext _localctx = new PrintContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_print);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(201);
			match(PRINTF);
			setState(202);
			match(LEFTBRAC);
			setState(203);
			match(STRING);
			setState(208);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(204);
				match(COMMA);
				setState(205);
				expression();
				}
				}
				setState(210);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(211);
			match(RIGHTBRAC);
			setState(212);
			match(SEMICOLON);
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

	public static class ConditionalContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(GoliteParser.IF, 0); }
		public TerminalNode LEFTBRAC() { return getToken(GoliteParser.LEFTBRAC, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RIGHTBRAC() { return getToken(GoliteParser.RIGHTBRAC, 0); }
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(GoliteParser.ELSE, 0); }
		public ConditionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditional; }
	}

	public final ConditionalContext conditional() throws RecognitionException {
		ConditionalContext _localctx = new ConditionalContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_conditional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(214);
			match(IF);
			setState(215);
			match(LEFTBRAC);
			setState(216);
			expression();
			setState(217);
			match(RIGHTBRAC);
			setState(218);
			block();
			setState(221);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(219);
				match(ELSE);
				setState(220);
				block();
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

	public static class LoopContext extends ParserRuleContext {
		public TerminalNode FOR() { return getToken(GoliteParser.FOR, 0); }
		public TerminalNode LEFTBRAC() { return getToken(GoliteParser.LEFTBRAC, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RIGHTBRAC() { return getToken(GoliteParser.RIGHTBRAC, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public LoopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loop; }
	}

	public final LoopContext loop() throws RecognitionException {
		LoopContext _localctx = new LoopContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_loop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(223);
			match(FOR);
			setState(224);
			match(LEFTBRAC);
			setState(225);
			expression();
			setState(226);
			match(RIGHTBRAC);
			setState(227);
			block();
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

	public static class ReturnContext extends ParserRuleContext {
		public TerminalNode RETURN() { return getToken(GoliteParser.RETURN, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ReturnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_return; }
	}

	public final ReturnContext return() throws RecognitionException {
		ReturnContext _localctx = new ReturnContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_return);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(229);
			match(RETURN);
			setState(231);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NUMBER) | (1L << MINUS) | (1L << NEW) | (1L << TRUE) | (1L << FALSE) | (1L << NIL) | (1L << LEFTBRAC) | (1L << EXCLAMATION) | (1L << ID))) != 0)) {
				{
				setState(230);
				expression();
				}
			}

			setState(233);
			match(SEMICOLON);
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

	public static class InvocationContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public InvocationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_invocation; }
	}

	public final InvocationContext invocation() throws RecognitionException {
		InvocationContext _localctx = new InvocationContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_invocation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(235);
			match(ID);
			setState(236);
			arguments();
			setState(237);
			match(SEMICOLON);
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

	public static class ArgumentsContext extends ParserRuleContext {
		public TerminalNode LEFTBRAC() { return getToken(GoliteParser.LEFTBRAC, 0); }
		public TerminalNode RIGHTBRAC() { return getToken(GoliteParser.RIGHTBRAC, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(GoliteParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(GoliteParser.COMMA, i);
		}
		public ArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arguments; }
	}

	public final ArgumentsContext arguments() throws RecognitionException {
		ArgumentsContext _localctx = new ArgumentsContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_arguments);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(239);
			match(LEFTBRAC);
			setState(248);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NUMBER) | (1L << MINUS) | (1L << NEW) | (1L << TRUE) | (1L << FALSE) | (1L << NIL) | (1L << LEFTBRAC) | (1L << EXCLAMATION) | (1L << ID))) != 0)) {
				{
				setState(240);
				expression();
				setState(245);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(241);
					match(COMMA);
					setState(242);
					expression();
					}
					}
					setState(247);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(250);
			match(RIGHTBRAC);
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

	public static class LValueContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(GoliteParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(GoliteParser.ID, i);
		}
		public List<TerminalNode> DOT() { return getTokens(GoliteParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(GoliteParser.DOT, i);
		}
		public LValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lValue; }
	}

	public final LValueContext lValue() throws RecognitionException {
		LValueContext _localctx = new LValueContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_lValue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(252);
			match(ID);
			setState(257);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(253);
				match(DOT);
				setState(254);
				match(ID);
				}
				}
				setState(259);
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

	public static class ExpressionContext extends ParserRuleContext {
		public List<BoolTermContext> boolTerm() {
			return getRuleContexts(BoolTermContext.class);
		}
		public BoolTermContext boolTerm(int i) {
			return getRuleContext(BoolTermContext.class,i);
		}
		public List<TerminalNode> OR() { return getTokens(GoliteParser.OR); }
		public TerminalNode OR(int i) {
			return getToken(GoliteParser.OR, i);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(260);
			boolTerm();
			setState(265);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==OR) {
				{
				{
				setState(261);
				match(OR);
				setState(262);
				boolTerm();
				}
				}
				setState(267);
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

	public static class BoolTermContext extends ParserRuleContext {
		public List<EqualTermContext> equalTerm() {
			return getRuleContexts(EqualTermContext.class);
		}
		public EqualTermContext equalTerm(int i) {
			return getRuleContext(EqualTermContext.class,i);
		}
		public List<TerminalNode> AND() { return getTokens(GoliteParser.AND); }
		public TerminalNode AND(int i) {
			return getToken(GoliteParser.AND, i);
		}
		public BoolTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_boolTerm; }
	}

	public final BoolTermContext boolTerm() throws RecognitionException {
		BoolTermContext _localctx = new BoolTermContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_boolTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(268);
			equalTerm();
			setState(273);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==AND) {
				{
				{
				setState(269);
				match(AND);
				setState(270);
				equalTerm();
				}
				}
				setState(275);
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

	public static class EqualTermContext extends ParserRuleContext {
		public RelationTermContext relationTerm() {
			return getRuleContext(RelationTermContext.class,0);
		}
		public List<OpRelationTermContext> opRelationTerm() {
			return getRuleContexts(OpRelationTermContext.class);
		}
		public OpRelationTermContext opRelationTerm(int i) {
			return getRuleContext(OpRelationTermContext.class,i);
		}
		public EqualTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_equalTerm; }
	}

	public final EqualTermContext equalTerm() throws RecognitionException {
		EqualTermContext _localctx = new EqualTermContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_equalTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(276);
			relationTerm();
			setState(280);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EQUALS || _la==NOTEQUALS) {
				{
				{
				setState(277);
				opRelationTerm();
				}
				}
				setState(282);
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

	public static class OpRelationTermContext extends ParserRuleContext {
		public Token op;
		public RelationTermContext rt;
		public RelationTermContext relationTerm() {
			return getRuleContext(RelationTermContext.class,0);
		}
		public TerminalNode EQUALS() { return getToken(GoliteParser.EQUALS, 0); }
		public TerminalNode NOTEQUALS() { return getToken(GoliteParser.NOTEQUALS, 0); }
		public OpRelationTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_opRelationTerm; }
	}

	public final OpRelationTermContext opRelationTerm() throws RecognitionException {
		OpRelationTermContext _localctx = new OpRelationTermContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_opRelationTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(283);
			((OpRelationTermContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==EQUALS || _la==NOTEQUALS) ) {
				((OpRelationTermContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(284);
			((OpRelationTermContext)_localctx).rt = relationTerm();
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

	public static class RelationTermContext extends ParserRuleContext {
		public SimpleTermContext simpleTerm() {
			return getRuleContext(SimpleTermContext.class,0);
		}
		public List<OpSimpleTermContext> opSimpleTerm() {
			return getRuleContexts(OpSimpleTermContext.class);
		}
		public OpSimpleTermContext opSimpleTerm(int i) {
			return getRuleContext(OpSimpleTermContext.class,i);
		}
		public RelationTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationTerm; }
	}

	public final RelationTermContext relationTerm() throws RecognitionException {
		RelationTermContext _localctx = new RelationTermContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_relationTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(286);
			simpleTerm();
			setState(290);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MORETHAN) | (1L << LESSTHAN) | (1L << GEQ) | (1L << LEQ))) != 0)) {
				{
				{
				setState(287);
				opSimpleTerm();
				}
				}
				setState(292);
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

	public static class OpSimpleTermContext extends ParserRuleContext {
		public Token op;
		public SimpleTermContext st;
		public SimpleTermContext simpleTerm() {
			return getRuleContext(SimpleTermContext.class,0);
		}
		public TerminalNode MORETHAN() { return getToken(GoliteParser.MORETHAN, 0); }
		public TerminalNode LESSTHAN() { return getToken(GoliteParser.LESSTHAN, 0); }
		public TerminalNode LEQ() { return getToken(GoliteParser.LEQ, 0); }
		public TerminalNode GEQ() { return getToken(GoliteParser.GEQ, 0); }
		public OpSimpleTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_opSimpleTerm; }
	}

	public final OpSimpleTermContext opSimpleTerm() throws RecognitionException {
		OpSimpleTermContext _localctx = new OpSimpleTermContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_opSimpleTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(293);
			((OpSimpleTermContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MORETHAN) | (1L << LESSTHAN) | (1L << GEQ) | (1L << LEQ))) != 0)) ) {
				((OpSimpleTermContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(294);
			((OpSimpleTermContext)_localctx).st = simpleTerm();
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

	public static class SimpleTermContext extends ParserRuleContext {
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public List<OpTermContext> opTerm() {
			return getRuleContexts(OpTermContext.class);
		}
		public OpTermContext opTerm(int i) {
			return getRuleContext(OpTermContext.class,i);
		}
		public SimpleTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simpleTerm; }
	}

	public final SimpleTermContext simpleTerm() throws RecognitionException {
		SimpleTermContext _localctx = new SimpleTermContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_simpleTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(296);
			term();
			setState(300);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==PLUS || _la==MINUS) {
				{
				{
				setState(297);
				opTerm();
				}
				}
				setState(302);
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

	public static class OpTermContext extends ParserRuleContext {
		public Token op;
		public TermContext t;
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public TerminalNode PLUS() { return getToken(GoliteParser.PLUS, 0); }
		public TerminalNode MINUS() { return getToken(GoliteParser.MINUS, 0); }
		public OpTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_opTerm; }
	}

	public final OpTermContext opTerm() throws RecognitionException {
		OpTermContext _localctx = new OpTermContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_opTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(303);
			((OpTermContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==PLUS || _la==MINUS) ) {
				((OpTermContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(304);
			((OpTermContext)_localctx).t = term();
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

	public static class TermContext extends ParserRuleContext {
		public UnaryTermContext unaryTerm() {
			return getRuleContext(UnaryTermContext.class,0);
		}
		public List<OpUnaryTermContext> opUnaryTerm() {
			return getRuleContexts(OpUnaryTermContext.class);
		}
		public OpUnaryTermContext opUnaryTerm(int i) {
			return getRuleContext(OpUnaryTermContext.class,i);
		}
		public TermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_term; }
	}

	public final TermContext term() throws RecognitionException {
		TermContext _localctx = new TermContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_term);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(306);
			unaryTerm();
			setState(310);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ASTERISK || _la==FSLASH) {
				{
				{
				setState(307);
				opUnaryTerm();
				}
				}
				setState(312);
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

	public static class OpUnaryTermContext extends ParserRuleContext {
		public Token op;
		public UnaryTermContext ut;
		public UnaryTermContext unaryTerm() {
			return getRuleContext(UnaryTermContext.class,0);
		}
		public TerminalNode ASTERISK() { return getToken(GoliteParser.ASTERISK, 0); }
		public TerminalNode FSLASH() { return getToken(GoliteParser.FSLASH, 0); }
		public OpUnaryTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_opUnaryTerm; }
	}

	public final OpUnaryTermContext opUnaryTerm() throws RecognitionException {
		OpUnaryTermContext _localctx = new OpUnaryTermContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_opUnaryTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(313);
			((OpUnaryTermContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==ASTERISK || _la==FSLASH) ) {
				((OpUnaryTermContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(314);
			((OpUnaryTermContext)_localctx).ut = unaryTerm();
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

	public static class UnaryTermContext extends ParserRuleContext {
		public TerminalNode EXCLAMATION() { return getToken(GoliteParser.EXCLAMATION, 0); }
		public SelectorTermContext selectorTerm() {
			return getRuleContext(SelectorTermContext.class,0);
		}
		public TerminalNode MINUS() { return getToken(GoliteParser.MINUS, 0); }
		public UnaryTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unaryTerm; }
	}

	public final UnaryTermContext unaryTerm() throws RecognitionException {
		UnaryTermContext _localctx = new UnaryTermContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_unaryTerm);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(321);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case EXCLAMATION:
				{
				setState(316);
				match(EXCLAMATION);
				setState(317);
				selectorTerm();
				}
				break;
			case MINUS:
				{
				setState(318);
				match(MINUS);
				setState(319);
				selectorTerm();
				}
				break;
			case NUMBER:
			case NEW:
			case TRUE:
			case FALSE:
			case NIL:
			case LEFTBRAC:
			case ID:
				{
				setState(320);
				selectorTerm();
				}
				break;
			default:
				throw new NoViableAltException(this);
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

	public static class SelectorTermContext extends ParserRuleContext {
		public FactorContext factor() {
			return getRuleContext(FactorContext.class,0);
		}
		public List<TerminalNode> DOT() { return getTokens(GoliteParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(GoliteParser.DOT, i);
		}
		public List<TerminalNode> ID() { return getTokens(GoliteParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(GoliteParser.ID, i);
		}
		public SelectorTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selectorTerm; }
	}

	public final SelectorTermContext selectorTerm() throws RecognitionException {
		SelectorTermContext _localctx = new SelectorTermContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_selectorTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(323);
			factor();
			setState(328);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(324);
				match(DOT);
				setState(325);
				match(ID);
				}
				}
				setState(330);
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

	public static class FactorContext extends ParserRuleContext {
		public TerminalNode LEFTBRAC() { return getToken(GoliteParser.LEFTBRAC, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RIGHTBRAC() { return getToken(GoliteParser.RIGHTBRAC, 0); }
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public TerminalNode NUMBER() { return getToken(GoliteParser.NUMBER, 0); }
		public TerminalNode NEW() { return getToken(GoliteParser.NEW, 0); }
		public TerminalNode TRUE() { return getToken(GoliteParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(GoliteParser.FALSE, 0); }
		public TerminalNode NIL() { return getToken(GoliteParser.NIL, 0); }
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public FactorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_factor; }
	}

	public final FactorContext factor() throws RecognitionException {
		FactorContext _localctx = new FactorContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_factor);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(345);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LEFTBRAC:
				{
				setState(331);
				match(LEFTBRAC);
				setState(332);
				expression();
				setState(333);
				match(RIGHTBRAC);
				}
				break;
			case ID:
				{
				setState(335);
				match(ID);
				setState(337);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LEFTBRAC) {
					{
					setState(336);
					arguments();
					}
				}

				}
				break;
			case NUMBER:
				{
				setState(339);
				match(NUMBER);
				}
				break;
			case NEW:
				{
				setState(340);
				match(NEW);
				setState(341);
				match(ID);
				}
				break;
			case TRUE:
				{
				setState(342);
				match(TRUE);
				}
				break;
			case FALSE:
				{
				setState(343);
				match(FALSE);
				}
				break;
			case NIL:
				{
				setState(344);
				match(NIL);
				}
				break;
			default:
				throw new NoViableAltException(this);
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

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3/\u015e\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\3\2\3\2\3\2\3\2\3\2"+
		"\3\3\7\3W\n\3\f\3\16\3Z\13\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3"+
		"\5\3\5\3\5\7\5i\n\5\f\5\16\5l\13\5\3\6\3\6\3\6\3\7\3\7\3\7\3\7\5\7u\n"+
		"\7\3\b\7\bx\n\b\f\b\16\b{\13\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\7\n\u0085"+
		"\n\n\f\n\16\n\u0088\13\n\3\13\7\13\u008b\n\13\f\13\16\13\u008e\13\13\3"+
		"\f\3\f\3\f\3\f\5\f\u0094\n\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\7\r\u009f"+
		"\n\r\f\r\16\r\u00a2\13\r\5\r\u00a4\n\r\3\r\3\r\3\16\3\16\3\17\7\17\u00ab"+
		"\n\17\f\17\16\17\u00ae\13\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3"+
		"\20\5\20\u00b9\n\20\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\23\3\23"+
		"\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\7\25\u00d1"+
		"\n\25\f\25\16\25\u00d4\13\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3"+
		"\26\3\26\5\26\u00e0\n\26\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\5\30"+
		"\u00ea\n\30\3\30\3\30\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\7\32\u00f6"+
		"\n\32\f\32\16\32\u00f9\13\32\5\32\u00fb\n\32\3\32\3\32\3\33\3\33\3\33"+
		"\7\33\u0102\n\33\f\33\16\33\u0105\13\33\3\34\3\34\3\34\7\34\u010a\n\34"+
		"\f\34\16\34\u010d\13\34\3\35\3\35\3\35\7\35\u0112\n\35\f\35\16\35\u0115"+
		"\13\35\3\36\3\36\7\36\u0119\n\36\f\36\16\36\u011c\13\36\3\37\3\37\3\37"+
		"\3 \3 \7 \u0123\n \f \16 \u0126\13 \3!\3!\3!\3\"\3\"\7\"\u012d\n\"\f\""+
		"\16\"\u0130\13\"\3#\3#\3#\3$\3$\7$\u0137\n$\f$\16$\u013a\13$\3%\3%\3%"+
		"\3&\3&\3&\3&\3&\5&\u0144\n&\3\'\3\'\3\'\7\'\u0149\n\'\f\'\16\'\u014c\13"+
		"\'\3(\3(\3(\3(\3(\3(\5(\u0154\n(\3(\3(\3(\3(\3(\3(\5(\u015c\n(\3(\2\2"+
		")\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>@BDF"+
		"HJLN\2\6\3\2%&\3\2\'*\3\2\4\5\3\2\7\b\2\u015f\2P\3\2\2\2\4X\3\2\2\2\6"+
		"[\3\2\2\2\bc\3\2\2\2\nm\3\2\2\2\ft\3\2\2\2\16y\3\2\2\2\20|\3\2\2\2\22"+
		"\u0081\3\2\2\2\24\u008c\3\2\2\2\26\u008f\3\2\2\2\30\u009a\3\2\2\2\32\u00a7"+
		"\3\2\2\2\34\u00ac\3\2\2\2\36\u00b8\3\2\2\2 \u00ba\3\2\2\2\"\u00be\3\2"+
		"\2\2$\u00c2\3\2\2\2&\u00c6\3\2\2\2(\u00cb\3\2\2\2*\u00d8\3\2\2\2,\u00e1"+
		"\3\2\2\2.\u00e7\3\2\2\2\60\u00ed\3\2\2\2\62\u00f1\3\2\2\2\64\u00fe\3\2"+
		"\2\2\66\u0106\3\2\2\28\u010e\3\2\2\2:\u0116\3\2\2\2<\u011d\3\2\2\2>\u0120"+
		"\3\2\2\2@\u0127\3\2\2\2B\u012a\3\2\2\2D\u0131\3\2\2\2F\u0134\3\2\2\2H"+
		"\u013b\3\2\2\2J\u0143\3\2\2\2L\u0145\3\2\2\2N\u015b\3\2\2\2PQ\5\4\3\2"+
		"QR\5\16\b\2RS\5\24\13\2ST\7\2\2\3T\3\3\2\2\2UW\5\6\4\2VU\3\2\2\2WZ\3\2"+
		"\2\2XV\3\2\2\2XY\3\2\2\2Y\5\3\2\2\2ZX\3\2\2\2[\\\7\13\2\2\\]\7+\2\2]^"+
		"\7\f\2\2^_\7\34\2\2_`\5\b\5\2`a\7\35\2\2ab\7,\2\2b\7\3\2\2\2cd\5\n\6\2"+
		"dj\7,\2\2ef\5\n\6\2fg\7,\2\2gi\3\2\2\2he\3\2\2\2il\3\2\2\2jh\3\2\2\2j"+
		"k\3\2\2\2k\t\3\2\2\2lj\3\2\2\2mn\7+\2\2no\5\f\7\2o\13\3\2\2\2pu\7\r\2"+
		"\2qu\7\16\2\2rs\7\7\2\2su\7+\2\2tp\3\2\2\2tq\3\2\2\2tr\3\2\2\2u\r\3\2"+
		"\2\2vx\5\20\t\2wv\3\2\2\2x{\3\2\2\2yw\3\2\2\2yz\3\2\2\2z\17\3\2\2\2{y"+
		"\3\2\2\2|}\7\17\2\2}~\5\22\n\2~\177\5\f\7\2\177\u0080\7,\2\2\u0080\21"+
		"\3\2\2\2\u0081\u0086\7+\2\2\u0082\u0083\7!\2\2\u0083\u0085\7+\2\2\u0084"+
		"\u0082\3\2\2\2\u0085\u0088\3\2\2\2\u0086\u0084\3\2\2\2\u0086\u0087\3\2"+
		"\2\2\u0087\23\3\2\2\2\u0088\u0086\3\2\2\2\u0089\u008b\5\26\f\2\u008a\u0089"+
		"\3\2\2\2\u008b\u008e\3\2\2\2\u008c\u008a\3\2\2\2\u008c\u008d\3\2\2\2\u008d"+
		"\25\3\2\2\2\u008e\u008c\3\2\2\2\u008f\u0090\7\20\2\2\u0090\u0091\7+\2"+
		"\2\u0091\u0093\5\30\r\2\u0092\u0094\5\32\16\2\u0093\u0092\3\2\2\2\u0093"+
		"\u0094\3\2\2\2\u0094\u0095\3\2\2\2\u0095\u0096\7\34\2\2\u0096\u0097\5"+
		"\16\b\2\u0097\u0098\5\34\17\2\u0098\u0099\7\35\2\2\u0099\27\3\2\2\2\u009a"+
		"\u00a3\7\36\2\2\u009b\u00a0\5\n\6\2\u009c\u009d\7!\2\2\u009d\u009f\5\n"+
		"\6\2\u009e\u009c\3\2\2\2\u009f\u00a2\3\2\2\2\u00a0\u009e\3\2\2\2\u00a0"+
		"\u00a1\3\2\2\2\u00a1\u00a4\3\2\2\2\u00a2\u00a0\3\2\2\2\u00a3\u009b\3\2"+
		"\2\2\u00a3\u00a4\3\2\2\2\u00a4\u00a5\3\2\2\2\u00a5\u00a6\7\37\2\2\u00a6"+
		"\31\3\2\2\2\u00a7\u00a8\5\f\7\2\u00a8\33\3\2\2\2\u00a9\u00ab\5\36\20\2"+
		"\u00aa\u00a9\3\2\2\2\u00ab\u00ae\3\2\2\2\u00ac\u00aa\3\2\2\2\u00ac\u00ad"+
		"\3\2\2\2\u00ad\35\3\2\2\2\u00ae\u00ac\3\2\2\2\u00af\u00b9\5\"\22\2\u00b0"+
		"\u00b9\5&\24\2\u00b1\u00b9\5(\25\2\u00b2\u00b9\5$\23\2\u00b3\u00b9\5*"+
		"\26\2\u00b4\u00b9\5,\27\2\u00b5\u00b9\5.\30\2\u00b6\u00b9\5 \21\2\u00b7"+
		"\u00b9\5\60\31\2\u00b8\u00af\3\2\2\2\u00b8\u00b0\3\2\2\2\u00b8\u00b1\3"+
		"\2\2\2\u00b8\u00b2\3\2\2\2\u00b8\u00b3\3\2\2\2\u00b8\u00b4\3\2\2\2\u00b8"+
		"\u00b5\3\2\2\2\u00b8\u00b6\3\2\2\2\u00b8\u00b7\3\2\2\2\u00b9\37\3\2\2"+
		"\2\u00ba\u00bb\7\25\2\2\u00bb\u00bc\5\64\33\2\u00bc\u00bd\7,\2\2\u00bd"+
		"!\3\2\2\2\u00be\u00bf\7\34\2\2\u00bf\u00c0\5\34\17\2\u00c0\u00c1\7\35"+
		"\2\2\u00c1#\3\2\2\2\u00c2\u00c3\7\26\2\2\u00c3\u00c4\5\66\34\2\u00c4\u00c5"+
		"\7,\2\2\u00c5%\3\2\2\2\u00c6\u00c7\5\64\33\2\u00c7\u00c8\7\6\2\2\u00c8"+
		"\u00c9\5\66\34\2\u00c9\u00ca\7,\2\2\u00ca\'\3\2\2\2\u00cb\u00cc\7\27\2"+
		"\2\u00cc\u00cd\7\36\2\2\u00cd\u00d2\7.\2\2\u00ce\u00cf\7!\2\2\u00cf\u00d1"+
		"\5\66\34\2\u00d0\u00ce\3\2\2\2\u00d1\u00d4\3\2\2\2\u00d2\u00d0\3\2\2\2"+
		"\u00d2\u00d3\3\2\2\2\u00d3\u00d5\3\2\2\2\u00d4\u00d2\3\2\2\2\u00d5\u00d6"+
		"\7\37\2\2\u00d6\u00d7\7,\2\2\u00d7)\3\2\2\2\u00d8\u00d9\7\21\2\2\u00d9"+
		"\u00da\7\36\2\2\u00da\u00db\5\66\34\2\u00db\u00dc\7\37\2\2\u00dc\u00df"+
		"\5\"\22\2\u00dd\u00de\7\22\2\2\u00de\u00e0\5\"\22\2\u00df\u00dd\3\2\2"+
		"\2\u00df\u00e0\3\2\2\2\u00e0+\3\2\2\2\u00e1\u00e2\7\23\2\2\u00e2\u00e3"+
		"\7\36\2\2\u00e3\u00e4\5\66\34\2\u00e4\u00e5\7\37\2\2\u00e5\u00e6\5\"\22"+
		"\2\u00e6-\3\2\2\2\u00e7\u00e9\7\24\2\2\u00e8\u00ea\5\66\34\2\u00e9\u00e8"+
		"\3\2\2\2\u00e9\u00ea\3\2\2\2\u00ea\u00eb\3\2\2\2\u00eb\u00ec\7,\2\2\u00ec"+
		"/\3\2\2\2\u00ed\u00ee\7+\2\2\u00ee\u00ef\5\62\32\2\u00ef\u00f0\7,\2\2"+
		"\u00f0\61\3\2\2\2\u00f1\u00fa\7\36\2\2\u00f2\u00f7\5\66\34\2\u00f3\u00f4"+
		"\7!\2\2\u00f4\u00f6\5\66\34\2\u00f5\u00f3\3\2\2\2\u00f6\u00f9\3\2\2\2"+
		"\u00f7\u00f5\3\2\2\2\u00f7\u00f8\3\2\2\2\u00f8\u00fb\3\2\2\2\u00f9\u00f7"+
		"\3\2\2\2\u00fa\u00f2\3\2\2\2\u00fa\u00fb\3\2\2\2\u00fb\u00fc\3\2\2\2\u00fc"+
		"\u00fd\7\37\2\2\u00fd\63\3\2\2\2\u00fe\u0103\7+\2\2\u00ff\u0100\7 \2\2"+
		"\u0100\u0102\7+\2\2\u0101\u00ff\3\2\2\2\u0102\u0105\3\2\2\2\u0103\u0101"+
		"\3\2\2\2\u0103\u0104\3\2\2\2\u0104\65\3\2\2\2\u0105\u0103\3\2\2\2\u0106"+
		"\u010b\58\35\2\u0107\u0108\7#\2\2\u0108\u010a\58\35\2\u0109\u0107\3\2"+
		"\2\2\u010a\u010d\3\2\2\2\u010b\u0109\3\2\2\2\u010b\u010c\3\2\2\2\u010c"+
		"\67\3\2\2\2\u010d\u010b\3\2\2\2\u010e\u0113\5:\36\2\u010f\u0110\7$\2\2"+
		"\u0110\u0112\5:\36\2\u0111\u010f\3\2\2\2\u0112\u0115\3\2\2\2\u0113\u0111"+
		"\3\2\2\2\u0113\u0114\3\2\2\2\u01149\3\2\2\2\u0115\u0113\3\2\2\2\u0116"+
		"\u011a\5> \2\u0117\u0119\5<\37\2\u0118\u0117\3\2\2\2\u0119\u011c\3\2\2"+
		"\2\u011a\u0118\3\2\2\2\u011a\u011b\3\2\2\2\u011b;\3\2\2\2\u011c\u011a"+
		"\3\2\2\2\u011d\u011e\t\2\2\2\u011e\u011f\5> \2\u011f=\3\2\2\2\u0120\u0124"+
		"\5B\"\2\u0121\u0123\5@!\2\u0122\u0121\3\2\2\2\u0123\u0126\3\2\2\2\u0124"+
		"\u0122\3\2\2\2\u0124\u0125\3\2\2\2\u0125?\3\2\2\2\u0126\u0124\3\2\2\2"+
		"\u0127\u0128\t\3\2\2\u0128\u0129\5B\"\2\u0129A\3\2\2\2\u012a\u012e\5F"+
		"$\2\u012b\u012d\5D#\2\u012c\u012b\3\2\2\2\u012d\u0130\3\2\2\2\u012e\u012c"+
		"\3\2\2\2\u012e\u012f\3\2\2\2\u012fC\3\2\2\2\u0130\u012e\3\2\2\2\u0131"+
		"\u0132\t\4\2\2\u0132\u0133\5F$\2\u0133E\3\2\2\2\u0134\u0138\5J&\2\u0135"+
		"\u0137\5H%\2\u0136\u0135\3\2\2\2\u0137\u013a\3\2\2\2\u0138\u0136\3\2\2"+
		"\2\u0138\u0139\3\2\2\2\u0139G\3\2\2\2\u013a\u0138\3\2\2\2\u013b\u013c"+
		"\t\5\2\2\u013c\u013d\5J&\2\u013dI\3\2\2\2\u013e\u013f\7\"\2\2\u013f\u0144"+
		"\5L\'\2\u0140\u0141\7\5\2\2\u0141\u0144\5L\'\2\u0142\u0144\5L\'\2\u0143"+
		"\u013e\3\2\2\2\u0143\u0140\3\2\2\2\u0143\u0142\3\2\2\2\u0144K\3\2\2\2"+
		"\u0145\u014a\5N(\2\u0146\u0147\7 \2\2\u0147\u0149\7+\2\2\u0148\u0146\3"+
		"\2\2\2\u0149\u014c\3\2\2\2\u014a\u0148\3\2\2\2\u014a\u014b\3\2\2\2\u014b"+
		"M\3\2\2\2\u014c\u014a\3\2\2\2\u014d\u014e\7\36\2\2\u014e\u014f\5\66\34"+
		"\2\u014f\u0150\7\37\2\2\u0150\u015c\3\2\2\2\u0151\u0153\7+\2\2\u0152\u0154"+
		"\5\62\32\2\u0153\u0152\3\2\2\2\u0153\u0154\3\2\2\2\u0154\u015c\3\2\2\2"+
		"\u0155\u015c\7\3\2\2\u0156\u0157\7\30\2\2\u0157\u015c\7+\2\2\u0158\u015c"+
		"\7\31\2\2\u0159\u015c\7\32\2\2\u015a\u015c\7\33\2\2\u015b\u014d\3\2\2"+
		"\2\u015b\u0151\3\2\2\2\u015b\u0155\3\2\2\2\u015b\u0156\3\2\2\2\u015b\u0158"+
		"\3\2\2\2\u015b\u0159\3\2\2\2\u015b\u015a\3\2\2\2\u015cO\3\2\2\2\35Xjt"+
		"y\u0086\u008c\u0093\u00a0\u00a3\u00ac\u00b8\u00d2\u00df\u00e9\u00f7\u00fa"+
		"\u0103\u010b\u0113\u011a\u0124\u012e\u0138\u0143\u014a\u0153\u015b";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}