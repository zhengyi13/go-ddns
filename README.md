# go-ddns
A GoDaddy (D)DNS client in Go

## Motivation

I want a simple client that I could (say) run out of a systemd timer
to update one or more GoDaddy domain A/CNAME records with whatever my
current IP address might be.

There are certainly other projects I might have used:

* https://github.com/oze4/godaddygo/tree/master seems like a great
  project
* https://github.com/kryptoslogic/godaddy-domainclient if I were of a
  mind to do things the OpenAPI way

... but really, my needs are simple enough, and I need the learning
experience.

## Status

I can get my current IP.

## Plans

1. Implement YAML config
1. Implement get/set A records
1. Implement get/set CNAME records
1. Implement alternative sources for ipinfo (i.e. what if ipinfo.io
   goes away)

## Reference

https://developer.godaddy.com/doc/endpoint/domains
