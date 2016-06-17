#!/usr/bin/env bash

#./dashboard_replace.sh mg ml "/home/christian/Documents/ELK/Dashboards/dashboard_test_mg"


folder="$3"

if [ "$2" == "*" ]
then
	new_folder="${folder/$1/all}"
else
	new_folder="${folder/$1/$2}"
fi

cp -R $folder $new_folder


if [ $# -eq 3 ] 
then
	for file in $new_folder/*/*.json
	do
		sed -i "s/logstash-canalv2-$1/logstash-canalv2-"$2"/g" "$file";
	done

	echo $new_folder/index-pattern/logstash-canalv2-"$1" $new_folder/index-pattern/logstash-canalv2-"$2"
	mv $new_folder/index-pattern/logstash-canalv2-"$1".json $new_folder/index-pattern/logstash-canalv2-"$2".json

else
	echo "./dashboard_replace.sh <country_code_initial_dashboard> <new_country_code> <dashboards_folder> "
fi


