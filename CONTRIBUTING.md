## Build image

```
make build
```

## Testing the plugin

```
docker run --rm \
  -e PLUGIN_TYPETALK_TOKEN=xxxxxxxx \
  -e PLUGIN_TOPIC_ID=12345 \
  -e PLUGIN_MESSAGE="hello world" \
  is2ei/drone-typetalk
```