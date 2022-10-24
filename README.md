# Config Master

This software runs as a web server, which could be supplied with different configuration files, all of which
should be placed info `configs` folder. Parameters are served
under `http://address:port/config_file_path/config_file_name/parameter_name`

## Supported config values types:

- Simple types: `number`, `bool`, `string`, `json`, `array`
- Selection types: `sequential selection`, `random selection`
- Sequence types: `arithmetic sequence`, `geometric sequence`

## Supported flags:

`-address`, `--address` - sets server's address. Default: `""` (implies `localhost`)

`-port`, `--port` - sets server's port. Default: `3333`

`-config-dir`, `--config-dir` - sets server's config directory. Default: `./configs`

`-strict`, `--strict` - sets if parameters must stay of the same type, so that user can not change type by update.
Default: `false`

## How to run:

```shell
make run
```

Or just use Docker image:

```shell
docker build -t config_master .
docker run -p 3333:3333 config_master
```

If used in Docker configs can be mounted to directory: `/config_master/configs`

## Getting parameter value:

Path to parameter can be obtained via combining parameter's name and path to containing json. So having file structure
like this:

```
├── configs
│   ├── first.json
│   ├── additional
│   │   ├── second.json
```

We will be able to access parameters located in `first.json` via `address:port/first/parameter_name`. To access
parameters from `second.json` we would use `address:port/additional/second/parameter_name`

To retrieve parameter value simply send GET request to its path:

```shell
curl http://localhost:3333/path/parameter_name
```

Response:

```json
{
  "parameter_name": "parameter value"
}
```

Trying to access parameter which does not exist will result in `404 page not found`

## Updating existing parameter:

Updates are done via sending PUT request to the same address from where parameter is retrieved

```shell
curl -X PUT \
  http://localhost:3333/path/parameter_name \
  -H 'content-type: application/json' \
  -d '{"value": "new value of parameter"}'
```

Note: if `strict` mode is activated type of field can not change, e.g. if `parameter_name` was used to store `string`
data it will not be possible to set it to `number`

Success response:

```json
{
  "result": "OK"
}
```

Failed response:

```json
{
  "error": "failed to set parameter_name, error: could not parse [1 2 3], got type []interface {} but string was expected"
}
```

## Possible config structure:

```json
{
  "first_param": "first_param_value",
  "second_param": true,
  "third_param": 3,
  "tenth_param": 3.1415926,
  "ninth_param": {
    "field1": "value 1",
    "field2": true,
    "field3": 1
  },
  "fourth_param": {
    "parameter_type": "arithmetic sequence",
    "value": 0,
    "increment": 1
  },
  "fifth_param": {
    "parameter_type": "geometric sequence",
    "value": 1,
    "multiplier": 2
  },
  "sixth_param": {
    "parameter_type": "random",
    "min": 0,
    "max": 10
  },
  "seventh_param": {
    "parameter_type": "sequential selection",
    "values": [
      "one",
      "two",
      "three"
    ]
  },
  "eighth_param": {
    "parameter_type": "random selection",
    "values": [
      "one",
      "two",
      "three"
    ]
  },
  "eleventh_param": {
    "parameter_type": "sequential selection",
    "values": [
      1,
      2,
      3
    ]
  },
  "twelfth_param": {
    "parameter_type": "sequential selection",
    "values": [
      {
        "some nested": "data"
      },
      2,
      3
    ]
  }
}
```

## Milestones:

- [x] Add non-strict types mode (now ON by default)
- [x] Automatic type deduction for simple parameters to reduce boilerplate in config files
- [x] Add parameters recursively from `configs` directory
- [x] Ability to update parameters in runtime
- [ ] Make big refactor so that code does not stink so much
- [ ] Ability to add parameters in runtime
- [ ] Ability to save added parameters into `configs` directory for later use
- [ ] Automatic type deduction for complex parameters (is it even necessary?)
- [ ] Ability to change simple parameter to selection / sequence and vice versa
- [ ] Make code actually Golang code and not try to panic everywhere (ongoing)
