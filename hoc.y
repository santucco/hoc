%{
package hoc

type _pair struct{
	inst	int
	nargs	int
}

%}
%union {
	sym			*symbol		// symbol pointer
	syms		symbols		// symbols table
	val			symval		// value 
	inst		int			// machine instruction position
	narg		int 		// number of arguments 
	formals 	*formal		// list of formal parameters 
	pair		_pair
	
}
%token	<val>	NUMBER STRING BOOL NULL
%token	<sym>	VAR PRINT PRINTF BLTIN UNDEF WHILE FOR IF ELSE LOCAL
%token	<sym>	FUNCTION PROCEDURE RETURN FUNC PROC READ COUNTOF
%type	<formals>	formals
%type	<inst>	expr stmt asgn prlist stmtlist localvar stmts
%type	<inst>	cond while for if begin end 
%type	<sym>	procname varname pair
%type	<narg>	arglist args
%type	<pair>	paramlist
%type 	<val> 	value array object elements members pairs
%right	'=' ADDEQ SUBEQ MULEQ DIVEQ MODEQ
%left ','
%left	OR
%left	AND
%left	GT GE LT LE EQ NE
%left	'+' '-'
%left	'*' '/' '%'
%left	UNARYMINUS NOT INC DEC
%right	'^'
%%
list:	  // nothing 
	| list term 
	| list defn term
	| list object term
	{
		this := yylex.(*Hoc)
		for  k, v := range $2.(symbols) {
			tracer.Trace(debLex, "list object: installing %v = %v", k, v)
			if sym := lookup(this.symbols, v.name); sym != nil {
				sym = v
			} else {
				this.symbols = install(this.symbols, v)
			}
		}
	
	}
	| list asgn term
	{ 
		this := yylex.(*Hoc)
		this.code((*Hoc).xpop); 
		this.code(nil); 
		return 1; 
	}
	| list stmt term
	{ 
		this := yylex.(*Hoc)
		this.code(nil); 
		return 1; 
	} 	
	| list expr term
	{ 
		this := yylex.(*Hoc)
		this.code((*Hoc).printtop); 
		this.code(nil); 
		return 1; 
	}
	;

term:
	';'
	| '\n'
	;

localdef: 
	LOCAL 
	{
		tracer.Trace(debLex, "localdef::LOCAL")

		this := yylex.(*Hoc)
		if this.locsyms == nil {
			yylex.Error("using local definition outside of function of procedure")
		}
	}
	;

localvar: 
	localdef VAR
	{
		tracer.Trace(debLex, "localvar:: localdef varname")
		this := yylex.(*Hoc)
		sym := lookup(this.locsyms, $2.name)
		if sym == nil {
			this.locsyms = install(this.locsyms, $2)
		}	
	}
	| localdef VAR '=' expr
	{
		tracer.Trace(debLex, "localvar::localdev varname '=' expr" )
		this := yylex.(*Hoc)
		sym := lookup(this.locsyms, $2.name)
		if sym == nil {
			this.locsyms = install(this.locsyms, $2)
			sym = $2
		}
		this.code((*Hoc).vpush)
		this.code(sym)
		this.code((*Hoc).assign)		
	}
	;

varname: 
	VAR
	{
		tracer.Trace(debLex, "VAR - %s", $1.name)
		this := yylex.(*Hoc)
		sym := this.lookup($1.name)
		if sym == nil {
			this.symbols = install(this.symbols, $1)
			sym = $1
		}
		$$ = sym
	}
	| varname '.' VAR
	{		
		tracer.Trace(debLex, "%s '.' %s ", $1.name, $3.name )
		if $1.val == nil {
			$1.val = newSymbols()
			$1.t = VAR
		}
		syms, ok := $1.val.(symbols)
		if !ok {
			yylex.Error( $1.name + " is not compound type")
		}
		sym := lookup(syms, $3.name)
		if sym == nil {
			syms = install(syms, $3)
			sym = $3
		}
		$$ = sym
	}
	;	
	
asgn:
	varname '=' expr 
	{ 
		tracer.Trace(debLex, "varname '=' expr" )
		this := yylex.(*Hoc)
		this.code((*Hoc).vpush)
		this.code($1)
		this.code((*Hoc).assign)
		$$=$3 
	}	  
	| varname ADDEQ expr 
	{ 
		tracer.Trace(debLex, "%s ADDEQ expr ", $1.name)
		this := yylex.(*Hoc)
		this.code((*Hoc).vpush)
		this.code($1)
		this.code((*Hoc).addeq) 
		$$=$3 
	}
	| varname SUBEQ expr 
	{ 
		tracer.Trace(debLex, "%s SUBEQ expr ", $1.name)
		this := yylex.(*Hoc)	
		this.code((*Hoc).vpush)
		this.code($1)
		this.code((*Hoc).subeq) 
		$$=$3 
	}
	| varname MULEQ expr 
	{ 
		tracer.Trace(debLex, "%s MULEQ expr ", $1.name)
		this := yylex.(*Hoc)
		this.code((*Hoc).vpush)
		this.code($1)
		this.code((*Hoc).muleq) 
		$$=$3 
	}
	| varname DIVEQ expr 
	{ 
		tracer.Trace(debLex, "%s DIVEQ expr ", $1.name)
		this := yylex.(*Hoc)
		this.code((*Hoc).vpush)
		this.code($1)
		this.code((*Hoc).diveq)
		$$=$3
	}
	| varname MODEQ expr 
	{ 
		tracer.Trace(debLex, "%s MODEQ expr ", $1.name)
		this := yylex.(*Hoc)
		this.code((*Hoc).vpush)
		this.code($1)
		this.code((*Hoc).modeq)
		$$=$3
	}
	;

stmt:	  
	localvar
	{
		tracer.Trace(debLex, "stmt::localvar")
		$$=$1
	}
	| expr 
	{ 
		tracer.Trace(debLex, "stmt::expr")
		this := yylex.(*Hoc)
		this.code((*Hoc).xpop)
	}
	| RETURN 
	{ 	
		tracer.Trace(debLex, "stmt::RETURN")	
		this := yylex.(*Hoc)
		this.defnonly("return") 
		this.code((*Hoc).procret) 
	}
	| RETURN expr 
	{ 
		tracer.Trace(debLex, "stmt::RETURN expr")	
		this := yylex.(*Hoc)
		this.defnonly("return") 
		$$=$2 
		this.code((*Hoc).funcret)
	}
	| PROCEDURE begin '(' args ')' 
	{ 
		tracer.Trace(debLex, "stmt::PROCEDURE begin")
		$$ = $2 
		this := yylex.(*Hoc)	
		this.code((*Hoc).call)
		this.code($1)
		this.code($4)
	}
	| PRINT
	{
		tracer.Trace(debLex, "stmt::PRINT")	
		this := yylex.(*Hoc)	
		this.code((*Hoc).printall) 
	}
	| PRINT prlist	
	{ 
		tracer.Trace(debLex, "stmt::PRINT prlist")
		$$ = $2
	}
	| PRINTF STRING paramlist 
	{ 
		tracer.Trace(debLex, "stmt::PRINTF STRING paramlist")
		this := yylex.(*Hoc)
		this.code((*Hoc).cpush)
		this.code($3.nargs) 
		this.code((*Hoc).cpush)
		this.code($2.(string)) 
		this.code((*Hoc).prfstr)
		$$ = $3.inst
	}
	| PRINTF varname paramlist 
	{ 
		tracer.Trace(debLex, "stmt::PRINTF STRING paramlist")
		this := yylex.(*Hoc)
		this.code((*Hoc).cpush)
		this.code($3.nargs) 
		this.code((*Hoc).vpush)
		this.code($2) 
		this.code((*Hoc).prfstr)
		$$ = $3.inst
	}	| while '(' cond ')' stmts end 
	{
		tracer.Trace(debLex, "stmt::while")
		tracer.Trace(debLex, "while '(' cond ')' stmts end ")
		this := yylex.(*Hoc)	
		this.prog[$1 + 1] = $5	// body of loop 
		this.prog[$1 + 2] = $6 	// end, if cond fails 
	}
	| for '(' cond ';' cond ';' cond ')' stmts end 
	{
		tracer.Trace(debLex, "stmt::for")
		this := yylex.(*Hoc)	
		this.prog[$1 + 1] = $5	// condition 
		this.prog[$1 + 2] = $7	// post loop 
		this.prog[$1 + 3] = $9	// body of loop 
		this.prog[$1 + 4] = $10	// end, if cond fails 
	}
	| if '(' cond ')' stmts end ELSE stmts end 
	{	// if with else 
		tracer.Trace(debLex, "stmt::if else") 
		tracer.Trace(debCode, "$1 %v $5 %v $8 %v $9 %v ", $1, $5, $8,$9)
		this := yylex.(*Hoc)	
		this.prog[$1 + 1] = $5	// thenpart 
		this.prog[$1 + 2] = $8	//elsepart 
		this.prog[$1 + 3] = $9 	// end, if cond fails 
	}
	| if '(' cond ')' stmts end 
	{	// else-less if 
		tracer.Trace(debLex, "stmt::if") 
		this := yylex.(*Hoc)	
		this.prog[$1 + 1] = $5	// thenpart  
		this.prog[$1 + 3] = $6 	// end, if cond fails 
	}
	;

stmts: 
	stmt
	{
		tracer.Trace(debLex, "stmts::stmt")
	}
	| '{' stmtlist '}'	
	{ 
		tracer.Trace(debLex, "stmt::stmtlist")
		$$ = $2
	}
	;

cond:	   
	expr 	
	{ 
		tracer.Trace(debLex, "cond")
		this := yylex.(*Hoc)	
		this.code(nil)
	}
	;

while:	  
	WHILE	
	{ 
		tracer.Trace(debLex, "WHILE")
		this := yylex.(*Hoc)
		$$ = this.code((*Hoc).whilecode)
		this.code(nil)
		this.code(nil) 
	}
	;

for:	  
	FOR	
	{ 
		tracer.Trace(debLex, "FOR")
		this := yylex.(*Hoc)
		$$ = this.code((*Hoc).forcode)
		this.code(nil)
		this.code(nil)
		this.code(nil) 
		this.code(nil)
	}
	;

if:	  
	IF	
	{ 
		tracer.Trace(debLex, "IF")
		this := yylex.(*Hoc)
		$$ = this.code((*Hoc).ifcode) 
		this.code(nil)
		this.code(nil)
		this.code(nil) 
	}
	;

begin:	  // nothing 
	{ 
		tracer.Trace(debLex, "begin")
		this := yylex.(*Hoc)	
		$$ = this.progp
	}
	;

end:	  // nothing 
	{ 
		tracer.Trace(debLex, "end")
		this := yylex.(*Hoc)	
		this.code(nil) 
		$$ = this.progp
	}
	;

stmtlist: // nothing 
	{ 
		tracer.Trace(debLex, "stmtlist")
		this := yylex.(*Hoc)	
		$$ = this.progp 
	}
	| stmtlist term
	| stmtlist stmts
	;

expr:	  
	NUMBER 
	{ 
		tracer.Trace(debLex, "expr::NUMBER")
		this := yylex.(*Hoc)	
		$$ = this.code((*Hoc).cpush)
		switch $1.(type) {
		case int:
			this.code($1.(int))
		case float64:
			this.code($1.(float64))
		default:
			tracer.Trace(debLex, "type of %v : %T", $1, $1)
			yylex.Error("wrong type for number constant") 
		}
	}
	| BOOL
	{
		tracer.Trace(debLex, "expr::NUMBER")
		this := yylex.(*Hoc)	
		$$ = this.code((*Hoc).cpush)
		this.code($1.(bool))	
	}
	| STRING
	{ 
		tracer.Trace(debLex, "expr::STRING")
		this := yylex.(*Hoc)	
		$$ = this.code((*Hoc).cpush)
		this.code($1.(string)) 
	}	
	| varname	 
	{ 
		tracer.Trace(debLex, "expr::VARNAME %s", $1.name)
		this := yylex.(*Hoc)	
		$$ = this.code((*Hoc).vpush)
		this.code($1)
		this.code((*Hoc).eval) 
	}
	| varname '[' expr ']'
	{
		tracer.Trace(debLex, "expr::VARNAME %s '[' expr ']'", $1.name)
		this := yylex.(*Hoc)	
		this.code((*Hoc).vpush)
		this.code($1)
		this.code((*Hoc).elemeval) 
		$$ = $3
	}
	| COUNTOF varname
	{
		tracer.Trace(debLex, "expr::COUNTOF %s", $1.name)
		this := yylex.(*Hoc)	
		$$ = this.code((*Hoc).vpush)
		this.code($2)
		this.code((*Hoc).countof) 
	}
	| asgn
	| FUNCTION begin '(' args ')'
	{ 
		tracer.Trace(debLex, "expr::FUNCTION")
		$$ = $2; 
		this := yylex.(*Hoc)	
		this.code((*Hoc).call)
		this.code($1)
		this.code($4) 
	}
/*	| READ '(' varname ')' 
	{ 
		this := yylex.(*Hoc)	
		$$ = this.code((*Hoc).varread)
		this.code($3) 
	}
*/
	| BLTIN '(' expr ')' 
	{ 
		tracer.Trace(debLex, "expr::BLTIN")
		$$=$3 
		this := yylex.(*Hoc)	
		this.code((*Hoc).bltin)
		this.code($1.val.(bltin))
	}
	| '(' expr ')'	
	{ 		tracer.Trace(debLex, "expr::'(' expr ')'")
		$$ = $2; 
	}
	| expr '+' expr	
	{ 
		tracer.Trace(debLex, "expr::expr '+' expr")
		this := yylex.(*Hoc)	
		this.code((*Hoc).add) 
	}
	| expr '-' expr	
	{ 
		tracer.Trace(debLex, "expr::expr '-' expr")
		this := yylex.(*Hoc)	
		this.code((*Hoc).sub) 
	}
	| expr '*' expr	
	{ 
		tracer.Trace(debLex, "expr::expr '*' expr")
		this := yylex.(*Hoc)	
		this.code((*Hoc).mul) 
	}
	| expr '/' expr	
	{ 
		tracer.Trace(debLex, "expr::expr '/' expr")
		this := yylex.(*Hoc)	
		this.code((*Hoc).div) 
	}
	| expr '%' expr	
	{ 
		tracer.Trace(debLex, "expr::expr '%' expr")
		this := yylex.(*Hoc)	
		this.code((*Hoc).mod) 
	}
	| expr '^' expr	
	{ 
		tracer.Trace(debLex, "expr::expr '^' expr")
		this := yylex.(*Hoc)	
		this.code ((*Hoc).power) 
	}
	| '-' expr   %prec UNARYMINUS   
	{ 
		tracer.Trace(debLex, "expr::'-' expr")
		$$=$2; 
		this := yylex.(*Hoc)	
		this.code((*Hoc).negate) 
	}
	| expr GT expr	
	{ 
		tracer.Trace(debLex, "expr::expr 'GT expr")
		this := yylex.(*Hoc)	
		this.code((*Hoc).gt)
	}
	| expr GE expr	
	{ 
		tracer.Trace(debLex, "expr::expr GE expr")
		this := yylex.(*Hoc)	
		this.code((*Hoc).ge) 
	}
	| expr LT expr	
	{ 		
		tracer.Trace(debLex, "expr::expr LT expr")
		this := yylex.(*Hoc)	
		this.code((*Hoc).lt)
	}
	| expr LE expr	
	{ 
		tracer.Trace(debLex, "expr::expr LE expr")
		this := yylex.(*Hoc)	
		this.code((*Hoc).le)
	}
	| expr EQ expr	
	{ 
		tracer.Trace(debLex, "expr::expr EQ expr")
		this := yylex.(*Hoc)	
		this.code((*Hoc).eq) 
	}
	| expr NE expr	
	{ 
		tracer.Trace(debLex, "expr::expr NE expr")
		this := yylex.(*Hoc)	
		this.code((*Hoc).ne) 
	}
	| expr AND expr	
	{ 
		tracer.Trace(debLex, "expr::expr AND expr")
		this := yylex.(*Hoc)	
		this.code((*Hoc).and)
	}
	| expr OR expr	
	{ 
		tracer.Trace(debLex, "expr::expr OR expr")
		this := yylex.(*Hoc)	
		this.code((*Hoc).or)
	}
	| NOT expr	
	{ 
		tracer.Trace(debLex, "expr::NOT expr")
		$$ = $2; 
		this := yylex.(*Hoc)	
		this.code((*Hoc).not) 
	}
	| INC varname	
	{ 
		tracer.Trace(debLex, "expr::INC expr")
		this := yylex.(*Hoc)
		$$ = this.code((*Hoc).vpush)
		this.code($2)	
		$$ = this.code((*Hoc).preinc)
	}
	| DEC varname	
	{ 
		tracer.Trace(debLex, "expr::DEC expr")
		this := yylex.(*Hoc)	
		$$ = this.code((*Hoc).vpush)
		this.code($2)		
		this.code((*Hoc).predec)
	}
	| varname INC	
	{ 
		tracer.Trace(debLex, "expr::expr INC")
		this := yylex.(*Hoc)	
		$$ = this.code((*Hoc).vpush)
		this.code($1)
		this.code((*Hoc).postinc)
	}
	| varname DEC	
	{ 
		tracer.Trace(debLex, "expr::expr DEC")
		this := yylex.(*Hoc)	
		$$ = this.code((*Hoc).vpush)
		this.code($1)
		this.code((*Hoc).postdec)
	}
	;

prlist:	 
	expr			
	{
		tracer.Trace(debLex, "prlist::expr")
		this := yylex.(*Hoc)	
		this.code((*Hoc).prexpr) 
	}
	| prlist ',' expr	
	{ 
		tracer.Trace(debLex, "prlist::prlist ',' expr")
		this := yylex.(*Hoc)	
		this.code((*Hoc).prexpr)
	}
	;

defn:	  
	FUNC procname 
	{ 
		tracer.Trace(debLex, "defn::FUNC %s", $2.name)
		this := yylex.(*Hoc)	
		sym := lookup(this.symbols, $2.name)
		if sym != nil {
			yylex.Error("symbol " + $2.name + " already defined")
		}
		$2.t = FUNCTION
		this.symbols = install(this.symbols, $2)
		this.locsyms = newSymbols()
	}
	'(' formals ')' 
	{ 
		tracer.Trace(debLex, "defn::FUNC %s formals", $2.name)	
	}
	stmts 
	{
		tracer.Trace(debLex, "defn::FUNC %s stmt", $2.name)
		this := yylex.(*Hoc)	
		this.code((*Hoc).procret); 
		this.define($2, $5) 
		this.locsyms = nil
	}
	| PROC procname 
	{ 
		tracer.Trace(debLex, "defn::PROC %s", $2.name)
		this := yylex.(*Hoc)	
		sym := lookup(this.symbols, $2.name)
		if sym != nil {
			yylex.Error("symbol " + $2.name + " already defined")
		}
		$2.t = PROCEDURE
		this.symbols = install(this.symbols, $2)
		this.locsyms = newSymbols()
	}
	'(' formals ')' 
	{
		tracer.Trace(debLex, "defn::PROC %s formals", $2.name)	
	}
	stmts 
	{ 
		tracer.Trace(debLex, "defn::PROC %s formals", $2.name)
		this := yylex.(*Hoc)	
		this.code((*Hoc).procret) 
		this.define($2, $5) 
		this.locsyms = nil
	}
	;

formals:	{ $$ = nil; }
	| varname			
	{ 
		tracer.Trace(debLex, "formals::%s", $1.name)
		tracer.Trace(debLex, "formal parameter %s", $1.name)
		$$ = formallist($1, nil) 
	}
	| varname ',' formals	
	{ 
		tracer.Trace(debLex, "formals::%s ',' formals", $1.name)
		$$ = formallist($1, $3) 
	}
	;

procname: VAR
	| FUNCTION
	| PROCEDURE
	;

args:	// nothing 
	{ 
		tracer.Trace(debLex, "args:: no arguments")
		$$ = 0 
	}
	|
	arglist
	{ 
		tracer.Trace(debLex, "args::arglist")	}	
	;

arglist: expr			
	{ 
		tracer.Trace(debLex, "arglist:: one argument")
		$$ = 1 
	}
	| arglist ',' expr	
	{ 
		tracer.Trace(debLex, "arglist:: yet another argument")
		$$ = $1 + 1 
	}
	;

paramlist:  // nothing
	{ 
		tracer.Trace(debLex, "paramlist")
		$$.nargs = 0
	}
	| ',' expr paramlist	
	{ 
		tracer.Trace(debLex, "paramlist:: expr")
		$$.nargs = $3.nargs + 1
		$$.inst = $2
	}
	;

delim: //nothing
	| '\n'
	;

object:	
	'{' delim members  delim '}' 
	{
		tracer.Trace(debLex, "object::members %v", $3)
		$$ = $3
	}
	;

members: // nothing
	{
		$$ = newSymbols()
	}
	| pairs
	{}
	;

pairs: 
	pair
	{
		syms := newSymbols()
		syms = install(syms, $1)
		tracer.Trace(debLex, "members::pair %v", syms)
		$$ = syms
	}
	| pairs ',' delim pair
	{
		tracer.Trace(debLex, "members::pair ',' members")
		syms := $1.(symbols)

		tracer.Trace(debLex, "syms: %v", syms)
	
		if s := lookup(syms, $4.name); s != nil {
				yylex.Error("dublicated value name: " + $4.name)
		}
		tracer.Trace(debLex, "pair %v", $4)
		syms = install(syms, $4)
		tracer.Trace(debLex, "members::pair ',' members %v", syms)
		$$ = syms	
	}
	;

pair:
	STRING ':' value
	{
		tracer.Trace(debLex, "pair")
		sym := newSymbol($1.(string), VAR, $3)
		tracer.Trace(debLex, "new symbol %v", sym)
		$$ = sym
	}
	;

array:
	'[' delim elements delim ']'
	{
		tracer.Trace(debLex, "array::elements %v", $3)
		$$ = $3
	}
	;

elements: //nothing
	{
		$$ = make([]symval, 0)
	}
	| value
	{
		tracer.Trace(debLex, "elements::value")
		e := make([]symval, 1)
		e[0] = $1
		$$ = e		
	}
	| value ',' delim elements
	{
		tracer.Trace(debLex, "elements::value ',' elements")
		e := make([]symval, len($4.([]symval)) + 1)
		e[0] = $1
		copy(e[1:], $4.([]symval))
		$$ = e
	}	
	;

value:
 	STRING
	| NUMBER
	| BOOL
	| NULL
	| object
	| array
	;

/*string:	'"' '"'
	| '"' chars '"'
	;
chars:	char
	| char chars
	;
char:	any-Unicode-character-except-"-or-\-or-control-character
	| '\' '"'
	| '\' '\'
	| '\' '/'
	| '\' 'b'
	| '\' 'f'
	| '\' 'n'
	| '\' 'r'
	| '\' 't'
	| '\' 'u' four-hex-digits 
	;
number:	int
	| int frac
	| int exp
	| int frac exp
	;
int:	digit
	| digit1-9 digits 
	| '-' digit
	| '-' digit1-9 digits 
	;
frac;	'.' digits
	;
exp:	e digits
	;
digits:	digit
	| digit digits
	;
e:	'e'
	| 'e' '+'
	| 'e' '-'
	| 'E'
	| 'E' '+'
	| 'E' '-'
	;
*/
%%
	// end of grammar

