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

## Example

Sample Bash script (`sample.sh`):
```bash
#!/bin/bash
VAR1="hello"
VAR2="world"
OUTPUT=$(ls)
if [ $VAR1 = "hello" ]; then
    echo $VAR1
fi
```

Running the linter:
```sh
./bash_linter sample.sh
```

Expected output:
```
Line 4: Unquoted command substitution found
Line 5: Unquoted variable $VAR1 found
Line 5: Consider using double square brackets for conditionals
Line 6: Unquoted variable $VAR1 found
Unused variable: VAR2
```
