#!/usr/bin/env bash

# ./import.sh .kibana_test pfa /home/christian/Téléchargements/dashboards 

# $1 : kibana index
# $2 : default index
# $3 : dashboards directory


# Loads previously saved Kibana dashboards from $3/ into the current Elasticsearch instance



if [ $# -eq 3 ] 
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
	echo "./export.sh <kibana_index> <default_index_pattern> <dashboards_folder>"
fi

#set default index-pattern
kibana_version=$(curl -s --noproxy 127.0.0.1 "127.0.0.1:9200/$1/config/_search?pretty=true&size=1000&fields=" | grep "_id" | sed -E 's/.*"_id" : "(.*)",/\1/');
curl --noproxy 127.0.0.1 -X PUT 127.0.0.1:9200/$1/config/$kibana_version -d '{"defaultIndex":"logstash-'$2'"}'