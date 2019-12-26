# Set up data

We assume a valid AWS account with a default project/region and this is in your default configuration stored in your `~/.aws`. Please perform all the configuration steps with AWS CLI as a prerequisite. Step by step explanations can be found [here](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html). 

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
* `docker` is installed and we can push to it.
To build the sample service:
``` sh
make
```
Note: the `Makefile` calls the `docker-machine` utility for MacOS. If you are running on a different environment, please update accordingly. 

# Run and tests
You can run the docker image that has been built in the previous section.
``` sh
./run_docker.sh
```
We use [robot framework](https://robotframework.org/) to run tests. Launch tests in a separate shell:
``` sh
 ./run_tests.sh
```