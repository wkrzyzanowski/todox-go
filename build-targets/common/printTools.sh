#!/bin/bash

# Available colors within the 'tput'
#Num  Colour    #define         R G B
#
#0    black     COLOR_BLACK     0,0,0
#1    red       COLOR_RED       1,0,0
#2    green     COLOR_GREEN     0,1,0
#3    yellow    COLOR_YELLOW    1,1,0
#4    blue      COLOR_BLUE      0,0,1
#5    magenta   COLOR_MAGENTA   1,0,1
#6    cyan      COLOR_CYAN      0,1,1
#7    white     COLOR_WHITE     1,1,1

no_color='tput sgr0'
empty_message="(No input message)"

printMessage() {
  local message="${1:-$empty_message}"
  echo "$($no_color)[MESSAGE]: ${message}$($no_color)"
}

printInfo() {
  local blue='tput setaf 4'
  local message="${1:-$empty_message}"
  echo "$($blue)[INFO]: ${message}$($no_color)"
}

printWarn() {
  local yellow='tput setaf 3'
  local message="${1:-$empty_message}"
  echo "$($yellow)[WARN]: ${message}$($no_color)"
}

printError() {
  local red='tput setaf 1'
  local message="${1:-$empty_message}"
  echo "$($red)[ERROR]: ${message}$($no_color)"
}
