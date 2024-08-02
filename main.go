package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

var (
	declaredVars    = map[string]bool{}
	declaredFuncs   = map[string]bool{}
	issuesFound     int
	checkShebang    bool
	checkUnusedVars bool
	errors          []string
	errorSuggestions = map[string]string{}
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

	// Interactive mode
	interactiveMode()
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
	checkFunctionDeclarations(trimmedLine)
	checkMissingKeywords(trimmedLine, lineNumber)
	checkVariableUsage(trimmedLine)
	checkNonPortableCommands(trimmedLine, lineNumber)
	checkDangerousCommands(trimmedLine, lineNumber)
	checkIndentation(trimmedLine, lineNumber)
	checkHardCodedPaths(trimmedLine, lineNumber)
	checkFunctionNaming(trimmedLine, lineNumber)
	checkExitCodes(trimmedLine, lineNumber)
	checkUnnecessaryCommands(trimmedLine, lineNumber)
	checkEfficientLoops(trimmedLine, lineNumber)
	checkCaseQuoting(trimmedLine, lineNumber)
	checkSudoUsage(trimmedLine, lineNumber)
	checkDocumentation(trimmedLine, lineNumber)
	checkErrorHandling(trimmedLine, lineNumber)
	checkScriptHeader(trimmedLine, lineNumber)
	checkShellBuiltIns(trimmedLine, lineNumber)
	checkSecurityVulnerabilities(trimmedLine, lineNumber)
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

func checkFunctionDeclarations(line string) {
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

func checkNonPortableCommands(line string, lineNumber int) {
	nonPortableCmds := []string{"which", "let", "source"}
	for _, cmd := range nonPortableCmds {
		if strings.Contains(line, cmd) {
			reportIssue(lineNumber, fmt.Sprintf("Non-portable command '%s' detected. Consider using a more portable alternative.", cmd))
		}
	}
}

func checkDangerousCommands(line string, lineNumber int) {
	dangerousCmds := []string{"rm -rf", "mkfs", ":(){ :|:& };:"}
	for _, cmd := range dangerousCmds {
		if strings.Contains(line, cmd) {
			reportIssue(lineNumber, fmt.Sprintf("Dangerous command '%s' detected. Ensure you have proper safeguards.", cmd))
		}
	}
}

func checkIndentation(line string, lineNumber int) {
	if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") {
		if strings.Contains(line, " ") && strings.Contains(line, "\t") {
			reportIssue(lineNumber, "Inconsistent indentation detected. Use either spaces or tabs consistently.")
		}
	}
}

func checkHardCodedPaths(line string, lineNumber int) {
	hardCodedPathRegex := regexp.MustCompile(`/[a-zA-Z0-9_/]+`)
	if hardCodedPathRegex.MatchString(line) {
		reportIssue(lineNumber, "Hard-coded path detected. Consider using variables or environment variables.")
	}
}

func checkFunctionNaming(line string, lineNumber int) {
	funcRegex := regexp.MustCompile(`^([a-zA-Z_][a-zA-Z0-9_]*)\s*\(\)`)
	match := funcRegex.FindStringSubmatch(line)
	if len(match) > 1 {
		funcName := match[1]
		if !strings.HasPrefix(funcName, "f_") {
			reportIssue(lineNumber, fmt.Sprintf("Function '%s' does not follow naming convention. Consider prefixing with 'f_'.", funcName))
		}
	}
}

func checkExitCodes(line string, lineNumber int) {
	if strings.Contains(line, "exit") && !strings.Contains(line, "exit 0") && !strings.Contains(line, "exit 1") {
		reportIssue(lineNumber, "Exit command without specific exit code detected. Consider using 'exit 0' for success or 'exit 1' for failure.")
	}
}

func checkUnnecessaryCommands(line string, lineNumber int) {
	unnecessaryCmds := []string{"cd -", "echo", "pwd"}
	for _, cmd := range unnecessaryCmds {
		if strings.Contains(line, cmd) {
			reportIssue(lineNumber, fmt.Sprintf("Unnecessary command '%s' detected. Consider removing it.", cmd))
		}
	}
}

func checkEfficientLoops(line string, lineNumber int) {
	if strings.Contains(line, "for ") && strings.Contains(line, "in") && strings.Contains(line, "seq") {
		reportIssue(lineNumber, "Inefficient loop detected. Consider using C-style loops for better performance.")
	}
}

func checkCaseQuoting(line string, lineNumber int) {
	caseRegex := regexp.MustCompile(`case\s+\$([a-zA-Z_][a-zA-Z0-9_]*)(\s*)in`)
	if caseRegex.MatchString(line) {
		reportIssue(lineNumber, "Unquoted variable in case statement detected. Consider quoting the variable.")
	}
}

func checkSudoUsage(line string, lineNumber int) {
	if strings.Contains(line, "sudo") && !strings.Contains(line, "&&") {
		reportIssue(lineNumber, "Insecure use of sudo detected. Consider using 'sudo -k && sudo' to ensure sudo permissions are reset.")
	}
}

func checkDocumentation(line string, lineNumber int) {
	if strings.HasPrefix(line, "function") && !strings.Contains(line, "#") {
		reportIssue(lineNumber, "Missing documentation for function. Consider adding comments to explain its purpose.")
	}
}

func checkErrorHandling(line string, lineNumber int) {
	if strings.Contains(line, "rm ") && !strings.Contains(line, "|| exit") {
		reportIssue(lineNumber, "Lack of error handling detected. Consider adding '|| exit' to critical commands.")
	}
}

func checkScriptHeader(line string, lineNumber int) {
	if lineNumber == 1 && !strings.HasPrefix(line, "#") {
		reportIssue(lineNumber, "Missing script header. Consider adding metadata like author, date, and purpose.")
	}
}

func checkShellBuiltIns(line string, lineNumber int) {
	shellBuiltIns := []string{"echo", "cd", "pwd", "let", "export", "unset"}
	for _, builtIn := range shellBuiltIns {
		if strings.Contains(line, builtIn) {
			reportIssue(lineNumber, fmt.Sprintf("Shell built-in '%s' detected. Ensure its usage is intentional.", builtIn))
		}
	}
}

func checkSecurityVulnerabilities(line string, lineNumber int) {
	if strings.Contains(line, "eval ") {
		reportIssue(lineNumber, "Potential security vulnerability detected with 'eval'. Consider refactoring to avoid using eval.")
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
	issue := ""
	if lineNumber > 0 {
		issue = fmt.Sprintf("Line %d: %s", lineNumber, message)
	} else {
		issue = message
	}
	errors = append(errors, issue)
	color.Red(issue)
	issuesFound++
	errorSuggestions[issue] = suggestFix(message)
}

func printSummaryReport() {
	if issuesFound == 0 {
		color.Green("No issues found.")
	} else {
		color.Red("%d issue(s) found.", issuesFound)
	}
}

func interactiveMode() {
	prompt := promptui.Select{
		Label: "Select an issue to view details or 'All Issues' to view all",
		Items: append([]string{"All Issues"}, errors...),
		Size:  10,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "All Issues" {
		for _, issue := range errors {
			color.Yellow(issue)
			fmt.Println("  Suggestion:", color.CyanString(errorSuggestions[issue]))
		}
	} else {
		color.Yellow(result)
		fmt.Println("  Suggestion:", color.CyanString(errorSuggestions[result]))
	}

	fmt.Println("To view more details about each issue, refer to the documentation or use the linter with detailed output enabled.")
}

func suggestFix(message string) string {
	switch {
	case strings.Contains(message, "Unquoted variable"):
		return "Quote the variable using \"${VAR}\" or '$VAR'."
	case strings.Contains(message, "command substitution"):
		return "Use $(...) instead of backticks."
	case strings.Contains(message, "double square brackets"):
		return "Use [[ ... ]] instead of [ ... ] for conditionals."
	case strings.Contains(message, "unused variable"):
		return "Remove the unused variable or use it appropriately."
	case strings.Contains(message, "dangerous command"):
		return "Add safeguards and ensure commands like 'rm -rf' are used with caution and proper checks."
	case strings.Contains(message, "inconsistent indentation"):
		return "Use either spaces or tabs consistently for indentation."
	case strings.Contains(message, "function naming"):
		return "Prefix function names with 'f_' for consistency."
	case strings.Contains(message, "shell built-in"):
		return "Ensure the usage of the shell built-in command is intentional and necessary."
	case strings.Contains(message, "exit code"):
		return "Specify an exit code (e.g., 'exit 0' for success, 'exit 1' for failure)."
	case strings.Contains(message, "hard-coded path"):
		return "Use variables or environment variables instead of hard-coded paths."
	case strings.Contains(message, "logical operators"):
		return "Use brackets around && and || for better readability."
	case strings.Contains(message, "empty variable declaration"):
		return "Initialize variables at the time of declaration."
	case strings.Contains(message, "loops and conditionals"):
		return "Use { } or ; do/done for better readability in loops and conditionals."
	case strings.Contains(message, "non-portable command"):
		return "Replace non-portable commands with more portable alternatives."
	case strings.Contains(message, "missing keyword"):
		return "Ensure that 'if' statements have a matching 'fi', and loops have matching 'done'."
	case strings.Contains(message, "documentation"):
		return "Add comments to explain the purpose of the function."
	case strings.Contains(message, "error handling"):
		return "Add error handling (e.g., '|| exit') to critical commands."
	case strings.Contains(message, "script header"):
		return "Add a script header with metadata like author, date, and purpose."
	case strings.Contains(message, "security vulnerability"):
		return "Refactor to avoid using potentially dangerous commands like 'eval'."
	default:
		return "Refer to best practices for resolving this issue."
	}
}
