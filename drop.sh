#!/bin/bash 
TABLE_NAME=or-table

aws dynamodb delete-table --table-name $TABLE_NAME
