# Config Master

This software runs as a web server, which could be supplied with different configuration files, all of which
should be placed info `configs` folder. Parameters are served
under `http://address:port/config_file_path/config_file_name/parameter_name`

## Supported config values types:

- Simple types: `number`, `bool`, `string`, `json`
- Selection types: `sequential selection`, `random selection`
- Sequence types: `arithmetic sequence`, `geometric sequence`

## Usage:

```shell
make run
```

Or just use Docker image:

```shell
docker build -t config_master .
docker run -p 3333:3333 config_master
```

If used in Docker configs can be mounted to directory: `/config_master/configs`

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

- [ ] Make code actually Golang code and not try to panic everywhere
- [x] Add non-strict types mode (now ON by default)
- [x] Automatic type deduction for simple parameters to reduce boilerplate in config files
- [x] Add parameters recursively from `configs` directory
- [ ] Ability to add parameters in runtime
- [ ] Ability to update parameters in runtime
- [ ] Ability to save added parameters into `configs` directory for later use
- [ ] Automatic type deduction for complex parameters (is it necessary?)