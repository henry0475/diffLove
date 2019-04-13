#!/bin/bash

go get -v "github.com/henry0475/diffLove/apis/dlAuthAPI"

go get -v "github.com/henry0475/diffLove/services/dlAuth"


nohup consul agent -dev > ~/dl_nohup/consul.nohup &
nohup micro --enable_stats --enable_tls --tls_cert_file=/root/fullchain.pem --tls_key_file=/root/privkey.pem api --address=0.0.0.0:443 --handler=http --namespace=zone.alumnus.api > ~/dl_nohup/apis.nohup &

# Services
nohup dlAuth > ~/dl_nohup/dlAuth.nohup &

# APIs
nohup dlAuthAPI > ~/dl_nohup/dlAuthAPI.nohup &
