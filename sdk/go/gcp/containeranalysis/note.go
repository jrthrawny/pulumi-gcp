// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containeranalysis

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a detailed description of a Note.
// 
// > **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
// 
// To get more information about Note, see:
// 
// * [API documentation](https://cloud.google.com/container-analysis/api/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/container-analysis/)
// 
// <div class = "oics-button" style="float: right; margin: 0 0 -15px">
//   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=container_analysis_note_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
//     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
//   </a>
// </div>
type Note struct {
	s *pulumi.ResourceState
}

// NewNote registers a new resource with the given unique name, arguments, and options.
func NewNote(ctx *pulumi.Context,
	name string, args *NoteArgs, opts ...pulumi.ResourceOpt) (*Note, error) {
	if args == nil || args.AttestationAuthority == nil {
		return nil, errors.New("missing required argument 'AttestationAuthority'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["attestationAuthority"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
	} else {
		inputs["attestationAuthority"] = args.AttestationAuthority
		inputs["name"] = args.Name
		inputs["project"] = args.Project
	}
	s, err := ctx.RegisterResource("gcp:containeranalysis/note:Note", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Note{s: s}, nil
}

// GetNote gets an existing Note resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNote(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NoteState, opts ...pulumi.ResourceOpt) (*Note, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["attestationAuthority"] = state.AttestationAuthority
		inputs["name"] = state.Name
		inputs["project"] = state.Project
	}
	s, err := ctx.ReadResource("gcp:containeranalysis/note:Note", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Note{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Note) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Note) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Note) AttestationAuthority() *pulumi.Output {
	return r.s.State["attestationAuthority"]
}

func (r *Note) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *Note) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// Input properties used for looking up and filtering Note resources.
type NoteState struct {
	AttestationAuthority interface{}
	Name interface{}
	Project interface{}
}

// The set of arguments for constructing a Note resource.
type NoteArgs struct {
	AttestationAuthority interface{}
	Name interface{}
	Project interface{}
}
