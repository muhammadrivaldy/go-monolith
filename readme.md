# Backend Service

## How to Running

Set up variable environment on your device.
export PATH_CONF = "config"
export FILE_CONF = "local.conf"

When you want to start this service, set up this aliases.
alias laundry-run = "gin --port 8080 --appPort 8081 --path . --build ./app run main.go"
and running with alias laundry-run
