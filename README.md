# Forecast

A weather forecast written in Go ready to deploy on google cloud functions.

![Wearher Forecast](https://66.media.tumblr.com/2af2ae862c72fe07a1822becad3fced5/tumblr_inline_prpwbaW2QP1r3zk59_1280.png)

## Resources

- [Openweather](https://openweather.co.uk/weather-api), Simple and fast APIs. Access to current weather, forecasts, maps and historical data in JSON, XML, and HTML formats. A variety of map layers are available including precipitation, clouds, pressure, temperature, wind, and many more.
- [Mapquest](https://developer.mapquest.com/), Build great location-based experiences and delight your customers with MapQuest for Business' geospatial solutions: mapping, geocoding, directions & search.
- [weather-icons](https://erikflowers.github.io/weather-icons/), **Weather Icons** is the only icon font and CSS with 222 weather themed icons, ready to be dropped right into Bootstrap, or any project that needs high quality weather, maritime, and meteorological based icons!
- [Google Cloud Functions](https://cloud.google.com/functions/docs/), **Serverless application** backends. Trigger your code from GCP services or call it directly from any web, mobile, or backend application.

## Cloud Functions

- Create a [Openweather](https://openweather.co.uk/weather-api) account, free plan offers 1,000 calls per day. **Get the API Key**
- Create a [Mapquest](https://developer.mapquest.com/) account, free plan offers FREE 15,000 transactions per month. **Get the API Key**
- Log on [Google cloud functions console](https://console.cloud.google.com/functions),then create a project and 2 cloud functions.

 Name    | Runtime | Memory allocated | Executed function | Environment variables    |
---------|---------|------------------|-------------------|--------------------------|
index    | Go 1.11 | 128 MB          | ForecastAPIV1     |                          |
forecast | Go 1.11 | 128 MB          | GetIndex          | mapquest_key = **YOUR KEY** openweather_key = **YOUR KEY**|

You can upload the **go code** to each cloud function or you can attach to a [Google source cloud](https://cloud.google.com/source-repositories/) repository.

Also you can create cloud functions using [gcloud sdk](https://cloud.google.com/sdk/) command line.

```bash
gcloud functions deploy index --entry-point "GetIndex" --runtime go111 --trigger-http --memory 128
gcloud functions deploy forecast --entry-point "ForecastAPIV1" --runtime go111 --trigger-http --memory 128 --set-env-vars "mapquest_key=**YOUR KEY**,openweather_key=**YOUR KEY**"
```

Now you project is ready on **<https://us-central1-PROJECT.cloudfunctions.net/index>**
