#!/bin/bash
usage () {
	echo "usage: setup.sh <import path>"
	echo "---------------------------------"
	echo "Example:"
	echo "	setup.sh 'github.com/JonathonGore/my-service'"
}

if [ $# -eq 0 ]
  then
	usage
	exit
fi

# Package required for renaming import paths
go get -u github.com/rogpeppe/govers

# Rename import paths referencing 
govers -d -m github.com/JonathonGore/go-service $1
