# quick-port-scanner

Basic port scanning tool for network auditing and analysis. Written in Go.

## Getting Started

Quick scan of common ports on a single IPv4 target:

```
./qps -a -ip 192.168.1.1 -timeout 250
Scanning ports...
OPEN PORT: 192.168.1.1:53 DNS
OPEN PORT: 192.168.1.1:8200 DLNA
OPEN PORT: 192.168.1.1:80 HTTP
OPEN PORT: 192.168.1.1:21 FTP
OPEN PORT: 192.168.1.1:5000 UPnP,AirPrint
OPEN PORT: 192.168.1.1:631 IPP,CUPS
```

## Reference

* https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers

* https://nullsec.us/top-1-000-tcp-and-udp-ports-nmap-default/
