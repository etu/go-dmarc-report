package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"net"
	"os"
)

func printDmarcFeedbacks(dfs []DmarcFeedback) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{
		"IP Address",
		"Hostname",
		"Message Count",
		"Disposition",
		"DKIM Domain",
		"DKIM Result",
		"SPF Domain",
		"SPF Result",
	})

	for _, df := range dfs {
		//fmt.Printf(
		//	"%s -- %s\n",
		//	time.Unix(df.ReportMetadata.DateRange.Begin, 0),
		//	time.Unix(df.ReportMetadata.DateRange.End, 0),
		//)

		for _, rec := range df.Record {
			hostnames, _ := net.LookupAddr(rec.Row.SourceIP)

			table.Append([]string{
				rec.Row.SourceIP,
				hostnames[0],
				fmt.Sprint(rec.Row.Count),
				rec.Row.PolicyEvaluated.Disposition,
				rec.AuthResults.DKIM.Domain,
				rec.AuthResults.DKIM.Result,
				rec.AuthResults.SPF.Domain,
				rec.AuthResults.SPF.Result,
			})
		}
	}

	table.Render()
}
