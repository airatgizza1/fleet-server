// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by dev-tools/cmd/buildlimits/buildlimits.go - DO NOT EDIT.

package config

import (
	"math"
	"runtime"
	"strings"
	"time"

	"github.com/elastic/elastic-agent/pkg/packer"
	"github.com/elastic/go-ucfg/yaml"
	"github.com/pbnjay/memory"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const (
	defaultCacheNumCounters = 500000           // 10x times expected count
	defaultCacheMaxCost     = 50 * 1024 * 1024 // 50MiB cache size

	defaultMaxConnections = 0 // no limit
	defaultPolicyThrottle = time.Millisecond * 5

	defaultCheckinInterval = time.Millisecond
	defaultCheckinBurst    = 1000
	defaultCheckinMax      = 0
	defaultCheckinMaxBody  = 1024 * 1024

	defaultArtifactInterval = time.Millisecond * 5
	defaultArtifactBurst    = 25
	defaultArtifactMax      = 50
	defaultArtifactMaxBody  = 0

	defaultEnrollInterval = time.Millisecond * 10
	defaultEnrollBurst    = 100
	defaultEnrollMax      = 50
	defaultEnrollMaxBody  = 1024 * 512

	defaultAckInterval = time.Millisecond * 10
	defaultAckBurst    = 100
	defaultAckMax      = 50
	defaultAckMaxBody  = 1024 * 1024 * 2

	defaultStatusInterval = time.Millisecond * 5
	defaultStatusBurst    = 25
	defaultStatusMax      = 50
	defaultStatusMaxBody  = 0
)

type valueRange struct {
	Min int `config:"min"`
	Max int `config:"max"`
}

type envLimits struct {
	Agents         valueRange           `config:"num_agents"`
	RecommendedRAM int                  `config:"recommended_min_ram"`
	Server         *serverLimitDefaults `config:"server_limits"`
	Cache          *cacheLimits         `config:"cache_limits"`
}

func defaultEnvLimits() *envLimits {
	return &envLimits{
		Agents: valueRange{
			Min: 0,
			Max: int(getMaxInt()),
		},
		Server: defaultserverLimitDefaults(),
		Cache:  defaultCacheLimits(),
	}
}

type cacheLimits struct {
	NumCounters int64 `config:"num_counters"`
	MaxCost     int64 `config:"max_cost"`
}

func defaultCacheLimits() *cacheLimits {
	return &cacheLimits{
		NumCounters: defaultCacheNumCounters,
		MaxCost:     defaultCacheMaxCost,
	}
}

type limit struct {
	Interval time.Duration `config:"interval"`
	Burst    int           `config:"burst"`
	Max      int64         `config:"max"`
	MaxBody  int64         `config:"max_body_byte_size"`
}

type serverLimitDefaults struct {
	PolicyThrottle time.Duration `config:"policy_throttle"`
	MaxConnections int           `config:"max_connections"`

	CheckinLimit  limit `config:"checkin_limit"`
	ArtifactLimit limit `config:"artifact_limit"`
	EnrollLimit   limit `config:"enroll_limit"`
	AckLimit      limit `config:"ack_limit"`
	StatusLimit   limit `config:"status_limit"`
}

func defaultserverLimitDefaults() *serverLimitDefaults {
	return &serverLimitDefaults{
		PolicyThrottle: defaultCacheNumCounters,
		MaxConnections: defaultCacheMaxCost,

		CheckinLimit: limit{
			Interval: defaultCheckinInterval,
			Burst:    defaultCheckinBurst,
			Max:      defaultCheckinMax,
			MaxBody:  defaultCheckinMaxBody,
		},
		ArtifactLimit: limit{
			Interval: defaultArtifactInterval,
			Burst:    defaultArtifactBurst,
			Max:      defaultArtifactMax,
			MaxBody:  defaultArtifactMaxBody,
		},
		EnrollLimit: limit{
			Interval: defaultEnrollInterval,
			Burst:    defaultEnrollBurst,
			Max:      defaultEnrollMax,
			MaxBody:  defaultEnrollMaxBody,
		},
		AckLimit: limit{
			Interval: defaultAckInterval,
			Burst:    defaultAckBurst,
			Max:      defaultAckMax,
			MaxBody:  defaultAckMaxBody,
		},
		StatusLimit: limit{
			Interval: defaultStatusInterval,
			Burst:    defaultStatusBurst,
			Max:      defaultStatusMax,
			MaxBody:  defaultStatusMaxBody,
		},
	}
}

var defaults []*envLimits

func init() {
	// Packed Files
	// internal/pkg/config/defaults/gt10000_limits.yml
	// internal/pkg/config/defaults/gt12500_limits.yml
	// internal/pkg/config/defaults/gt5000_limits.yml
	// internal/pkg/config/defaults/gt50_limits.yml
	// internal/pkg/config/defaults/gt7500_limits.yml
	// internal/pkg/config/defaults/lte50_limits.yml
	// internal/pkg/config/defaults/max_limits.yml
	unpacked := packer.MustUnpack("eJzsl12Psjgbx8+fjzHHT3YAxY2bzEEVYXDSGglS4GRDYUSwvOQW5GWz331TfImjOED2TjbZ7OEk48XV//Xrr+0fL0Gcff6IHfqa7v1XN4m3gf/qfW6dnGaHVz/jOY7jfqdBFGSHX6qIvvz2QmIjMwU5trHIuXXyoc5nmYPHE1VaFFACxTwAPsHyeBWAEoZrcbVOKxujI8E8JZFGbXmaOVjcupGcrQIwhvq+mkdoZwk7agplSjC91M09zG8tYcqTWKNuDNk3BPYNCBJfVfjdpzwNyQhxqwBUUNqPoL6poAQ+XMGovMio7uoVRChTa0S3nrKryEjjiLKZqJKRucVNPUHMbYw4B09zt058WKtsXez3oaMYoSPQ/FJ35Se+Oge+g0XONpeChYuJKm0K+G7ULAd1Pgs8c1l7UuKjS9/z23w2BVyzb8uVp9DIwoi71g7YN4GvKjT3FLYe+bAKAA+lGd/0Owe+FRuVO1pPVGl97hGc1iEl/uqak5G70fRA5OnBwXzKejn3lp5ytXdESnyoz7LHuov7nqtzFjtLyLZEoZljridt9ZAOCs+E556WvBs3c+IujDBuLNOfqJLf9P7y/04YBXEAjBVr4AbGaqXvu2AsUQjHCKShhVFiG+fA3mGzwNO30NE7Dbs+/T+4wJFZ5o7B2gwbhkBE9YKHISjmMWKhUJd+recq7G8kmiMtcaMp5yklZX2SJrRrvZxERugp9EiaDQCrS4Bsw1hClhJ6HezfAgbqmwJKi2a4bqyldiSHXgc0Q4eszmeUxMsjUcorPCpo20SLgjzWLR96Ds9ZYPTjJ23KcdO7//bWDaQ4QI6bG05YHi630vcfbmSEROAzG4t7G2tbgmlujpY7ojdr41b9WeSeiJG/ihGkNcvDNpfn7LW6yep9diQKDT+NKee8L4/eu3aw9RshPBOjDgaKcdkumLuZonqwFKsnrNwxqP7DDPZbP6sHQdKLv/703ZiQkSfO4yW1hGlGsJHbirE3BZ6ZpLIwd+oqXH9Ygswm+5UWReSJwegymgk0JHRbcOCRvCm+t6A7kDyjJfnN/c7noQTKZpLmknOwvbNGWgchvepeSbSxWBGhPHTRDBX+0VT6va37Uqf177GX9X4dcAq7/Ffr7cUu6yEJiMjvbb3xE+uVUFqPke4KP/06GA61ntySv9py+oLBpy/Ej5yc7XE7W24gf2Ubfxfevti0yV0OHf7bG0jZYtHq0c7rXtdAmn320d65ZIfmwDDFffPqGHrJg1L34Troktc2snDTrozeelsUT1C4rzvoUG3V2/pORz3Ral/3w4HaiVXklH2dBk/P3G/BghWqrXH/Z+3pVv8vedY2L6//nrXNi+ft5c///RUAAP//y2rXXQ==")

	for f, v := range unpacked {
		cfg, err := yaml.NewConfig(v, DefaultOptions...)
		if err != nil {
			panic(errors.Wrap(err, "Cannot read spec from "+f))
		}

		l := defaultEnvLimits()
		if err := cfg.Unpack(&l, DefaultOptions...); err != nil {
			panic(errors.Wrap(err, "Cannot unpack spec from "+f))
		}

		defaults = append(defaults, l)
	}
}

func loadLimits(agentLimit int) *envLimits {
	return loadLimitsForAgents(agentLimit)
}

func loadLimitsForAgents(agentLimit int) *envLimits {
	for _, l := range defaults {
		// get nearest limits for configured agent numbers
		if l.Agents.Min < agentLimit && agentLimit <= l.Agents.Max {
			log.Info().Msgf("Using system limits for %d to %d agents for a configured value of %d agents", l.Agents.Min, l.Agents.Max, agentLimit)
			ramSize := int(memory.TotalMemory() / 1024 / 1024)
			if ramSize < l.RecommendedRAM {
				log.Warn().Msgf("Detected %d MB of system RAM, which is lower than the recommended amount (%d MB) for the configured agent limit", ramSize, l.RecommendedRAM)
			}
			return l
		}
	}
	log.Info().Msgf("No applicable limit for %d agents, using default.", agentLimit)
	return defaultEnvLimits()
}

func getMaxInt() int64 {
	if strings.HasSuffix(runtime.GOARCH, "64") {
		return math.MaxInt64
	}
	return math.MaxInt32
}
