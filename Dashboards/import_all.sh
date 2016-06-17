#!/usr/bin/env bash
declare -a country=("mg" "ci" "sn" "cm" "ml" "ne" "gn" "cd" "*")
# declare -a country=("ne" "*")

#./import_all.sh .kibana_test logstash-canalv2 4.4.1 /home/christian/Documents/ELK/Dashboards/dashboard_test 
if [ $# -eq 4 ] 
	then
	# now loop through the above array
	for i in "${country[@]}"
	do

	   echo -e "\n===================="
	   echo -e "       $i           "
	   echo -e "===================="
	   if [ "$i" == "*" ]
	   then
	   		./import.sh "$1_all" "$2-*" "$4_all" "$3"
	   else
	   		./import.sh "$1_$i" "$2-$i" "$4_$i" "$3"
	   fi
	   # or do whatever with individual element of the array
	done
else
	echo "./dashboard_all_country.sh <base_kibana_index> <base_logstash_index> <kibana_version> <base_dashboards_folder> "
fi