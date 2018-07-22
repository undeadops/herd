package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/elb"

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
	cfg, err := external.LoadDefaultAWSConfig(
		external.WithSharedConfigProfile(awsProfile),
	)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	// TODO: Revisit this... what happens when profile isnt defaulted to us-east-1
	cfg.Region = awsRegion

	var elbList []map[string]string

	//elbsvc := elb.New(sess)
	svc := elb.New(cfg)

	// I don't want to filter results of ELB's
	filter := &elb.DescribeLoadBalancersInput{}

	req := svc.DescribeLoadBalancersRequest(filter)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elb.ErrCodeAccessPointNotFoundException:
				fmt.Println(elb.ErrCodeAccessPointNotFoundException, aerr.Error())
			case elb.ErrCodeDependencyThrottleException:
				fmt.Println(elb.ErrCodeDependencyThrottleException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return err
	}

	for _, lb := range result.LoadBalancerDescriptions {
		elbfilter := &elb.DescribeInstanceHealthInput{
			LoadBalancerName: lb.LoadBalancerName,
		}
		req := svc.DescribeInstanceHealthRequest(elbfilter)
		result, err := req.Send()
		if err != nil {
			fmt.Println("There was and error retreiving ELB Health", err.Error())
			log.Fatal(err.Error())
		}

		noservice := 0
		inservice := 0

		for _, inst := range result.InstanceStates {
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
