#!/bin/bash

# Unused variables
unusedVar1="hello"
unusedVar2="world"

# Variable declarations
VAR1= 
VAR2="sample"

# Hard-coded paths
tempDir=/tmp/mydir

# Function declaration
myFunc() {
    echo "This is my function"
}

# Unused function
unusedFunc() {
    echo "This function is never used"
}

# Deprecated syntax for command substitution
output=`ls -l`

# Unquoted variables and command substitution
echo $output
echo ${VAR2}

# Dangerous command without safeguard
rm -rf $tempDir

# Loop with inefficient construct
for i in $(seq 1 10); do
    echo $i
done

# Inconsistent indentation
    if [ $VAR2 = "sample" ]; then
    echo "Sample variable is set"
    fi

# Missing fi and done
if [[ $VAR2 == "sample" ]]
    echo "Sample variable is set"
fi

for i in 1 2 3
    echo $i
done

# Non-portable command
which ls

# Use of sudo without safeguard
sudo rm -rf /important/dir

# Exit in if statement
if [ -f "/path/to/file" ]; then
    exit 1
fi

# Use of eval (potential security risk)
eval "ls -l"

# Case statement without proper quoting
case $VAR2 in
    sample)
        echo "It is a sample"
        ;;
    *)
        echo "Not a sample"
        ;;
esac

# Missing error handling
cd /nonexistentdir

# Unnecessary command
cd -

# Exit command without specific code
exit

# Function without documentation
undocumentedFunc() {
    echo "No documentation"
}

# Inconsistent indentation
if [ $VAR2 = "sample" ]; then
	echo "This line is indented with a tab"
  echo "This line is indented with spaces"
fi
