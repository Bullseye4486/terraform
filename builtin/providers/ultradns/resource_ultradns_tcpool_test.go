package ultradns

import (
	"fmt"
	"testing"

	"github.com/Ensighten/udnssdk"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccUltradnsTcpool(t *testing.T) {
	var record udnssdk.RRSet
	domain := "ultradns.phinze.com"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccTcpoolCheckDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testCfgTcpoolMinimal, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUltradnsRecordExists("ultradns_tcpool.it", &record),
					// Specified
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "zone", domain),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "name", "test-tcpool-minimal"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "ttl", "300"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.0.host", "10.6.0.1"),
					// Defaults
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "act_on_probes", "true"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "description", "Minimal TC Pool"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "max_to_lb", "0"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "run_probes", "true"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.0.failover_delay", "0"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.0.priority", "1"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.0.run_probes", "true"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.0.state", "NORMAL"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.0.threshold", "1"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.0.weight", "2"),
					// Generated
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "id", "test-tcpool-minimal.ultradns.phinze.com"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "hostname", "test-tcpool-minimal.ultradns.phinze.com."),
				),
			},
			resource.TestStep{
				Config: fmt.Sprintf(testCfgTcpoolMaximal, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUltradnsRecordExists("ultradns_tcpool.it", &record),
					// Specified
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "zone", domain),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "name", "test-tcpool-maximal"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "ttl", "300"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "description", "traffic controller pool with all settings tuned"),

					resource.TestCheckResourceAttr("ultradns_tcpool.it", "act_on_probes", "false"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "max_to_lb", "2"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "run_probes", "false"),

					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.0.host", "10.6.1.1"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.0.failover_delay", "30"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.0.priority", "1"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.0.run_probes", "true"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.0.state", "ACTIVE"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.0.threshold", "1"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.0.weight", "2"),

					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.1.host", "10.6.1.2"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.1.failover_delay", "30"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.1.priority", "2"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.1.run_probes", "true"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.1.state", "INACTIVE"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.1.threshold", "1"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.1.weight", "4"),

					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.2.host", "10.6.1.3"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.2.failover_delay", "30"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.2.priority", "3"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.2.run_probes", "false"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.2.state", "NORMAL"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.2.threshold", "1"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "rdata.2.weight", "8"),
					// Generated
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "id", "test-tcpool-maximal.ultradns.phinze.com"),
					resource.TestCheckResourceAttr("ultradns_tcpool.it", "hostname", "test-tcpool-maximal.ultradns.phinze.com."),
				),
			},
		},
	})
}

const testCfgTcpoolMinimal = `
resource "ultradns_tcpool" "it" {
  zone        = "%s"
  name        = "test-tcpool-minimal"
  ttl         = 300
  description = "Minimal TC Pool"

  rdata {
    host = "10.6.0.1"
  }
}
`

const testCfgTcpoolMaximal = `
resource "ultradns_tcpool" "it" {
  zone        = "%s"
  name        = "test-tcpool-maximal"
  ttl         = 300
  description = "traffic controller pool with all settings tuned"

  act_on_probes = false
  max_to_lb     = 2
  run_probes    = false

  rdata {
    host = "10.6.1.1"

    failover_delay = 30
    priority       = 1
    run_probes     = true
    state          = "ACTIVE"
    threshold      = 1
    weight         = 2
  }

  rdata {
    host = "10.6.1.2"

    failover_delay = 30
    priority       = 2
    run_probes     = true
    state          = "INACTIVE"
    threshold      = 1
    weight         = 4
  }

  rdata {
    host = "10.6.1.3"

    failover_delay = 30
    priority       = 3
    run_probes     = false
    state          = "NORMAL"
    threshold      = 1
    weight         = 8
  }

  backup_record_rdata          = "10.6.1.4"
  backup_record_failover_delay = 30
}
`