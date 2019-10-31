# gofwc (Go Web Frameworks Comparison)

This is the repository that contains the code for my Master's Thesis in University of Macedonia with title **"Distributed Computing in Go: Comparison and Benchmarking of Web Applications Frameworks"**

## Setup
In order to run the code provided, you should have installed at least the version ```go1.12``` of the [Go language](https://golang.org/dl) in your computer . 

The project should be included in the ```src``` directory of your ```GOPATH```.

The frameworks' source code has not been included in this repository. You can download and install them by visiting their websites and following the installation instructions:
- Beego [https://beego.me/](https://beego.me/)
- Buffalo [https://gobuffalo.io](https://gobuffalo.io)
- Echo [https://echo.labstack.com](https://echo.labstack.com)
- Gin [https://gin-gonic.com](https://gin-gonic.com)
- Iris [https://iris-go.com](https://iris-go.com)

All the subprojects store the application data in an SQLite3 database. To be able to access the database through the application you should also install an SQLite3 driver. I installed the [go-sqlite3](https://github.com/mattn/go-sqlite3) driver. Please refer to this repository's installation instructions.
