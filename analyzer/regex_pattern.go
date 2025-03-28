package analyzer

import "regexp"

var (
	// Matches blank lines (empty or only whitespace)
	BlankLineRegex = regexp.MustCompile(`(?m)^\s*$`)

	// Matches C-style single-line (//) and multi-line (/* ... */) comments
	CCommentRegex = regexp.MustCompile(`(?:^|[^"'/])(?://.*|/\*[\s\S]*?\*/)`)

	// Matches CSS-style multi-line (/* ... */) comments
	CSSCommentRegex = regexp.MustCompile(`(?:^|[^"'/])(?:/\*[\s\S]*?\*/)`)

	// Matches HTML-style comments (<!-- ... -->)
	HTMLCommentRegex = regexp.MustCompile(`(^|[^\\])<!--[\s\S]*?-->`)

	// Matches Bash-style comments (#...) and alternate form using ':'
	BashCommentRegex = regexp.MustCompile(`#.*|:[ \t]*'[\s\S]*?'`)

	// Matches Windows Batch (Shell) comments using REM and ::
	ShellCommentRegex = regexp.MustCompile(`(?:REM\s.*)|(?:::.*)`)

	// Matches Python-style comments (#...) and multi-line ('''...''' or """...""")
	PythonCommentRegex = regexp.MustCompile(`#.*|'{3}[\s\S]*?'{3}|"{3}[\s\S]*?"{3}`)

	// Matches Assembly-style comments (;...)
	AssemblyCommentRegex = regexp.MustCompile(`(?:^|[^"';])(?:;.*)`)

	// Matches SQL-style comments (--...) and multi-line (/* ... */)
	SQLCommentRegex = regexp.MustCompile(`--.*|/\*[\s\S]*?\*/`)

	// Matches Lua-style comments (--...) and block comments (--[[ ... ]])
	LuaCommentRegex = regexp.MustCompile(`--(?:\[\[[\s\S]*?\]\])?.*$`)

	// Matches Haskell-style comments (--...) and multi-line ({- ... -})
	HaskellCommentRegex = regexp.MustCompile(`--.*|\{-[\s\S]*?-\}`)

	// Matches PowerShell-style comments (#...) and multi-line (<# ... #>)
	PowerShellCommentRegex = regexp.MustCompile(`#.*|<#[\s\S]*?#>`)

	// Matches VB-style comments ('...) and block comments (/* ... */)
	VBCommentRegex = regexp.MustCompile(`'.*|/\*[\s\S]*?\*/`)

	// Matches Clojure-style comments (;...) and block comments (* ... *)
	ClojureCommentRegex = regexp.MustCompile(`;.*|\(\*[\s\S]*?\*\)`)

	// Matches Julia-style comments (#...) and block comments (#= ... =#)
	JuliaCommentRegex = regexp.MustCompile(`#.*|#=[\s\S]*?=#`)

	// Matches Fortran-style comments (!...) and block comments (/* ... */)
	FortranCommentRegex = regexp.MustCompile(`!.*|/\*[\s\S]*?\*/`)

	// Matches Elixir-style comments (#...) and documentation comments using @moduledoc and @doc
	ElixirCommentRegx = regexp.MustCompile(`#.*|@moduledoc\s*"""[\s\S]*?"""|@doc\s*"""[\s\S]*?"""`)

	// Matches Ruby-style comments (#...) and block comments (=begin...=end)
	RubyCommentRegex = regexp.MustCompile(`#.*|\=begin[\s\S]*?\=end`)

	// Matches Pascal-style comments ({...}), (*...*), and single-line (//...)
	PascalCommentRegex = regexp.MustCompile(`\{[\s\S]*?\}|\(\*[\s\S]*?\*\)|//.*`)
)
