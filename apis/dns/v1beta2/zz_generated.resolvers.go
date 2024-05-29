// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

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

func (mg *ManagedZone) ResolveReferences( // ResolveReferences of this ManagedZone.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.PeeringConfig != nil {
		if mg.Spec.ForProvider.PeeringConfig.TargetNetwork != nil {
			{
				m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PeeringConfig.TargetNetwork.NetworkURL),
					Extract:      common.SelfLinkExtractor(),
					Reference:    mg.Spec.ForProvider.PeeringConfig.TargetNetwork.NetworkURLRef,
					Selector:     mg.Spec.ForProvider.PeeringConfig.TargetNetwork.NetworkURLSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.PeeringConfig.TargetNetwork.NetworkURL")
			}
			mg.Spec.ForProvider.PeeringConfig.TargetNetwork.NetworkURL = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.PeeringConfig.TargetNetwork.NetworkURLRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.PrivateVisibilityConfig != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.PrivateVisibilityConfig.GkeClusters); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("container.gcp.upbound.io", "v1beta2", "Cluster", "ClusterList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateVisibilityConfig.GkeClusters[i4].GkeClusterName),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.PrivateVisibilityConfig.GkeClusters[i4].GkeClusterNameRef,
					Selector:     mg.Spec.ForProvider.PrivateVisibilityConfig.GkeClusters[i4].GkeClusterNameSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.PrivateVisibilityConfig.GkeClusters[i4].GkeClusterName")
			}
			mg.Spec.ForProvider.PrivateVisibilityConfig.GkeClusters[i4].GkeClusterName = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.PrivateVisibilityConfig.GkeClusters[i4].GkeClusterNameRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.PrivateVisibilityConfig != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.PrivateVisibilityConfig.Networks); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateVisibilityConfig.Networks[i4].NetworkURL),
					Extract:      common.SelfLinkExtractor(),
					Reference:    mg.Spec.ForProvider.PrivateVisibilityConfig.Networks[i4].NetworkURLRef,
					Selector:     mg.Spec.ForProvider.PrivateVisibilityConfig.Networks[i4].NetworkURLSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.PrivateVisibilityConfig.Networks[i4].NetworkURL")
			}
			mg.Spec.ForProvider.PrivateVisibilityConfig.Networks[i4].NetworkURL = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.PrivateVisibilityConfig.Networks[i4].NetworkURLRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.PeeringConfig != nil {
		if mg.Spec.InitProvider.PeeringConfig.TargetNetwork != nil {
			{
				m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PeeringConfig.TargetNetwork.NetworkURL),
					Extract:      common.SelfLinkExtractor(),
					Reference:    mg.Spec.InitProvider.PeeringConfig.TargetNetwork.NetworkURLRef,
					Selector:     mg.Spec.InitProvider.PeeringConfig.TargetNetwork.NetworkURLSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.PeeringConfig.TargetNetwork.NetworkURL")
			}
			mg.Spec.InitProvider.PeeringConfig.TargetNetwork.NetworkURL = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.PeeringConfig.TargetNetwork.NetworkURLRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.PrivateVisibilityConfig != nil {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.PrivateVisibilityConfig.GkeClusters); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("container.gcp.upbound.io", "v1beta2", "Cluster", "ClusterList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PrivateVisibilityConfig.GkeClusters[i4].GkeClusterName),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.PrivateVisibilityConfig.GkeClusters[i4].GkeClusterNameRef,
					Selector:     mg.Spec.InitProvider.PrivateVisibilityConfig.GkeClusters[i4].GkeClusterNameSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.PrivateVisibilityConfig.GkeClusters[i4].GkeClusterName")
			}
			mg.Spec.InitProvider.PrivateVisibilityConfig.GkeClusters[i4].GkeClusterName = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.PrivateVisibilityConfig.GkeClusters[i4].GkeClusterNameRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.PrivateVisibilityConfig != nil {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.PrivateVisibilityConfig.Networks); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PrivateVisibilityConfig.Networks[i4].NetworkURL),
					Extract:      common.SelfLinkExtractor(),
					Reference:    mg.Spec.InitProvider.PrivateVisibilityConfig.Networks[i4].NetworkURLRef,
					Selector:     mg.Spec.InitProvider.PrivateVisibilityConfig.Networks[i4].NetworkURLSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.PrivateVisibilityConfig.Networks[i4].NetworkURL")
			}
			mg.Spec.InitProvider.PrivateVisibilityConfig.Networks[i4].NetworkURL = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.PrivateVisibilityConfig.Networks[i4].NetworkURLRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this Policy.
func (mg *Policy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Networks); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Networks[i3].NetworkURL),
				Extract:      common.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Networks[i3].NetworkRef,
				Selector:     mg.Spec.ForProvider.Networks[i3].NetworkSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Networks[i3].NetworkURL")
		}
		mg.Spec.ForProvider.Networks[i3].NetworkURL = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Networks[i3].NetworkRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Networks); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Networks[i3].NetworkURL),
				Extract:      common.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Networks[i3].NetworkRef,
				Selector:     mg.Spec.InitProvider.Networks[i3].NetworkSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Networks[i3].NetworkURL")
		}
		mg.Spec.InitProvider.Networks[i3].NetworkURL = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Networks[i3].NetworkRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this RecordSet.
func (mg *RecordSet) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("dns.gcp.upbound.io", "v1beta2", "ManagedZone", "ManagedZoneList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ManagedZone),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ManagedZoneRef,
			Selector:     mg.Spec.ForProvider.ManagedZoneSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ManagedZone")
	}
	mg.Spec.ForProvider.ManagedZone = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ManagedZoneRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.RoutingPolicy != nil {
		if mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup != nil {
			if mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary != nil {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta2", "ForwardingRule", "ForwardingRuleList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].IPAddress),
							Extract:      resource.ExtractParamPath("ip_address", false),
							Reference:    mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].IPAddressRef,
							Selector:     mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].IPAddressSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].IPAddress")
					}
					mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].IPAddress = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].IPAddressRef = rsp.ResolvedReference

				}
			}
		}
	}
	if mg.Spec.ForProvider.RoutingPolicy != nil {
		if mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup != nil {
			if mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary != nil {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].NetworkURL),
							Extract:      resource.ExtractResourceID(),
							Reference:    mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].NetworkURLRef,
							Selector:     mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].NetworkURLSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].NetworkURL")
					}
					mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].NetworkURL = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].NetworkURLRef = rsp.ResolvedReference

				}
			}
		}
	}
	if mg.Spec.ForProvider.RoutingPolicy != nil {
		if mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup != nil {
			if mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary != nil {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta2", "ForwardingRule", "ForwardingRuleList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].Project),
							Extract:      resource.ExtractParamPath("project", false),
							Reference:    mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].ProjectRef,
							Selector:     mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].ProjectSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].Project")
					}
					mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].Project = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].ProjectRef = rsp.ResolvedReference

				}
			}
		}
	}
	if mg.Spec.ForProvider.RoutingPolicy != nil {
		if mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup != nil {
			if mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary != nil {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta2", "ForwardingRule", "ForwardingRuleList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].Region),
							Extract:      resource.ExtractParamPath("region", false),
							Reference:    mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].RegionRef,
							Selector:     mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].RegionSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].Region")
					}
					mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].Region = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].RegionRef = rsp.ResolvedReference

				}
			}
		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("dns.gcp.upbound.io", "v1beta2", "ManagedZone", "ManagedZoneList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ManagedZone),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ManagedZoneRef,
			Selector:     mg.Spec.InitProvider.ManagedZoneSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ManagedZone")
	}
	mg.Spec.InitProvider.ManagedZone = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ManagedZoneRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.RoutingPolicy != nil {
		if mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup != nil {
			if mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary != nil {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta2", "ForwardingRule", "ForwardingRuleList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].IPAddress),
							Extract:      resource.ExtractParamPath("ip_address", false),
							Reference:    mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].IPAddressRef,
							Selector:     mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].IPAddressSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].IPAddress")
					}
					mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].IPAddress = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].IPAddressRef = rsp.ResolvedReference

				}
			}
		}
	}
	if mg.Spec.InitProvider.RoutingPolicy != nil {
		if mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup != nil {
			if mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary != nil {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].NetworkURL),
							Extract:      resource.ExtractResourceID(),
							Reference:    mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].NetworkURLRef,
							Selector:     mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].NetworkURLSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].NetworkURL")
					}
					mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].NetworkURL = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].NetworkURLRef = rsp.ResolvedReference

				}
			}
		}
	}
	if mg.Spec.InitProvider.RoutingPolicy != nil {
		if mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup != nil {
			if mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary != nil {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta2", "ForwardingRule", "ForwardingRuleList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].Project),
							Extract:      resource.ExtractParamPath("project", false),
							Reference:    mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].ProjectRef,
							Selector:     mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].ProjectSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].Project")
					}
					mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].Project = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].ProjectRef = rsp.ResolvedReference

				}
			}
		}
	}
	if mg.Spec.InitProvider.RoutingPolicy != nil {
		if mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup != nil {
			if mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary != nil {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta2", "ForwardingRule", "ForwardingRuleList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].Region),
							Extract:      resource.ExtractParamPath("region", false),
							Reference:    mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].RegionRef,
							Selector:     mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].RegionSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].Region")
					}
					mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].Region = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.RoutingPolicy.PrimaryBackup.Primary.InternalLoadBalancers[i6].RegionRef = rsp.ResolvedReference

				}
			}
		}
	}

	return nil
}