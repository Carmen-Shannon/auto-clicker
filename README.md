To build for a specific OS and architecture, you can pass the OS_ARCH variable to the make command. For example:

To build for Windows (amd64):
`make OS_ARCH=windows/amd64`

To build for Linux (amd64):
`make OS_ARCH=linux/amd64`

To build for macOS (amd64):
`make OS_ARCH=darwin/amd64`

Find your build for your system in the `build` directory, run the application.
