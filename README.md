# agent-tester

<div align="center">

[![Go](https://github.com/mixdjoker/agent-tester/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/mixdjoker/agent-tester/actions/workflows/go.yml)
[![CodeQL](https://github.com/mixdjoker/agent-tester/actions/workflows/codeql.yml/badge.svg)](https://github.com/mixdjoker/agent-tester/actions/workflows/codeql.yml)

</div>

Agent service that run tests, save last results in JSON, send result to agent-connector

## REST-API

## Under-the-hood

## *UDP Buffer Sizes*
Experiments have shown that QUIC transfers on high-bandwidth connections can be limited by the size of the UDP receive and send buffer. The receive buffer holds packets that have been received by the kernel, but not yet read by the application (quic-go in this case). The send buffer holds packets that have been sent by quic-go, but not sent out by the kernel. In both cases, once these buffers fill up, the kernel will drop any new incoming packet.

It is recommended to increase the maximum buffer size by running:

```
sysctl -w net.core.rmem_max=7500000
sysctl -w net.core.wmem_max=7500000
```

This command would increase the maximum send and the receive buffer size to roughly 7.5 MB. Note that these settings are not persisted across reboots.

*Persistent changes*:
In file `/etc/sysctl.d/99-quic.conf` add or change:

```
net.core.rmem_max=7500000
net.core.wmem_max=7500000
```

Reread new configuration:

```
sudo sysctl --system
```
