#!/bin/bash 
TABLE_NAME=or-table

aws dynamodb create-table --cli-input-json file://table.json
