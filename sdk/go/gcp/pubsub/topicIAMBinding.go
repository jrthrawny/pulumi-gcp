// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Three different resources help you manage your IAM policy for pubsub topic. Each of these resources serves a different use case:
// 
// * `google_pubsub_topic_iam_policy`: Authoritative. Sets the IAM policy for the topic and replaces any existing policy already attached.
// * `google_pubsub_topic_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the topic are preserved.
// * `google_pubsub_topic_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the topic are preserved.
// 
// > **Note:** `google_pubsub_topic_iam_policy` **cannot** be used in conjunction with `google_pubsub_topic_iam_binding` and `google_pubsub_topic_iam_member` or they will fight over what your policy should be.
// 
// > **Note:** `google_pubsub_topic_iam_binding` resources **can be** used in conjunction with `google_pubsub_topic_iam_member` resources **only if** they do not grant privilege to the same role.
type TopicIAMBinding struct {
	s *pulumi.ResourceState
}

// NewTopicIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewTopicIAMBinding(ctx *pulumi.Context,
	name string, args *TopicIAMBindingArgs, opts ...pulumi.ResourceOpt) (*TopicIAMBinding, error) {
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.Topic == nil {
		return nil, errors.New("missing required argument 'Topic'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["members"] = nil
		inputs["project"] = nil
		inputs["role"] = nil
		inputs["topic"] = nil
	} else {
		inputs["members"] = args.Members
		inputs["project"] = args.Project
		inputs["role"] = args.Role
		inputs["topic"] = args.Topic
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:pubsub/topicIAMBinding:TopicIAMBinding", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TopicIAMBinding{s: s}, nil
}

// GetTopicIAMBinding gets an existing TopicIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TopicIAMBindingState, opts ...pulumi.ResourceOpt) (*TopicIAMBinding, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["members"] = state.Members
		inputs["project"] = state.Project
		inputs["role"] = state.Role
		inputs["topic"] = state.Topic
	}
	s, err := ctx.ReadResource("gcp:pubsub/topicIAMBinding:TopicIAMBinding", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TopicIAMBinding{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *TopicIAMBinding) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *TopicIAMBinding) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// (Computed) The etag of the topic's IAM policy.
func (r *TopicIAMBinding) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

func (r *TopicIAMBinding) Members() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["members"])
}

// The project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *TopicIAMBinding) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The role that should be applied. Only one
// `google_pubsub_topic_iam_binding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *TopicIAMBinding) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// The topic name or id to bind to attach IAM policy to.
func (r *TopicIAMBinding) Topic() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["topic"])
}

// Input properties used for looking up and filtering TopicIAMBinding resources.
type TopicIAMBindingState struct {
	// (Computed) The etag of the topic's IAM policy.
	Etag interface{}
	Members interface{}
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `google_pubsub_topic_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// The topic name or id to bind to attach IAM policy to.
	Topic interface{}
}

// The set of arguments for constructing a TopicIAMBinding resource.
type TopicIAMBindingArgs struct {
	Members interface{}
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `google_pubsub_topic_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// The topic name or id to bind to attach IAM policy to.
	Topic interface{}
}
