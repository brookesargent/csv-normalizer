# CSV Normalizer

### To Build and Run:
* Run `make build-mac` or `make build-linux` from the root directory depending on your OS
* Run `make run` from the root directory

### Notes
To stay within the 4 hour time limit on this project, I definitely had to make technical choices and tradeoffs. I was able to get most of the core requirements of the normalizer in place but ran out of time when it came to validating the UTF-8 characters in the freeform text fields. I used TDD when writing the normalizations I did get to in the `normalizations` package. Given more time I would have liked to:
* Write a normalization for the Unicode Character Replacement to fully meet all the requirements
* Write Github Actions to utilize the included Makefile and run linting and testing on push
* Add in distributed tracing for observability. My current favorite tool for this is Honeycomb so I likely would have done this with their Go Beeline.
* Write integration or e2e tests to build confidence that the full process still works when changes are being made. I found unit tests to be valuable in writing the individual normalization functions and they are quicker/less expensive to write than an e2e test so that's why I chose to focus effort there to start with.
* If the above had been possible I also would have liked to include a Github Action that builds the binary and executes the e2e test before merging a feature branch into main
