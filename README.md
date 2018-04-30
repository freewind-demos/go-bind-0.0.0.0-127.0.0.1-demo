Go Net Binding Demo
===================

Question about binding `0.0.0.0:9999` and `127.0.0.1:9999`

```
brew install go
go run hello.go
```

You can try different address to connect the port `9999`, to see the response:

- `telnet 127.0.0.1 9999`
- `telnet localhost 9999`
- `telnet 0.0.0.0 9999`
- `telnet <my-internal-ip> 9999`, where `my-internal-ip` can be found by `ifconfig | grep "inet " | grep -Fv 127.0.0.1 | awk '{print $2}'`

And since the code is quite simple, you will have to restart it after each try.

It's strange for me to understand the response, since:

- `telnet 127.0.0.1 9999`: `127.0.0.1` (OK)
- `telnet localhost 9999`: `0.0.0.0` (WHY?!)
- `telnet 0.0.0.0 9999`: `127.0.0.1` (WHY?!)
- `telnet <my-internal-ip> 9999`: `0.0.0.0` (OK)

I'm confused about the responses of `telnet localhost 9999` and `telnet 0.0.0.0 9999`

Question: <https://stackoverflow.com/questions/50096683/puzzled-about-the-telnet-localhost-and-telnet-0-0-0-0>