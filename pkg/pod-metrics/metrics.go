/*
 * This file is part of the kubevirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2017-2018 Red Hat, Inc.
 *
 */

package metrics

import (
	"path/filepath"

	"kubevirt.io/kubevirt/pkg/api/v1"
	diskutils "kubevirt.io/kubevirt/pkg/ephemeral-disk-utils"
	"kubevirt.io/kubevirt/pkg/log"
	"kubevirt.io/kubevirt/pkg/precond"
)

func MetricsFileDirectory(baseDir string) string {
	return filepath.Join(baseDir, "metrics-files")
}

func MetricsFileFromNamespaceName(baseDir string, namespace string, name string) string {
	metricsFile := namespace + "_" + name
	return filepath.Join(baseDir, "metrics-files", metricsFile)
}

func MetricsFileRemove(baseDir string, vmi *v1.VirtualMachineInstance) error {
	namespace := precond.MustNotBeEmpty(vmi.GetObjectMeta().GetNamespace())
	domain := precond.MustNotBeEmpty(vmi.GetObjectMeta().GetName())

	file := MetricsFileFromNamespaceName(baseDir, namespace, domain)

	return diskutils.RemoveFile(file)
}

func MetricsFileExists(baseDir string, vmi *v1.VirtualMachineInstance) (bool, error) {
	namespace := precond.MustNotBeEmpty(vmi.GetObjectMeta().GetNamespace())
	domain := precond.MustNotBeEmpty(vmi.GetObjectMeta().GetName())

	filePath := MetricsFileFromNamespaceName(baseDir, namespace, domain)
	exists, err := diskutils.FileExists(filePath)
	if err != nil {
		log.Log.Reason(err).Errorf("Error encountered while attempting to verify if metrics file at path %s exists.", filePath)

		return false, err
	}
	return exists, nil
}
