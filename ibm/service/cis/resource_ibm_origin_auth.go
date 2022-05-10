// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

import (
        "fmt"
        "github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
        "github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
        "github.com/IBM/networking-go-sdk/authenticatedoriginpullapiv1"
        "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
      cisOriginAuthID     = "cis_origin_auth_id"
      cisOriginAuthLevel  = "zone-or-host"
      cisOriginAuthHost   = "host-name"
      cisOriginAuthOption = "On-off"
      cisOriginAuthZoneID = 1
)

func ResourceIBMCISOriginAuthPull() *schema.Resource {
     return &schema.Resource{
        Create:   ResourceIBMCISOriginAuthPullCreate,
        Read:     ResourceIBMCISOriginAuthPullRead,
        Update:   ResourceIBMCISOriginAuthPullUpdate,
        Delete:   ResourceIBMCISOriginAuthPullDelete,
        Importer: &schema.ResourceImporter{},
        Schema: map[string]*schema.Schema{
              cisID: {
                   Type:        schema.TypeString,
                   Description: "CIS instance crn",
                   Required:    true,
              },
              cisOriginAuthLevel: {
                   Type:        schema.TypeString,
                   Description: "Authentication type zone or host",
                   Required:    true,
              },
              cisOriginAuthHost: {
                   Type:        schema.TypeString,
                   Description: "Host Name",
                   optional:    true,
              },
              cisOriginAuthOption: {
                   Type:        schema.TypeString,
                   Description: "Enable or disable origin auth",
                   optional:    true,
              },
              cisOriginAuthZoneID: {
                   Type:        schema.TypeInt,
                   Description: "Enable or disable origin auth",
                   Required:    true,
              },

        }, // rg : end of scheme

    } // rg : end of return


} // rg : end of function 


func ResourceIBMCISOriginAuthPullCreate(d *schema.ResourceData, meta interface{}) error {
      sess, err := meta.(conns.ClientSession).CisOrigAuthSession()()
      if err != nil {
                return fmt.Errorf("[ERROR] Error while getting the CisOrigAuthSession %s", err)
      }

      crn := d.Get(cisID).(string)
      sess.Crn = core.StringPtr(crn)
      DefaultServiceName := "default Service origin auth"
      serviceURL :=""
      zoneID:=1
      globalOptions := &AuthenticatedOriginPullApiV1Options{
                       ServiceName: DefaultServiceName,
                       URL: serviceURL,
                       Authenticator: authenticator,
                       Crn: &crn,
                       ZoneIdentifier: &zoneId,
      }
     
     service, serviceErr := NewAuthenticatedOriginPullApiV1(globalOptions)
     if serviceErr != nil {
		fmt.Println(serviceErr) 
     }

     return dataIBMCISOriginAuthRead(d, meta)
}

func ResourceIBMCISOriginAuthPullRead(d *schema.ResourceData, meta interface{}) error {
      sess, err := meta.(conns.ClientSession).CisOrigAuthSession()()
      if err != nil {
                return fmt.Errorf("[ERROR] Error while getting the CisOrigAuthSession %s", err)
      }



}

func ResourceIBMCISOriginAuthPullUpdate(d *schema.ResourceData, meta interface{}) error {
      sess, err := meta.(conns.ClientSession).CisOrigAuthSession()()
      if err != nil {
                return fmt.Errorf("[ERROR] Error while getting the CisOrigAuthSession %s", err)
      }
}

func ResourceIBMCISOriginAuthPullDelete(d *schema.ResourceData, meta interface{}) error {
      sess, err := meta.(conns.ClientSession).CisOrigAuthSession()()
      if err != nil {
                return fmt.Errorf("[ERROR] Error while getting the CisOrigAuthSession %s", err)
      }

}

