# Go Trial Simulator

A courtroom trial simulator that tests lawyers by recreating all the steps of an actual trial in a "Choose Your Own Adventure" style format. The system presents various stages, including opening statements, witness testimonies, and closing arguments, in a dynamic and interactive way. By incorporating user input, Gemini generates realistic courtroom scenarios and adapts to the strategies and actions of the lawyers, providing a challenging and immersive trial experience where each decision can influence the outcome.

## Run the sample
Get a Gemini 2 API key

## Launch Google AI Studio: https://aistudio.google.com/
Click Get API Key
Set the API Key in the API_KEY environment variable

export API_KEY=<your_api_key>
Compile and run the program:

go run .
When asked "What type of trial would you like to simulate?", provide a courtroom scenario.

For example, type: I want to simulate a criminal trial


## Set the build environment for Windows:

set GOOS=windows
set GOARCH=amd64
go build -o app-windows.exe

#!/bin/bash

## set 
export API_KEY=""

## run
./app-windows.exe
