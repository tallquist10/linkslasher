# LinkSlasher

LinkSlasher is a url shortener meant to be just as easy on its developers as it is on its users. The interface is dead simple to use, and it is built to use either SQLite or MySQL, with the ability to easily add support for other engines.

## Setting up

After cloning the project, you can build the project locally using `build.sh`, which takes the Go minor version (e.g. 1.23) your system is using as a required parameter, and then optionally the name you want your output binary to use. 

```shell
./build.sh 1.23 # creates the output binary at ./linkslasher using Go 1.23

./build.sh 1.24 output # creates the output binary at ./output using Go 1.24
```

Once the binary is built, run the `run.sh` script to run locally, setting the `SQL_MODE` environment variable to set whether you want to use a MySQL or SQLite database.

### Building the Docker image

Building the Docker image for the project performs the following tasks:
1. Build the application using the Go version described in the Dockerfile.
2. Create and populate the SQLite database file with the tables necessary to run the application based on the contents of `sql/init.sql`.
3. Generate the output binary with necessary files and expose the correct port to be run on a Docker engine.

## Frequently Asked Questions
### When I open the project for the first time, I see a bunch of errors. Why is that?
Don't fret! Because this project does not have a `go.sum` file (explained below), it doesn't have the packages pulled in. All you have to do is either build the Docker image or run `./build.sh` locally as described in [Setting up](#setting-up), and all of those errors will magically disappear.

### How does this project not have a go.sum file?

This is a great question! The idea behind this project is that you should be able to run it on your local machine without having to download Go versions to match. This is done using the `go mod edit` command to change the `go.mod` file to use your local Go version, and then build the project based on that version. If you use Docker, you can build the application into Docker using one of the Dockerfiles in the project. If you feel like upgrading the Go version to a new version than what exists in the project, please feel free to do so! Simply make the new Dockerfile and name it as `Dockerfile{version}`, where `{version}` is the Go minor version being used, such as 1.23.

### What are some future improvements?