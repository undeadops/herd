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

func printElb(elb []map[string]string) error {
	for _, e := range elb {
		fmt.Printf(
			"%-45s\t%-3s\t%-2s:%-2s\t%-15s\t%-50s\n",
			color.YellowString(e["name"]), e["numInstances"], color.GreenString(e["instances_in"]),
			color.RedString(e["instances_out"]), e["scheme"], e["dns"])
	}
	return nil
}

func elbList(c *cli.Context) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           awsProfile,
	}))

	var elbList []map[string]string

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

		elbInstance := map[string]string{
			"name":          *lb.LoadBalancerName,
			"numInstances":  strconv.Itoa(len(lb.Instances)),
			"instances_in":  strconv.Itoa(inservice),
			"instances_out": strconv.Itoa(noservice),
			"scheme":        *lb.Scheme,
			"dns":           *lb.DNSName,
		}

		elbList = append(elbList, elbInstance)
	}
	printElb(elbList)
	return nil
}
