#!/bin/bash
VAR1="hello"
VAR2="world"
OUTPUT=$(ls)
if [ $VAR1 = "hello" ]; then
    echo $VAR1
fi
for i in 1 2 3; do
    echo $i
done
if [[ $VAR1 == "hello" ]]; then
    echo "This is a test"
fi
