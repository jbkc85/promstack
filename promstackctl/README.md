# PromStackCTL

## Adding an Exporter

```sh
curl -XPUT consul.endpoint:8500/v1/kv/promstack/exporters/node-exporter -d '{"port":9100,"tags":["exporter"]}'
```
