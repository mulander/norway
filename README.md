norway
======

CVS standalone web-based code review / commit stream tool

installation
------------

In order to compile the standalone app:

`go get github.com/mulander/norway/norway`

Then run:
`$GOPATH/bin/norway -base $NORWAY -cvs $CVS`

Then open your browser to:

`http://localhost:8080`

In order to use the CVS interop features in your code just:

`import "github.com/mulander/norway"

status
------

In active development and maintained. Bugs/feature request reported on the github issue tracker will be actively fixed.

testing
-------

In order to run the tests install goconvey:

`go get github.com/`

Then run goconvey in the checked-out code directory:

`$GOPATH/bin/goconvey`

Then open your browser to:

`http://localhost:8080`

goals
-----
* work with a repository checkout - don't assume central repository access
* easy setup - checkout the repository and start the service to get going
* stream committs and filter down to the interesting leafs
* down to the guts - easy to ask for/perform a code review
