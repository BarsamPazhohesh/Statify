
@echo off
REM This script demonstrates the use of comments in a batch file.


(
REM This looks like a 
REM multi-line comment
REM But it's more of a command grouping
)

REM Display a welcome message
echo Welcome to my Batch script!

REM Creating a new folder named "MyFolder"
mkdir MyFolder

REM Navigate into the newly created folder
cd MyFolder

REM Create a text file inside "MyFolder"
echo This is a sample file. > sample.txt

REM Display the contents of the file
type sample.txt

REM Go back to the previous directory
cd ..

REM Pause the script so the user can see the output before exiting
pause
