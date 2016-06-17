#!/usr/bin/env bash
declare -a country=("mg" "ci" "sn" "cm" "ml" "ne" "gn" "cd" "*")
# declare -a country=("*")


if [ $# -eq 2 ] 
	then
	# now loop through the above array
	for i in "${country[@]}"
	do
	   echo "$i"
	   if [ $1 != "$i" ]
	   then
	   		./create_dashboard.sh $1 "$i" $2
	   fi
	   # or do whatever with individual element of the array
	done
else
	echo "./dashboard_all_country.sh <country_code_initial_dashboard> <dashboards_folder> "
fi