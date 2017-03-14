# Message Service


## Installation
1. [Install Glide](#install-glide)
2. [Install Dependency](#install-dependency)

### Install Glide

[Glide](https://github.com/Masterminds/glide) is used to is a package manager for Go

On Mac OS X you can install the latest release via Homebrew:
~~~
$ brew update
$ brew install glide
~~~

On Ubuntu Precise(12.04), Trusty (14.04), Wily (15.10) or Xenial (16.04) you can install from our PPA:
~~~
sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update
sudo apt-get install glide
~~~

### Install Dependency
run commands below under project folder
~~~
$ glide install
$ go get -u golang.org/x/net/html
~~~

## How To Do

### How to Run Unit Test
Under the project folder, just simple run 
~~~
$ go test -v $(glide novendor)
~~~

### How to Run the Service

Get this repo locally, and build it!
~~~
$ cd $GOPATH/src
$ git clone git@github.com:JunjieQi/code_exercise.git
$ cd code_exercise
$ glide install
$ go get -u golang.org/x/net/html
$ go build
$ ./code_exercise
~~~

### How to Test the Endpoint
This service will listen port 3000, and it supports POST method for the endpoint `http://localhost:3000/v1/message`;
the post body looks like below:
~~~
@bob @john (success) such a cool feature; https://twitter.com/jdorfman/status/430511497475670016
~~~
And the response is JSON format
~~~
{
  "mentions": [
    "bob",
    "john"
  ],
  "emoticons": [
    "success"
  ],
  "links": [
    {
      "url": "https://twitter.com/jdorfman/status/430511497475670016",
      "title": "Justin Dorfman on Twitter: \"nice @littlebigdetail from @HipChat (shows hex colors when pasted in chat). http://t.co/7cI6Gjy5pq\""
    }
  ]
}
~~~

## Need To Do
1. There are still some corner cases needed to be covered
2. Need to do some load test
3. need to have a logger to log each request and error information