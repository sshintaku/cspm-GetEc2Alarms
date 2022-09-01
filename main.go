package main

import (
	"log"

	"github.com/sshintaku/prisma_session"
)

func main() {
	var params = prisma_session.ReadParameters()
	session := prisma_session.Session{ApiUrl: params.ApiUrl}
	session.CreateSession()
	dict := session.DescribeEC2Instances(params)
	//session.GetHostInfo()
	result, resultError := session.GetCSPMAlerts(params.AlertQueryParameters)
	if resultError != nil {
		log.Fatalln(resultError)
	}
	session.GenerateAlertSummary(result, dict)

}
