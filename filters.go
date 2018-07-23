package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func generateFilter() []ec2.Filter {

	filters := []ec2.Filter{
		ec2.Filter{
			Name:   aws.String("instance-state-name"),
			Values: []string{"running", "pending"},
		},
	}
	if awsFilterName != "" {
		filters = append(filters, ec2.Filter{
			Name:   aws.String("tag:Name"),
			Values: []string{awsFilterName},
		})
	}

	if awsFilterTags != "" {
		result := strings.Split(awsFilterTags, ",")
		if len(result) > 1 {
			for i := range result {
				tags := result[i]
				t := strings.Split(tags, "=")
				if len(t) == 2 {
					filters = append(filters, ec2.Filter{
						Name:   aws.String("tag:" + t[0]),
						Values: []string{t[1]},
					})
				}
			}
		} else {
			fmt.Println("Inside else")
			t := strings.Split(awsFilterTags, "=")
			filters = append(filters, ec2.Filter{
				Name:   aws.String("tag:" + t[0]),
				Values: []string{t[1]},
			})
		}
	}
	return filters
}
