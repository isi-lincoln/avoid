# Running a docker-compose test

The idea here is to have a small environment that can test just the dns switching capability.

Components:
- etcd
- coredns
- dns-service
- gw(0/1)
- ue

Purpose:

`etcd` is our data store for storing data in a persistent key/value data store.

`coredns` is running coredns with the avoid plugin.
The way this works is that we load in a `Corefile` which tells coredns how to operate.
Running `git submodule update --init --recursive` will pull our version of coredns with our plugins.
This is the same version that is running in the docker container.
The first section of the file has something along the lines of `example.com:53`.
Meaning any requests targeting this domain, follow a set of actions.
In our case the first action is to use the avoid plugin, which passes in an argument `avoid dns-service`.
This value (currently not used) would tell coredns where the dns backend is located at.
This value is currentlt hardcoded to `avoid`.
So a static `/etc/hosts` value for avoid would allow this to work more "programmically".
The coredns code will then use a grpc protobuf service to request from that server any dns entries.

`dns-service` is our code in this repo (`dns/main.go`).
It is responsible for updating etcd from the grpc protobuf service.

`gw0/1` act as our dummies to indicate moving dns/servers.

`ue` is our endpoint which will be requesting from coredns a server.


## Commands

First on gw0, gw1, ue, coredns (the `overlay interface`) we need to get ip addr information:

`coredns: 172.22.0.2`
`gw0: 172.22.0.4`
`gw1: 172.22.0.5`
`ue: 172.22.0.3`

Then on `dns-service`:

```
root@0686e855f989:/# avoid-dns-cli update cli 172.22.0.3 gw.example.com 0 ATEST --a 172.22.0.4 --aaaa aaaa::0001
INFO[0000] In check record                              
INFO[0000] check record okay                            
test cli
sent request: entries:{ue:"172.22.0.3"  name:"gw.example.com"  arecords:"172.22.0.4"  aaaarecords:"aaaa::0001"  txt:"ATEST"}
code:1
```

So now we have a dns entry specifically for ue, when it requests dns entries for `gw.example.com` it gets `172.22.0.4`.

What we should see BEFORE we add an entry:

```
root@2431005e448e:/# dig gw.example.com @dns

; <<>> DiG 9.18.24-0ubuntu0.22.04.1-Ubuntu <<>> gw0.example.com @dns
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: SERVFAIL, id: 6317
;; flags: qr rd; QUERY: 1, ANSWER: 0, AUTHORITY: 0, ADDITIONAL: 1
;; WARNING: recursion requested but not available

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 1232
; COOKIE: 875d61e66b1aee9a (echoed)
;; QUESTION SECTION:
;gw.example.com.               IN      A

;; Query time: 8 msec
;; SERVER: 172.22.0.2#53(dns) (UDP)
;; WHEN: Wed Jun 12 16:57:05 UTC 2024
;; MSG SIZE  rcvd: 56
```

```
root@2431005e448e:/# dig gw.example.com @dns

; <<>> DiG 9.18.24-0ubuntu0.22.04.1-Ubuntu <<>> gw.example.com @dns
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 31903
;; flags: qr aa rd; QUERY: 1, ANSWER: 2, AUTHORITY: 0, ADDITIONAL: 1
;; WARNING: recursion requested but not available

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 1232
; COOKIE: 76f3f7dd85a2fd36 (echoed)
;; QUESTION SECTION:
;gw.example.com.                        IN      A

;; ANSWER SECTION:
gw.example.com.         0       IN      A       172.22.0.4
gw.example.com.         0       IN      AAAA    aaaa::1

;; Query time: 8 msec
;; SERVER: 172.22.0.2#53(dns) (UDP)
;; WHEN: Wed Jun 12 17:22:15 UTC 2024
;; MSG SIZE  rcvd: 127
```

Our logs will show something similar:

```
coredns_1      | time="2024-06-12T16:57:05Z" level=info msg="avoid: Received query gw.example.com. from 172.22.0.3\n"
coredns_1      | time="2024-06-12T16:57:05Z" level=info msg="avoid: Rewritten query gw.example.com\n"
coredns_1      | time="2024-06-12T16:57:05Z" level=info msg="avoid: requesting: 172.22.0.3/gw.example.com from dns-service:9000"
dns-service_1  | time="2024-06-12T16:57:05Z" level=debug msg="Lookup DNS Item" Name=gw.example.com Ue=172.22.0.3
coredns_1      | time="2024-06-12T16:57:05Z" level=error msg="%s: Show(): %vavoidrpc error: code = Unknown desc = 172.22.0.3/gw.example.com: Show(): not found"
coredns_1      | [ERROR] Recovered from panic in server: "dns://:53" runtime error: invalid memory address or nil pointer dereference
dns-service_1  | time="2024-06-12T17:20:48Z" level=info msg=Update A="[172.22.0.4]" AAAA="[aaaa::0001]" Identity=172.22.0.3 Index=0 Name=gw.example.com
dns-service_1  | time="2024-06-12T17:20:48Z" level=info msg="In check record\n"
dns-service_1  | time="2024-06-12T17:20:48Z" level=info msg="check record okay\n"
dns-service_1  | time="2024-06-12T17:22:06Z" level=info msg="List DNS Entry Keys"
coredns_1      | time="2024-06-12T17:22:15Z" level=info msg="avoid: Received query gw.example.com. from 172.22.0.3\n"
coredns_1      | time="2024-06-12T17:22:15Z" level=info msg="avoid: Rewritten query gw.example.com\n"
coredns_1      | time="2024-06-12T17:22:15Z" level=info msg="avoid: requesting: 172.22.0.3/gw.example.com from dns-service:9000"
dns-service_1  | time="2024-06-12T17:22:15Z" level=debug msg="Lookup DNS Item" Name=gw.example.com Ue=172.22.0.3
dns-service_1  | time="2024-06-12T17:22:15Z" level=debug msg="Show Found" Entry="ue:\"172.22.0.3\"  name:\"gw.example.com\"  arecords:\"172.22.0.4\"  aaaarecords:\"aaaa::0001\"  txt:\"ATEST\"  version:1"
coredns_1      | time="2024-06-12T17:22:15Z" level=info msg="avoid: Response: [gw.example.com.\t0\tIN\tA\t172.22.0.4 gw.example.com.\t0\tIN\tAAAA\taaaa::1]"
coredns_1      | [INFO] 172.22.0.3:59724 - 31903 "A IN gw.example.com. udp 55 false 1232" NOERROR qr,aa,rd 104 0.006601697s
```

Now, lets change it to gw1.


On `dns-service`:

```
root@0686e855f989:/# avoid-dns-cli update cli 172.22.0.3 gw.example.com 0 ATEST --a 172.22.0.5 --aaaa aaaa::0002
INFO[0000] In check record                              
INFO[0000] check record okay                            
test cli
sent request: entries:{ue:"172.22.0.3"  name:"gw.example.com"  arecords:"172.22.0.5"  aaaarecords:"aaaa::0002"  txt:"ATEST"}
code:1
```

And then on `ue`:

```
root@2431005e448e:/# dig gw.example.com @dns

; <<>> DiG 9.18.24-0ubuntu0.22.04.1-Ubuntu <<>> gw.example.com @dns
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 29248
;; flags: qr aa rd; QUERY: 1, ANSWER: 2, AUTHORITY: 0, ADDITIONAL: 1
;; WARNING: recursion requested but not available

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 1232
; COOKIE: de5ff7b6b8bc5462 (echoed)
;; QUESTION SECTION:
;gw.example.com.                        IN      A

;; ANSWER SECTION:
gw.example.com.         0       IN      A       172.22.0.5
gw.example.com.         0       IN      AAAA    aaaa::2

;; Query time: 4 msec
;; SERVER: 172.22.0.2#53(dns) (UDP)
;; WHEN: Wed Jun 12 17:24:45 UTC 2024
;; MSG SIZE  rcvd: 127
```

And our logs show the same:

```
dns-service_1  | time="2024-06-12T17:24:36Z" level=info msg=Update A="[172.22.0.5]" AAAA="[aaaa::0002]" Identity=172.22.0.3 Index=0 Name=gw.example.com
dns-service_1  | time="2024-06-12T17:24:36Z" level=info msg="In check record\n"
dns-service_1  | time="2024-06-12T17:24:36Z" level=info msg="check record okay\n"
coredns_1      | time="2024-06-12T17:24:45Z" level=info msg="avoid: Received query gw.example.com. from 172.22.0.3\n"
coredns_1      | time="2024-06-12T17:24:45Z" level=info msg="avoid: Rewritten query gw.example.com\n"
coredns_1      | time="2024-06-12T17:24:45Z" level=info msg="avoid: requesting: 172.22.0.3/gw.example.com from dns-service:9000"
dns-service_1  | time="2024-06-12T17:24:45Z" level=debug msg="Lookup DNS Item" Name=gw.example.com Ue=172.22.0.3
dns-service_1  | time="2024-06-12T17:24:45Z" level=debug msg="Show Found" Entry="ue:\"172.22.0.3\"  name:\"gw.example.com\"  arecords:\"172.22.0.5\"  aaaarecords:\"aaaa::0002\"  txt:\"ATEST\"  version:2"
coredns_1      | time="2024-06-12T17:24:45Z" level=info msg="avoid: Response: [gw.example.com.\t0\tIN\tA\t172.22.0.5 gw.example.com.\t0\tIN\tAAAA\taaaa::2]"
coredns_1      | [INFO] 172.22.0.3:40384 - 29248 "A IN gw.example.com. udp 55 false 1232" NOERROR qr,aa,rd 104 0.004800326s
```
