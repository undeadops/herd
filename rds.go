package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/fatih/color"
	"github.com/gosuri/uitable"
	"github.com/urfave/cli"
)

func printRds(dbs []map[string]string) error {
	table := uitable.New()
	table.AddRow("NAME", "ENGINE", "VERSION", "IS_PUBLIC", "ENDPOINT")
	for _, r := range dbs {
		// fmt.Printf(
		// 	"%-28s\t%s:%s\t%-18s\t%-50s\n",
		// 	color.YellowString(r["name"]), r["engine"], r["engine_version"],
		// 	r["isPublic"], r["endpoint"])
		table.AddRow(color.YellowString(r["name"]), r["engine"], r["engine_version"], r["isPublic"], r["endpoint"])
	}
	fmt.Println(table)
	return nil
}

func rdsList(c *cli.Context) error {

	cfg, err := external.LoadDefaultAWSConfig(
		external.WithSharedConfigProfile(awsProfile),
	)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	// TODO: Revisit this... what happens when profile isnt defaulted to us-east-1
	cfg.Region = awsRegion

	rdsFilter := &rds.DescribeDBInstancesInput{}

	var rdsList []map[string]string

	svc := rds.New(cfg)
	req := svc.DescribeDBInstancesRequest(rdsFilter)
	resp, err := req.Send()
	if err != nil {
		fmt.Println("There was an error listing RDS Instances", err.Error())
		log.Fatal(err.Error())
	}

	for _, dbinst := range resp.DBInstances {
		r := map[string]string{
			"name":           *dbinst.DBInstanceIdentifier,
			"id":             *dbinst.DbiResourceId,
			"endpoint":       fmt.Sprintf("%s:%d", *dbinst.Endpoint.Address, *dbinst.Endpoint.Port),
			"engine":         *dbinst.Engine,
			"engine_version": *dbinst.EngineVersion,
			//   These Aren't working, will revisit...
			//			"security_groups":  *rds.DBSecurityGroups.DBSecurityGroupName,
			//"parameter_groups": strings.Join(*dbinst.DBParameterGroups.DBParameterGroupName, ","),
			"subnet_group": *dbinst.DBSubnetGroup.DBSubnetGroupName,
		}
		if *dbinst.MultiAZ == true {
			r["multi_az"] = "Yes"
		} else {
			r["multi_az"] = "No"
		}
		if *dbinst.PubliclyAccessible == true {
			r["isPublic"] = color.CyanString("Public")
		} else {
			r["isPublic"] = color.GreenString("NotPublic")
		}

		rdsList = append(rdsList, r)
	}
	printRds(rdsList)
	return nil
}
