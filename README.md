# admiral-cli

Admiral Docker Registry administration and authentication system CLI

## Usage

```
Admiral CLI is the Admiral commander

Usage:
  admiral-cli [flags]
  admiral-cli [command]

Available Commands:
  image       Image command
  user        User command
  version     Print the version number of Admiral

Flags:
  -a, --address string    Admiral address (default "localhost")
  -P, --password string   Admiral password
  -p, --port int          Admiral port (default 3000)
  -U, --username string   Admiral username

Use "admiral-cli [command] --help" for more information about a command.
```

## Images
```
Image command

Usage:
  admiral-cli image [command]

Available Commands:
  delete      Delete the given image from the registry
  list        List images available in my namespace
  set-private Set the given image as private
  set-public  Set the given image as public (pull only)
```

## User
```
User command

Usage:
  admiral-cli user [command]

Available Commands:
  create      Creates a new user if it does not exist
```
