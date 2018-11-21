# Host detect

This program aims to keep scaning a list of configured  hosts and ports (or default ports) and exports the results to Prometheus.

With this, you can track your hostnames for open ports that shouldn't be open, configure alarms and so on.

## How to configure your yaml

A very simple example of the accepted structure is:
```
timeout: 1s
interval: 1m0s
hosts:
- network: tcp
  address: www.google.com:80
- network: tcp
  address: www.google.com:23
- network: tcp
  address: zebeleo.com.br:80
```

`timeout` is the amount of time the code will wait until it gives up trying to connect.
Keep in mind that a port will be considered `closed` *only* when a timeout happens.
Otherwise, the state will be considered as `failure`.

`interval` is the time interval that the program follows between each cycle of checking the hosts.

The acceptable inputs for both `timeout` and `interval` are described in Golang's docs for [time.ParseDuration](https://golang.org/pkg/time/#ParseDuration).

`hosts` is where you describe which hosts you wish to check whether the port is open or not.
The acceptable inputs for network and address are described in Golang's docs for [net.Dial](https://golang.org/pkg/net/#Dial)

## How to run

Hostdetect uses Go modules, therefore it depends on Go 1.11 and newer.
Also, it means it doesn't necessarily have to be contained inside your `$GOPATH/src` folder.
Once you've cloned the repo locally, you may run `go mod verify` to make sure all dependencies are present.


## Kubernetes integration

TODO

## How to contribute

1. Fork the project & clone locally.
2. Create an upstream remote and sync your local copy before you branch.
3. Branch for each separate piece of work.
4. Do the work, write good commit messages, and read the CONTRIBUTING file if there is one.
5. Push to your origin repository.
6. Create a new PR in GitHub.
7. Respond to any code review feedback.

If you want to contribute and have fun with us, we'll be glad to have you onboard. 

