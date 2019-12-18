#!/bin/bash 
TABLE_NAME=or-table

for i in {1..1000} 
do
    for j in {1..3}
    do
        read -d '' payload << EOM
        {
            "ID": {"S": "$i"},
            "SubKey": {"S": "S-$j"},
            "Data": {"S": "payload - $i/$j"}
        } 
EOM
        aws dynamodb put-item --table-name $TABLE_NAME --item "$payload"
        echo $payload
    done
done