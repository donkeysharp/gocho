Building Instructions
=====================

## Requirements
You need these tools installed in order to develop Gocho:

* Yarn (ui)
* NodeJS (ui)
* GNU Make (for general process building)
* Go (for main code compilation)
* go-bindata (for embedding UI inside final binary)

## Service Development
In order to run Gocho service and start development on it there are some things to consider: the next steps need to be run only the first time if you don't want to modify the UI.

#### Step 1
Create the UI files e.g. html, javascript and css

    $ make ui

It will create a `ui/build` directory with the resulting files for the UI

#### Step 2
As UI files are embedded inside the final binary, we use [go-binddata](https://github.com/jteeuwen/go-bindata) to achieve this. The `generate` command in the Makefile creates an `assets/assets_gen.go` file with the embedded UI code.

    $ make generate

So far the previous steps need to be run only the first time unless you are modifying UI, in that case check the `UI Development` section.

To build `gocho` binary and test it while you do changes run:

    $ make build-dev

Which will create the `gocho` binary at `$GOPATH/bin/gocho` as that command runs `go install github.com/donkeysharp/gocho/cmd/gocho`

## UI Development
Gocho UI uses React and was intialized using [Create React App](https://github.com/facebook/create-react-app).

All React code for the dashboard is located in the `ui` directory. This directory has the package.json and the yarn.lock file, in order to develop and test the UI you need to run the next

    $ cd ui
    # Install UI dependencies
    $ yarn install
    # Start development server
    $ yarn start

That will bring up a development server at http://localhost:3000 with the UI so it can be developed. The development server for UI is configured to use a proxy to `http://localhost:1337` (see `package.json`) so it's important to have a backend running, check `Service Development` section.
