# Litewallet SDK

## Install gomobile

``` plain
$ go get -v golang.org/x/mobile/cmd/gomobile

$ gomobile init
```

## Build

``` plain
$ cd github.com/QOSGroup/litewallet/

$ gomobile bind --target=android ./litewallet/mobile/
```

## API

``` sh
//get all the validators

$ ./gaiad query staking validators -o json
```

## too large could not be uploaded to github, use internal downloading server instead:

- iOS  
  framwork:vagrant@192.168.1.176:/home/vagrant/mobile_frameworks/framework.tar.gz  

- Android  
  vagrant@192.168.1.176:/home/vagrant/mobile_frameworks/  
