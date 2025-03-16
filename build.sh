#!/bin/bash
project=$(go list -m)
build="target"
echo "Building project $project"
if [ -d $build ] 
then
    rm $build/$project
else
	mkdir $build/
fi
go build -ldflags "-s -w"
echo "Copying files"
mv "$project" $build/
