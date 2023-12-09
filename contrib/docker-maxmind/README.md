# Getting started

## Requirements

- Docker
- Docker Compose
- MaxMind License Key. Get it from https://www.maxmind.com/en/my_license_key

## Building the image

Get the latest version number from https://github.com/superstes/geoip-lookup-service/releases

Update the version number in `Dockerfile`

Run


    docker compose build

## Running the image

Update `geoip.env` with your Account ID and License Key

Run


    docker compose up -d

Check the image status and logs


    docker logs -f geoip-update
    docker logs -f geoip-lookup-service

Example logs from `geoip-update`


    # STATE: Running geoipupdate
    # STATE: Sleeping for 72 hours

Example logs from `geoip-lookup-service`


    Version: 1
    by Superstes (GPLv3)
    
    Listening on http://0.0.0.0:10069

(Optional) Test if the service is working (Requires `curl`)


    curl "http://127.0.0.1:10069/?lookup=country&ip=1.1.1.1"
Expected Output


    {"registered_country":{"geoname_id":2077456,"iso_code":"AU","names":{"de":"Australien","en":"Australia","es":"Australia","fr":"Australie","ja":"オーストラリア","pt-BR":"Austrália","ru":"Австралия","zh-CN":"澳大利亚"}}}

# Removing the service

Run


    docker compose down

(Optional) Remove data volume


    docker volume rm geoip-lookup-service_geoipupdate_data