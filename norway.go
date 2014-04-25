/*
Package norway implements methods for interacting with a checked-out CVS repository

There are two types of interactions supported:
1. Parsing the on-disk stored CVS data like the CVS/Entries file
2. Talking with the CVS server using a binary executable CVS client to obtain
   information that is not stored on the file system
*/
package norway
