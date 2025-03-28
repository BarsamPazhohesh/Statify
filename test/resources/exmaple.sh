#!/bin/bash

# This is a single-line comment
echo "Hello, World!" # This is another single-line comment

# Below is a multi-line comment
: '
This is a multi-line comment.
You can write multiple lines of comments here.
'

# Perform some operations
count=5
while [ $count -gt 0 ]; do
  echo "Count: $count"
  ((count--)) # Decrement the count
done

# Exit the script
exit 0
