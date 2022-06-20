// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	cisMtlsID            = "mtls_id"
	cisMtlsCert          = "cert_mtls"
	cisMtlsCertName      = "cert_name"
	cisMtlsHostName      = "host_name"
	cisMtlsCertCreatedAt = "created_at"
	cisMtlsCertUpdatedAt = "updated_at"
	cisMtlsCertExpireOn  = "expires_on"
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
			cisMtlsID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Mtls transaction ID",
			},
			cisMtlsCert: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Certificate contents",
			},
			cisMtlsCertName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Certificate name",
			},
			cisMtlsHostName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Certificate host name",
			},
			cisMtlsCertCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate Created At",
			},
			cisMtlsCertUpdatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate Updated At",
			},
			cisMtlsCertExpireOn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate Expires on",
			},
		},
	}
}
func ResourceIBMCISMtlsCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisMtlsSession() %s %v", err, sess)
	}
	crn := d.Get(cisID).(string)
	zoneID := d.Get(cisDomainID).(string)
	sess.Crn = core.StringPtr(crn)

	options := sess.NewCreateAccessCertificateOptions(zoneID)
	if name, ok := d.GetOk(cisMtlsCertName); ok {
		options.SetName(name.(string))
	}

	if cert_val, ok := d.GetOk(cisMtlsCert); ok {
		options.SetCertificate(cert_val.(string))
	}

	if host_val, ok := d.GetOk(cisMtlsHostName); ok {
		options.SetAssociatedHostnames([]string{host_val.(string)})
	}

	result, resp, err := sess.CreateAccessCertificate(options)
	if err != nil || result == nil {
		return fmt.Errorf("[ERROR] Error creating MTLS access certificate %v", resp)
	}

	d.SetId(flex.ConvertCisToTfThreeVar(*result.Result.ID, zoneID, crn))
	return ResourceIBMCISMtlsRead(d, meta)

}

func ResourceIBMCISMtlsRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisMtlsSession() %s %v", err, sess)
	}
	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	certID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(d.Id())
	getOptions := sess.NewGetAccessCertificateOptions(zoneID, certID)
	result, response, err := sess.GetAccessCertificate(getOptions)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Error While reading MTLS access certificate %v:%v:%v", err, response, result)
	}

	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)
	d.Set(cisMtlsID, *result.Result.ID)
	d.Set(cisMtlsCertCreatedAt, *result.Result.CreatedAt)
	d.Set(cisMtlsCertUpdatedAt, *result.Result.UpdatedAt)
	d.Set(cisMtlsCertExpireOn, *result.Result.ExpiresOn)

	return nil
}

func ResourceIBMCISMtlsUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisMtlsSession() %s %v", err, sess)
	}

	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	certID, zoneID, _, _ := flex.ConvertTfToCisThreeVar(d.Id())

	if d.HasChange(cisMtlsCertName) ||
		d.HasChange(cisMtlsHostName) {

		updateOption := sess.NewUpdateAccessCertificateOptions(zoneID, certID)
		if host_val, ok := d.GetOk(cisMtlsHostName); ok {
			updateOption.SetAssociatedHostnames([]string{host_val.(string)})
		}

		if name, ok := d.GetOk(cisMtlsCertName); ok {
			updateOption.SetName(name.(string))
		}

		updateResult, updateResp, updateErr := sess.UpdateAccessCertificate(updateOption)
		if updateErr != nil {
			if updateResp != nil && updateResp.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			return fmt.Errorf("[ERROR] Error while updating the MTLS cert options %v", updateResult)
		}
	}

	return ResourceIBMCISMtlsRead(d, meta)
}

func ResourceIBMCISMtlsDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisMtlsSession() %s %v", err, sess)
	}

	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	zoneID := d.Get(cisDomainID).(string)

	listOpt := sess.NewListAccessCertificatesOptions(zoneID)
	listResult, listResp, listErr := sess.ListAccessCertificates(listOpt)

	if listErr != nil {
		if listResp != nil && listResp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Error While deleting the MTLS cert")
	}

	for _, certId := range listResult.Result {
		delOpt := sess.NewDeleteAccessCertificateOptions(zoneID, *certId.ID)
		_, delResp, delErr := sess.DeleteAccessCertificate(delOpt)
		if delErr != nil {
			if delResp != nil && delResp.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			return fmt.Errorf("[ERROR] Error While deleting the MTLS cert")
		}

	}

	return nil

}
