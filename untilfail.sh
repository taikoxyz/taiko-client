#!/bin/bash

count=0
while make test; 
do 
    :; 
    count=$(expr $count + 1)
    echo "loop #$count"
done