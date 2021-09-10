// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package qcmon

const (
	PERIOD             = 60
	UNIT_AVERAGE       = "Average"
	DEFAULT_STATISTICS = "Average,Minimum,Maximum"
	UNIT_PERCENT       = "Percent"
	UNIT_BPS           = "bps"
	UNIT_MBPS          = "Mbps"
	UNIT_BYTEPS        = "Bps"
	UNIT_CPS           = "cps"
	UNIT_COUNT         = "count"
	UNIT_MEM           = "byte"
	UNIT_MSEC          = "ms"
	UNIT_COUNT_SEC     = "count/s"

	//ESC监控指标
	INFLUXDB_FIELD_CPU_USAGE           = "vm_cpu.usage_active"
	INFLUXDB_FIELD_MEM_USAGE           = "vm_mem.used_percent"
	INFLUXDB_FIELD_DISK_READ_BPS       = "vm_diskio.read_bps"
	INFLUXDB_FIELD_DISK_WRITE_BPS      = "vm_diskio.write_bps"
	INFLUXDB_FIELD_DISK_READ_IOPS      = "vm_diskio.read_iops"
	INFLUXDB_FIELD_DISK_WRITE_IOPS     = "vm_diskio.write_iops"
	INFLUXDB_FIELD_NET_BPS_RX          = "vm_netio.bps_recv"
	INFLUXDB_FIELD_NET_BPS_RX_INTERNET = INFLUXDB_FIELD_NET_BPS_RX + ",net_type=internet"
	INFLUXDB_FIELD_NET_BPS_RX_INTRANET = INFLUXDB_FIELD_NET_BPS_RX + ",net_type=intranet"
	INFLUXDB_FIELD_NET_BPS_TX          = "vm_netio.bps_sent"
	INFLUXDB_FIELD_NET_BPS_TX_INTERNET = INFLUXDB_FIELD_NET_BPS_TX + ",net_type=internet"
	INFLUXDB_FIELD_NET_BPS_TX_INTRANET = INFLUXDB_FIELD_NET_BPS_TX + ",net_type=intranet"
	INFLUXDB_FIELD_WANOUTTRAFFIC       = "vm_eipio.bps_out"
	INFLUXDB_FIELD_WANINTRAFFIC        = "vm_eipio.bps_in"
	INFLUXDB_FIELD_WANOUTPKG           = "vm_eipio.pps_out"
	INFLUXDB_FIELD_WANINPKG            = "vm_eipio.pps_in"

	//RDS监控指标
	INFLUXDB_FIELD_RDS_CPU_USAGE            = "rds_cpu.usage_active"
	INFLUXDB_FIELD_RDS_MEM_USAGE            = "rds_mem.used_percent"
	INFLUXDB_FIELD_RDS_NET_BPS_RX           = "rds_netio.bps_recv"
	INFLUXDB_FIELD_RDS_NET_BPS_RX_INTRANET  = INFLUXDB_FIELD_RDS_NET_BPS_RX + ",net_type=intranet"
	INFLUXDB_FIELD_RDS_NET_BPS_RX_MYSQL     = INFLUXDB_FIELD_RDS_NET_BPS_RX + ",server_type=mysql"
	INFLUXDB_FIELD_RDS_NET_BPS_RX_SQLSERVER = INFLUXDB_FIELD_RDS_NET_BPS_RX + ",server_type=sqlserver"
	INFLUXDB_FIELD_RDS_NET_BPS_TX           = "rds_netio.bps_sent"
	INFLUXDB_FIELD_RDS_NET_BPS_TX_INTRANET  = INFLUXDB_FIELD_RDS_NET_BPS_TX + ",net_type=intranet"
	INFLUXDB_FIELD_RDS_NET_BPS_TX_MYSQL     = INFLUXDB_FIELD_RDS_NET_BPS_TX + ",server_type=mysql"
	INFLUXDB_FIELD_RDS_NET_BPS_TX_SQLSERVER = INFLUXDB_FIELD_RDS_NET_BPS_TX + ",server_type=sqlserver"
	INFLUXDB_FIELD_RDS_DISK_USAGE           = "rds_disk.used_percent"
	INFLUXDB_FIELD_RDS_DISK_READ_BPS        = "rds_diskio.read_bps"
	INFLUXDB_FIELD_RDS_DISK_WRITE_BPS       = "rds_diskio.write_bps"
	INFLUXDB_FIELD_RDS_CONN_COUNT           = "rds_conn.used_count"
	INFLUXDB_FIELD_RDS_CONN_USAGE           = "rds_conn.used_percent"

	INFLUXDB_FIELD_RDS_QPS              = "rds_qps.query_qps"
	INFLUXDB_FIELD_RDS_TPS              = "rds_tps.trans_qps"
	INFLUXDB_FIELD_RDS_INNODB_REDA_BPS  = "rds_innodb.read_bps"
	INFLUXDB_FIELD_RDS_INNODB_WRITE_BPS = "rds_innodb.write_bps"

	//REDIS监控指标
	INFLUXDB_FIELD_REDIS_CPU_USAGE      = "dcs_cpu.usage_active"
	INFLUXDB_FIELD_REDIS_MEM_USAGE      = "dcs_mem.used_percent"
	INFLUXDB_FIELD_REDIS_NET_BPS_RX     = "dcs_netio.bps_recv"
	INFLUXDB_FIELD_REDIS_NET_BPS_TX     = "dcs_netio.bps_sent"
	INFLUXDB_FIFLD_REDIS_CONN_USAGE     = "dcs_conn.used_percent"
	INFLUXDB_FIFLD_REDIS_OPT_SES        = "dcs_instantopt.opt_sec"
	INFLUXDB_FIFLD_REDIS_CACHE_KEYS     = "dcs_cachekeys.key_count"
	INFLUXDB_FIFLD_REDIS_CACHE_EXP_KEYS = INFLUXDB_FIFLD_REDIS_CACHE_KEYS + ",exp=expire"
	INFLUXDB_FIFLD_REDIS_DATA_MEM_USAGE = "dcs_datamem.used_byte"

	//对象存储OSS监控指标
	INFLUXDB_FIELD_OSS_NET_BPS_RX          = "oss_netio.bps_recv"
	INFLUXDB_FIELD_OSS_NET_BPS_RX_INTERNET = INFLUXDB_FIELD_OSS_NET_BPS_RX + ",net_type=internet"
	INFLUXDB_FIELD_OSS_NET_BPS_RX_INTRANET = INFLUXDB_FIELD_OSS_NET_BPS_RX + ",net_type=intranet"
	INFLUXDB_FIELD_OSS_NET_BPS_TX          = "oss_netio.bps_sent"
	INFLUXDB_FIELD_OSS_NET_BPS_TX_INTERNET = INFLUXDB_FIELD_OSS_NET_BPS_TX + ",net_type=internet"
	INFLUXDB_FIELD_OSS_NET_BPS_TX_INTRANET = INFLUXDB_FIELD_OSS_NET_BPS_TX + ",net_type=intranet"
	INFLUXDB_FIELD_OSS_LATECY              = "oss_latency.req_late"
	INFLUXDB_FIELD_OSS_LATECY_GET          = INFLUXDB_FIELD_OSS_LATECY + ",request=get"
	INFLUXDB_FIELD_OSS_LATECY_POST         = INFLUXDB_FIELD_OSS_LATECY + ",request=post"
	INFLUXDB_FIELD_OSS_REQ_COUNT           = "oss_req.req_count"
	INFLUXDB_FIELD_OSS_REQ_COUNT_GET       = INFLUXDB_FIELD_OSS_REQ_COUNT + ",request=get"
	INFLUXDB_FIELD_OSS_REQ_COUNT_POST      = INFLUXDB_FIELD_OSS_REQ_COUNT + ",request=post"
	INFLUXDB_FIELD_OSS_REQ_COUNT_5XX       = INFLUXDB_FIELD_OSS_REQ_COUNT + ",request=5xx"
	INFLUXDB_FIELD_OSS_REQ_COUNT_4XX       = INFLUXDB_FIELD_OSS_REQ_COUNT + ",request=4xx"

	//负载均衡监控指标
	INFLUXDB_FIELD_ELB_NET_BPS_RX     = "haproxy.bin"
	INFLUXDB_FIELD_ELB_NET_BPS_TX     = "haproxy.bout"
	INFLUXDB_FIELD_ELB_REQ_RATE       = "haproxy.req_rate,request=http"
	INFLUXDB_FIELD_ELB_CONN_RATE      = "haproxy.conn_rate,request=tcp"
	INFLUXDB_FIELD_ELB_DREQ_COUNT     = "haproxy.dreq,request=http"
	INFLUXDB_FIELD_ELB_DCONN_COUNT    = "haproxy.dcon,request=tcp"
	INFLUXDB_FIELD_ELB_HRSP_COUNT     = "haproxy.hrsp_Nxx"
	INFLUXDB_FIELD_ELB_HRSP_COUNT_2XX = INFLUXDB_FIELD_ELB_HRSP_COUNT + ",request=2xx"
	INFLUXDB_FIELD_ELB_HRSP_COUNT_3XX = INFLUXDB_FIELD_ELB_HRSP_COUNT + ",request=3xx"
	INFLUXDB_FIELD_ELB_HRSP_COUNT_4XX = INFLUXDB_FIELD_ELB_HRSP_COUNT + ",request=4xx"
	INFLUXDB_FIELD_ELB_HRSP_COUNT_5XX = INFLUXDB_FIELD_ELB_HRSP_COUNT + ",request=5xx"
	INFLUXDB_FIELD_ELB_CHC_STATUS     = "haproxy.check_status"
	INFLUXDB_FIELD_ELB_CHC_CODE       = "haproxy.check_code"
	INFLUXDB_FIELD_ELB_LAST_CHC       = "haproxy.last_chk"

	KEY_VMS   = "vms"
	KEY_CPUS  = "cpus"
	KEY_MEMS  = "mems"
	KEY_DISKS = "disks"

	KEY_LIMIT  = "limit"
	KEY_ADMIN  = "admin"
	KEY_USABLE = "usable"
)

var tecentMetricSpecs = map[string][]string{
	"CPUUsage":      {DEFAULT_STATISTICS, UNIT_PERCENT, INFLUXDB_FIELD_CPU_USAGE},
	"MemUsage":      {DEFAULT_STATISTICS, UNIT_PERCENT, INFLUXDB_FIELD_MEM_USAGE},
	"lanOuttraffic": {DEFAULT_STATISTICS, UNIT_MBPS, INFLUXDB_FIELD_NET_BPS_TX_INTRANET},
	"WanOuttraffic": {DEFAULT_STATISTICS, UNIT_MBPS, INFLUXDB_FIELD_NET_BPS_TX_INTERNET},
	"lanIntraffic":  {DEFAULT_STATISTICS, UNIT_MBPS, INFLUXDB_FIELD_NET_BPS_RX_INTRANET},
	"WanIntraffic":  {DEFAULT_STATISTICS, UNIT_MBPS, INFLUXDB_FIELD_NET_BPS_RX_INTERNET},
}

var tecentRedisMetricSpecs = map[string][]string{
	"CpuUsMin":         {DEFAULT_STATISTICS, UNIT_PERCENT, INFLUXDB_FIELD_REDIS_CPU_USAGE},
	"StorageUsMin":     {DEFAULT_STATISTICS, UNIT_PERCENT, INFLUXDB_FIELD_REDIS_MEM_USAGE},
	"InFlowMin":        {DEFAULT_STATISTICS, UNIT_BPS, INFLUXDB_FIELD_REDIS_NET_BPS_RX},
	"OutFlowMin":       {DEFAULT_STATISTICS, UNIT_BPS, INFLUXDB_FIELD_REDIS_NET_BPS_TX},
	"ConnectionsUsMin": {DEFAULT_STATISTICS, UNIT_COUNT, INFLUXDB_FIFLD_REDIS_CONN_USAGE},
	"QpsMin":           {DEFAULT_STATISTICS, UNIT_COUNT, INFLUXDB_FIFLD_REDIS_OPT_SES},
	"KeysMin":          {DEFAULT_STATISTICS, UNIT_COUNT, INFLUXDB_FIFLD_REDIS_CACHE_KEYS},
	"ExpiredKeysMin":   {DEFAULT_STATISTICS, UNIT_COUNT, INFLUXDB_FIFLD_REDIS_CACHE_EXP_KEYS},
	"StorageMin":       {DEFAULT_STATISTICS, UNIT_MEM, INFLUXDB_FIFLD_REDIS_DATA_MEM_USAGE},
}

var tecentRdsMetricSpecs = map[string][]string{
	"CPUUseRate":        {DEFAULT_STATISTICS, UNIT_PERCENT, INFLUXDB_FIELD_RDS_CPU_USAGE},
	"MemoryUseRate":     {DEFAULT_STATISTICS, UNIT_PERCENT, INFLUXDB_FIELD_RDS_MEM_USAGE},
	"BytesSent":         {DEFAULT_STATISTICS, UNIT_BYTEPS, INFLUXDB_FIELD_RDS_NET_BPS_TX_INTRANET},
	"BytesReceived":     {DEFAULT_STATISTICS, UNIT_BYTEPS, INFLUXDB_FIELD_RDS_NET_BPS_RX_INTRANET},
	"VolumeRate":        {DEFAULT_STATISTICS, UNIT_PERCENT, INFLUXDB_FIELD_RDS_DISK_USAGE},
	"ThreadsConnected":  {DEFAULT_STATISTICS, UNIT_COUNT, INFLUXDB_FIELD_RDS_CONN_COUNT},
	"ConnectionUseRate": {DEFAULT_STATISTICS, UNIT_PERCENT, INFLUXDB_FIELD_RDS_CONN_USAGE},
	"QPS":               {DEFAULT_STATISTICS, UNIT_COUNT_SEC, INFLUXDB_FIELD_RDS_QPS},
	"TPS":               {DEFAULT_STATISTICS, UNIT_COUNT_SEC, INFLUXDB_FIELD_RDS_TPS},
	"InnodbDataRead":    {DEFAULT_STATISTICS, UNIT_BYTEPS, INFLUXDB_FIELD_RDS_INNODB_REDA_BPS},
	"InnodbDataWritten": {DEFAULT_STATISTICS, UNIT_BYTEPS, INFLUXDB_FIELD_RDS_INNODB_WRITE_BPS},
}