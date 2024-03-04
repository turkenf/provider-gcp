/*
Copyright 2021 The Crossplane Authors.

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
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	common "github.com/upbound/provider-gcp/config/common"
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *CAPoolIAMMember) ResolveReferences( // ResolveReferences of this CAPoolIAMMember.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("privateca.gcp.upbound.io", "v1beta1", "CAPool", "CAPoolList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CAPool),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.CAPoolRef,
			Selector:     mg.Spec.ForProvider.CAPoolSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CAPool")
	}
	mg.Spec.ForProvider.CAPool = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CAPoolRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("privateca.gcp.upbound.io", "v1beta1", "CAPool", "CAPoolList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CAPool),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.CAPoolRef,
			Selector:     mg.Spec.InitProvider.CAPoolSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CAPool")
	}
	mg.Spec.InitProvider.CAPool = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CAPoolRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Certificate.
func (mg *Certificate) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("privateca.gcp.upbound.io", "v1beta1", "CertificateAuthority", "CertificateAuthorityList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CertificateAuthority),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CertificateAuthorityRef,
			Selector:     mg.Spec.ForProvider.CertificateAuthoritySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CertificateAuthority")
	}
	mg.Spec.ForProvider.CertificateAuthority = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CertificateAuthorityRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("privateca.gcp.upbound.io", "v1beta1", "CertificateTemplate", "CertificateTemplateList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CertificateTemplate),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.CertificateTemplateRef,
			Selector:     mg.Spec.ForProvider.CertificateTemplateSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CertificateTemplate")
	}
	mg.Spec.ForProvider.CertificateTemplate = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CertificateTemplateRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("privateca.gcp.upbound.io", "v1beta1", "CAPool", "CAPoolList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Pool),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.PoolRef,
			Selector:     mg.Spec.ForProvider.PoolSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Pool")
	}
	mg.Spec.ForProvider.Pool = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PoolRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("privateca.gcp.upbound.io", "v1beta1", "CertificateAuthority", "CertificateAuthorityList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CertificateAuthority),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.CertificateAuthorityRef,
			Selector:     mg.Spec.InitProvider.CertificateAuthoritySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CertificateAuthority")
	}
	mg.Spec.InitProvider.CertificateAuthority = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CertificateAuthorityRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("privateca.gcp.upbound.io", "v1beta1", "CertificateTemplate", "CertificateTemplateList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CertificateTemplate),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.CertificateTemplateRef,
			Selector:     mg.Spec.InitProvider.CertificateTemplateSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CertificateTemplate")
	}
	mg.Spec.InitProvider.CertificateTemplate = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CertificateTemplateRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this CertificateAuthority.
func (mg *CertificateAuthority) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("privateca.gcp.upbound.io", "v1beta1", "CAPool", "CAPoolList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Pool),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.PoolRef,
			Selector:     mg.Spec.ForProvider.PoolSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Pool")
	}
	mg.Spec.ForProvider.Pool = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PoolRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SubordinateConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("privateca.gcp.upbound.io", "v1beta1", "CertificateAuthority", "CertificateAuthorityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubordinateConfig[i3].CertificateAuthority),
				Extract:      resource.ExtractParamPath("name", true),
				Reference:    mg.Spec.ForProvider.SubordinateConfig[i3].CertificateAuthorityRef,
				Selector:     mg.Spec.ForProvider.SubordinateConfig[i3].CertificateAuthoritySelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SubordinateConfig[i3].CertificateAuthority")
		}
		mg.Spec.ForProvider.SubordinateConfig[i3].CertificateAuthority = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SubordinateConfig[i3].CertificateAuthorityRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.SubordinateConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("privateca.gcp.upbound.io", "v1beta1", "CertificateAuthority", "CertificateAuthorityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubordinateConfig[i3].CertificateAuthority),
				Extract:      resource.ExtractParamPath("name", true),
				Reference:    mg.Spec.InitProvider.SubordinateConfig[i3].CertificateAuthorityRef,
				Selector:     mg.Spec.InitProvider.SubordinateConfig[i3].CertificateAuthoritySelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.SubordinateConfig[i3].CertificateAuthority")
		}
		mg.Spec.InitProvider.SubordinateConfig[i3].CertificateAuthority = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.SubordinateConfig[i3].CertificateAuthorityRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this CertificateTemplateIAMMember.
func (mg *CertificateTemplateIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("privateca.gcp.upbound.io", "v1beta1", "CertificateTemplate", "CertificateTemplateList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CertificateTemplate),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.CertificateTemplateRef,
			Selector:     mg.Spec.ForProvider.CertificateTemplateSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CertificateTemplate")
	}
	mg.Spec.ForProvider.CertificateTemplate = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CertificateTemplateRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("privateca.gcp.upbound.io", "v1beta1", "CertificateTemplate", "CertificateTemplateList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CertificateTemplate),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.CertificateTemplateRef,
			Selector:     mg.Spec.InitProvider.CertificateTemplateSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CertificateTemplate")
	}
	mg.Spec.InitProvider.CertificateTemplate = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CertificateTemplateRef = rsp.ResolvedReference

	return nil
}
