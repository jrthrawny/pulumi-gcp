# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Image(pulumi.CustomResource):
    create_timeout: pulumi.Output[int]
    """
    Configurable timeout in minutes for creating images. Default is 4 minutes.
    """
    description: pulumi.Output[str]
    """
    The description of the image to be created
    """
    family: pulumi.Output[str]
    """
    The name of the image family to which this image belongs.
    """
    label_fingerprint: pulumi.Output[str]
    """
    The fingerprint of the assigned labels.
    """
    labels: pulumi.Output[dict]
    """
    A set of key/value label pairs to assign to the image.
    """
    licenses: pulumi.Output[list]
    """
    A list of license URIs to apply to this image. Changing this
    forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    A unique name for the resource, required by GCE.
    Changing this forces a new resource to be created.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    raw_disk: pulumi.Output[dict]
    """
    The raw disk that will be used as the source of the image.
    Changing this forces a new resource to be created. Structure is documented
    below.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    source_disk: pulumi.Output[str]
    """
    The URL of a disk that will be used as the source of the
    image. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, create_timeout=None, description=None, family=None, labels=None, licenses=None, name=None, project=None, raw_disk=None, source_disk=None, __name__=None, __opts__=None):
        """
        Creates a bootable VM image resource for Google Compute Engine from an existing
        tarball. For more information see [the official documentation](https://cloud.google.com/compute/docs/images) and
        [API](https://cloud.google.com/compute/docs/reference/latest/images).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] create_timeout: Configurable timeout in minutes for creating images. Default is 4 minutes.
        :param pulumi.Input[str] description: The description of the image to be created
        :param pulumi.Input[str] family: The name of the image family to which this image belongs.
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to the image.
        :param pulumi.Input[list] licenses: A list of license URIs to apply to this image. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] name: A unique name for the resource, required by GCE.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[dict] raw_disk: The raw disk that will be used as the source of the image.
               Changing this forces a new resource to be created. Structure is documented
               below.
        :param pulumi.Input[str] source_disk: The URL of a disk that will be used as the source of the
               image. Changing this forces a new resource to be created.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['create_timeout'] = create_timeout

        __props__['description'] = description

        __props__['family'] = family

        __props__['labels'] = labels

        __props__['licenses'] = licenses

        __props__['name'] = name

        __props__['project'] = project

        __props__['raw_disk'] = raw_disk

        __props__['source_disk'] = source_disk

        __props__['label_fingerprint'] = None
        __props__['self_link'] = None

        super(Image, __self__).__init__(
            'gcp:compute/image:Image',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

