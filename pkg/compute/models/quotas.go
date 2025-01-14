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

package models

import (
	"context"
	"fmt"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/tristate"

	api "yunion.io/x/onecloud/pkg/apis/compute"
	identityapi "yunion.io/x/onecloud/pkg/apis/identity"
	"yunion.io/x/onecloud/pkg/cloudcommon/db"
	"yunion.io/x/onecloud/pkg/cloudcommon/db/quotas"
	commonOptions "yunion.io/x/onecloud/pkg/cloudcommon/options"
	"yunion.io/x/onecloud/pkg/compute/options"
	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/mcclient/auth"
	"yunion.io/x/onecloud/pkg/mcclient/utils"
	"yunion.io/x/onecloud/pkg/util/rbacutils"
)

type SQuotaManager struct {
	quotas.SQuotaBaseManager
}

var (
	Quota                    SQuota
	QuotaManager             *SQuotaManager
	QuotaUsageManager        *SQuotaManager
	QuotaPendingUsageManager *SQuotaManager
)

func init() {
	Quota = SQuota{}

	QuotaUsageManager = &SQuotaManager{
		SQuotaBaseManager: quotas.NewQuotaUsageManager(Quota,
			rbacutils.ScopeProject,
			"quota_usage_tbl",
			"quota_usage",
			"quota_usages",
		),
	}
	QuotaPendingUsageManager = &SQuotaManager{
		SQuotaBaseManager: quotas.NewQuotaUsageManager(Quota,
			rbacutils.ScopeProject,
			"quota_pending_usage_tbl",
			"quota_pending_usage",
			"quota_pending_usages",
		),
	}
	QuotaManager = &SQuotaManager{
		SQuotaBaseManager: quotas.NewQuotaBaseManager(Quota,
			rbacutils.ScopeProject,
			"quota_tbl",
			QuotaPendingUsageManager,
			QuotaUsageManager,
			"quota",
			"quotas",
		),
	}
	quotas.Register(QuotaManager)
}

type SQuota struct {
	quotas.SQuotaBase

	SComputeResourceKeys

	// 主机数量配额
	Count int `default:"-1" allow_zero:"true" json:"count"`
	// 主机CPU核数量配额
	Cpu int `default:"-1" allow_zero:"true" json:"cpu"`
	// 主机内存容量配额
	Memory int `default:"-1" allow_zero:"true" json:"memory"`
	// 主机存储容量配额
	Storage int `default:"-1" allow_zero:"true" json:"storage"`

	// 主机组配额
	Group int `default:"-1" allow_zero:"true" json:"group"`
	// 直通设备(GPU)配额
	IsolatedDevice int `default:"-1" allow_zero:"true" json:"isolated_device"`
}

func (self *SQuota) GetKeys() quotas.IQuotaKeys {
	return self.SComputeResourceKeys
}

func (self *SQuota) SetKeys(keys quotas.IQuotaKeys) {
	self.SComputeResourceKeys = keys.(SComputeResourceKeys)
}

func (self *SQuota) FetchSystemQuota() {
	keys := self.SComputeResourceKeys
	base := 0
	switch options.Options.DefaultQuotaValue {
	case commonOptions.DefaultQuotaUnlimit:
		base = -1
	case commonOptions.DefaultQuotaZero:
		base = 0
		if keys.Scope() == rbacutils.ScopeDomain { // domain level quota
			base = 10
		} else if keys.DomainId == identityapi.DEFAULT_DOMAIN_ID && keys.ProjectId == auth.AdminCredential().GetProjectId() {
			base = 1
		}
	case commonOptions.DefaultQuotaDefault:
		base = 1
		if keys.Scope() == rbacutils.ScopeDomain {
			base = 10
		}
	}
	defaultValue := func(def int) int {
		if base < 0 {
			return -1
		} else {
			return def * base
		}
	}
	self.Count = defaultValue(options.Options.DefaultServerQuota)
	self.Cpu = defaultValue(options.Options.DefaultCpuQuota)
	self.Memory = defaultValue(options.Options.DefaultMemoryQuota)
	self.Storage = defaultValue(options.Options.DefaultStorageQuota)
	self.Group = defaultValue(options.Options.DefaultGroupQuota)
	self.IsolatedDevice = defaultValue(options.Options.DefaultIsolatedDeviceQuota)
}

func (self *SQuota) FetchUsage(ctx context.Context) error {
	keys := self.SComputeResourceKeys

	scope := keys.Scope()
	ownerId := keys.OwnerId()

	rangeObjs := make([]db.IStandaloneModel, 0)
	if len(keys.ManagerId) > 0 {
		obj, err := CloudproviderManager.FetchById(keys.ManagerId)
		if err != nil {
			return errors.Wrap(err, "CloudproviderManager.FetchById")
		}
		rangeObjs = append(rangeObjs, obj.(db.IStandaloneModel))
	} else if len(keys.AccountId) > 0 {
		obj, err := CloudaccountManager.FetchById(keys.AccountId)
		if err != nil {
			return errors.Wrap(err, "CloudaccountManager.FetchById")
		}
		rangeObjs = append(rangeObjs, obj.(db.IStandaloneModel))
	}

	if len(keys.ZoneId) > 0 {
		obj, err := ZoneManager.FetchById(keys.ZoneId)
		if err != nil {
			return errors.Wrap(err, "ZoneManager.FetchById")
		}
		rangeObjs = append(rangeObjs, obj.(db.IStandaloneModel))
	} else if len(keys.RegionId) > 0 {
		obj, err := CloudregionManager.FetchById(keys.RegionId)
		if err != nil {
			return errors.Wrap(err, "CloudregionManager.FetchById")
		}
		rangeObjs = append(rangeObjs, obj.(db.IStandaloneModel))
	}
	var hypervisors []string
	if len(keys.Hypervisor) > 0 {
		hypervisors = []string{keys.Hypervisor}
	}
	var providers []string
	if len(keys.Provider) > 0 {
		providers = []string{keys.Provider}
	}
	var brands []string
	if len(keys.Brand) > 0 {
		brands = []string{keys.Brand}
	}

	diskSize := totalDiskSize(scope, ownerId, tristate.None, tristate.None, false, false, rangeObjs, providers, brands, keys.CloudEnv, hypervisors)

	guest := usageTotalGuestResouceCount(scope, ownerId, rangeObjs, nil, hypervisors, false, false, nil, nil, providers, brands, keys.CloudEnv, nil, rbacutils.SPolicyResult{})

	self.Count = guest.TotalGuestCount
	self.Cpu = guest.TotalCpuCount
	self.Memory = guest.TotalMemSize
	self.Storage = diskSize
	self.Group = 0
	self.IsolatedDevice = guest.TotalIsolatedCount
	return nil
}

func (self *SQuota) ResetNegative() {
	if self.Count < 0 {
		self.Count = 0
	}
	if self.Cpu < 0 {
		self.Cpu = 0
	}
	if self.Memory < 0 {
		self.Memory = 0
	}
	if self.Storage < 0 {
		self.Storage = 0
	}
	if self.Group < 0 {
		self.Group = 0
	}
	if self.IsolatedDevice < 0 {
		self.IsolatedDevice = 0
	}
}

func (self *SQuota) IsEmpty() bool {
	if self.Count > 0 {
		return false
	}
	if self.Cpu > 0 {
		return false
	}
	if self.Memory > 0 {
		return false
	}
	if self.Storage > 0 {
		return false
	}
	if self.Group > 0 {
		return false
	}
	if self.IsolatedDevice > 0 {
		return false
	}
	return true
}

func (self *SQuota) Add(quota quotas.IQuota) {
	squota := quota.(*SQuota)
	self.Count = self.Count + quotas.NonNegative(squota.Count)
	self.Cpu = self.Cpu + quotas.NonNegative(squota.Cpu)
	self.Memory = self.Memory + quotas.NonNegative(squota.Memory)
	self.Storage = self.Storage + quotas.NonNegative(squota.Storage)
	self.Group = self.Group + quotas.NonNegative(squota.Group)
	self.IsolatedDevice = self.IsolatedDevice + quotas.NonNegative(squota.IsolatedDevice)
}

func nonNegative(val int) int {
	return quotas.NonNegative(val)
}

func (self *SQuota) Sub(quota quotas.IQuota) {
	squota := quota.(*SQuota)
	self.Count = nonNegative(self.Count - squota.Count)
	self.Cpu = nonNegative(self.Cpu - squota.Cpu)
	self.Memory = nonNegative(self.Memory - squota.Memory)
	self.Storage = nonNegative(self.Storage - squota.Storage)
	self.Group = nonNegative(self.Group - squota.Group)
	self.IsolatedDevice = nonNegative(self.IsolatedDevice - squota.IsolatedDevice)
}

func (self *SQuota) Allocable(request quotas.IQuota) int {
	squota := request.(*SQuota)
	cnt := -1
	if self.Count >= 0 && squota.Count > 0 && (cnt < 0 || cnt > self.Count/squota.Count) {
		cnt = self.Count / squota.Count
	}
	if self.Cpu >= 0 && squota.Cpu > 0 && (cnt < 0 || cnt > self.Cpu/squota.Cpu) {
		cnt = self.Cpu / squota.Cpu
	}
	if self.Memory >= 0 && squota.Memory > 0 && (cnt < 0 || cnt > self.Memory/squota.Memory) {
		cnt = self.Memory / squota.Memory
	}
	if self.Storage >= 0 && squota.Storage > 0 && (cnt < 0 || cnt > self.Storage/squota.Storage) {
		cnt = self.Storage / squota.Storage
	}
	if self.Group >= 0 && squota.Group > 0 && (cnt < 0 || cnt > self.Group/squota.Group) {
		cnt = self.Group / squota.Group
	}
	if self.IsolatedDevice >= 0 && squota.IsolatedDevice > 0 && (cnt < 0 || cnt > self.IsolatedDevice/squota.IsolatedDevice) {
		cnt = self.IsolatedDevice / squota.IsolatedDevice
	}
	return cnt
}

func (self *SQuota) Update(quota quotas.IQuota) {
	squota := quota.(*SQuota)
	if squota.Count > 0 {
		self.Count = squota.Count
	}
	if squota.Cpu > 0 {
		self.Cpu = squota.Cpu
	}
	if squota.Memory > 0 {
		self.Memory = squota.Memory
	}
	if squota.Storage > 0 {
		self.Storage = squota.Storage
	}
	if squota.Group > 0 {
		self.Group = squota.Group
	}
	if squota.IsolatedDevice > 0 {
		self.IsolatedDevice = squota.IsolatedDevice
	}
}

func (used *SQuota) Exceed(request quotas.IQuota, quota quotas.IQuota) error {
	err := quotas.NewOutOfQuotaError()
	sreq := request.(*SQuota)
	squota := quota.(*SQuota)
	if quotas.Exceed(used.Count, sreq.Count, squota.Count) {
		err.Add(used, "count", squota.Count, used.Count, sreq.Count)
	}
	if quotas.Exceed(used.Cpu, sreq.Cpu, squota.Cpu) {
		err.Add(used, "cpu", squota.Cpu, used.Cpu, sreq.Cpu)
	}
	if quotas.Exceed(used.Memory, sreq.Memory, squota.Memory) {
		err.Add(used, "memory", squota.Memory, used.Memory, sreq.Memory)
	}
	if quotas.Exceed(used.Storage, sreq.Storage, squota.Storage) {
		err.Add(used, "storage", squota.Storage, used.Storage, sreq.Storage)
	}
	if quotas.Exceed(used.Group, sreq.Group, squota.Group) {
		err.Add(used, "group", squota.Group, used.Group, sreq.Group)
	}
	if quotas.Exceed(used.IsolatedDevice, sreq.IsolatedDevice, squota.IsolatedDevice) {
		err.Add(used, "isolated_device", squota.IsolatedDevice, used.IsolatedDevice, sreq.IsolatedDevice)
	}
	if err.IsError() {
		return err
	} else {
		return nil
	}
}

func keyName(prefix, name string) string {
	if len(prefix) > 0 {
		return fmt.Sprintf("%s.%s", prefix, name)
	} else {
		return name
	}
}

func (self *SQuota) ToJSON(prefix string) jsonutils.JSONObject {
	ret := jsonutils.NewDict()
	ret.Add(jsonutils.NewInt(int64(self.Count)), keyName(prefix, "count"))
	ret.Add(jsonutils.NewInt(int64(self.Cpu)), keyName(prefix, "cpu"))
	ret.Add(jsonutils.NewInt(int64(self.Memory)), keyName(prefix, "memory"))
	ret.Add(jsonutils.NewInt(int64(self.Storage)), keyName(prefix, "storage"))
	ret.Add(jsonutils.NewInt(int64(self.Group)), keyName(prefix, "group"))
	ret.Add(jsonutils.NewInt(int64(self.IsolatedDevice)), keyName(prefix, "isolated_device"))
	return ret
}

func (manager *SQuotaManager) FetchIdNames(ctx context.Context, idMap map[string]map[string]string) (map[string]map[string]string, error) {
	for field := range idMap {
		switch field {
		case "domain_id":
			fieldIdMap, err := utils.FetchDomainNames(ctx, idMap[field])
			if err != nil {
				return nil, errors.Wrap(err, "utils.FetchDomainNames")
			}
			idMap[field] = fieldIdMap
		case "tenant_id":
			fieldIdMap, err := utils.FetchTenantNames(ctx, idMap[field])
			if err != nil {
				return nil, errors.Wrap(err, "utils.FetchTenantNames")
			}
			idMap[field] = fieldIdMap
		case "region_id":
			fieldIdMap, err := fetchRegionNames(idMap[field])
			if err != nil {
				return nil, errors.Wrap(err, "fetchRegionNames")
			}
			idMap[field] = fieldIdMap
		case "zone_id":
			fieldIdMap, err := fetchZoneNames(idMap[field])
			if err != nil {
				return nil, errors.Wrap(err, "fetchZoneNames")
			}
			idMap[field] = fieldIdMap
		case "account_id":
			fieldIdMap, err := fetchAccountNames(idMap[field])
			if err != nil {
				return nil, errors.Wrap(err, "fetchAccountNames")
			}
			idMap[field] = fieldIdMap
		case "manager_id":
			fieldIdMap, err := fetchManagerNames(idMap[field])
			if err != nil {
				return nil, errors.Wrap(err, "fetchManagerNames")
			}
			idMap[field] = fieldIdMap
		}
	}
	return idMap, nil
}

func fetchRegionNames(idMap map[string]string) (map[string]string, error) {
	return db.FetchIdNameMap(CloudregionManager, idMap)
}

func fetchZoneNames(idMap map[string]string) (map[string]string, error) {
	return db.FetchIdNameMap(ZoneManager, idMap)
}

func fetchAccountNames(idMap map[string]string) (map[string]string, error) {
	return db.FetchIdNameMap(CloudaccountManager, idMap)
}

func fetchManagerNames(idMap map[string]string) (map[string]string, error) {
	return db.FetchIdNameMap(CloudproviderManager, idMap)
}

type SComputeResourceKeys struct {
	quotas.SZonalCloudResourceKeys

	// 主机配额适用的主机类型，参考主机List的Hypervisor列表
	Hypervisor string `width:"16" charset:"ascii" nullable:"false" primary:"true" list:"user"`
}

func (k SComputeResourceKeys) Fields() []string {
	return append(k.SZonalCloudResourceKeys.Fields(), "hypervisor")
}

func (k SComputeResourceKeys) Values() []string {
	return append(k.SZonalCloudResourceKeys.Values(), k.Hypervisor)
}

func (k1 SComputeResourceKeys) Compare(ik quotas.IQuotaKeys) int {
	k2 := ik.(SComputeResourceKeys)
	r := k1.SZonalCloudResourceKeys.Compare(k2.SZonalCloudResourceKeys)
	if r != 0 {
		return r
	}
	if k1.Hypervisor < k2.Hypervisor {
		return -1
	} else if k1.Hypervisor > k2.Hypervisor {
		return 1
	}
	return 0
}

func fetchCloudQuotaKeys(scope rbacutils.TRbacScope, ownerId mcclient.IIdentityProvider, manager *SCloudprovider) quotas.SCloudResourceKeys {
	keys := quotas.SCloudResourceKeys{}
	keys.SBaseProjectQuotaKeys = quotas.OwnerIdProjectQuotaKeys(scope, ownerId)
	if manager != nil {
		account := manager.GetCloudaccount()
		keys.Provider = account.Provider
		keys.Brand = account.Brand
		keys.CloudEnv = account.GetCloudEnv()
		keys.AccountId = account.Id
		keys.ManagerId = manager.Id
	} else {
		keys.Provider = api.CLOUD_PROVIDER_ONECLOUD
		keys.Brand = api.ONECLOUD_BRAND_ONECLOUD
		keys.CloudEnv = api.CLOUD_ENV_ON_PREMISE
	}
	return keys
}

func fetchRegionalQuotaKeys(scope rbacutils.TRbacScope, ownerId mcclient.IIdentityProvider, region *SCloudregion, manager *SCloudprovider) quotas.SRegionalCloudResourceKeys {
	keys := quotas.SRegionalCloudResourceKeys{}
	keys.SCloudResourceKeys = fetchCloudQuotaKeys(scope, ownerId, manager)
	if region != nil {
		keys.RegionId = region.Id
	}
	return keys
}

func fetchZonalQuotaKeys(scope rbacutils.TRbacScope, ownerId mcclient.IIdentityProvider, zone *SZone, manager *SCloudprovider) quotas.SZonalCloudResourceKeys {
	keys := quotas.SZonalCloudResourceKeys{}
	keys.SCloudResourceKeys = fetchCloudQuotaKeys(scope, ownerId, manager)
	if zone != nil {
		keys.RegionId = zone.CloudregionId
		keys.ZoneId = zone.Id
	}
	return keys
}

func fetchComputeQuotaKeys(scope rbacutils.TRbacScope, ownerId mcclient.IIdentityProvider, zone *SZone, manager *SCloudprovider, hypervisor string) SComputeResourceKeys {
	keys := SComputeResourceKeys{}
	keys.SZonalCloudResourceKeys = fetchZonalQuotaKeys(scope, ownerId, zone, manager)
	keys.Hypervisor = hypervisor
	return keys
}
