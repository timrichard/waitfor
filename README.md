### waitfor

Small golang utility I've produced to block until a host is pingable.
Designed to be used in expressions such as ```waitfor 192.168.1.2 60 && /bin/dosomething```


#### Usage

Usage: waitfor <ip_addr> <timeout_secs>

It sends one ping in a loop until either a response is received, or <timeout> seconds have elapsed with no response.
It ends with an exit code 0 if a response was received, otherwise exit code 1.
