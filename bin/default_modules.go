package bin

import (
	"github.com/dmin12/zgrab_tls1.3"
	"github.com/dmin12/zgrab_tls1.3/modules"
	"github.com/dmin12/zgrab_tls1.3/modules/bacnet"
	"github.com/dmin12/zgrab_tls1.3/modules/banner"
	"github.com/dmin12/zgrab_tls1.3/modules/dnp3"
	"github.com/dmin12/zgrab_tls1.3/modules/fox"
	"github.com/dmin12/zgrab_tls1.3/modules/ftp"
	"github.com/dmin12/zgrab_tls1.3/modules/http"
	"github.com/dmin12/zgrab_tls1.3/modules/imap"
	"github.com/dmin12/zgrab_tls1.3/modules/ipp"
	"github.com/dmin12/zgrab_tls1.3/modules/modbus"
	"github.com/dmin12/zgrab_tls1.3/modules/mongodb"
	"github.com/dmin12/zgrab_tls1.3/modules/mssql"
	"github.com/dmin12/zgrab_tls1.3/modules/mysql"
	"github.com/dmin12/zgrab_tls1.3/modules/ntp"
	"github.com/dmin12/zgrab_tls1.3/modules/oracle"
	"github.com/dmin12/zgrab_tls1.3/modules/pop3"
	"github.com/dmin12/zgrab_tls1.3/modules/postgres"
	"github.com/dmin12/zgrab_tls1.3/modules/redis"
	"github.com/dmin12/zgrab_tls1.3/modules/siemens"
	"github.com/dmin12/zgrab_tls1.3/modules/smb"
	"github.com/dmin12/zgrab_tls1.3/modules/smtp"
	"github.com/dmin12/zgrab_tls1.3/modules/telnet"
)

var defaultModules zgrab2.ModuleSet

func init() {
	defaultModules = map[string]zgrab2.ScanModule{
		"bacnet":   &bacnet.Module{},
		"banner":   &banner.Module{},
		"dnp3":     &dnp3.Module{},
		"fox":      &fox.Module{},
		"ftp":      &ftp.Module{},
		"http":     &http.Module{},
		"imap":     &imap.Module{},
		"ipp":      &ipp.Module{},
		"modbus":   &modbus.Module{},
		"mongodb":  &mongodb.Module{},
		"mssql":    &mssql.Module{},
		"mysql":    &mysql.Module{},
		"ntp":      &ntp.Module{},
		"oracle":   &oracle.Module{},
		"pop3":     &pop3.Module{},
		"postgres": &postgres.Module{},
		"redis":    &redis.Module{},
		"siemens":  &siemens.Module{},
		"smb":      &smb.Module{},
		"smtp":     &smtp.Module{},
		"ssh":      &modules.SSHModule{},
		"telnet":   &telnet.Module{},
		"tls":      &modules.TLSModule{},
	}
}

// NewModuleSetWithDefaults returns a newly allocated ModuleSet containing all
// ScanModules implemented by the ZGrab2 framework.
func NewModuleSetWithDefaults() zgrab2.ModuleSet {
	out := zgrab2.ModuleSet{}
	defaultModules.CopyInto(out)
	return out
}
