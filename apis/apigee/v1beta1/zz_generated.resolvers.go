// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-gcp/config/common"
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *AddonsConfig) ResolveReferences( // ResolveReferences of this AddonsConfig.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apigee.gcp.upbound.io", "v1beta2", "Organization", "OrganizationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Org),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.ForProvider.OrgRef,
			Selector:     mg.Spec.ForProvider.OrgSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Org")
	}
	mg.Spec.ForProvider.Org = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrgRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigee.gcp.upbound.io", "v1beta2", "Organization", "OrganizationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Org),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.InitProvider.OrgRef,
			Selector:     mg.Spec.InitProvider.OrgSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Org")
	}
	mg.Spec.InitProvider.Org = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.OrgRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EndpointAttachment.
func (mg *EndpointAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apigee.gcp.upbound.io", "v1beta2", "Organization", "OrganizationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.OrgIDRef,
			Selector:     mg.Spec.ForProvider.OrgIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrgIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "ServiceAttachment", "ServiceAttachmentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAttachment),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ServiceAttachmentRef,
			Selector:     mg.Spec.ForProvider.ServiceAttachmentSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAttachment")
	}
	mg.Spec.ForProvider.ServiceAttachment = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAttachmentRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "ServiceAttachment", "ServiceAttachmentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceAttachment),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ServiceAttachmentRef,
			Selector:     mg.Spec.InitProvider.ServiceAttachmentSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceAttachment")
	}
	mg.Spec.InitProvider.ServiceAttachment = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAttachmentRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Envgroup.
func (mg *Envgroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apigee.gcp.upbound.io", "v1beta2", "Organization", "OrganizationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.OrgIDRef,
			Selector:     mg.Spec.ForProvider.OrgIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrgIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EnvgroupAttachment.
func (mg *EnvgroupAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apigee.gcp.upbound.io", "v1beta1", "Envgroup", "EnvgroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EnvgroupID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.EnvgroupIDRef,
			Selector:     mg.Spec.ForProvider.EnvgroupIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EnvgroupID")
	}
	mg.Spec.ForProvider.EnvgroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EnvgroupIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigee.gcp.upbound.io", "v1beta2", "Environment", "EnvironmentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Environment),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.EnvironmentRef,
			Selector:     mg.Spec.ForProvider.EnvironmentSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Environment")
	}
	mg.Spec.ForProvider.Environment = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EnvironmentRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigee.gcp.upbound.io", "v1beta1", "Envgroup", "EnvgroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EnvgroupID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.EnvgroupIDRef,
			Selector:     mg.Spec.InitProvider.EnvgroupIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EnvgroupID")
	}
	mg.Spec.InitProvider.EnvgroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EnvgroupIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigee.gcp.upbound.io", "v1beta2", "Environment", "EnvironmentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Environment),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.EnvironmentRef,
			Selector:     mg.Spec.InitProvider.EnvironmentSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Environment")
	}
	mg.Spec.InitProvider.Environment = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EnvironmentRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Environment.
func (mg *Environment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apigee.gcp.upbound.io", "v1beta1", "Organization", "OrganizationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.OrgIDRef,
			Selector:     mg.Spec.ForProvider.OrgIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrgIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Instance.
func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta2", "CryptoKey", "CryptoKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DiskEncryptionKeyName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DiskEncryptionKeyNameRef,
			Selector:     mg.Spec.ForProvider.DiskEncryptionKeyNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DiskEncryptionKeyName")
	}
	mg.Spec.ForProvider.DiskEncryptionKeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DiskEncryptionKeyNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigee.gcp.upbound.io", "v1beta2", "Organization", "OrganizationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.OrgIDRef,
			Selector:     mg.Spec.ForProvider.OrgIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrgIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta2", "CryptoKey", "CryptoKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DiskEncryptionKeyName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.DiskEncryptionKeyNameRef,
			Selector:     mg.Spec.InitProvider.DiskEncryptionKeyNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DiskEncryptionKeyName")
	}
	mg.Spec.InitProvider.DiskEncryptionKeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DiskEncryptionKeyNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this InstanceAttachment.
func (mg *InstanceAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apigee.gcp.upbound.io", "v1beta2", "Environment", "EnvironmentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Environment),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.EnvironmentRef,
			Selector:     mg.Spec.ForProvider.EnvironmentSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Environment")
	}
	mg.Spec.ForProvider.Environment = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EnvironmentRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigee.gcp.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.InstanceIDRef,
			Selector:     mg.Spec.ForProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigee.gcp.upbound.io", "v1beta2", "Environment", "EnvironmentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Environment),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.EnvironmentRef,
			Selector:     mg.Spec.InitProvider.EnvironmentSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Environment")
	}
	mg.Spec.InitProvider.Environment = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EnvironmentRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigee.gcp.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.InstanceIDRef,
			Selector:     mg.Spec.InitProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NATAddress.
func (mg *NATAddress) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apigee.gcp.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.InstanceIDRef,
			Selector:     mg.Spec.ForProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Organization.
func (mg *Organization) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AuthorizedNetwork),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.AuthorizedNetworkRef,
			Selector:     mg.Spec.ForProvider.AuthorizedNetworkSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AuthorizedNetwork")
	}
	mg.Spec.ForProvider.AuthorizedNetwork = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AuthorizedNetworkRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta1", "CryptoKey", "CryptoKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RuntimeDatabaseEncryptionKeyName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.RuntimeDatabaseEncryptionKeyNameRef,
			Selector:     mg.Spec.ForProvider.RuntimeDatabaseEncryptionKeyNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RuntimeDatabaseEncryptionKeyName")
	}
	mg.Spec.ForProvider.RuntimeDatabaseEncryptionKeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RuntimeDatabaseEncryptionKeyNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AuthorizedNetwork),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.AuthorizedNetworkRef,
			Selector:     mg.Spec.InitProvider.AuthorizedNetworkSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AuthorizedNetwork")
	}
	mg.Spec.InitProvider.AuthorizedNetwork = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AuthorizedNetworkRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta1", "CryptoKey", "CryptoKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RuntimeDatabaseEncryptionKeyName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.RuntimeDatabaseEncryptionKeyNameRef,
			Selector:     mg.Spec.InitProvider.RuntimeDatabaseEncryptionKeyNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RuntimeDatabaseEncryptionKeyName")
	}
	mg.Spec.InitProvider.RuntimeDatabaseEncryptionKeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RuntimeDatabaseEncryptionKeyNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SyncAuthorization.
func (mg *SyncAuthorization) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apigee.gcp.upbound.io", "v1beta2", "Organization", "OrganizationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Name),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.ForProvider.NameRef,
			Selector:     mg.Spec.ForProvider.NameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Name")
	}
	mg.Spec.ForProvider.Name = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NameRef = rsp.ResolvedReference

	return nil
}
