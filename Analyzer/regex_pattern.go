package Analyzer

type commentRegexs struct {
	BlankLine         string
	CComment          string
	CSSComment        string
	HTMLComment       string
	BashComment       string
	ShellComment      string
	PythonComment     string
	AssemblyComment   string
	SQLComment        string
	LuaComment        string
	HaskellComment    string
	PowerShellComment string
	MatlabComment     string
	VBComment         string
	ClojureComment    string
	JuliaComment      string
	FortranComment    string
	ElixirComment     string
	RubyComment       string
	PascalComment     string
}

var CommentRegexs = commentRegexs{
	BlankLine:         `(?m)^\s*$`,
	CComment:          `(?:^|[^"'/])(?://.*|/\*[\s\S]*?\*/)`,
	CSSComment:        `(?:^|[^"'/])(?:/\*[\s\S]*?\*/)`,
	HTMLComment:       `(^|[^\\])`,
	BashComment:       `#.*|:[ \t]*'[\s\S]*?'`,
	ShellComment:      `(?:REM\s.*)|(?:::.*)`,
	PythonComment:     `#.*|'{3}[\s\S]*?'{3}|"{3}[\s\S]*?"{3}`,
	AssemblyComment:   `(?:^|[^"';])(?:;.*)`,
	SQLComment:        `--.*|/\*[\s\S]*?\*/`,
	LuaComment:        `--(?:\[\[[\s\S]*?\]\])?.*$`,
	HaskellComment:    `--.*|\{-[\s\S]*?-\}`,
	PowerShellComment: `#.*|<#[\s\S]*?#>`,
	MatlabComment:     `(?:^|[^"'/])(?:%.*|/\*[\s\S]*?\*/)`,
	VBComment:         `'.*|/\*[\s\S]*?\*/`,
	ClojureComment:    `;.*|\(\*[\s\S]*?\*\)`,
	JuliaComment:      `#.*|#=[\s\S]*?=#`,
	FortranComment:    `!.*|/\*[\s\S]*?\*/`,
	ElixirComment:     `#.*|@moduledoc\s*"""[\s\S]*?"""|@doc\s*"""[\s\S]*?"""`,
	RubyComment:       `#.*|\=begin[\s\S]*?\=end`,
	PascalComment:     `\{[\s\S]*?\}|\(\*[\s\S]*?\*\)|//.*`,
}
