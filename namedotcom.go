package namedotcom

import (
	"github.com/libdns/namedotcom"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

// Provider wraps the provider implementation as a Caddy module.
type Provider struct{ *namedotcom.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.namedotcom",
		New: func() caddy.Module {
			return &Provider{new(namedotcom.Provider)}
		},
	}
}

// Provision implements the Provisioner interface to initialize the Namedotcom client
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()
	p.Provider.APIToken = repl.ReplaceAll(p.Provider.APIToken, "")
	p.Provider.User = repl.ReplaceAll(p.Provider.User, "")
	p.Provider.APIUrl = repl.ReplaceAll(p.Provider.APIUrl, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// namedotcom  {
//     api_token <string>
//     user <string>
//     endpoint <string>
// }
//
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {

		if d.NextArg() {
			return d.ArgErr()
		}

		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_token":
				if d.NextArg() {
					p.Provider.APIToken = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "user_name":
				if d.NextArg() {
					p.Provider.User = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "api_url":
				if d.NextArg() {
					p.Provider.APIUrl = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}

	if p.APIToken == "" {
		return d.Err("field api_token cannot be empty")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
