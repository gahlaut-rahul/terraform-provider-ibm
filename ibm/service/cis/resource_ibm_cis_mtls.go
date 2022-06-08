// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	cisMtlsID       = "mtls_id"
	cisMtlsCert     = "cert_mtls"
	cisMtlsCertName = "cert_name"
)

func ResourceIBMCISMtls() *schema.Resource {
	return &schema.Resource{
		Create:   ResourceIBMCISMtlsCreate,
		Read:     ResourceIBMCISMtlsRead,
		Update:   ResourceIBMCISMtlsUpdate,
		Delete:   ResourceIBMCISMtlsDelete,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Description:      "Associated CIS domain",
				Required:         true,
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisMtlsCert: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Certificate content",
			},
			cisMtlsCertName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Certificate Name",
			},
		},
	}
}
func ResourceIBMCISMtlsCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisMtlsSession() %s %v", err, sess)
	}
	// crn := d.Get(cisID).(string)
	// zoneID, _, _ := flex.ConvertTftoCisTwoVar(d.Get(cisDomainID).(string))

	return ResourceIBMCISMtlsRead(d, meta)

}
func ResourceIBMCISMtlsRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisMtlsSession() %s %v", err, sess)
	}

	return nil
}
func ResourceIBMCISMtlsUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisMtlsSession() %s %v", err, sess)
	}

	return ResourceIBMCISMtlsRead(d, meta)
}
func ResourceIBMCISMtlsDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisMtlsSession() %s %v", err, sess)
	}

	return nil

}
