# Tornimo Grafana Datasource

[Tornimo](https://tornimo.io) is a real-time monitoring platform, which provides graphite-compatible API with many additional functions, stability and scalability.

Datasource is forked from Graphite datasource located in [Grafana](https://github.com/grafana/grafana) repo.

For any suggestions or bug reports post to [tornimo github](https://github.com/tornimo/tornimo-grafana-datasource/issues).

## Installation

To install this plugin you need to add artifact to your grafana's plugins directory, and then restart grafana server. Steps:
- SSH to your grafana server
- Go to grafana plugins directory(default=`/var/lib/grafana/plugins`)
- Download tornimo-datasource zip file from [releases](https://github.com/tornimo/tornimo-grafana-datasource/releases).
`wget https://github.com/tornimo/tornimo-grafana-datasource/releases/download/<version>/tornimo-datasource.zip`
- unpack it here
`unzip tornimo-datasource.zip`
- remove no longer needed archive
`rm tornimo-datasource.zip`
- restart grafana

### Docker
To run docker with this plugin installed you should run:

```
docker run \
    -d -p 3000:3000 \
    --name=grafana \
    -e "GF_INSTALL_PLUGINS=https://github.com/tornimo/tornimo-grafana-datasource/releases/download/<version>/tornimo-datasource.zip;tornimo-datasource" \
    grafana/grafana
```

## How to use
- Open grafana
- Go to configuration -> Data sources -> Add data source
- Choose Tornimo
- As URL put your Grafana backend URL (you can find it in your admin panel)
- As Access put Server
- As Token put your Tornimo's api key (you can generate it in your admin panel)
- choose Bucket size (60s is default)
