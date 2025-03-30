package analyzer

import (
	"regexp"
	"strings"
)

type ExtractComment func(source string) []string

// CommentSyntax represents symbols used for single-line and multi-line comments.
type CommentSyntax struct {
	ExtractComment ExtractComment
}

// Mapping of programming languages to their respective comment syntax.
var languageToCommentSyntax = map[Language]CommentSyntax{
	Go:         {ExtractComment: ExtractCComments},
	C:          {ExtractComment: ExtractCComments},
	CPlusPlus:  {ExtractComment: ExtractCComments},
	CSharp:     {ExtractComment: ExtractCComments},
	Rust:       {ExtractComment: ExtractCComments},
	JavaScript: {ExtractComment: ExtractCComments},
	TypeScript: {ExtractComment: ExtractCComments},
	Java:       {ExtractComment: ExtractCComments},
	Kotlin:     {ExtractComment: ExtractCComments},
	Swift:      {ExtractComment: ExtractCComments},
	PHP:        {ExtractComment: ExtractCComments},
	Dart:       {ExtractComment: ExtractCComments},
	Scala:      {ExtractComment: ExtractCComments},
	ObjectiveC: {ExtractComment: ExtractCComments},
	Python:     {ExtractComment: ExtractPythonComments},
	R:          {ExtractComment: ExtractPythonComments},
	Ruby:       {ExtractComment: ExtractRubyComments},
	Perl:       {ExtractComment: ExtractRubyComments},
	Pascal:     {ExtractComment: ExtractPascalComments},
	FSharp:     {ExtractComment: ExtractPascalComments},
	Elixir:     {ExtractComment: ExtractElixirComments},
	HTML:       {ExtractComment: ExtractHTMLComments},
	CSS:        {ExtractComment: ExtractCSSComments},
	SQL:        {ExtractComment: ExtractSQLComments},
	Lua:        {ExtractComment: ExtractLuaComments},
	Haskell:    {ExtractComment: ExtractHaskellComments},
	Assembly:   {ExtractComment: ExtractAssemblyComments},
	Bash:       {ExtractComment: ExtractBashComments},
	Shell:      {ExtractComment: ExtractShellComments},
	PowerShell: {ExtractComment: ExtractPowerShellComments},
	Matlab:     {ExtractComment: ExtractMatlabComments},
	VB:         {ExtractComment: ExtractVBComments},
	Clojure:    {ExtractComment: ExtractClojureComments},
	Julia:      {ExtractComment: ExtractJuliaComments},
	Fortran:    {ExtractComment: ExtractFortranComments},
	Unknown:    {ExtractComment: ExtractCComments}, // Default for unknown languages
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
func extractComments(source string, pattern string) []string {
	regex := regexp.MustCompile(pattern)
	comments := regex.FindAllString(source, -1)
	for i := range comments {
		comment := strings.TrimLeft(comments[i], "\n\t ")

		// Split comment into lines
		lines := strings.Split(comment, "\n")

		// Filter out empty lines
		var cleanedLines []string
		for _, line := range lines {
			// Trim spaces and check if the line is non-empty
			if strings.TrimSpace(line) != "" {
				cleanedLines = append(cleanedLines, line)
			}
		}
		comments[i] = strings.Join(cleanedLines, "\n")
	}
	return comments
}

// ExtractCComments extracts comments from C, C++, and Java code.
func ExtractCComments(source string) []string {
	return extractComments(source, CommentRegexs.CComment)
}

// ExtractCSSComments extracts comments from CSS files.
func ExtractCSSComments(source string) []string {
	return extractComments(source, CommentRegexs.CSSComment)
}

// ExtractHTMLComments extracts comments from HTML files.
func ExtractHTMLComments(source string) []string {
	return extractComments(source, CommentRegexs.HTMLComment)
}

// ExtractBashComments extracts comments from Bash scripts.
func ExtractBashComments(source string) []string {
	return extractComments(source, CommentRegexs.BashComment)
}

// ExtractShellComments extracts comments from Windows Batch scripts.
func ExtractShellComments(source string) []string {
	return extractComments(source, CommentRegexs.ShellComment)
}

// ExtractPythonComments extracts comments from Python scripts, including multi-line docstrings.
func ExtractPythonComments(source string) []string {
	return extractComments(source, CommentRegexs.PythonComment)
}

// ExtractAssemblyComments extracts comments from Assembly code.
func ExtractAssemblyComments(source string) []string {
	return extractComments(source, CommentRegexs.AssemblyComment)
}

// ExtractSQLComments extracts comments from SQL queries.
func ExtractSQLComments(source string) []string {
	return extractComments(source, CommentRegexs.SQLComment)
}

// ExtractLuaComments extracts comments from Lua scripts.
func ExtractLuaComments(source string) []string {
	return extractComments(source, CommentRegexs.LuaComment)
}

// ExtractHaskellComments extracts comments from Haskell code.
func ExtractHaskellComments(source string) []string {
	return extractComments(source, CommentRegexs.HaskellComment)
}

// ExtractPowerShellComments extracts comments from PowerShell scripts.
func ExtractPowerShellComments(source string) []string {
	return extractComments(source, CommentRegexs.PowerShellComment)
}

// ExtractVBComments extracts comments from Visual Basic code.
func ExtractVBComments(source string) []string {
	return extractComments(source, CommentRegexs.VBComment)
}

// ExtractClojureComments extracts comments from Clojure code.
func ExtractClojureComments(source string) []string {
	return extractComments(source, CommentRegexs.ClojureComment)
}

// ExtractJuliaComments extracts comments from Julia code.
func ExtractJuliaComments(source string) []string {
	return extractComments(source, CommentRegexs.JuliaComment)
}

// ExtractFortranComments extracts comments from Fortran code.
func ExtractFortranComments(source string) []string {
	return extractComments(source, CommentRegexs.FortranComment)
}

// ExtractElixirComments extracts Elixir-style (# and documentation comments).
func ExtractElixirComments(source string) []string {
	return extractComments(source, CommentRegexs.ElixirComment)
}

// ExtractRubyComments extracts Ruby-style (# and =begin...=end) comments.
func ExtractRubyComments(source string) []string {
	return extractComments(source, CommentRegexs.RubyComment)
}

// ExtractPascalComments extracts Pascal-style ({...}, (*...*), and //...) comments.
func ExtractPascalComments(source string) []string {
	return extractComments(source, CommentRegexs.PascalComment)
}

// ExtractMatlabComments extracts Pascal-style (% and {/*...*/}) comments.
func ExtractMatlabComments(source string) []string {
	return extractComments(source, CommentRegexs.MatlabComment)
}
