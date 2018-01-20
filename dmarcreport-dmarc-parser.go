package main

import (
	"encoding/xml"
	"log"
)

type DmarcFeedback struct {
	ReportMetadata struct {
		OrgName          string `xml:"org_name"`
		Email            string `xml:"email"`
		ExtraContactInfo string `xml:"extra_contact_info"`
		ReportID         uint   `xml:"report_id"`
		DateRange        struct {
			Begin int64 `xml:"begin"`
			End   int64 `xml:"end"`
		} `xml:"date_range"`
	} `xml:"report_metadata"`

	PolicyPublished struct {
		Domain          string `xml:"domain"`
		DKIM            string `xml:"adkim"`
		SPF             string `xml:"aspf"`
		Policy          string `xml:"p"`
		SubdomainPolicy string `xml:"sp"`
		Percentage      uint   `xml:"pct"`
	} `xml:"policy_published"`

	Record []struct {
		Row struct {
			SourceIP        string `xml:"source_ip"`
			Count           uint   `xml:"count"`
			PolicyEvaluated struct {
				Disposition string `xml:"disposition"`
				DKIM        string `xml:"dkim"`
				SPF         string `xml:"spf"`
			} `xml:"policy_evaluated"`
		} `xml:"row"`

		Identifiers struct {
			HeaderFrom string `xml:"header_from"`
		} `xml:"identifiers"`

		AuthResults struct {
			DKIM struct {
				Domain   string `xml:"domain"`
				Result   string `xml:"result"`
				Selector string `xml:"selector"`
			} `xml:"dkim"`
			SPF struct {
				Domain string `xml:"domain"`
				Result string `xml:"result"`
			} `xml:"spf"`
		} `xml:"auth_results"`
	} `xml:"record"`
}

func parseDmarcXML(contents []byte) DmarcFeedback {
	var df DmarcFeedback

	// Parse XML file
	err := xml.Unmarshal(contents, &df)
	if err != nil {
		log.Fatal(err)
	}

	return df
}
