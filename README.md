# Herd

Because naming is hard.

AWS CLI for getting to the most needed info I use everyday
but don't want to open the AWS Console to look up.

Modeled after projects such as acli and jungle.  Both python
projects.  Since they both are incomplete to what I wanted
and I was looking to learn more Go.  I wrote herd as a way to
do both.  

## Installing

### MacOS
```sh
brew install undeadops/tap/herd
```

## Usage

List all ec2 instances on your default AWS profile

```sh
herd ec2 ls
```

List all ec2 instances off an AWS profile named development
```sh
herd --profile development ec2 ls
```

There is some basic regexing built-in
Searching EC2 instances by the name "*web*" on the development profile
```sh
herd --profile development ec2 list -f '*web*'
```

I also have built-in the ability to pull up instances based of of tag names
Searching EC@ for instances with the cost_center that equals elasticsearch
```sh
herd ec2 ls -t 'cost_center=elasticsearch'
```

These should also all still work with the rds and elb types as well. 
