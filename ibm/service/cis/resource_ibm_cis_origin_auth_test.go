// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisOriginAuth_Basic(t *testing.T) {
	name := "ibm_cis_origin_auth." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisOriginAuthsBasic1("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "certificate", "-----BEGIN CERTIFICATE-----\nMIIDpjCCAo4CCQDiw+/u+5c4eTANBgkqhkiG9w0BAQsFADCBlDELMAkGA1UEBhMC\nSU4xEjAQBgNVBAgMCUtBUk5BVEFLQTESMBAGA1UEBwwJQkFOR0FMT1JFMQwwCgYD\nVQQKDANJQk0xEjAQBgNVBAsMCUlCTSBDTE9VRDEXMBUGA1UEAwwOaWJtdGVzdG1h\nY2hpbmUxIjAgBgkqhkiG9w0BCQEWE2RhcnVueWEuZC5jQGlibS5jb20wHhcNMjIw\nNDA0MTM1ODE2WhcNMjMwNDA0MTM1ODE2WjCBlDELMAkGA1UEBhMCSU4xEjAQBgNV\nBAgMCUtBUk5BVEFLQTESMBAGA1UEBwwJQkFOR0FMT1JFMQwwCgYDVQQKDANJQk0x\nEjAQBgNVBAsMCUlCTSBDTE9VRDEXMBUGA1UEAwwOaWJtdGVzdG1hY2hpbmUxIjAg\nBgkqhkiG9w0BCQEWE2RhcnVueWEuZC5jQGlibS5jb20wggEiMA0GCSqGSIb3DQEB\nAQUAA4IBDwAwggEKAoIBAQCxg0xZgI+JExNCL41AK7FSphsHGP18/RsmrVHiQxGS\nONnh4pBtMJ+/HnnqEoko52L9GGWadu9494JG4vb1Oz3jBJx6vyOBAfJX9EIO0JCz\n/bDdOgyAl9L4MzXF0T5Mc511jHcwMH8jLgczC7hPVm2Acz68z3OFajViLEq7d3+a\n3pC1YV93P3BWn0tNCnHMfUmiXTg40iCVSl1BDpg1hwQnY/L6zAAF+k2jhCJ6W8Ny\nCukSbZ0sEzrhNteYASzWS9k2KMJT8PxoX6bmDWiVVIGHW08YnOC9OZjxHG8fagFs\npEn2FDFc0KNpH7lpLc1qMfWI/i/7cOkHjpqahuD6z9xLAgMBAAEwDQYJKoZIhvcN\nAQELBQADggEBAJIMKN23ChGVU6v+2GT3nnUh5IcZO5qb2bEJrvlyb30uVD8FoBkP\nh7dXlOGsh6tReLB0HLGOz9bnDO1Xzls73So8Ep3M2Xk/42JdzKwXL+Bw3KKTEHP/\n9QUijuwqFTW13KIwX2PWfpYpZTkQwWpi6FS7io+JtEAfO/c5MuwjaWLBEGm7t+HX\nIG21Z2TyIMhFfFoprZG98BSJA4bdqW5gZL2gijoKEtXYpkx65u+4txV566jg2dDr\nKwnFm3A0zHZ3ObRWNt6Vat0SUqOnMOeb0yGNNoxgnoc2NSXlg3+PH9e9FBs5uKE8\npfOqqBCXtdq9QUKjIJnujw/CsYWW4vqLNRI=\n-----END CERTIFICATE-----\n"),
					resource.TestCheckResourceAttr(name, "private_key", "-----BEGIN RSA PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCxg0xZgI+JExNC\nL41AK7FSphsHGP18/RsmrVHiQxGSONnh4pBtMJ+/HnnqEoko52L9GGWadu9494JG\n4vb1Oz3jBJx6vyOBAfJX9EIO0JCz/bDdOgyAl9L4MzXF0T5Mc511jHcwMH8jLgcz\nC7hPVm2Acz68z3OFajViLEq7d3+a3pC1YV93P3BWn0tNCnHMfUmiXTg40iCVSl1B\nDpg1hwQnY/L6zAAF+k2jhCJ6W8NyCukSbZ0sEzrhNteYASzWS9k2KMJT8PxoX6bm\nDWiVVIGHW08YnOC9OZjxHG8fagFspEn2FDFc0KNpH7lpLc1qMfWI/i/7cOkHjpqa\nhuD6z9xLAgMBAAECggEAbO29NE8HxX3HG55Cd1ZYgfccLsbPBpvqxVkmHko5xhjM\n2yhEqDxmSslQ1qp5MHiM7fLCpn7FhN2dPBKaqPGpkF2MCGaySr//Dqn8v0qNAWZz\n2c19TovcEiKapME6EYAA59lCanfYDKZ6FIDkoQrQNzqBDSvgH8aE67FySoeR7l4O\nX2ltn2iEDdfbx7oFRdvA4mq2JfnIEK7faaVF+AybwDdn0WGyzcvsju85uEtHF/SS\n75BtwFZzyTNuEcjhBWevA9dTjBFMcnpWx2HgxyO/oKcNuiG8KFQMOoPLlfKw8Lwt\ntqhYmoV1ATmLsdsp9v3d2alvO3WlOrzWcBCS91jqQQKBgQDrAbLflXNOtFjwChEQ\nh6JzFbFGCeQw+avkh+h4qMPKpuX+kMzOQRZb30PDY+zxSDD5lJXqRxTNa5TsByVj\n8GAGWvy/84pLmB6KR/ujVJ+DH0OlqyEQTPttrTtrEkI0uw6cO/AGP6Fb8Gp0/3QN\n0np1She1iptPHxi9HFUhjFONvwKBgQDBXsp56yDmRAnl7twrCNvQTCB1dq7nPZE6\n/7N+Vpznon1k9At92rHErgYo9Rib21o/hPNeQJTIaey70ODB7q0BRv2e2PXwfyN3\nJrJwGYRRO8vLO/zrhHVxzjDt405bXR5R/j2IpR2pgrLXqpx+PfzujWFayQFkbzHf\ncqQBEwzsdQKBgGvzztBIHbzEuaoiZa5bL/N/vnw25PzeY+jJya9Ljw0TV8l1iK8i\nVPwE9mLWDyzTBbRQXgFNf6/RQIqfybw72lBxEXO3kwqgqT7KTDy+Dbw062U51Clh\nw4mhLw9DRuhkGRUJr3ufVScfrDdsdUo4Koqga324WxmgZkPQtQaBKIyPAoGAWudN\n9jyj7bwEjzRYCl8Svvxasf3GQWz/DiZQ4k6jWn1Xx5K2qEacFWLeAHkgRXy8E2pT\n4nYnu4OYR77tOh4S9KvD5N4H2DRcntHxRqOoQWwD5RnhT3Kop4SQGfUmy+qdq1wC\n328H371Sh/JruSk484hBQSWHYwinAG1rThn/lFUCgYA3okRpQKVulm7xg/g7pKwy\nxvBVlNVa0BxtqFO+vn//kfW7CMNSJQO5m+X0y7wYoX0yYsXHAX/PusuCsJGuvjnd\nnvKzVqYTUQ5HMA0rk80TizM34loHbrDnCwMS0WJTNyyt/QC+KRcrKHXvuHYgu2J6\nSjci4j+SSooywqCcC4Liww==\n-----END RSA PRIVATE KEY-----\n"), // pragma: whitelist secret
				),
			},
		},
	})
}

func testAccCheckCisOriginAuthsBasic1(id string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_origin_auth" "%[1]s" {
		cis_id                    = data.ibm_cis.cis.id
		domain_id                 = data.ibm_cis_domain.cis_domain.domain_id
		certificate               = "-----BEGIN CERTIFICATE-----\nMIIDpjCCAo4CCQDiw+/u+5c4eTANBgkqhkiG9w0BAQsFADCBlDELMAkGA1UEBhMC\nSU4xEjAQBgNVBAgMCUtBUk5BVEFLQTESMBAGA1UEBwwJQkFOR0FMT1JFMQwwCgYD\nVQQKDANJQk0xEjAQBgNVBAsMCUlCTSBDTE9VRDEXMBUGA1UEAwwOaWJtdGVzdG1h\nY2hpbmUxIjAgBgkqhkiG9w0BCQEWE2RhcnVueWEuZC5jQGlibS5jb20wHhcNMjIw\nNDA0MTM1ODE2WhcNMjMwNDA0MTM1ODE2WjCBlDELMAkGA1UEBhMCSU4xEjAQBgNV\nBAgMCUtBUk5BVEFLQTESMBAGA1UEBwwJQkFOR0FMT1JFMQwwCgYDVQQKDANJQk0x\nEjAQBgNVBAsMCUlCTSBDTE9VRDEXMBUGA1UEAwwOaWJtdGVzdG1hY2hpbmUxIjAg\nBgkqhkiG9w0BCQEWE2RhcnVueWEuZC5jQGlibS5jb20wggEiMA0GCSqGSIb3DQEB\nAQUAA4IBDwAwggEKAoIBAQCxg0xZgI+JExNCL41AK7FSphsHGP18/RsmrVHiQxGS\nONnh4pBtMJ+/HnnqEoko52L9GGWadu9494JG4vb1Oz3jBJx6vyOBAfJX9EIO0JCz\n/bDdOgyAl9L4MzXF0T5Mc511jHcwMH8jLgczC7hPVm2Acz68z3OFajViLEq7d3+a\n3pC1YV93P3BWn0tNCnHMfUmiXTg40iCVSl1BDpg1hwQnY/L6zAAF+k2jhCJ6W8Ny\nCukSbZ0sEzrhNteYASzWS9k2KMJT8PxoX6bmDWiVVIGHW08YnOC9OZjxHG8fagFs\npEn2FDFc0KNpH7lpLc1qMfWI/i/7cOkHjpqahuD6z9xLAgMBAAEwDQYJKoZIhvcN\nAQELBQADggEBAJIMKN23ChGVU6v+2GT3nnUh5IcZO5qb2bEJrvlyb30uVD8FoBkP\nh7dXlOGsh6tReLB0HLGOz9bnDO1Xzls73So8Ep3M2Xk/42JdzKwXL+Bw3KKTEHP/\n9QUijuwqFTW13KIwX2PWfpYpZTkQwWpi6FS7io+JtEAfO/c5MuwjaWLBEGm7t+HX\nIG21Z2TyIMhFfFoprZG98BSJA4bdqW5gZL2gijoKEtXYpkx65u+4txV566jg2dDr\nKwnFm3A0zHZ3ObRWNt6Vat0SUqOnMOeb0yGNNoxgnoc2NSXlg3+PH9e9FBs5uKE8\npfOqqBCXtdq9QUKjIJnujw/CsYWW4vqLNRI=\n-----END CERTIFICATE-----\n"
		private_key               = "-----BEGIN RSA PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCxg0xZgI+JExNC\nL41AK7FSphsHGP18/RsmrVHiQxGSONnh4pBtMJ+/HnnqEoko52L9GGWadu9494JG\n4vb1Oz3jBJx6vyOBAfJX9EIO0JCz/bDdOgyAl9L4MzXF0T5Mc511jHcwMH8jLgcz\nC7hPVm2Acz68z3OFajViLEq7d3+a3pC1YV93P3BWn0tNCnHMfUmiXTg40iCVSl1B\nDpg1hwQnY/L6zAAF+k2jhCJ6W8NyCukSbZ0sEzrhNteYASzWS9k2KMJT8PxoX6bm\nDWiVVIGHW08YnOC9OZjxHG8fagFspEn2FDFc0KNpH7lpLc1qMfWI/i/7cOkHjpqa\nhuD6z9xLAgMBAAECggEAbO29NE8HxX3HG55Cd1ZYgfccLsbPBpvqxVkmHko5xhjM\n2yhEqDxmSslQ1qp5MHiM7fLCpn7FhN2dPBKaqPGpkF2MCGaySr//Dqn8v0qNAWZz\n2c19TovcEiKapME6EYAA59lCanfYDKZ6FIDkoQrQNzqBDSvgH8aE67FySoeR7l4O\nX2ltn2iEDdfbx7oFRdvA4mq2JfnIEK7faaVF+AybwDdn0WGyzcvsju85uEtHF/SS\n75BtwFZzyTNuEcjhBWevA9dTjBFMcnpWx2HgxyO/oKcNuiG8KFQMOoPLlfKw8Lwt\ntqhYmoV1ATmLsdsp9v3d2alvO3WlOrzWcBCS91jqQQKBgQDrAbLflXNOtFjwChEQ\nh6JzFbFGCeQw+avkh+h4qMPKpuX+kMzOQRZb30PDY+zxSDD5lJXqRxTNa5TsByVj\n8GAGWvy/84pLmB6KR/ujVJ+DH0OlqyEQTPttrTtrEkI0uw6cO/AGP6Fb8Gp0/3QN\n0np1She1iptPHxi9HFUhjFONvwKBgQDBXsp56yDmRAnl7twrCNvQTCB1dq7nPZE6\n/7N+Vpznon1k9At92rHErgYo9Rib21o/hPNeQJTIaey70ODB7q0BRv2e2PXwfyN3\nJrJwGYRRO8vLO/zrhHVxzjDt405bXR5R/j2IpR2pgrLXqpx+PfzujWFayQFkbzHf\ncqQBEwzsdQKBgGvzztBIHbzEuaoiZa5bL/N/vnw25PzeY+jJya9Ljw0TV8l1iK8i\nVPwE9mLWDyzTBbRQXgFNf6/RQIqfybw72lBxEXO3kwqgqT7KTDy+Dbw062U51Clh\nw4mhLw9DRuhkGRUJr3ufVScfrDdsdUo4Koqga324WxmgZkPQtQaBKIyPAoGAWudN\n9jyj7bwEjzRYCl8Svvxasf3GQWz/DiZQ4k6jWn1Xx5K2qEacFWLeAHkgRXy8E2pT\n4nYnu4OYR77tOh4S9KvD5N4H2DRcntHxRqOoQWwD5RnhT3Kop4SQGfUmy+qdq1wC\n328H371Sh/JruSk484hBQSWHYwinAG1rThn/lFUCgYA3okRpQKVulm7xg/g7pKwy\nxvBVlNVa0BxtqFO+vn//kfW7CMNSJQO5m+X0y7wYoX0yYsXHAX/PusuCsJGuvjnd\nnvKzVqYTUQ5HMA0rk80TizM34loHbrDnCwMS0WJTNyyt/QC+KRcrKHXvuHYgu2J6\nSjci4j+SSooywqCcC4Liww==\n-----END RSA PRIVATE KEY-----\n" # pragma: whitelist secret
	  }
`, id)
}
