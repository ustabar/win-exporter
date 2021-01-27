// +build windows

package collector

import (
	"errors"

	"github.com/StackExchange/wmi"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

func init() {
	registerCollector("aspnet", NewASPNETCollector)
}

type ASPNETCollector struct {
	ApplicationRestarts      *prometheus.Desc
	ApplicationsRunning      *prometheus.Desc
	AuditFailureEventsRaised *prometheus.Desc
	AuditSuccessEventsRaised *prometheus.Desc

	ErrorEventsRaised               *prometheus.Desc
	Frequency_Object                *prometheus.Desc
	Frequency_PerfTime              *prometheus.Desc
	Frequency_Sys100NS              *prometheus.Desc
	InfrastructureErrorEventsRaised *prometheus.Desc

	RequestErrorEventsRaised     *prometheus.Desc
	RequestExecutionTime         *prometheus.Desc
	RequestsCurrent              *prometheus.Desc
	RequestsDisconnected         *prometheus.Desc
	RequestsInNativeQueue        *prometheus.Desc
	RequestsQueued               *prometheus.Desc
	RequestsRejected             *prometheus.Desc
	RequestWaitTime              *prometheus.Desc
	StateServerSessionsAbandoned *prometheus.Desc
	StateServerSessionsActive    *prometheus.Desc
	StateServerSessionsTimedOut  *prometheus.Desc
	StateServerSessionsTotal     *prometheus.Desc
	Timestamp_Object             *prometheus.Desc
	Timestamp_PerfTime           *prometheus.Desc
	Timestamp_Sys100NS           *prometheus.Desc
	WorkerProcessesRunning       *prometheus.Desc
	WorkerProcessRestarts        *prometheus.Desc
}

// NewASPNETCollector ...
func NewASPNETCollector() (Collector, error) {
	const subsystem = "aspnet"

	buildASPNET := &ASPNETCollector{
		// Websites
		// Gauges
		ApplicationRestarts: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "aspnet_application_restarts"),
			"Number of times the application has been restarted during the web server's lifetime.",
			[]string{"site01"},
			nil,
		),

		ApplicationsRunning: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site02"},
			nil,
		),

		AuditFailureEventsRaised: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site03"},
			nil,
		),

		AuditSuccessEventsRaised: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site04"},
			nil,
		),

		ErrorEventsRaised: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site05"},
			nil,
		),

		Frequency_Object: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site06"},
			nil,
		),

		Frequency_PerfTime: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site07"},
			nil,
		),

		Frequency_Sys100NS: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site08"},
			nil,
		),

		InfrastructureErrorEventsRaised: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site09"},
			nil,
		),

		RequestErrorEventsRaised: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site10"},
			nil,
		),

		RequestExecutionTime: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site11"},
			nil,
		),
		RequestsCurrent: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site12"},
			nil,
		),
		RequestsDisconnected: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site13"},
			nil,
		),
		RequestsInNativeQueue: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site14"},
			nil,
		),
		RequestsQueued: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site15"},
			nil,
		),
		RequestsRejected: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site16"},
			nil,
		),
		RequestWaitTime: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site17"},
			nil,
		),
		StateServerSessionsAbandoned: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site18"},
			nil,
		),

		StateServerSessionsActive: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site19"},
			nil,
		),

		StateServerSessionsTimedOut: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site20"},
			nil,
		),

		StateServerSessionsTotal: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site21"},
			nil,
		),

		Timestamp_Object: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site22"},
			nil,
		),

		Timestamp_PerfTime: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site23"},
			nil,
		),

		Timestamp_Sys100NS: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site24"},
			nil,
		),

		WorkerProcessesRunning: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site25"},
			nil,
		),

		WorkerProcessRestarts: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "current_anonymous_users"),
			"Number of users who currently have an anonymous connection using the Web service (WebService.CurrentAnonymousUsers)",
			[]string{"site26"},
			nil,
		),
	}

	return buildASPNET, nil
}

// Collect sends the metric values for each metric
// to the provided prometheus Metric channel.
func (c *ASPNETCollector) Collect(ctx *ScrapeContext, ch chan<- prometheus.Metric) error {
	if desc, err := c.collect(ch); err != nil {
		log.Error("failed collecting aspnet metrics:", desc, err)
		return err
	}
	return nil
}

type Win32_PerfRawData_ASPNET4030319_ASPNETv4030319 struct {
	Name string

	ApplicationRestarts      uint32
	ApplicationsRunning      uint32
	AuditFailureEventsRaised uint32
	AuditSuccessEventsRaised uint32

	ErrorEventsRaised               uint32
	Frequency_Object                uint64
	Frequency_PerfTime              uint64
	Frequency_Sys100NS              uint64
	InfrastructureErrorEventsRaised uint32

	RequestErrorEventsRaised     uint32
	RequestExecutionTime         uint32
	RequestsCurrent              uint32
	RequestsDisconnected         uint32
	RequestsInNativeQueue        uint32
	RequestsQueued               uint32
	RequestsRejected             uint32
	RequestWaitTime              uint32
	StateServerSessionsAbandoned uint32
	StateServerSessionsActive    uint32
	StateServerSessionsTimedOut  uint32
	StateServerSessionsTotal     uint32
	Timestamp_Object             uint64
	Timestamp_PerfTime           uint64
	Timestamp_Sys100NS           uint64
	WorkerProcessesRunning       uint32
	WorkerProcessRestarts        uint32
}

func (c *ASPNETCollector) collect(ch chan<- prometheus.Metric) (*prometheus.Desc, error) {
	var dst []Win32_PerfRawData_ASPNET4030319_ASPNETv4030319

	q := queryAll(&dst)
	if err := wmi.Query(q, &dst); err != nil {
		return nil, err
	}

	if len(dst) == 0 {
		return nil, errors.New("WMI query returned empty result set")
	}

	for _, aspnetitem := range dst {
		// Gauges
		// ch <- prometheus.MustNewConstMetric(
		// 	c.ApplicationRestarts,
		// 	prometheus.GaugeValue,
		// 	float64(aspnetitem.ApplicationRestarts),
		// 	aspnetitem.Name,
		// )
		// Counters
		ch <- prometheus.MustNewConstMetric(
			c.ApplicationRestarts,
			prometheus.CounterValue,
			float64(aspnetitem.ApplicationRestarts),
			aspnetitem.Name,
		)
	}

	return nil, nil
}
