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
* `go` is installed on your computer. With MacOS, the best is to use [Brew](https://brew.sh/) and then run `brew install go`.
* `VirtualBox` is installed on your computer. With MacOS, the best is to use [Brew](https://brew.sh/) and then run ` brew cask install virtualbox`.
* `docker` is installed and we can push to it. With MacOS, the best is to use [Brew](https://brew.sh/) and then run `brew install docker`. Once docker is installed, create a docker VM with `docker-machine create <name of your VM>`. 
* `python3` is installed.  With MacOS, the best is to use [Brew](https://brew.sh/) and then run `brew install docker`. Once docker is installed, create a docker VM with `brew install python3`
* `RobotFramework` is installed with `pip`: `pip3 install robotframework`.

Before compiling: 
* change the variable `DOCKER_VM` in the `Makefile` and set it with the name of your docker machine VM.
* run your docker VM with `docker-machine start <name of your docker VM>

To build the sample service:
``` sh
make
```
Note: the `Makefile` calls the `docker-machine` utility for MacOS. If you are running on a different environment, please update accordingly.

# Run and test
You can run the docker image that has been built in the previous section.
``` sh
./run_docker.sh
```
We use [robot framework](https://robotframework.org/) to run tests. Launch tests in a separate shell:
``` sh
 ./run_tests.sh
```