
use std::vec::Vec;

use crate::scanner::*;

%%{
    machine flux;

    alphtype u8;

    include WChar "unicode.rl";

    action advance_line {
        // We do this for every newline we find.
        // This allows us to return correct line/column for each token
        // back to the caller.
        *cur_line += 1;
        *last_newline = fpc + 1;
    }

    action advance_line_between_tokens {
        // We do this for each newline we find in the whitespace between tokens,
        // so we can record the location of the first byte of a token.
        last_newline_before_token = *last_newline;
        cur_line_token_start = *cur_line;
    }

    newline = '\n' @advance_line;
    any_count_line = any | newline;

    identifier = ( ualpha | "_" ) ( ualnum | "_" )*;

    decimal_lit = (digit - "0") digit*;
    int_lit = "0" | decimal_lit;

    float_lit = (digit+ "." digit*) | ("." digit+);

    duration_unit = "y" | "mo" | "w" | "d" | "h" | "m" | "s" | "ms" | "us" | "µs" | "ns";
    duration_lit = ( int_lit duration_unit )+;

    date = digit{4} "-" digit{2} "-" digit{2};
    time_offset = "Z" | (("+" | "-") digit{2} ":" digit{2});
    time = digit{2} ":" digit{2} ":" digit{2} ( "." digit* )? time_offset?;
    date_time_lit = date ( "T" time )?;

    escaped_char = "\\" ( "n" | "r" | "t" | "\\" | '"' | "${" );
    unicode_value = (any_count_line - [\\$]) | escaped_char;
    byte_value = "\\x" xdigit{2};
    dollar_value = "$" ( any_count_line - "{" );
    string_lit_char = ( unicode_value | byte_value | dollar_value );
    string_lit = '"' string_lit_char* "$"? :> '"';

    regex_escaped_char = "\\" ( "/" | "\\");
    regex_unicode_value = (any_count_line - "/") | regex_escaped_char;
    regex_lit = "/" ( regex_unicode_value | byte_value )+ "/";

    # The newline is optional so that a comment at the end of a file is considered valid.
    single_line_comment = "//" [^\n]* newline?;

    # Whitespace is standard ws and control codes->
    # (Note that newlines are handled separately; see notes above)
    whitespace = (space - '\n')+;

    # The regex literal is not compatible with division so we need two machines->
    # One machine contains the full grammar and is the main one, the other is used to scan when we are
    # in the middle of an expression and we are potentially expecting a division operator.
    main_with_regex := |*
        # If we see a regex literal, we accept that and do not go to the other scanner.
        regex_lit => { tok = TOK_REGEX; fbreak; };

        # We have to specify whitespace here so that leading whitespace doesn't cause a state transition.
        whitespace;

        newline => advance_line_between_tokens;

        # Any other character we transfer to the main state machine that defines the entire language.
        any => { fhold; fgoto main; };
    *|;

    # This machine does not contain the regex literal.
    main := |*
        single_line_comment => { tok = TOK_COMMENT; fbreak; };

        "and" => { tok = TOK_AND; fbreak; };
        "or" => { tok = TOK_OR; fbreak; };
        "not" => { tok = TOK_NOT; fbreak; };
        "empty" => { tok = TOK_EMPTY; fbreak; };
        "in" => { tok = TOK_IN; fbreak; };
        "import" => { tok = TOK_IMPORT; fbreak; };
        "package" => { tok = TOK_PACKAGE; fbreak; };
        "return" => { tok = TOK_RETURN; fbreak; };
        "option" => { tok = TOK_OPTION; fbreak; };
        "builtin" => { tok = TOK_BUILTIN; fbreak; };
        "test" => { tok = TOK_TEST; fbreak; };
        "if" => { tok = TOK_IF; fbreak; };
        "then" => { tok = TOK_THEN; fbreak; };
        "else" => { tok = TOK_ELSE; fbreak; };
        "exists" => { tok = TOK_EXISTS; fbreak; };

        identifier => { tok = TOK_IDENT; fbreak; };
        int_lit => { tok = TOK_INT; fbreak; };
        float_lit => { tok = TOK_FLOAT; fbreak; };
        duration_lit => { tok = TOK_DURATION; fbreak; };
        date_time_lit => { tok = TOK_TIME; fbreak; };
        string_lit => { tok = TOK_STRING; fbreak; };

        "+" => { tok = TOK_ADD; fbreak; };
        "-" => { tok = TOK_SUB; fbreak; };
        "*" => { tok = TOK_MUL; fbreak; };
        "/" => { tok = TOK_DIV; fbreak; };
        "%" => { tok = TOK_MOD; fbreak; };
        "^" => { tok = TOK_POW; fbreak; };
        "==" => { tok = TOK_EQ; fbreak; };
        "<" => { tok = TOK_LT; fbreak; };
        ">" => { tok = TOK_GT; fbreak; };
        "<=" => { tok = TOK_LTE; fbreak; };
        ">=" => { tok = TOK_GTE; fbreak; };
        "!=" => { tok = TOK_NEQ; fbreak; };
        "=~" => { tok = TOK_REGEXEQ; fbreak; };
        "!~" => { tok = TOK_REGEXNEQ; fbreak; };
        "=" => { tok = TOK_ASSIGN; fbreak; };
        "=>" => { tok = TOK_ARROW; fbreak; };
        "<-" => { tok = TOK_PIPE_RECEIVE; fbreak; };
        "(" => { tok = TOK_LPAREN; fbreak; };
        ")" => { tok = TOK_RPAREN; fbreak; };
        "[" => { tok = TOK_LBRACK; fbreak; };
        "]" => { tok = TOK_RBRACK; fbreak; };
        "{" => { tok = TOK_LBRACE; fbreak; };
        "}" => { tok = TOK_RBRACE; fbreak; };
        ":" => { tok = TOK_COLON; fbreak; };
        "|>" => { tok = TOK_PIPE_FORWARD; fbreak; };
        "," => { tok = TOK_COMMA; fbreak; };
        "." => { tok = TOK_DOT; fbreak; };
        '"' => { tok = TOK_QUOTE; fbreak; };
        '?' => { tok = TOK_QUESTION_MARK; fbreak; };

        whitespace;

        newline => advance_line_between_tokens;
    *|;

    # This is the scanner used when parsing a string expression.
    string_expr := |*
        "${" => { tok = TOK_STRINGEXPR; fbreak; };
        '"' => { tok = TOK_QUOTE; fbreak; };
        (string_lit_char - "\"")+ => { tok = TOK_TEXT; fbreak; };
    *|;
}%%

%% write data nofinal;

pub fn scan(
    data: &[u8],
    mode: i32,
    pp: &mut i32,
    _data: i32,
    pe: i32,
    eof: i32,
    last_newline: &mut i32,
    cur_line: &mut i32,
    token: &mut u32,
    token_start: &mut i32,
    token_start_line: &mut i32,
    token_start_col: &mut i32,
    token_end: &mut i32,
    token_end_line: &mut i32,
    token_end_col: &mut i32 ) -> u32
{
    let mut cs = flux_start;
    match mode {
        0 => { cs = flux_en_main },
        1 => { cs = flux_en_main_with_regex },
        2 => { cs = flux_en_string_expr },
        _ => {},
    }
    let mut p: i32 = *pp;

    let mut act: i32 = 0;
    let mut ts: i32 = 0;
    let mut te: i32 = 0;
    let mut tok: TOK = TOK_ILLEGAL;

    let mut last_newline_before_token: i32 = *last_newline;
    let mut cur_line_token_start: i32 = *cur_line;

    // alskdfj
    %% write init nocs;
    %% write exec;

    // Update output args.
    *token = tok;

    *token_start = ts - _data;
    *token_start_line = cur_line_token_start;
    *token_start_col = ts - last_newline_before_token + 1;

    *token_end = te - _data;

    if (*last_newline > te) {
        // te (the token end pointer) will only be less than last_newline
        // (pointer to the last newline the scanner saw) if we are trying
        // to find a multi-line token (either string or regex literal)
        // but don't find the closing `/` or `"`.
        // In that case we need to reset last_newline and cur_line.
        *cur_line = cur_line_token_start;
        *last_newline = last_newline_before_token;
    }

    *token_end_line = *cur_line;
    *token_end_col = te - *last_newline + 1;

    *pp = p;
    if cs == flux_error {
        return 1
    } else {
        return 0;
    }
}
