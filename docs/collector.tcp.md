# tcp collector

The tcp collector exposes metrics about the TCP/IPv4 network stack.

|||
-|-
Metric name prefix  | `tcp`
Classes             | [`Win32_PerfRawData_Tcpip_TCPv4`](https://msdn.microsoft.com/en-us/library/aa394341(v=vs.85).aspx)
Enabled by default? | No

## Flags

None

## Metrics

Name | Description | Type | Labels
-----|-------------|------|-------
`wmi_tcp_connection_failures` | Number of times TCP connections have made a direct transition to the CLOSED state from the SYN-SENT state or the SYN-RCVD state, plus the number of times TCP connections have made a direct transition from the SYN-RCVD state to the LISTEN state | counter | None
`wmi_tcp_connections_active` |  Number of times TCP connections have made a direct transition from the CLOSED state to the SYN-SENT state.| counter | None
`wmi_tcp_connections_established` | Number of TCP connections for which the current state is either ESTABLISHED or CLOSE-WAIT. | counter | None
`wmi_tcp_connections_passive` | Number of times TCP connections have made a direct transition from the LISTEN state to the SYN-RCVD state. | counter | None
`wmi_tcp_connections_reset` | Number of times TCP connections have made a direct transition from the LISTEN state to the SYN-RCVD state. | counter | None
`wmi_tcp_segments_total` | Total segments sent or received using the TCP protocol | counter | None
`wmi_tcp_segments_received_total` | Total segments received, including those received in error. This count includes segments received on currently established connections | counter | None
`wmi_tcp_segments_retransmitted_total` | Total segments retransmitted. That is, segments transmitted that contain one or more previously transmitted bytes | counter | None
`wmi_tcp_segments_sent_total` | Total segments sent, including those on current connections, but excluding those containing *only* retransmitted bytes | counter | None

### Example metric
_This collector does not yet have explained examples, we would appreciate your help adding them!_

## Useful queries
_This collector does not yet have any useful queries added, we would appreciate your help adding them!_

## Alerting examples
_This collector does not yet have alerting examples, we would appreciate your help adding them!_
