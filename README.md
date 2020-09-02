# dobind
CLI for getting zone records in BIND format from DigitalOcean.

### Installation

Grab the latest binary from [releases](https://github.com/pellizzetti/digitalocean-dns-bind-export/releases) page for your architecture and place it somewhere on your `PATH`.

### Usage

Set the access token (get one [here](https://cloud.digitalocean.com/settings/api/tokens)):

```sh
$ export DO_ACCESS_TOKEN="WqvcUVyGEzmMFo28ukKRVFW8rqzAQY9uijSpaPrB8mCXZE6wQ7csNrRa6ZaHHBgo"
```

Extract BIND information for the zone `example.com` and write to file `example.com.zone`:

```sh
$ dobind -z example.com > example.com.zone
```

output:

```bind
$ORIGIN example.com.
@                      3600 SOA   ns1.p30.dynect.net. (
                              zone-admin.dyndns.com.     ; address of responsible party
                              2016072701                 ; serial number
                              3600                       ; refresh period
                              600                        ; retry period
                              604800                     ; expire time
                              1800                     ) ; minimum ttl
                      86400 NS    ns1.p30.dynect.net.
                      86400 NS    ns2.p30.dynect.net.
                      86400 NS    ns3.p30.dynect.net.
                      86400 NS    ns4.p30.dynect.net.
                       3600 MX    10 mail.example.com.
                       3600 MX    20 vpn.example.com.
                       3600 MX    30 mail.example.com.
                         60 A     204.13.248.106
                       3600 TXT   "v=spf1 includespf.dynect.net ~all"
mail                  14400 A     204.13.248.106
vpn                      60 A     216.146.45.240
webapp                   60 A     216.146.46.10
webapp                   60 A     216.146.46.11
www                   43200 CNAME example.com.
```