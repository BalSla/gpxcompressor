# GPX compressor

> The purpose of this utility is to reduce the load on rendering components, such as [Leaflet](https://leafletjs.com/), when they render multiple files simultaneously.

## To do
- [ ] Change date in GPX to today (by default) using data 
- [ ] Build output file name and path basing on date provided in parametes plus `root folder`
- [ ] Add `root folder` to config file
- [ ] Read `root folder` from config file
- [ ] Add to this file description how to register executable file in MacOS

## Compressing

```bash
gpxcomopressor -input <name> -output <new_name>  -epsilon <epsilon_value>
```

Where:
**input** - gpx name
**output** - compressed gpx (output)
**epsilon** - compression accuracy (50m accuracy - approximately 0.0000045 )

## Batch compressing

```bash
find YOUR_FOLDER -type f -name "*.gpx" -exec sh -c 'go run . "$0" NEW_FOLDER/c_$(basename "$0") 0.0000045' {} \;
```

## Credits

gpxcompressor is based on:

[https://github.com/tkrajina/gpxgo](https://github.com/tkrajina/gpxgo)

## License

gpxcompressor is licensed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)