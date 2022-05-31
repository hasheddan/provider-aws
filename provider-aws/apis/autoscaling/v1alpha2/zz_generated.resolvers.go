/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha22 "github.com/upbound/official-providers/provider-aws/apis/ec2/v1alpha2"
	v1alpha2 "github.com/upbound/official-providers/provider-aws/apis/elbv2/v1alpha2"
	v1alpha21 "github.com/upbound/official-providers/provider-aws/apis/iam/v1alpha2"
	common "github.com/upbound/official-providers/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Attachment.
func (mg *Attachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ALBTargetGroupArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.ALBTargetGroupArnRef,
		Selector:     mg.Spec.ForProvider.ALBTargetGroupArnSelector,
		To: reference.To{
			List:    &v1alpha2.LBTargetGroupList{},
			Managed: &v1alpha2.LBTargetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ALBTargetGroupArn")
	}
	mg.Spec.ForProvider.ALBTargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ALBTargetGroupArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AutoscalingGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AutoscalingGroupNameRef,
		Selector:     mg.Spec.ForProvider.AutoscalingGroupNameSelector,
		To: reference.To{
			List:    &AutoscalingGroupList{},
			Managed: &AutoscalingGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AutoscalingGroupName")
	}
	mg.Spec.ForProvider.AutoscalingGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AutoscalingGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this AutoscalingGroup.
func (mg *AutoscalingGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceLinkedRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.ServiceLinkedRoleArnRef,
		Selector:     mg.Spec.ForProvider.ServiceLinkedRoleArnSelector,
		To: reference.To{
			List:    &v1alpha21.RoleList{},
			Managed: &v1alpha21.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceLinkedRoleArn")
	}
	mg.Spec.ForProvider.ServiceLinkedRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceLinkedRoleArnRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCZoneIdentifier),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VPCZoneIdentifierRefs,
		Selector:      mg.Spec.ForProvider.VPCZoneIdentifierSelector,
		To: reference.To{
			List:    &v1alpha22.SubnetList{},
			Managed: &v1alpha22.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCZoneIdentifier")
	}
	mg.Spec.ForProvider.VPCZoneIdentifier = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VPCZoneIdentifierRefs = mrsp.ResolvedReferences

	return nil
}
