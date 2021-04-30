
name.com module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with name.com.

## Caddy module name

```
dns.providers.namedotcom
```

## Authenticating

See [the associated README in the libdns package](https://github.com/libdns/namedotcom) for important information about credentials.

## Building

To compile this Caddy module, follow the steps describe at the [Caddy Build from Source](https://github.com/caddyserver/caddy#build-from-source) instructions and import the `github.com/caddy-dns/namedotcom` plugin
## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
    "module": "acme", 
      "challenges": {
          "dns": {
              "provider": {
                 "name":   "namedotcom",
                 "token":  "{env.NAMEDOTCOM_TOKEN}",
                 "user":   "{env.NAMEDOTCOM_USER}",
                 "server": "{env.NAMEDOTCOM_SERVER}"
              }
          }
      }
}
```

or with the Caddyfile:

```
tls {
  dns namedotcom {
    token {$NAMEDOTCOM_TOKEN}
    server {$NAMEDOTCOM_SERVER}
    user {$NAMEDOTCOM_USER}
  }
}
```