package Analyzer

import (
	"fmt"
	"regexp"
	"statfiy/FileManager"
	"strings"
)

// Language represents a programming language type
type Language int

// List of supported programming languages
const (
	Go Language = iota
	C
	CPlusPlus
	CSharp
	Rust
	JavaScript
	TypeScript
	Python
	Java
	Kotlin
	Swift
	HTML
	CSS
	SQL
	PHP
	Ruby
	Dart
	Lua
	Perl
	Scala
	Haskell
	Assembly
	Bash
	R
	Matlab
	VB
	ObjectiveC
	Shell
	Pascal
	Elixir
	Clojure
	FSharp
	Julia
	PowerShell
	Fortran
	Zig
	Unknown // Default for unknown or unsupported languages
)

// languageNames maps Language enums to their string representations
var languageNames = map[Language]string{
	Go:         "Go",
	C:          "C",
	CPlusPlus:  "C++",
	CSharp:     "C#",
	Rust:       "Rust",
	JavaScript: "JavaScript",
	TypeScript: "TypeScript",
	Python:     "Python",
	Java:       "Java",
	Kotlin:     "Kotlin",
	Swift:      "Swift",
	HTML:       "HTML",
	CSS:        "CSS",
	SQL:        "SQL",
	PHP:        "PHP",
	Ruby:       "Ruby",
	Dart:       "Dart",
	Lua:        "Lua",
	Perl:       "Perl",
	Scala:      "Scala",
	Haskell:    "Haskell",
	Assembly:   "Assembly",
	Bash:       "Bash",
	R:          "R",
	Matlab:     "MATLAB",
	VB:         "Visual Basic",
	ObjectiveC: "Objective-C",
	Shell:      "Shell",
	Pascal:     "Pascal",
	Elixir:     "Elixir",
	Clojure:    "Clojure",
	FSharp:     "F#",
	Julia:      "Julia",
	PowerShell: "PowerShell",
	Fortran:    "Fortran",
	Zig:        "Zig",
	Unknown:    "Unknown",
}

// extensionToLanguage maps file extensions to Language enums
var extensionToLanguage = map[string]Language{
	".go":    Go,
	".c":     C,
	".h":     C,
	".cpp":   CPlusPlus,
	".cc":    CPlusPlus,
	".cxx":   CPlusPlus,
	".hpp":   CPlusPlus,
	".cs":    CSharp,
	".rs":    Rust,
	".js":    JavaScript,
	".ts":    TypeScript,
	".py":    Python,
	".java":  Java,
	".kt":    Kotlin,
	".swift": Swift,
	".html":  HTML,
	".css":   CSS,
	".sql":   SQL,
	".php":   PHP,
	".rb":    Ruby,
	".dart":  Dart,
	".lua":   Lua,
	".pl":    Perl,
	".scala": Scala,
	".hs":    Haskell,
	".asm":   Assembly,
	".sh":    Bash,
	".r":     R,
	".m":     Matlab, // Also used for Objective-C, requires extra check
	".vb":    VB,
	".mm":    ObjectiveC, // Objective-C++
	".bat":   Shell,
	".ps1":   PowerShell,
	".p":     Pascal,
	".ex":    Elixir,
	".clj":   Clojure,
	".fs":    FSharp,
	".jl":    Julia,
	".zig":   Zig,
}

var GitHubLanguageColors = map[Language]string{
	Go:         "#00ADD8",
	C:          "#555555",
	CPlusPlus:  "#F34B7D",
	CSharp:     "#178600",
	Rust:       "#DEA584",
	JavaScript: "#F1E05A",
	TypeScript: "#3178c6",
	Python:     "#3572A5",
	Java:       "#B07219",
	Kotlin:     "#F18E33",
	Swift:      "#FFAC45",
	HTML:       "#E34C26",
	CSS:        "#563D7C",
	SQL:        "#438EFF",
	PHP:        "#777BB4",
	Ruby:       "#701516",
	Dart:       "#00B4AB",
	Lua:        "#000080",
	Perl:       "#0298C3",
	Scala:      "#c22d40",
	Haskell:    "#5e5086",
	Assembly:   "#6E4C13",
	Bash:       "#89E051",
	R:          "#198CE7",
	Matlab:     "#0076A8",
	VB:         "#945DB7",
	ObjectiveC: "#438EFF",
	Shell:      "#89E051",
	Pascal:     "#E31C3D",
	Elixir:     "#6e4a7e",
	Clojure:    "#db5855",
	FSharp:     "#B845FC",
	Julia:      "#A93939",
	PowerShell: "#012456",
	Fortran:    "#4d41b1",
	Zig:        "#EC915C",
}

// String returns the string representation of a Language
func (l Language) String() string {
	if name, exists := languageNames[l]; exists {
		return name
	}
	return "Unknown"
}

func (l Language) GetColor() string {
	if color, exists := GitHubLanguageColors[l]; exists {
		return color
	}
	return GitHubLanguageColors[Unknown]
}

// GetLanguage determines the programming language based on file extension
func GetLanguage(metadata FileMetadata) Language {
	if lang, exists := extensionToLanguage[metadata.Extension]; exists {
		if metadata.Extension == ".m" {
			return DetectMFileType(metadata)
		}
		return lang
	}
	return Unknown // Default if extension is not recognized
}

func DetectMFileType(metadata FileMetadata) Language {
	objcPatterns := []*regexp.Regexp{
		regexp.MustCompile(`@interface`),
		regexp.MustCompile(`@implementation`),
		regexp.MustCompile(`@property`),
		regexp.MustCompile(`#import`),
		regexp.MustCompile(`NS[A-Z][a-zA-Z]+`), // Typical Objective-C class prefix
	}

	matlabPatterns := []*regexp.Regexp{
		regexp.MustCompile(`function`),
		regexp.MustCompile(`%`), // Matlab comment
		regexp.MustCompile(`linspace`),
		regexp.MustCompile(`zeros\(`),
		regexp.MustCompile(`ones\(`),
	}

	detectedType := Matlab
	linesRead := 0

	err := FileManager.ReadLinesLimit(metadata.Path, 20, func(line string) error {
		// Check for Objective-C patterns
		for _, pattern := range objcPatterns {
			if pattern.MatchString(line) {
				detectedType = ObjectiveC
				return fmt.Errorf("stop processing")
			}
		}

		// Check for Matlab patterns
		for _, pattern := range matlabPatterns {
			if pattern.MatchString(line) {
				detectedType = Matlab
				return fmt.Errorf("stop processing")
			}
		}

		linesRead++
		return nil
	})

	// Handle any processing errors (except our intentional stop)
	if err != nil && !strings.Contains(err.Error(), "stop processing") {
		return Matlab
	}

	return detectedType
}
