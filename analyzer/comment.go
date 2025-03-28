package analyzer

import (
	"regexp"
	"strings"
)

type ExtractComment func(source string) []string

// CommentSyntax represents symbols used for single-line and multi-line comments.
type CommentSyntax struct {
	SingleLine     string // Symbol for single-line comments (e.g., "//")
	MultiLineOpen  string // Opening symbol for multi-line comments (e.g., "/*")
	MultiLineClose string // Closing symbol for multi-line comments (e.g., "*/")
	ExtractComment ExtractComment
}

// Mapping of programming languages to their respective comment syntax.
var languageToCommentSyntax = map[Language]CommentSyntax{
	Go: {
		SingleLine:     "//",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractCComments,
	},

	C: {
		SingleLine:     "//",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractCComments,
	},

	CPlusPlus: {
		SingleLine:     "//",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractCComments,
	},
	CSharp: {
		SingleLine:     "//",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractCComments,
	},
	Rust: {
		SingleLine:     "//",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractCComments,
	},
	JavaScript: {
		SingleLine:     "//",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractCComments,
	},
	TypeScript: {
		SingleLine:     "//",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractCComments,
	},
	Java: {
		SingleLine:     "//",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractCComments,
	},
	Kotlin: {
		SingleLine:     "//",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractCComments,
	},
	Swift: {
		SingleLine:     "//",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractCComments,
	},
	PHP: {
		SingleLine:     "//",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractCComments,
	},
	Dart: {
		SingleLine:     "//",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractCComments,
	},
	Scala: {
		SingleLine:     "//",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractCComments,
	},
	ObjectiveC: {
		SingleLine:     "//",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractCComments,
	},
	Python: {
		SingleLine:     "#",
		MultiLineOpen:  "\"\"\"",
		MultiLineClose: "\"\"\"",
		ExtractComment: ExtractPythonComments,
	},
	R: {
		SingleLine:     "#",
		MultiLineOpen:  "\"\"\"",
		MultiLineClose: "\"\"\"",
		ExtractComment: ExtractPythonComments,
	},
	Ruby: {
		SingleLine:     "#",
		MultiLineOpen:  "=begin",
		MultiLineClose: "=end",
		ExtractComment: ExtractRubyComments,
	},
	Perl: {
		SingleLine:     "#",
		MultiLineOpen:  "=begin",
		MultiLineClose: "=end",
		ExtractComment: ExtractRubyComments,
	},
	Pascal: {
		SingleLine:     "//",
		MultiLineOpen:  "(*",
		MultiLineClose: "*)",
		ExtractComment: ExtractPascalComments,
	},
	FSharp: {
		SingleLine:     "//",
		MultiLineOpen:  "(*",
		MultiLineClose: "*)",
		ExtractComment: ExtractPascalComments,
	},
	Elixir: {
		SingleLine:     "#",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractElixirComments,
	},
	HTML: {
		SingleLine:     "<!--",
		MultiLineOpen:  "<!--",
		MultiLineClose: "-->",
		ExtractComment: ExtractHTMLComments,
	},
	CSS: {
		SingleLine:     "",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractCSSComments,
	},
	SQL: {
		SingleLine:     "--",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractSQLComments,
	},

	Lua: {
		SingleLine:     "--",
		MultiLineOpen:  "--[[",
		MultiLineClose: "--]]",
		ExtractComment: ExtractLuaComments,
	},

	Haskell: {
		SingleLine:     "--",
		MultiLineOpen:  "{-",
		MultiLineClose: "-}",
		ExtractComment: ExtractHaskellComments,
	},
	Assembly: {
		SingleLine:     ";",
		MultiLineOpen:  "",
		MultiLineClose: "",
		ExtractComment: ExtractAssemblyComments,
	},
	Bash: {
		SingleLine:     "#",
		MultiLineOpen:  ": '",
		MultiLineClose: "';",
		ExtractComment: ExtractBashComments,
	},
	Shell: {
		SingleLine:     "REM:",
		MultiLineOpen:  "::",
		MultiLineClose: "::",
		ExtractComment: ExtractShellComments,
	},
	PowerShell: {
		SingleLine:     "#",
		MultiLineOpen:  "<#",
		MultiLineClose: "#>",
		ExtractComment: ExtractPowerShellComments,
	},
	Matlab: {
		SingleLine:     "%",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
	},
	VB: {
		SingleLine:     "'",
		MultiLineOpen:  "",
		MultiLineClose: "",
		ExtractComment: ExtractVBComments,
	},

	Clojure: {
		SingleLine:     ";",
		MultiLineOpen:  "#|",
		MultiLineClose: "|#",
		ExtractComment: ExtractClojureComments,
	},

	Julia: {
		SingleLine:     "#",
		MultiLineOpen:  "#=",
		MultiLineClose: "=#",
		ExtractComment: ExtractJuliaComments,
	},

	Fortran: {
		SingleLine:     "!",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractFortranComments,
	},
	Unknown: {
		SingleLine:     "//",
		MultiLineOpen:  "/*",
		MultiLineClose: "*/",
		ExtractComment: ExtractCComments,
	},
}

// ExtractCommentsByLanguage extracts all comments from the given source code based on the specified programming language.
//
// Arguments:
//   - source: The source code as a string.
//   - lang: The programming language used in the source code.
//
// Returns:
//   - []string: A slice of extracted comments.

func ExtractCommentsByLanguage(source string, lang Language) []string {
	syntax, found := languageToCommentSyntax[lang]
	if !found || lang == Unknown {
		return nil
	}

	groups := syntax.ExtractComment(source)

	// for i, group := range groups {
	// 	fmt.Printf("=========================\ngroup: %v\n", i)
	// 	fmt.Println(group)
	// }
	return groups
}

// extractComments extracts comments from the given source using a regex pattern.
// This function ensures that extracted comments do not have leading newlines or tabs.
func extractComments(source string, pattern *regexp.Regexp) []string {
	comments := pattern.FindAllString(source, -1)
	for i := range comments {
		comments[i] = strings.TrimLeft(comments[i], "\n\t ")
	}
	return comments
}

// ExtractCComments extracts comments from C, C++, and Java code.
func ExtractCComments(source string) []string {
	return extractComments(source, CCommentRegex)
}

// ExtractCSSComments extracts comments from CSS files.
func ExtractCSSComments(source string) []string {
	return extractComments(source, CSSCommentRegex)
}

// ExtractHTMLComments extracts comments from HTML files.
func ExtractHTMLComments(source string) []string {
	return extractComments(source, HTMLCommentRegex)
}

// ExtractBashComments extracts comments from Bash scripts.
func ExtractBashComments(source string) []string {
	return extractComments(source, BashCommentRegex)
}

// ExtractShellComments extracts comments from Windows Batch scripts.
func ExtractShellComments(source string) []string {
	return extractComments(source, ShellCommentRegex)
}

// ExtractPythonComments extracts comments from Python scripts, including multi-line docstrings.
func ExtractPythonComments(source string) []string {
	return extractComments(source, PythonCommentRegex)
}

// ExtractAssemblyComments extracts comments from Assembly code.
func ExtractAssemblyComments(source string) []string {
	return extractComments(source, AssemblyCommentRegex)
}

// ExtractSQLComments extracts comments from SQL queries.
func ExtractSQLComments(source string) []string {
	return extractComments(source, SQLCommentRegex)
}

// ExtractLuaComments extracts comments from Lua scripts.
func ExtractLuaComments(source string) []string {
	return extractComments(source, LuaCommentRegex)
}

// ExtractHaskellComments extracts comments from Haskell code.
func ExtractHaskellComments(source string) []string {
	return extractComments(source, HaskellCommentRegex)
}

// ExtractPowerShellComments extracts comments from PowerShell scripts.
func ExtractPowerShellComments(source string) []string {
	return extractComments(source, PowerShellCommentRegex)
}

// ExtractVBComments extracts comments from Visual Basic code.
func ExtractVBComments(source string) []string {
	return extractComments(source, VBCommentRegex)
}

// ExtractClojureComments extracts comments from Clojure code.
func ExtractClojureComments(source string) []string {
	return extractComments(source, ClojureCommentRegex)
}

// ExtractJuliaComments extracts comments from Julia code.
func ExtractJuliaComments(source string) []string {
	return extractComments(source, JuliaCommentRegex)
}

// ExtractFortranComments extracts comments from Fortran code.
func ExtractFortranComments(source string) []string {
	return extractComments(source, FortranCommentRegex)
}

// ExtractElixirComments extracts Elixir-style (# and documentation comments).
func ExtractElixirComments(source string) []string {
	return extractComments(source, ElixirCommentRegx)
}

// ExtractRubyComments extracts Ruby-style (# and =begin...=end) comments.
func ExtractRubyComments(source string) []string {
	return extractComments(source, RubyCommentRegex)
}

// ExtractPascalComments extracts Pascal-style ({...}, (*...*), and //...) comments.
func ExtractPascalComments(source string) []string {
	return extractComments(source, PascalCommentRegex)
}
