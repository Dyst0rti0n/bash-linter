# Bash Script Linter

A simple Bash script linter written in Go to detect common issues in Bash scripts.

## Features

- Detects unquoted variables.
- Detects unquoted command substitutions.
- Suggests using double square brackets for conditionals.
- Detects unused variables.

## Usage

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/bash_linter.git
   cd bash_linter
   ```

2. Build the linter:
   ```sh
   go build -o bash_linter main.go
   ```

3. Run the linter on a Bash script:
   ```sh
   ./bash_linter <script.sh>
   ```
**Alternatively, you could run `go run main.go <script.sh>`**

## Example

Sample Bash script (`sample.sh`)

Running the linter:
```sh
./bash_linter sample.sh
```

Expected output:
```
Line 6: Unused variable: unusedVar1
Line 7: Unused variable: unusedVar2
Line 10: Variable VAR1 declared but not initialized
Line 13: Hard-coded path detected. Consider using variables or environment variables.
Line 16: Unused function: unusedFunc
Line 19: Consider using $(...) instead of backticks for command substitution
Line 22: Unquoted variable $output found
Line 23: Unquoted variable ${VAR2} found
Line 26: Dangerous command 'rm -rf' detected. Ensure you have proper safeguards.
Line 29: Inefficient loop detected. Consider using C-style loops for better performance.
Line 32: Inconsistent indentation detected. Use either spaces or tabs consistently.
Line 33: Consider using double square brackets for conditionals
Line 39: Possible missing 'fi' for 'if' statement
Line 43: Possible missing 'done' for 'for' statement
Line 46: Non-portable command 'which' detected. Consider using a more portable alternative.
Line 49: Insecure use of sudo detected. Consider using 'sudo -k && sudo' to ensure sudo permissions are reset.
Line 52: Usage of 'exit' inside 'if' statement detected. Consider refactoring.
Line 56: Potential security vulnerability detected with 'eval'. Consider refactoring to avoid using eval.
Line 59: Unquoted variable in case statement detected. Consider quoting the variable.
Line 66: Lack of error handling detected. Consider adding '|| exit' to critical commands.
Line 70: Unnecessary command 'cd -' detected. Consider removing it.
Line 73: Exit command without specific exit code detected. Consider using 'exit 0' for success or 'exit 1' for failure.
Line 76: Function 'undocumentedFunc' does not follow naming convention. Consider prefixing with 'f_'.
Line 79: Missing documentation for function. Consider adding comments to explain its purpose.
Line 82: Inconsistent indentation detected. Use either spaces or tabs consistently.
Line 83: Inconsistent indentation detected. Use either spaces or tabs consistently.
28 issue(s) found.
```
