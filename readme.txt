aws dynamodb create-table --cli-input-json file://table.json
aws dynamodb put-item --table-name orichaud-Test --item file://put.json
aws dynamodb scan --table-name orichaud-Test
aws dynamodb get-item --table-name orichaud-Test --key file://get.json
aws dynamodb delete-table --table-name orichaud-Test