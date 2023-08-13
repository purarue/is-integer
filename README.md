## is-integer

```
Usage: is-integer [-h] [-d] <number>
	-d: keep decimal part of number
    <number>: number to check if it is an integer

If it is an integer, it will be printed as such
If its a float and -d is not specified, the decimal part will be truncated

If the number is not a valid number, this will exit with a exit code of 1
```

If you want to use this in a script, you can do something like this:

```bash
read -r somevar  # ask user for input
if parsed="$(is-integer "$somevar")"; then
	some_other_program "$parsed"
else
    echo "Error: $somevar is not an integer" >&2
    exit 1
fi
```

I had initially tried to do this in pure bash like [described in the myriad of options here](https://stackoverflow.com/questions/806906/how-do-i-test-if-a-variable-is-a-number-in-bash), but those typically have some edge case I'm not a fan of. I also tried using perls `looks_like_number`, but the rounding when truncating floats was messy and would round half-to-even instead of just truncating.

So, I decided to write it in golang instead, so I could use a nice type system, and it handles all the edge cases I wanted nicely.

This uses the stdlib bigint package, so it can handle arbitrarily large integers without overflowing.

## Installation

Using `go install` to put it on your `$GOBIN`:

```
go get github.com/seanbreckenridge/is-integer
```

Manually:

```bash
git clone https://github.com/seanbreckenridge/is-integer
cd ./is-integer
go build .
# copy binary somewhere on your $PATH
sudo cp ./is-integer /usr/local/bin
```

## Examples

```bash
$ is-integer 5
5
$ is-integer 5.540
5
$ is-integer -d 5.540
5.540
$ is-integer -d -4.5
-4.5
$ is-integer -4.5
-4
$ is-integer ffff
$ echo $?
1
$ is-integer 555fff
$ echo $?
1
```
