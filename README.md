# Bash Script Linter

A simple Bash script linter written in Go to detect common issues in Bash scripts.

## Features

- Detects unquoted variables.
- Detects unquoted command substitutions.
- Suggests using double square brackets for conditionals.
- Detects unused variables.
- Detects hard-coded paths.
- Detects improper function naming conventions.
- Detects unnecessary commands.
- Detects dangerous commands.
- Checks for missing keywords (e.g., `fi`, `done`).
- Checks for non-portable commands.
- Checks for indentation consistency.
- Checks for insecure use of `sudo`.
- Checks for missing or incorrect shebang.
- Checks for exit codes.
- Provides suggestions for improvements.

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

```bash
#!/bin/bash

# Example script to showcase linter issues

VAR1=
unusedVar1="This is an unused variable"
unusedVar2="This is another unused variable"
echo "This is a sample script"
myFunc() {
  echo "This function does something"
}
unusedFunc() {
  echo "This function is not used"
}
output=`ls -l`
echo $output
echo ${VAR2}
tempDir="/tmp/mydir"
rm -rf $tempDir
for i in `seq 1 10`
do
  echo $i
done
if [ $VAR2 -eq 1 ]
then
  echo "VAR2 is equal to 1"
fi
which bash
sudo rm -rf /root/*
eval "ls -l"
case $VAR2 in
  1)
    echo "One"
    ;;
  2)
    echo "Two"
    ;;
esac
cd -
exit
undocumentedFunc() {
  echo "This function is undocumented"
}
```

Running the linter:
```sh
./bash_linter sample.sh
```

Expected output:
```
Line 8: Variable VAR1 declared but not initialized
Line 12: Hard-coded path detected. Consider using variables or environment variables.
Line 15: Function 'myFunc' does not follow naming convention. Consider prefixing with 'f_'.
Line 16: Unnecessary command 'echo' detected. Consider removing it.
Line 16: Shell built-in 'echo' detected. Ensure its usage is intentional.
Line 20: Function 'unusedFunc' does not follow naming convention. Consider prefixing with 'f_'.
Line 21: Unnecessary command 'echo' detected. Consider removing it.
Line 21: Shell built-in 'echo' detected. Ensure its usage is intentional.
Line 25: Unquoted command substitution found
Line 25: Consider using $(...) instead of backticks for command substitution
Line 28: Unquoted variable $output found
Line 28: Unnecessary command 'echo' detected. Consider removing it.
Line 28: Shell built-in 'echo' detected. Ensure its usage is intentional.
Line 29: Unquoted variable ${VAR2} found
Line 29: Unnecessary command 'echo' detected. Consider removing it.
Line 29: Shell built-in 'echo' detected. Ensure its usage is intentional.
Line 32: Unquoted variable $tempDir found
Line 32: Dangerous command 'rm -rf' detected. Ensure you have proper safeguards.
Line 32: Lack of error handling detected. Consider adding '|| exit' to critical commands.
Line 35: Unquoted command substitution found
Line 35: Possible missing 'done' for 'for' statement
Line 35: Inefficient loop detected. Consider using C-style loops for better performance.
Line 36: Unquoted variable $i found
Line 36: Unnecessary command 'echo' detected. Consider removing it.
Line 36: Shell built-in 'echo' detected. Ensure its usage is intentional.
Line 40: Unquoted variable $VAR2 found
Line 40: Consider using double square brackets for conditionals
Line 40: Possible missing 'fi' for 'if' statement
Line 41: Unnecessary command 'echo' detected. Consider removing it.
Line 41: Shell built-in 'echo' detected. Ensure its usage is intentional.
Line 45: Unquoted variable $VAR2 found
Line 45: Consider using { } or ; do/done for loops and conditionals for better readability
Line 45: Possible missing 'fi' for 'if' statement
Line 46: Unnecessary command 'echo' detected. Consider removing it.
Line 46: Shell built-in 'echo' detected. Ensure its usage is intentional.
Line 49: Consider using { } or ; do/done for loops and conditionals for better readability
Line 49: Possible missing 'done' for 'for' statement
Line 50: Unquoted variable $i found
Line 50: Unnecessary command 'echo' detected. Consider removing it.
Line 50: Shell built-in 'echo' detected. Ensure its usage is intentional.
Line 54: Non-portable command 'which' detected. Consider using a more portable alternative.
Line 57: Dangerous command 'rm -rf' detected. Ensure you have proper safeguards.
Line 57: Hard-coded path detected. Consider using variables or environment variables.
Line 57: Insecure use of sudo detected. Consider using 'sudo -k && sudo' to ensure sudo permissions are reset.
Line 57: Lack of error handling detected. Consider adding '|| exit' to critical commands.
Line 60: Consider using double square brackets for conditionals
Line 60: Hard-coded path detected. Consider using variables or environment variables.
Line 65: Potential security vulnerability detected with 'eval'. Consider refactoring to avoid using eval.
Line 68: Unquoted variable $VAR2 found
Line 68: Unquoted variable in case statement detected. Consider quoting the variable.
Line 70: Unnecessary command 'echo' detected. Consider removing it.
Line 70: Shell built-in 'echo' detected. Ensure its usage is intentional.
Line 73: Unnecessary command 'echo' detected. Consider removing it.
Line 73: Shell built-in 'echo' detected. Ensure its usage is intentional.
Line 78: Hard-coded path detected. Consider using variables or environment variables.
Line 78: Shell built-in 'cd' detected. Ensure its usage is intentional.
Line 81: Unnecessary command 'cd -' detected. Consider removing it.
Line 81: Shell built-in 'cd' detected. Ensure its usage is intentional.
Line 84: Exit command without specific exit code detected. Consider using 'exit 0' for success or 'exit 1' for failure.
Line 87: Function 'undocumentedFunc' does not follow naming convention. Consider prefixing with 'f_'.
Line 88: Unnecessary command 'echo' detected. Consider removing it.
Line 88: Shell built-in 'echo' detected. Ensure its usage is intentional.
Line 92: Unquoted variable $VAR2 found
Line 92: Consider using double square brackets for conditionals
Line 92: Possible missing 'fi' for 'if' statement
Line 93: Unnecessary command 'echo' detected. Consider removing it.
Line 93: Shell built-in 'echo' detected. Ensure its usage is intentional.
Line 94: Unnecessary command 'echo' detected. Consider removing it.
Line 94: Shell built-in 'echo' detected. Ensure its usage is intentional.
Unused variable: unusedVar1
Unused variable: unusedVar2
Unused variable: VAR1
72 issue(s) found.
v All Issues

Line 8: Variable VAR1 declared but not initialized
  Suggestion: Initialize the variable at the time of declaration.

Line 12: Hard-coded path detected. Consider using variables or environment variables.
  Suggestion: Use variables or environment variables instead of hard-coded paths.

Line 15: Function 'myFunc' does not follow naming convention. Consider prefixing with 'f_'.
  Suggestion: Prefix function names with 'f_' for consistency.

Line 16: Unnecessary command 'echo' detected. Consider removing it.
  Suggestion: Remove the unnecessary command or ensure its usage is intentional.

Line 16: Shell built-in 'echo' detected. Ensure its usage is intentional.
  Suggestion: Remove the unnecessary command or ensure its usage is intentional.

... (more suggestions for each issue)

To view more details about each issue, refer to the documentation or use the linter with detailed output enabled.
```