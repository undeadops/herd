# Herd

Because naming is hard.

AWS CLI for getting to the most needed info I use everyday 
but dont want to open the AWS Console to look up. 

Modeled after projects such as acli and jungle.  Both python
projects.  Since they both are incomplete to what I wanted
and I was looking for a project to learn Go.  I wrote herd
as a way to do both.  

## Install MacOS

```sh
brew install undeadops/tap/herd
```

## Build

Building should be simply

    go build -o herd *.go

### Current Status

Currently ec2 searching works.  I will start adding more 
and more services as time goes on.  I will try to stay
on top of keeping the code base clean.  Pull requests 
welcome. 

### Using

Running it in is current form should be fairly simple.

    ./herd ec2 ls -f '*web*'

Will search EC2 instances by the name "*web*"

    ./herd --profile devacct ec2 list -f '*web*'

Utilizing multiple AWS profiles.  This will preform the same
list but using the devacct AWS profile.



