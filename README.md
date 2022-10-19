# Config Master

This software runs as a standalone web server, which could be supplied with different configuration files, all of which
should be place info `configs` folder

Supported config values types:

- Simple types: `int`, `float`, `bool`, `string`, `json`
- Selection types: `sequential selection`, `random selection`
- Sequence types: `arithmetic sequence`, `geometric sequence`

Usage:

```shell
make run
```

Or just use Docker image:

```shell
docker build -t config_master .
docker run -p 3333:3333 config_master
```

If used in Docker configs can be mounted to directory: `/config_master/configs`

Configs may look like this:

```json
{
  "first_param": {
    "type": "string",
    "value": "first_param_value"
  },
  "second_param": {
    "type": "bool",
    "value": true
  },
  "third_param": {
    "type": "int",
    "value": 3
  },
  "fourth_param": {
    "type": "arithmetic sequence",
    "value": 0,
    "increment": 1
  },
  "fifth_param": {
    "type": "geometric sequence",
    "value": 1,
    "multiplier": 2
  },
  "sixth_param": {
    "type": "random",
    "min": 0,
    "max": 10
  },
  "seventh_param": {
    "type": "sequential selection",
    "values": [
      "one",
      "two",
      "three"
    ]
  },
  "eighth_param": {
    "type": "random selection",
    "values": [
      "one",
      "two",
      "three"
    ]
  },
  "ninth_param": {
    "type": "json",
    "value": {
      "field1": "value 1",
      "field2": true,
      "field3": 1
    }
  },
  "tenth_param": {
    "type": "float",
    "value": 3.1415926
  },
  "eleventh_param": {
    "type": "sequential selection",
    "values": [
      1,
      2,
      3
    ]
  },
  "twelfth_param": {
    "type": "sequential selection",
    "values": [
      {
        "one": 1
      },
      2,
      3
    ]
  }
}
```