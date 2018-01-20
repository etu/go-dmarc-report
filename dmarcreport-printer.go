package main

import (
	"fmt"
	"net"
	"time"
)

func printDmarcFeedbacks(dfs []DmarcFeedback) {
	for _, df := range dfs {
		fmt.Printf(
			"%s -- %s\n",
			time.Unix(df.ReportMetadata.DateRange.Begin, 0),
			time.Unix(df.ReportMetadata.DateRange.End, 0),
		)

		for _, rec := range df.Record {
			hostnames, _ := net.LookupAddr(rec.Row.SourceIP)

			fmt.Printf(
				"%15s | %s | %d | %s | %30s | %s | %30s | %s\n",
				rec.Row.SourceIP,
				hostnames[0],
				rec.Row.Count,
				rec.Row.PolicyEvaluated.Disposition,
				rec.AuthResults.DKIM.Domain,
				rec.AuthResults.DKIM.Result,
				rec.AuthResults.SPF.Domain,
				rec.AuthResults.SPF.Result,
			)
		}
	}
}
