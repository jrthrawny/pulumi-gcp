# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Lien(pulumi.CustomResource):
    create_time: pulumi.Output[str]
    name: pulumi.Output[str]
    origin: pulumi.Output[str]
    parent: pulumi.Output[str]
    reason: pulumi.Output[str]
    restrictions: pulumi.Output[list]
    def __init__(__self__, __name__, __opts__=None, origin=None, parent=None, reason=None, restrictions=None):
        """
        A Lien represents an encumbrance on the actions that can be performed on a resource.
        
        
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] origin
        :param pulumi.Input[str] parent
        :param pulumi.Input[str] reason
        :param pulumi.Input[list] restrictions
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not origin:
            raise TypeError('Missing required property origin')
        __props__['origin'] = origin

        if not parent:
            raise TypeError('Missing required property parent')
        __props__['parent'] = parent

        if not reason:
            raise TypeError('Missing required property reason')
        __props__['reason'] = reason

        if not restrictions:
            raise TypeError('Missing required property restrictions')
        __props__['restrictions'] = restrictions

        __props__['create_time'] = None
        __props__['name'] = None

        super(Lien, __self__).__init__(
            'gcp:resourcemanager/lien:Lien',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

