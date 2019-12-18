package core

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

const (
	MAX        = 100
	TABLE_NAME = "or-table"
)

// This is the Item returned
type Item struct {
	ID     string `json:ID`
	SubKey string `json:SubKey`
	Data   string `json:Data`
}

type Items struct {
	Count int     `json:Count`
	Items []*Item `json:Items`
}

func (item *Item) String() string {
	return fmt.Sprintf("[ID=%s, SubKey=%s, Data=%s]", item.ID, item.SubKey, item.Data)
}

// Fetch an item by ID and SubKey
func HandleScanItems(context *Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	max := MAX
	if maxval, present := vars["max"]; present {
		value, err := strconv.Atoi(maxval)
		if err != nil {
			glog.Errorf("Invalid parameter - max=%s - %s", maxval, err)
			w.WriteHeader(http.StatusBadRequest)
			context.OnFailure()
			return
		}
		if value < 1 {
			glog.Errorf("Invalid non positive parameter - max=%s - %d", maxval, value)
			w.WriteHeader(http.StatusBadRequest)
			context.OnFailure()
			return
		}
		if value > MAX {
			glog.Errorf("max limit exceeding %d - max=%d", MAX, value)
			w.WriteHeader(http.StatusBadRequest)
			context.OnFailure()
			return
		}
		max = value
	}

	glog.Infof("Scan items requested - max=%d", max)

	params := &dynamodb.ScanInput{
		TableName: aws.String(TABLE_NAME),
		Limit:     aws.Int64(int64(max))}
	if result, err := context.DynamoClient.Scan(params); err != nil {
		glog.Errorf("Cannot scan %s: %s", TABLE_NAME, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		scannedItems := []*Item{}
		count := 0
		for _, i := range result.Items {
			item := &Item{}
			if err := dynamodbattribute.UnmarshalMap(i, item); err != nil {
				glog.Errorf("Cannot decode result from DB: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				context.OnFailure()
				return
			}
			scannedItems = append(scannedItems, item)
			count += 1
			if count > max {
				glog.Infof("Cannot return more than %d items", MAX)
				break
			}
		}
		glog.Infof("Items retrieved: %d", count)
		items := &Items{Items: scannedItems, Count: count}
		if err := Send(items, w); err != nil {
			glog.Errorf("Cannot transfer JSON payload: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			context.OnFailure()
		} else {
			context.OnSuccess()
		}
	}
}

// Fetch an item by ID and SubKey
func HandleGetItem(context *Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	subKey := vars["subkey"]
	glog.Infof("item %s/%s requested", id, subKey)

	input := &dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"ID":     {S: aws.String(id)},
			"SubKey": {S: aws.String(subKey)}}}

	if result, err := context.DynamoClient.GetItem(input); err != nil {
		glog.Errorf("Cannot get item %s/%s, from DB: %s", id, subKey, err.Error())
		w.WriteHeader(http.StatusNotFound)
		context.OnFailure()
		return
	} else {
		item := &Item{}
		if err := dynamodbattribute.UnmarshalMap(result.Item, item); err != nil {
			glog.Errorf("Cannot decode result from DB: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			context.OnFailure()
			return
		}
		glog.Infof("item %s/%s unmarshalled: %s", id, subKey, item)
		if item.ID == "" || item.SubKey == "" {
			glog.Errorf("Cannot get item %s/%s from DB", id, subKey)
			w.WriteHeader(http.StatusNotFound)
			context.OnFailure()
			return
		}

		if err := Send(item, w); err != nil {
			glog.Errorf("Cannot transfer JSON payload: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			context.OnFailure()
		} else {
			context.OnSuccess()
		}
	}
}

// Fetch an item by ID and SubKey
func HandleGetItems(context *Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	glog.Infof("items %s requested", id)

	condition := expression.Key("ID").Equal(expression.Value(id))
	builder := expression.NewBuilder().WithKeyCondition(condition)
	expr, _ := builder.Build()
	input := &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		TableName:                 aws.String(TABLE_NAME),
	}

	if result, err := context.DynamoClient.Query(input); err != nil {
		glog.Errorf("Cannot get items %s from DB: %s", id, err.Error())
		w.WriteHeader(http.StatusNotFound)
		context.OnFailure()
		return
	} else {
		allItems := []*Item{}
		for _, value := range result.Items {
			item := &Item{}
			err = dynamodbattribute.UnmarshalMap(value, &item)
			allItems = append(allItems, item)
			if len(allItems) > MAX {
				glog.Infof("Cannot return more than %d items", MAX)
				break
			}
		}

		items := &Items{
			Count: len(allItems),
			Items: allItems}
		if err := Send(items, w); err != nil {
			glog.Errorf("Cannot transfer JSON payload: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			context.OnFailure()
		} else {
			context.OnSuccess()
		}
	}
}
