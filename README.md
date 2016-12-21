# Demo for a [funker-go](https://github.com/bfirsh/funker-go) issue

[![Build Status](https://travis-ci.org/AkihiroSuda/demo-funker-issue.svg?branch=master)](https://travis-ci.org/AkihiroSuda/demo-funker-issue)

Steps:

 * Create a swarm cluster (just a single node is enough)
 * Execute `docker build -t akihirosuda/demo-funker-issue .`.
 * Execute `./test.sh 55s`.
 * Execute `./test.sh 65s`.

In my experiment, the result is:

| Duration |Result              |
|----------|--------------------|
| 55s      |Pass                |
| 65s      |Hangs forever       |

System:

 * Docker 1.14.0-dev ([2cf32ee](https://github.com/docker/docker/commit/2cf32ee04360add15b4431ba55032824ac089349))
 * Ubuntu 16.04
 * Kernel 4.4.0-53-generic #74-Ubuntu

Even reproducible on Travis CI (Ubuntu 14.04) plus Docker 1.13.0-rc4
