# GPX compressor

> The purpose of this utility is to reduce the load on rendering components, such as [Leaflet](https://leafletjs.com/), when they render multiple files simultaneously.

## Compressing

```bash
gpxcomopress <name> <new_name>  <epsilon_value>
```

Where:
**name** - gpx name
**new_name** - compressed gpx (output)
**epsilon_value** - compression coefficient (50m accuracy - approximately 0.0000045 )

## Batch compressing

```bash
find YOUR_FOLDER -type f -name "*.gpx" -exec sh -c 'go run . "$0" NEW_FOLDER/c_$(basename "$0") 0.0000045' {} \;
```

## Credits

gpxcompressor is based on:

[https://github.com/tkrajina/gpxgo](https://github.com/tkrajina/gpxgo)

## License

gpxcompressor is licensed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)