# Litewallet SDK

## Install gomobile

``` plain
// install

$ go get -v golang.org/x/mobile/cmd/gomobile

$ gomobile init
```

## Build

``` plain
// bind

$ cd github.com/QOSGroup/litewallet/

$ gomobile bind --target=android ./litewallet/mobile/
```

## API

``` sh
// sdk_apis_cosmos.go

// CosmosGetAllValidators

$ ./gaiad query staking validators -o json
$ ./gaiad query staking validators -o json --limit 200 --page 3
```

## too large could not be uploaded to github, use internal downloading server instead:

- iOS  
  framwork:vagrant@192.168.1.176:/home/vagrant/mobile_frameworks/framework.tar.gz  

- Android  
  vagrant@192.168.1.176:/home/vagrant/mobile_frameworks/  
