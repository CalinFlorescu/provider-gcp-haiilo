/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this TransferJob
func (mg *TransferJob) GetTerraformResourceType() string {
	return "google_storage_transfer_job"
}

// GetConnectionDetailsMapping for this TransferJob
func (tr *TransferJob) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"transfer_spec[*].aws_s3_data_source[*].aws_access_key[*].access_key_id": "spec.forProvider.transferSpec[*].awsS3DataSource[*].awsAccessKey[*].accessKeyIdSecretRef", "transfer_spec[*].aws_s3_data_source[*].aws_access_key[*].secret_access_key": "spec.forProvider.transferSpec[*].awsS3DataSource[*].awsAccessKey[*].secretAccessKeySecretRef", "transfer_spec[*].azure_blob_storage_data_source[*].azure_credentials[*].sas_token": "spec.forProvider.transferSpec[*].azureBlobStorageDataSource[*].azureCredentials[*].sasTokenSecretRef"}
}

// GetObservation of this TransferJob
func (tr *TransferJob) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this TransferJob
func (tr *TransferJob) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this TransferJob
func (tr *TransferJob) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this TransferJob
func (tr *TransferJob) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this TransferJob
func (tr *TransferJob) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this TransferJob using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *TransferJob) LateInitialize(attrs []byte) (bool, error) {
	params := &TransferJobParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *TransferJob) GetTerraformSchemaVersion() int {
	return 0
}