#!/usr/bin/env bash

# ./import.sh .kibana_test logstash-canalv2-mg /home/christian/Documents/ELK/Dashboards/dashboard_mg 


# $1 : kibana index
# $2 : default index-pattern
# $3 : dashboards directory
# $4 : kibana version


# Loads previously saved Kibana dashboards from $3/ into the current Elasticsearch instance



if [ $# -eq 4 ] 
then
	#dashboard
	for file in $3/dashboard/*.json
	do
		filename=$(basename $file)
		name=${filename%.*}
		curl --noproxy 127.0.0.1 -X PUT 127.0.0.1:9200/$1/dashboard/${name} -T $3/dashboard/${filename} --silent
		echo
	done


	#visualization
	for file in $3/visualization/*.json
	do
		filename=$(basename $file)
		name=${filename%.*}
		curl --noproxy 127.0.0.1 -X PUT 127.0.0.1:9200/$1/visualization/${name} -T $3/visualization/${filename} --silent
		echo
	done


	#search
	for file in $3/search/*.json
	do
		filename=$(basename $file)
		name=${filename%.*}
		curl --noproxy 127.0.0.1 -X PUT 127.0.0.1:9200/$1/search/${name} -T $3/search/${filename} --silent
		echo
	done

	#index-pattern
	for file in $3/index-pattern/*.json
	do
		filename=$(basename $file)
		name=${filename%.*}
		curl --noproxy 127.0.0.1 -X PUT 127.0.0.1:9200/$1/index-pattern/${name} -T $3/index-pattern/${filename} --silent
		echo
	done
else
	echo "./export.sh <kibana_index> <default_index_pattern> <dashboards_folder> <kibana_version>"
fi

#set default index-pattern
# kibana_version=$(curl -s --noproxy 127.0.0.1 "127.0.0.1:9200/$1/config/_search?pretty=true&size=1000&fields=" | grep "_id" | sed -E 's/.*"_id" : "(.*)",/\1/');
# echo $kibana_version

curl --noproxy 127.0.0.1 -X PUT 127.0.0.1:9200/$1/config/$4 -d '{"defaultIndex":"'$2'"}'
# echo "curl --noproxy 127.0.0.1 -X PUT 127.0.0.1:9200/"$1"/config/"$kibana_version" -d "'{"defaultIndex":"'$2'"}'
# curl --noproxy 127.0.0.1 -X PUT 127.0.0.1:9200/.kibana_test/config/4.4.1 -d '{"defaultIndex":"logstash-canalv2-mg"}'


#curl --noproxy 127.0.0.1 -X PUT 127.0.0.1:9200/$1/config/4.4.1 -d '{"buildNum":9693,"defaultIndex":"'$2'"}'


#if [ kibana_version ]

# curl -XPUT 127.0.0.1:9200/$1/config/4.4.1 -d '{"defaultIndex" : "logstash-*"}