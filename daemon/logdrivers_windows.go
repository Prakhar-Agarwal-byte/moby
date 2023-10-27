package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	// Importing packages here only to make sure their init gets called and
	// therefore they register themselves to the logdriver factory.
	_ "github.com/Prakhar-Agarwal-byte/moby/daemon/logger/awslogs"
	_ "github.com/Prakhar-Agarwal-byte/moby/daemon/logger/etwlogs"
	_ "github.com/Prakhar-Agarwal-byte/moby/daemon/logger/fluentd"
	_ "github.com/Prakhar-Agarwal-byte/moby/daemon/logger/gcplogs"
	_ "github.com/Prakhar-Agarwal-byte/moby/daemon/logger/gelf"
	_ "github.com/Prakhar-Agarwal-byte/moby/daemon/logger/jsonfilelog"
	_ "github.com/Prakhar-Agarwal-byte/moby/daemon/logger/logentries"
	_ "github.com/Prakhar-Agarwal-byte/moby/daemon/logger/loggerutils/cache"
	_ "github.com/Prakhar-Agarwal-byte/moby/daemon/logger/splunk"
	_ "github.com/Prakhar-Agarwal-byte/moby/daemon/logger/syslog"
)
