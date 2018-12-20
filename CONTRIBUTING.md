## Build image

```
docker build -t is2ei/drone-typetalk .
```

## Testing the plugin

```
docker run --rm \
  -e TYPETALK_TOKEN=xxxxxxxx \
  -e TOPIC_ID=12345 \
  -e MESSAGE="hello world" \
  is2ei/drone-typetalk
```