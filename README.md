# Set up data

Create Dynamodb table:
``` sh
./create.sh
```
Fill the table with synthetic data. This takes a while.
``` sh 
./load.sh
```
Drop table and data:
``` sh
./drop.sh
```

# Build the sample application
We assume:
* `go` is install on your computer.
* docker is installed and we can push to it.
To build the sample service:
``` sh
make
```

