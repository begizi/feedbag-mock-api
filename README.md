FeedBag Mock Api
================

## Why?

The answer is pretty simple. Not everyone knows go who wants to work on FeedBag.
This tool is a simple compiled binary that points to some static json files with the
expected FeedBag endpoints. Developers can simply execute the binary, tweek the json
files to their liking, and begin working on the frontend as if their were hitting
a live server.

## Why is it written in go if it's job is to keep people from having to work with go?

Good question. Mainly because I want to work as much with go as possible to learn as
much as I can about the language. Another main reason is that go easily compiles for
many types of systems without really thinking twice, and what's nicer than just executing
a file and then the server is up in the blink of an eye?

## Why not something like Apiary for a mock server?

2 reasons.

1. People may want to tweek the responses without having to make their own mock endpoints.
2. I will be adding fake socket support to mimic a live environment. Not so easily done with
those internet based tools.

## Installation

```
$ git clone git@github.com:begizi/feedbag-mock-api.git

// For Mac
$ ./feedbag-mock-api/serve -port 3000

// For Linux
$ ./feedbag-mock-api/serve.linux -port 3000

//For Windows
$ ./feedbag-mock-api/serve.exe -port 3000
```

Thats it.

*Force updated for Feedbag testing*
