# gopp

Just a simple map[string]interface{} parsed as json as a package because I'm lazy. Transforms this

```go
&{File: Log:{Level:info Pretty:true Color:true} Debug:{Addr:0.0.0.0:9104 Token: Pprof:false Zpages:false} HTTP:{Addr:localhost:9100 Root:/ Namespace:com.owncloud.web} Tracing:{Enabled:false Type:jaeger Endpoint: Collector: Service:phoenix} Asset:{Path:} OIDC:{MetadataURL: Authority: ClientID: ResponseType: Scope:} Phoenix:{Path: Config:{Server:http://localhost:9135 Theme:owncloud Version:0.1.0 OpenIDConnect:{MetadataURL:http://localhost:9130/.well-known/openid-configuration Authority:http://localhost:9130 ClientID:phoenix ResponseType:code Scope:openid profile email} Apps:[draw-io files markdown-viewer media-viewer pdf-viewer] ExternalApps:[]}}}
```

into this

```json
{
  "File": "",
  "Log": {
    "Level": "info",
    "Pretty": true,
    "Color": true
  },
  "Debug": {
    "Addr": "0.0.0.0:9989",
    "Token": "",
    "Pprof": false,
    "Zpages": false
  },
  "HTTP": {
    "Addr": "localhost:9100",
    "Root": "/",
    "Namespace": "com.test"
  },
  "Tracing": {
    "Enabled": false,
    "Type": "jaeger",
    "Endpoint": "",
    "Collector": "",
    "Service": "phoenix"
  },
  "Asset": {
    "Path": ""
  },
  "OIDC": {},
  "Phoenix": {
    "Path": "",
    "Config": {
      "server": "http://localhost:9135",
      "theme": "owncloud",
      "version": "0.1.0",
      "openIdConnect": {
        "metadata_url": "http://localhost:9130/.well-known/openid-configuration",
        "authority": "http://localhost:9130",
        "client_id": "phoenix",
        "response_type": "code",
        "scope": "openid profile email"
      },
      "apps": [
        "draw-io",
        "files",
        "markdown-viewer",
        "media-viewer",
        "pdf-viewer"
      ]
    }
  }
}
```
