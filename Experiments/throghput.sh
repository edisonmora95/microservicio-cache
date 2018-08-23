#!/bin/bash

# Define a timestamp function
timestamp() {
  date +"%T"
}



counter2=0

while [ $counter2 -lt 10 ]
do
    counter=0
    timestamp1=$(date +%s)
    timestamp1=$(($timestamp1 +10))
    while true
    do

        timestamp2=$(date +%s)

        if [ $timestamp2 -gt $timestamp1 ]
        then
            break
        fi
        curl -O http://ec2-18-223-117-144.us-east-2.compute.amazonaws.com/api/gif/gif-0 > /dev/null 2>&1
        counter=$(($counter +1))
    done

    echo $counter
    counter2=$(($counter2 +1))
done