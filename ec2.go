package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/urfave/cli"
)

func ec2List(c *cli.Context) error {
	fmt.Println("EC2 List Instances")
	fmt.Println("AWS Config Profile: ", awsProfile)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           awsProfile,
	}))

	ec2svc := ec2.New(sess)
	params := &ec2.DescribeInstancesInput{
		Filters: generateFilter(),
	}
	resp, err := ec2svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("there was an error listing instances in", err.Error())
		log.Fatal(err.Error())
	}

	for res, _ := range resp.Reservations {
		for _, instances := range resp.Reservations[res].Instances {
			name := "None"
			costCenter := "None"
			env := "None"
			for _, keys := range instances.Tags {
				if *keys.Key == "Name" {
					name = url.QueryEscape(*keys.Value)
				}
				if *keys.Key == "cost_center" {
					costCenter = url.QueryEscape(*keys.Value)
				}
				if *keys.Key == "env" {
					env = url.QueryEscape(*keys.Value)
				}
			}
			i := map[string]string{
				"instanceId":      *instances.InstanceId,
				"name":            name,
				"ipAddress":       *instances.PrivateIpAddress,
				"instanceType":    *instances.InstanceType,
				"publicIpAddress": *instances.PublicIpAddress,
				"keyName":         *instances.KeyName,
				"state":           *instances.State.Name,
				"arch":            *instances.Architecture,
				"cost_center":     costCenter,
				"env":             env,
			}
			fmt.Println(i["name"] + "  " + i["state"] + "  " + i["instanceId"] + "  " + i["instanceType"] + " " + i["ipAddress"] + "  " + i["publicIpAddress"])
		}
	}
	return nil
}
