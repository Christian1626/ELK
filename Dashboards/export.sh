#!/usr/bin/env bash

# ./export.sh .kibana /home/christian/Téléchargements/dashboards
# $1 : kibana index
# $2 : dashboards directory


#
# Saves all dashboards in Kibana 

if [ $# -eq 2 ] 
then
	# dashboard
	curl -s --noproxy 127.0.0.1 "127.0.0.1:9200/$1/dashboard/_search?pretty=true&size=1000&fields=" | grep "_id" | sed -E 's/.*"_id" : "(.*)",/\1/' | while read -r line; do curl --noproxy 127.0.0.1 -s -X GET 127.0.0.1:9200/$1/dashboard/$line/_source > /home/christian/Téléchargements/dashboards/dashboard/$line.json; done

	# visualization
	curl -s --noproxy 127.0.0.1 "127.0.0.1:9200/$1/visualization/_search?pretty=true&size=1000&fields=" | grep "_id" | sed -E 's/.*"_id" : "(.*)",/\1/' | while read -r line; do curl --noproxy 127.0.0.1 -s -X GET 127.0.0.1:9200/$1/visualization/$line/_source > $2/visualization/$line.json; done

	# search
	curl -s --noproxy 127.0.0.1 "127.0.0.1:9200/$1/search/_search?pretty=true&size=1000&fields=" | grep "_id" | sed -E 's/.*"_id" : "(.*)",/\1/' | while read -r line; do curl --noproxy 127.0.0.1 -s -X GET 127.0.0.1:9200/$1/search/$line/_source > $2/search/$line.json; done

	# index-pattern
	curl -s --noproxy 127.0.0.1 "127.0.0.1:9200/$1/index-pattern/_search?pretty=true&size=1000&fields=" | grep "_id" | sed -E 's/.*"_id" : "(.*)",/\1/' | while read -r line; do curl --noproxy 127.0.0.1 -s -X GET 127.0.0.1:9200/$1/index-pattern/$line/_source > $2/index-pattern/$line.json; done
else
	echo "./export.sh <kibana_index> <dashboards_folder>"
fi