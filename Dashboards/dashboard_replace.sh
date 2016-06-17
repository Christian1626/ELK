#!/usr/bin/env bash

#./dashboard_replace.sh mg ml "/home/christian/Documents/ELK/Dashboards/dashboard_test_mg"


folder="$3"

if [ "$2" == "*" ]
then
	echo "OUI"
	new_folder="${folder/$1/all}"
else
	new_folder="${folder/$1/$2}"
fi

cp -R $folder $new_folder


if [ $# -eq 3 ] 
	then
for file in $new_folder/*/*.json
	do
		echo "var1:$1 var2:$2";
		sed -i "s/logstash-canalv2-$1/logstash-canalv2-"$2"/g" "$file";
	done
else
	echo "./dashboard_replace.sh <country_code_initial_dashboard> <new_country_code> <dashboards_folder> "
fi

