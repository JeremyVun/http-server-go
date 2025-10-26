#!/bin/bash

name=result
route=hello

while getopts ":r:n:" opt; do
    case $opt in
        r) route="$OPTARG" ;;
        n) name="$OPTARG" ;;
        \?) echo "Invalid option"; exit 1 ;;
        :) echo "Option -$OPTARG requires an arg"; exit 1 ;;
    esac
done

echo "Benching localhost:$port/$route with $threads threads"

~/tools/apache-jmeter/bin/jmeter -n -t perf_test.jmx -l results/$name.csv -f -e -o reports/$name -Jroute=$route
