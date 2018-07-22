package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"

	"github.com/urfave/cli"
)

//{{$ec2.Name}}\t\t{{$ec2.state}}\t{{$ec2.id}}\t{{$ec2.role}}\t{{$ec2.cost_center}}\t{{$ec2.env}}\t{{$ec2.instancetype}}\t{{$ec2.ipaddress}}\t{{$ec2.publicIpaddress}}

func printEc2Instances(instances []map[string]string) error {
	for _, i := range instances {
		fmt.Printf(
			"\x1b[33m%-25s\x1b[0m\t%-20s\t%-7s\t%-26s\t%-15s\t%-15s\t%-15s\t%-10s\n",
			i["name"], i["id"], i["state"], i["role"], i["cost_center"], i["ipAddress"], i["publicIpAddress"], i["ami"])
	}
	return nil
}

func ec2List(c *cli.Context) error {
	fmt.Println("EC2 List Instances")
	fmt.Println("AWS Config Profile: ", awsProfile)

	cfg, err := external.LoadDefaultAWSConfig(
		external.WithSharedConfigProfile(awsProfile),
	)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	// TODO: Revisit this... what happens when profile isnt defaulted to us-east-1
	cfg.Region = awsRegion

	var ec2instances []map[string]string

	params := &ec2.DescribeInstancesInput{
		Filters: generateFilter(),
	}

	svc := ec2.New(cfg)

	req := svc.DescribeInstancesRequest(params)
	resp, err := req.Send()
	if err == nil {
		fmt.Println("there was an error listing instances in", err.Error())
		log.Fatal(err.Error())
	}

	for res := range resp.Reservations {
		for _, instances := range resp.Reservations[res].Instances {
			name := "None"
			costCenter := "None"
			env := "None"
			role := "None"
			for _, keys := range instances.Tags {
				if *keys.Key == "Name" {
					name = url.QueryEscape(*keys.Value)
				} else {
					name = "None"
				}
				if *keys.Key == "cost_center" {
					costCenter = url.QueryEscape(*keys.Value)
				} else {
					costCenter = "Not Defined"
				}
				if *keys.Key == "env" {
					env = url.QueryEscape(*keys.Value)
				} else {
					env = "unknown"
				}
				if *keys.Key == "role" {
					role = url.QueryEscape(*keys.Value)
				} else {
					role = "Not Listed"
				}
			}

			instanceType, _ := instances.InstanceType.MarshalValue()
			stateName, _ := instances.InstanceType.MarshalValue()
			arch, _ := instances.Architecture.MarshalValue()

			i := map[string]string{
				"id":           *instances.InstanceId,
				"name":         name,
				"ipAddress":    *instances.PrivateIpAddress,
				"instanceType": instanceType,
				"keyName":      *instances.KeyName,
				"state":        stateName,
				"arch":         arch,
				"cost_center":  costCenter,
				"env":          env,
				"role":         role,
				"ami":          *instances.ImageId,
			}
			if instances.PublicIpAddress != nil {
				i["publicIpAddress"] = *instances.PublicIpAddress
			} else {
				i["publicIpAddress"] = "N/A"
			}
			ec2instances = append(ec2instances, i)
		}
	}
	err = printEc2Instances(ec2instances)
	if err != nil {
		log.Fatal("Execute: ", err)
		return err
	}
	return nil
}
