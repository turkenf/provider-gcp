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
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Application) ResolveReferences( // ResolveReferences of this Application.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "Project", "ProjectList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Project),
			Extract:      resource.ExtractParamPath("project_id", false),
			Reference:    mg.Spec.ForProvider.ProjectRef,
			Selector:     mg.Spec.ForProvider.ProjectSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Project")
	}
	mg.Spec.ForProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "Project", "ProjectList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Project),
			Extract:      resource.ExtractParamPath("project_id", false),
			Reference:    mg.Spec.InitProvider.ProjectRef,
			Selector:     mg.Spec.InitProvider.ProjectSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Project")
	}
	mg.Spec.InitProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProjectRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ApplicationURLDispatchRules.
func (mg *ApplicationURLDispatchRules) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DispatchRules); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("appengine.gcp.upbound.io", "v1beta1", "StandardAppVersion", "StandardAppVersionList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DispatchRules[i3].Service),
				Extract:      resource.ExtractParamPath("service", false),
				Reference:    mg.Spec.ForProvider.DispatchRules[i3].ServiceRef,
				Selector:     mg.Spec.ForProvider.DispatchRules[i3].ServiceSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DispatchRules[i3].Service")
		}
		mg.Spec.ForProvider.DispatchRules[i3].Service = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DispatchRules[i3].ServiceRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.DispatchRules); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("appengine.gcp.upbound.io", "v1beta1", "StandardAppVersion", "StandardAppVersionList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DispatchRules[i3].Service),
				Extract:      resource.ExtractParamPath("service", false),
				Reference:    mg.Spec.InitProvider.DispatchRules[i3].ServiceRef,
				Selector:     mg.Spec.InitProvider.DispatchRules[i3].ServiceSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DispatchRules[i3].Service")
		}
		mg.Spec.InitProvider.DispatchRules[i3].Service = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DispatchRules[i3].ServiceRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this FirewallRule.
func (mg *FirewallRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appengine.gcp.upbound.io", "v1beta1", "Application", "ApplicationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Project),
			Extract:      resource.ExtractParamPath("project", false),
			Reference:    mg.Spec.ForProvider.ProjectRef,
			Selector:     mg.Spec.ForProvider.ProjectSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Project")
	}
	mg.Spec.ForProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appengine.gcp.upbound.io", "v1beta1", "Application", "ApplicationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Project),
			Extract:      resource.ExtractParamPath("project", false),
			Reference:    mg.Spec.InitProvider.ProjectRef,
			Selector:     mg.Spec.InitProvider.ProjectSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Project")
	}
	mg.Spec.InitProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProjectRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServiceNetworkSettings.
func (mg *ServiceNetworkSettings) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appengine.gcp.upbound.io", "v1beta1", "StandardAppVersion", "StandardAppVersionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Service),
			Extract:      resource.ExtractParamPath("service", false),
			Reference:    mg.Spec.ForProvider.ServiceRef,
			Selector:     mg.Spec.ForProvider.ServiceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Service")
	}
	mg.Spec.ForProvider.Service = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appengine.gcp.upbound.io", "v1beta1", "StandardAppVersion", "StandardAppVersionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Service),
			Extract:      resource.ExtractParamPath("service", false),
			Reference:    mg.Spec.InitProvider.ServiceRef,
			Selector:     mg.Spec.InitProvider.ServiceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Service")
	}
	mg.Spec.InitProvider.Service = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StandardAppVersion.
func (mg *StandardAppVersion) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccount),
			Extract:      resource.ExtractParamPath("email", true),
			Reference:    mg.Spec.ForProvider.ServiceAccountRef,
			Selector:     mg.Spec.ForProvider.ServiceAccountSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccount")
	}
	mg.Spec.ForProvider.ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceAccount),
			Extract:      resource.ExtractParamPath("email", true),
			Reference:    mg.Spec.InitProvider.ServiceAccountRef,
			Selector:     mg.Spec.InitProvider.ServiceAccountSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceAccount")
	}
	mg.Spec.InitProvider.ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountRef = rsp.ResolvedReference

	return nil
}
