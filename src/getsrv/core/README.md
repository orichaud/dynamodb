# Set up data

Create Dynamodb table:
``` sh
./create.sh
```
Fill the table with synthetic data:
``` sh 
./load.sh
```
Drop table and data:
``` sh
./drop.sh
```

# Build the sample application
We assume `go` is install on your computer.
``` sh
make
```