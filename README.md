### waitfor

Small golang utility I've produced to block until a host is pingable.

Designed to be used in expressions such as `waitfor 192.168.1.2 60 && /bin/dosomething`


#### Usage

Usage: waitfor <ip_addr> <timeout_secs>

It sends one ping in a loop until either a response is received, or <timeout> seconds have elapsed with no response.
It ends with an exit code 0 if a response was received, otherwise exit code 1.

#### Installation

##### MacOS and Linux

After deploying the binary to a suitable directory like `/usr/local/bin`, ensure the permissions are correct.

*Example* : `sudo chmod u=rwx,go=rx /usr/local/bin/waitfor`

##### Then for Linux only...


Also as root, use `setcap` to allow the binary to use raw sockets so that it can send and receive ICMP packets.

*Example* : `sudo setcap cap_net_raw=+ep /usr/local/bin/waitfor`