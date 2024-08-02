package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

var (
	declaredVars    = map[string]bool{}
	declaredFuncs   = map[string]bool{}
	issuesFound     int
	checkShebang    bool
	checkUnusedVars bool
)

func init() {
	flag.BoolVar(&checkShebang, "check-shebang", true, "Enable/disable shebang check")
	flag.BoolVar(&checkUnusedVars, "check-unused-vars", true, "Enable/disable unused variable check")
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: bash_linter [options] <script.sh>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	filePath := flag.Args()[0]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	firstLine := true
	for scanner.Scan() {
		line := scanner.Text()
		if firstLine {
			if checkShebang {
				checkShebangLine(line, lineNumber)
			}
			firstLine = false
		}
		lintLine(line, lineNumber)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	if checkUnusedVars {
		// Check for unused variables after reading all lines
		checkUnusedVarsFunc()
	}

	// Check for unused functions after reading all lines
	checkUnusedFuncs()

	// Print a summary report
	printSummaryReport()
}

func checkShebangLine(line string, lineNumber int) {
	if !strings.HasPrefix(line, "#!/bin/bash") {
		reportIssue(lineNumber, "Missing or incorrect shebang. Consider adding '#!/bin/bash' at the top of the script")
	}
}

func lintLine(line string, lineNumber int) {
	// Skip comments and empty lines
	trimmedLine := strings.TrimSpace(line)
	if strings.HasPrefix(trimmedLine, "#") || trimmedLine == "" {
		return
	}

	checkUnquotedVariables(trimmedLine, lineNumber)
	checkUnquotedCommandSubstitutions(trimmedLine, lineNumber)
	checkSingleSquareBrackets(trimmedLine, lineNumber)
	checkLogicalOperators(trimmedLine, lineNumber)
	checkEmptyVariableDeclarations(trimmedLine, lineNumber)
	checkLoopsAndConditionals(trimmedLine, lineNumber)
	checkCommandSubstitutionBestPractice(trimmedLine, lineNumber)
	checkExitInIfStatement(trimmedLine, lineNumber)
	checkFunctionDeclarations(trimmedLine, lineNumber)
	checkMissingKeywords(trimmedLine, lineNumber)
	checkVariableUsage(trimmedLine)
}

func checkUnquotedVariables(line string, lineNumber int) {
	unquotedVarRegex := regexp.MustCompile(`\$\{?[a-zA-Z_][a-zA-Z0-9_]*\}?`)
	matches := unquotedVarRegex.FindAllString(line, -1)
	for _, match := range matches {
		if !isQuoted(line, match) {
			reportIssue(lineNumber, fmt.Sprintf("Unquoted variable %s found", match))
		}
	}
}

func checkUnquotedCommandSubstitutions(line string, lineNumber int) {
	unquotedCmdSubRegex := regexp.MustCompile(`\$\([^)]+\)|` + "`[^`]+`")
	if unquotedCmdSubRegex.MatchString(line) {
		reportIssue(lineNumber, "Unquoted command substitution found")
	}
}

func checkSingleSquareBrackets(line string, lineNumber int) {
	singleBracketRegex := regexp.MustCompile(`\[[^]]+\]`)
	if singleBracketRegex.MatchString(line) && !strings.Contains(line, "[[") {
		reportIssue(lineNumber, "Consider using double square brackets for conditionals")
	}
}

func checkLogicalOperators(line string, lineNumber int) {
	if (strings.Contains(line, "&&") || strings.Contains(line, "||")) &&
		!(strings.Contains(line, "(") && strings.Contains(line, ")")) {
		reportIssue(lineNumber, "Consider using brackets around && and || for clarity")
	}
}

func checkEmptyVariableDeclarations(line string, lineNumber int) {
	if strings.Contains(line, "=") && !strings.Contains(line, "==") {
		parts := strings.SplitN(line, "=", 2)
		varName := strings.TrimSpace(parts[0])
		varValue := strings.TrimSpace(parts[1])
		if isValidVarName(varName) {
			declaredVars[varName] = false
		}
		if varValue == "" {
			reportIssue(lineNumber, fmt.Sprintf("Variable %s declared but not initialized", varName))
		}
	}
}

func checkLoopsAndConditionals(line string, lineNumber int) {
	if strings.HasPrefix(line, "for ") || strings.HasPrefix(line, "while ") || strings.HasPrefix(line, "if ") {
		if !strings.Contains(line, ";") && !strings.Contains(line, "{") {
			reportIssue(lineNumber, "Consider using { } or ; do/done for loops and conditionals for better readability")
		}
	}
}

func checkCommandSubstitutionBestPractice(line string, lineNumber int) {
	backtickCmdSubRegex := regexp.MustCompile("`[^`]+`")
	if backtickCmdSubRegex.MatchString(line) {
		reportIssue(lineNumber, "Consider using $(...) instead of backticks for command substitution")
	}
}

func checkExitInIfStatement(line string, lineNumber int) {
	if strings.HasPrefix(strings.TrimSpace(line), "if ") && strings.Contains(line, "exit") {
		reportIssue(lineNumber, "Usage of 'exit' inside 'if' statement detected. Consider refactoring.")
	}
}

func checkFunctionDeclarations(line string, lineNumber int) {
	funcRegex := regexp.MustCompile(`^([a-zA-Z_][a-zA-Z0-9_]*)\s*\(\)`)
	match := funcRegex.FindStringSubmatch(line)
	if len(match) > 1 {
		funcName := match[1]
		declaredFuncs[funcName] = false
	}
}

func checkMissingKeywords(line string, lineNumber int) {
	keywords := map[string]string{
		"if":   "fi",
		"for":  "done",
		"while":"done",
	}
	for start, end := range keywords {
		if strings.HasPrefix(line, start) && !strings.Contains(line, end) {
			reportIssue(lineNumber, fmt.Sprintf("Possible missing '%s' for '%s' statement", end, start))
		}
	}
}

func checkVariableUsage(line string) {
	for varName := range declaredVars {
		if strings.Contains(line, "$"+varName) || strings.Contains(line, "${"+varName+"}") {
			declaredVars[varName] = true
		}
	}
	for funcName := range declaredFuncs {
		if strings.Contains(line, funcName+"(") {
			declaredFuncs[funcName] = true
		}
	}
}

func isQuoted(line, variable string) bool {
	quotedVarRegex := regexp.MustCompile(`(['"][^'"]*\$` + regexp.QuoteMeta(variable) + `[^'"]*['"])|(\$` + regexp.QuoteMeta(variable) + `[^a-zA-Z0-9_])`)
	return quotedVarRegex.MatchString(line)
}

func isValidVarName(varName string) bool {
	varNameRegex := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
	return varNameRegex.MatchString(varName)
}

func checkUnusedVarsFunc() {
	for varName, used := range declaredVars {
		if !used {
			reportIssue(-1, fmt.Sprintf("Unused variable: %s", varName))
		}
	}
}

func checkUnusedFuncs() {
	for funcName, used := range declaredFuncs {
		if !used {
			reportIssue(-1, fmt.Sprintf("Unused function: %s", funcName))
		}
	}
}

func reportIssue(lineNumber int, message string) {
	if lineNumber > 0 {
		color.Red("Line %d: %s", lineNumber, message)
	} else {
		color.Red(message)
	}
	issuesFound++
}

func printSummaryReport() {
	if issuesFound == 0 {
		color.Green("No issues found.")
	} else {
		color.Red("%d issue(s) found.", issuesFound)
	}
}
