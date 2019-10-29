// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2019 Datadog, Inc.

// +build docker

package ecs

import (
	"github.com/DataDog/datadog-agent/pkg/diagnose/diagnosis"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

func init() {
	diagnosis.Register("ECS Metadata availability", diagnoseECS)
	diagnosis.Register("ECS Metadata with tags availability", diagnoseECSTags)
	diagnosis.Register("ECS Fargate Metadata availability", diagnoseFargate)
	diagnosis.Register("ECS Fargate Metadata with tags availability", diagnoseFargateTags)
}

// diagnose the ECS metadata API availability
func diagnoseECS() error {
	client, err := MetaV1()
	if err != nil {
		log.Error(err)
	}
	_, err = client.GetTasks()
	return err
}

// diagnose the ECS metadata with tags API availability
func diagnoseECSTags() error {
	client, err := MetaV3InCurrentTask()
	if err != nil {
		log.Error(err)
	}
	_, err = client.GetTaskWithTags()
	return err
}

// diagnose the ECS Fargate metadata API availability
func diagnoseFargate() error {
	_, err := MetaV2().GetTask()
	if err != nil {
		log.Error(err)
	}
	return err
}

// diagnose the ECS Fargate metadata with tags API availability
func diagnoseFargateTags() error {
	_, err := MetaV2().GetTaskWithTags()
	if err != nil {
		log.Error(err)
	}
	return err
}
