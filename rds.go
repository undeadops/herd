package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func printRds(dbs []map[string]string) error {
	for _, r := range dbs {
		fmt.Printf(
			"%-28s\t%s:%s\t%-18s\t%-50s\n",
			color.YellowString(r["name"]), r["engine"], r["engine_version"],
			r["isPublic"], r["endpoint"])
	}
	return nil
}

func rdsList(c *cli.Context) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           awsProfile,
	}))

	var rdsList []map[string]string

	rdsFilter := &rds.DescribeDBInstancesInput{}

	rdsSvc := rds.New(sess)
	rdsInstances, err := rdsSvc.DescribeDBInstances(rdsFilter)
	if err != nil {
		fmt.Println("There was an error listing RDS Instances", err.Error())
		log.Fatal(err.Error())
	}
	for _, dbinst := range rdsInstances.DBInstances {
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
