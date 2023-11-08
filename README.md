# go-ddns
A GoDaddy (D)DNS client in Go

* Motivation

I want a simple client that I could (say) run out of a systemd timer
to update one or more GoDaddy domain A/CNAME records with whatever my
current IP address might be.

* Status

I can get my current IP.

* Plans

0. Implement YAML config
1. Implement get/set A records
2. Implement get/set CNAME records
3. Implement alternative sources for ipinfo (i.e. what if ipinfo.io goes away)

* Reference

https://developer.godaddy.com/doc/endpoint/domains
