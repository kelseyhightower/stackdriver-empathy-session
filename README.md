# Stackdriver Empathy Session

This repository holds the source code for each of the components required to participate in the Stackdriver empathy session.

## Components

* `frontend` - An HTTP service that processes HTTP requests by sending ping messages to a `backend` service over pubsub.
* `backend` - Reads messages from a Pub/Sub subscription and prints the message data to standard out.
* `client` - Submits HTTP requests to the `frontend` service, because cURL would have been too easy.
