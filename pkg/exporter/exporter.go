/*
Copyright 2020, Staffbase GmbH and contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package exporter

import (
	"time"

	"github.com/syseleven/syseleven-exporter/pkg/api"
	"github.com/syseleven/syseleven-exporter/pkg/auth"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

type Exporter struct {
	Username    string
	Password    string
	ProjectID   string
	UseAppCreds bool
}

func New(projectID string, useAppCreds bool, username string, password string) (*Exporter, error) {
	return &Exporter{
		ProjectID:   projectID,
		UseAppCreds: useAppCreds,
		Username:    username,
		Password:    password,
	}, nil
}

// set utils for API v3 for Quota and Usage Information

func SetUtilsV3(quota map[string]api.QuotaV3, usage map[string]api.CurrentUsageV3, exporter *Exporter) {
	computeCoresTotal.Reset()
	computeInstancesTotal.Reset()
	computeRamTotalMegabytes.Reset()
	dnsZonesTotal.Reset()
	networkFloatingIPsTotal.Reset()
	networkLoadbalancersTotal.Reset()
	s3SpaceTotalBytes.Reset()
	volumeSpaceTotalGigabytes.Reset()
	volumeVolumesTotalGigabytes.Reset()
	computeCoresUsed.Reset()
	computeInstancesUsed.Reset()
	computeRamUsedMegabytes.Reset()
	dnsZonesUsed.Reset()
	networkFloatingIPsUsed.Reset()
	networkLoadbalancersUsed.Reset()
	s3SpaceUsedBytes.Reset()
	volumeSpaceUsedGigabytes.Reset()
	volumeVolumesUsedGigabytes.Reset()
	computeFlavorsUsed.Reset()

	for k, v := range quota {
		computeCoresTotal.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.ComputeCores)
		computeInstancesTotal.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.ComputeInstances)
		computeRamTotalMegabytes.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.ComputeRAMMb)
		dnsZonesTotal.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.DNSZones)
		networkFloatingIPsTotal.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.NetworkFloatingips)
		networkLoadbalancersTotal.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.LoadbalancerLoadbalancers)
		volumeSpaceTotalGigabytes.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.VolumeSpaceGb)
		volumeVolumesTotalGigabytes.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.VolumeVolumes)

		for _, os := range v.Objectstorage {
			s3SpaceTotalBytes.With(prometheus.Labels{"region": k, "project": exporter.ProjectID, "type": os.Type}).Set(os.SpaceBytes)
		}

	}

	for k, v := range usage {
		computeCoresUsed.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.ComputeCores)
		computeInstancesUsed.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.ComputeInstances)
		computeRamUsedMegabytes.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.ComputeRAMMb)
		dnsZonesUsed.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.DNSZones)
		networkFloatingIPsUsed.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.NetworkFloatingips)
		networkLoadbalancersUsed.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.LoadbalancerLoadbalancers)
		volumeSpaceUsedGigabytes.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.VolumeSpaceGb)
		volumeVolumesUsedGigabytes.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.VolumeVolumes)

		for flavor := range v.ComputeFlavors {
			computeFlavorsUsed.With(prometheus.Labels{"region": k, "project": exporter.ProjectID, "flavor": flavor}).Set(v.ComputeFlavors[flavor])
		}

		for _, os := range v.Objectstorage {
			s3SpaceUsedBytes.With(prometheus.Labels{"region": k, "project": exporter.ProjectID, "type": os.Type}).Set(os.SpaceBytes)
		}
	}


}

// set utils for API v1 for Quota and Usage Information

func SetUtilsV1(quota map[string]api.QuotaV1, usage map[string]api.CurrentUsageV1, exporter *Exporter) {
	computeCoresTotal.Reset()
	computeInstancesTotal.Reset()
	computeRamTotalMegabytes.Reset()
	dnsZonesTotal.Reset()
	networkFloatingIPsTotal.Reset()
	networkLoadbalancersTotal.Reset()
	s3SpaceTotalBytes.Reset()
	volumeSpaceTotalGigabytes.Reset()
	volumeVolumesTotalGigabytes.Reset()
	computeCoresUsed.Reset()
	computeInstancesUsed.Reset()
	computeRamUsedMegabytes.Reset()
	dnsZonesUsed.Reset()
	networkFloatingIPsUsed.Reset()
	networkLoadbalancersUsed.Reset()
	s3SpaceUsedBytes.Reset()
	volumeSpaceUsedGigabytes.Reset()
	volumeVolumesUsedGigabytes.Reset()
	computeFlavorsUsed.Reset()

	for k, v := range quota {
		computeCoresTotal.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.ComputeCores)
		computeInstancesTotal.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.ComputeInstances)
		computeRamTotalMegabytes.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.ComputeRAMMb)
		dnsZonesTotal.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.DNSZones)
		networkFloatingIPsTotal.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.NetworkFloatingips)
		networkLoadbalancersTotal.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.NetworkLoadbalancers)
		s3SpaceTotalBytes.With(prometheus.Labels{"region": k, "project": exporter.ProjectID, "type": "quobyte"}).Set(v.S3SpaceBytes)
		volumeSpaceTotalGigabytes.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.VolumeSpaceGb)
		volumeVolumesTotalGigabytes.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.VolumeVolumes)
	}

	for k, v := range usage {
		computeCoresUsed.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.ComputeCores)
		computeInstancesUsed.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.ComputeInstances)
		computeRamUsedMegabytes.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.ComputeRAMMb)
		dnsZonesUsed.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.DNSZones)
		networkFloatingIPsUsed.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.NetworkFloatingips)
		networkLoadbalancersUsed.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.NetworkLoadbalancers)
		s3SpaceUsedBytes.With(prometheus.Labels{"region": k, "project": exporter.ProjectID, "type": "quobyte"}).Set(v.S3SpaceBytes)
		volumeSpaceUsedGigabytes.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.VolumeSpaceGb)
		volumeVolumesUsedGigabytes.With(prometheus.Labels{"region": k, "project": exporter.ProjectID}).Set(v.VolumeVolumes)

		for flavor := range v.ComputeFlavors {
			computeFlavorsUsed.With(prometheus.Labels{"region": k, "project": exporter.ProjectID, "flavor": flavor}).Set(v.ComputeFlavors[flavor])
		}
	}
}

func Run(interval int64, apiVersion string, exporter *Exporter) {
	for {
		var token string
		var err error

		log.Infof("Scrape Quota and Usage Metrics")
		if !(exporter.UseAppCreds) {
			token, err = auth.GetToken(exporter.ProjectID, exporter.Username, exporter.Password)
		} else {
			token, err = auth.GetTokenAppCreds(exporter.ProjectID, exporter.Username, exporter.Password)
		}
		if err != nil {
			log.WithError(err).Error("Could not get API Token")
			time.Sleep(60 * time.Second)
			continue
		}

		switch {
		case apiVersion == "v3":
			quota, err := api.GetQuotaV3(exporter.ProjectID, token)
			if err != nil {
				log.WithError(err).Error("Could not get quota")
				time.Sleep(60 * time.Second)
				continue
			}

			usage, err := api.GetCurrentUsageV3(exporter.ProjectID, token)
			if err != nil {
				log.WithError(err).Error("Could not get current usage")
				time.Sleep(60 * time.Second)
				continue
			}

			SetUtilsV3(quota, usage, exporter)

		default:
			quota, err := api.GetQuotaV1(exporter.ProjectID, token)
			if err != nil {
				log.WithError(err).Error("Could not get quota")
				time.Sleep(60 * time.Second)
				continue
			}

			usage, err := api.GetCurrentUsageV1(exporter.ProjectID, token)
			if err != nil {
				log.WithError(err).Error("Could not get current usage")
				time.Sleep(60 * time.Second)
				continue
			}

			SetUtilsV1(quota, usage, exporter)
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}
