# Forecast
A weather forecast written in Go ready to deploy on google cloud functions.

## Resources
- [Apixu](https://www.apixu.com/), free, fast, simple and fully managed JSON and XML **Weather API** and Geo Developer API
- [weather-icons](https://erikflowers.github.io/weather-icons/), **Weather Icons** is the only icon font and CSS with 222 weather themed icons, ready to be dropped right into Bootstrap, or any project that needs high quality weather, maritime, and meteorological based icons!
- [Google Cloud Functions](https://cloud.google.com/functions/docs/), **Serverless application** backends. Trigger your code from GCP services or call it directly from any web, mobile, or backend application.

## Configure
- Create a [Apixu](https://www.apixu.com/signup.aspx) account, free plan offers 10.000 API calls per month. **Get the API Key**
- Log on [Google cloud functions console](https://console.cloud.google.com/functions),then create a project and 2 cloud functions.

 Name    | Runtime | Memory allocated | Executed function | Environment variables    |
---------|---------|------------------|-------------------|--------------------------|
index    | Go 1.11 | 128 MB 	        | ForecastAPIV1     |                          |
forecast | Go 1.11 | 128 MB 	        | GetIndex          | apixu_key = **YOUR KEY** |

You can upload the **go code** to each cloud function or you can attach to a [Google source cloud](https://cloud.google.com/source-repositories/) repository.
