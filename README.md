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

## Getting available parameters:

You can get available parameters from `/` as well as from `/config_name`

```shell
curl http://localhost:3333/
```

Response:

```json
{
  "/": {
    "/1": {
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
      "fifth_param": {
        "multiplier": 2,
        "parameter_type": "geometric sequence",
        "value": 1
      },
      "first_param": "first_param_value",
      "fourth_param": {
        "increment": 1,
        "parameter_type": "arithmetic sequence",
        "value": 0
      },
      "ninth_param": {
        "field1": "value 1",
        "field2": true,
        "field3": 1
      },
      "second_param": true,
      "seventh_param": {
        "parameter_type": "sequential selection",
        "values": [
          "one",
          "two",
          "three"
        ]
      },
      "sixth_param": {
        "max": 10,
        "min": 0,
        "parameter_type": "random"
      },
      "tenth_param": 3.1415926,
      "third_param": 3,
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
    },
    "/second/2": {
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
      "fifth_param": {
        "multiplier": 2,
        "parameter_type": "geometric sequence",
        "value": 1
      },
      "first_param": "first_param_value",
      "fourth_param": {
        "increment": 1,
        "parameter_type": "arithmetic sequence",
        "value": 0
      },
      "ninth_param": {
        "field1": "value 1",
        "field2": true,
        "field3": 1
      },
      "second_param": true,
      "seventh_param": {
        "parameter_type": "sequential selection",
        "values": [
          "one",
          "two",
          "three"
        ]
      },
      "sixth_param": {
        "max": 10,
        "min": 0,
        "parameter_type": "random"
      },
      "tenth_param": 3.1415926,
      "third_param": 3,
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
  }
}
```

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
  "error": "failed to set parameter_name to 1 due to type mismatch (got float64, expected string)"
}
```

## Extending selection parameters:

To extend selection parameters values it is enough to send POST request with value or array of values to be added to a selection

```shell
curl -X POST \
  http://localhost:3333/path/parameter_name \
  -H 'content-type: application/json' \
  -d '{"value": ["additional selection option 1", "additional selection option 2"]}'
```

Success response:

```json
{
  "result": "OK"
}
```

## Adding new parameter:

Updates are done via sending POST request to the config files from where parameters are retrieved. Data sent in post
request should contain data in same format as in json configuration files (which also means that multiple new parameters
may be created with a single request)

```shell
curl -X POST \
  http://localhost:3333/path/config \
  -H 'content-type: application/json' \
  -d '{"first": {"parameter_type": "arithmetic sequence","value": 0,"increment": 1}}'
```

Note: if `strict` mode is activated type of field can not change, e.g. if `parameter_name` was used to store `string`
data it will not be possible to set it to `number`

Success response:

```json
{
  "result": "OK"
}
```

Possible failed responses:

```json
{
  "error": "unexpected end of JSON input"
}
```

```json
{
  "error": "http: multiple registrations for /path/config/parameter_name"
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
- [x] Make big parameter refactor so that code does not stink so much
- [x] Ability to view available parameters in file
- [x] Ability to add parameters in runtime
- [x] Ability to add additional selection options for selection parameters in runtime
- [x] Make code actually Golang code and not try to panic everywhere
- [ ] Ability to save added parameters into `configs` directory for later use
- [ ] Rewrite server part so that code does not stink so much
- [ ] Add some Web UI and/or CLI for simpler use and navigation
- [ ] Automatic type deduction for complex parameters (is it even necessary?)
- [ ] Ability to change simple parameter to selection / sequence and vice versa
