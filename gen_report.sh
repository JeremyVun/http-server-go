#!/bin/bash

name=result

while getopts ":n:" opt; do
    case $opt in
        n) name="$OPTARG" ;;
        \?) echo "Invalid option"; exit 1 ;;
        :) echo "Option -$OPTARG requires an arg"; exit 1 ;;
    esac
done

echo "Generating report for $name"

~/tools/apache-jmeter/bin/jmeter -g results/$name.csv -o reports/$name
