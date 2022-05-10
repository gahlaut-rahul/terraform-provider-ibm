// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"fmt"
	"log"
	"strings"
	"time"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/networking-go-sdk/authenticatedoriginpullapiv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	cisOriginAuthZone                   = "ZoneVal"
	cisOriginAuthHost                   = "Host"
	cisOriginAuthEnable                 = "on"
)


func DataSourceIBMCISOriginAuthPull() *schema.Resource {
   return &schema.Resource{
      Read: dataIBMCISOriginAuthRead,
      Schema: map[string]*schema.Schema{
         cisID: {
            Type:        schema.TypeString,
            Description: "CIS instance crn",
            Required:    true,  
         },
         cisOriginAuthZone: {        
            Type:        schema.TypeString,
            Description: "CIS Origin Auth Zone",
            Required:    true,
         },
         cisOriginAuthHost: {        
            Type:        schema.TypeString,
            Description: "CIS Origina Auth Host name",
            optional:    true,
         },
        cisOriginAuthEnable: {
            Type:        schema.TypeString,
            Description: "CIS instance crn",
            Required:    true,
         },
      }, // rg : end of schema  
   } // rg : end of return 


} // rg : end of DataSourceIBMCISOriginAuthPull

func dataIBMCISOriginAuthRead(d *schema.ResourceData, meta interface{}) error {
  sess, err := meta.(conns.ClientSession).CisOrigAuthSession()()
   if err != nil {
           return err
   }

  crn := d.Get(cisID).(string)
  sess.Crn = core.StringPtr(crn) 


   d.Set(cisOriginAuthEnable, "on/off") // rg : will be fetached and displayed
   return nil
} 
