package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func elbList(c *cli.Context) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           awsProfile,
	}))

	elbsvc := elb.New(sess)

	// I don't want to filter results of ELB's
	filter := &elb.DescribeLoadBalancersInput{}

	loadbalancers, err := elbsvc.DescribeLoadBalancers(filter)
	if err != nil {
		fmt.Println("There was an error listing ELB Instances", err.Error())
		log.Fatal(err.Error())
	}

	for _, lb := range loadbalancers.LoadBalancerDescriptions {
		elbfilter := &elb.DescribeInstanceHealthInput{
			LoadBalancerName: lb.LoadBalancerName,
		}
		elbhealth, err := elbsvc.DescribeInstanceHealth(elbfilter)
		if err != nil {
			fmt.Println("There was and error retreiving ELB Health", err.Error())
			log.Fatal(err.Error())
		}

		noservice := 0
		inservice := 0

		for _, inst := range elbhealth.InstanceStates {
			switch state := *inst.State; state {
			case "InService":
				inservice++
			case "OutOfService":
				noservice++
			case "Unknown":
				noservice++
			}
		}

		elbState := fmt.Sprintf("In: %-2s / Out: %-2s", color.GreenString(strconv.Itoa(inservice)), color.RedString(strconv.Itoa(noservice)))

		fmt.Printf(
			"%-45s\t%-4d\t%-15s\t%-15s\t%-50s\n",
			color.YellowString(*lb.LoadBalancerName), len(lb.Instances), elbState, *lb.Scheme, *lb.DNSName)
	}
	return nil
}
