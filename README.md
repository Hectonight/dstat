# dstat

## About
dstat is a command line tool to help you quickly find statistics pertaining to a set of data.
You can either you a file or stdin as input.

## Installation
Download the binary found [here](https://github.com/Hectonight/dstat/releases/tag/1.0.0)
and add it to your PATH.

## Usage
```
Usage: dstat [OPTIONS]... [FILES]...
josh@josh-XPS-15-9530:~/GolandProjects/dstat$ dstat -h | xclip -sel clip
Usage: dstat [OPTIONS]... [FILES]...
If FILES is not set read from standard input.
  -n, --count               Size of the data set.
  -q, --first-quartile      Find the first quartile of the data.
  -h, --help                Display this message.
  -I, --ignore string       Ignore characters in this string.
  -i, --iqr                 Find the IQR of the data.
      --max                 Find the maximum of the data.
      --mean                Find the mean of the data.
      --median              Find the median of the data.
      --min                 Find the minimum of the data.
  -p, --range               Find the range of the data.
      --stdev               Find the standard deviation of the data while treating it as a sample.
      --stdevp              Find the standard deviation of the data while treating it as a population.
      --sum                 Find the sum of the data
  -s, --summary             Give a 5 number summary of the data
  -Q, --third-quartile      Find the third quartile of the data.
      --var                 Find the variance of the data while treating it as a sample.
      --varp                Find the variance of the data while treating it as a population.
  -W, --whitespace string   Treat characters in this string as whitespace.
```