# cloud-cli
### Multi-Cloud CLI Tool

***Work In Progress***

#### Motivation

I use both AWS and GCP for various things, but mainly for uploading files on a regular basis to both S3 and GCP storage buckets.

I'm not particularly fond of the AWS CLI and the GCP one, I can never remember all the commands.

So decided to start writing this and at the same time play around with the Cobra CLI library. 

I wanted a tool that would allow me to use both S3 and GCP Buckets within one tool. 

For now, only AWS S3 is in the process of being implemented

#### Settings

Create a yaml file and place it in `$HOME/.cloud-cli-settings.yaml`   
Alternatively you can have the file anywhere you want but will have to pass the a `--config` flag for each command with the location of the file.

See the example `settings.yaml` file in this directory

#### Commands

`cloud-cli --help` for help

#### Installation

Run `make install` to install
