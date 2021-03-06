// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > **Warning:** These resources are in beta, and should be used with the terraform-provider-google-beta provider.
 * See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
 * 
 * Three different resources help you manage your IAM policy for GCE subnetwork. Each of these resources serves a different use case:
 * 
 * * `google_compute_subnetwork_iam_policy`: Authoritative. Sets the IAM policy for the subnetwork and replaces any existing policy already attached.
 * * `google_compute_subnetwork_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subnetwork are preserved.
 * * `google_compute_subnetwork_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subnetwork are preserved.
 * 
 * > **Note:** `google_compute_subnetwork_iam_policy` **cannot** be used in conjunction with `google_compute_subnetwork_iam_binding` and `google_compute_subnetwork_iam_member` or they will fight over what your policy should be.
 * 
 * > **Note:** `google_compute_subnetwork_iam_binding` resources **can be** used in conjunction with `google_compute_subnetwork_iam_member` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_compute\_subnetwork\_iam\_policy
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const admin = pulumi.output(gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         members: ["user:jane@example.com"],
 *         role: "roles/editor",
 *     }],
 * }));
 * const subnet = new gcp.compute.SubnetworkIAMPolicy("subnet", {
 *     policyData: admin.apply(admin => admin.policyData),
 *     subnetwork: "your-subnetwork-id",
 * });
 * ```
 * 
 * ## google\_compute\_subnetwork\_iam\_binding
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const subnet = new gcp.compute.SubnetworkIAMBinding("subnet", {
 *     members: ["user:jane@example.com"],
 *     role: "roles/compute.networkUser",
 *     subnetwork: "your-subnetwork-id",
 * });
 * ```
 * 
 * ## google\compute\_subnetwork\_iam\_member
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const subnet = new gcp.compute.SubnetworkIAMMember("subnet", {
 *     member: "user:jane@example.com",
 *     role: "roles/compute.networkUser",
 *     subnetwork: "your-subnetwork-id",
 * });
 * ```
 */
export class SubnetworkIAMPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SubnetworkIAMPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubnetworkIAMPolicyState, opts?: pulumi.CustomResourceOptions): SubnetworkIAMPolicy {
        return new SubnetworkIAMPolicy(name, <any>state, { ...opts, id: id });
    }

    /**
     * (Computed) The etag of the subnetwork's IAM policy.
     */
    public /*out*/ readonly etag: pulumi.Output<string>;
    /**
     * The policy data generated by
     * a `google_iam_policy` data source.
     */
    public readonly policyData: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * The region of the subnetwork. If
     * unspecified, this defaults to the region configured in the provider.
     */
    public readonly region: pulumi.Output<string>;
    /**
     * The name of the subnetwork.
     */
    public readonly subnetwork: pulumi.Output<string>;

    /**
     * Create a SubnetworkIAMPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetworkIAMPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubnetworkIAMPolicyArgs | SubnetworkIAMPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: SubnetworkIAMPolicyState = argsOrState as SubnetworkIAMPolicyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["subnetwork"] = state ? state.subnetwork : undefined;
        } else {
            const args = argsOrState as SubnetworkIAMPolicyArgs | undefined;
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            if (!args || args.subnetwork === undefined) {
                throw new Error("Missing required property 'subnetwork'");
            }
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["subnetwork"] = args ? args.subnetwork : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        super("gcp:compute/subnetworkIAMPolicy:SubnetworkIAMPolicy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SubnetworkIAMPolicy resources.
 */
export interface SubnetworkIAMPolicyState {
    /**
     * (Computed) The etag of the subnetwork's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `google_iam_policy` data source.
     */
    readonly policyData?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region of the subnetwork. If
     * unspecified, this defaults to the region configured in the provider.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The name of the subnetwork.
     */
    readonly subnetwork?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SubnetworkIAMPolicy resource.
 */
export interface SubnetworkIAMPolicyArgs {
    /**
     * The policy data generated by
     * a `google_iam_policy` data source.
     */
    readonly policyData: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region of the subnetwork. If
     * unspecified, this defaults to the region configured in the provider.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The name of the subnetwork.
     */
    readonly subnetwork: pulumi.Input<string>;
}
