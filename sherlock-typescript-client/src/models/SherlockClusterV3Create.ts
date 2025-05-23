/* tslint:disable */
/* eslint-disable */
/**
 * Sherlock
 * The Data Science Platform\'s source-of-truth service. Note: this API will try to load and return associations in responses, so clients won\'t need to make as many requests. This behavior isn\'t recursive, though, so associations of associations are *not* fully loaded (even if it might seem that way from looking at the data types).
 *
 * The version of the OpenAPI document: development
 * Contact: dsp-devops@broadinstitute.org
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface SherlockClusterV3Create
 */
export interface SherlockClusterV3Create {
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockClusterV3Create
     */
    address?: string;
    /**
     * Required when creating if provider is 'azure'
     * @type {string}
     * @memberof SherlockClusterV3Create
     */
    azureSubscription?: string;
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockClusterV3Create
     */
    base?: string;
    /**
     * Required when creating if provider is 'google'
     * @type {string}
     * @memberof SherlockClusterV3Create
     */
    googleProject?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockClusterV3Create
     */
    helmfileRef?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockClusterV3Create
     */
    location?: string;
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockClusterV3Create
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockClusterV3Create
     */
    provider?: SherlockClusterV3CreateProviderEnum;
    /**
     * If present, requires membership in the given role for mutations. Set to an empty string to clear.
     * @type {string}
     * @memberof SherlockClusterV3Create
     */
    requiredRole?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockClusterV3Create
     */
    requiresSuitability?: boolean;
}


/**
 * @export
 */
export const SherlockClusterV3CreateProviderEnum = {
    Google: 'google',
    Azure: 'azure'
} as const;
export type SherlockClusterV3CreateProviderEnum = typeof SherlockClusterV3CreateProviderEnum[keyof typeof SherlockClusterV3CreateProviderEnum];


/**
 * Check if a given object implements the SherlockClusterV3Create interface.
 */
export function instanceOfSherlockClusterV3Create(value: object): value is SherlockClusterV3Create {
    return true;
}

export function SherlockClusterV3CreateFromJSON(json: any): SherlockClusterV3Create {
    return SherlockClusterV3CreateFromJSONTyped(json, false);
}

export function SherlockClusterV3CreateFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockClusterV3Create {
    if (json == null) {
        return json;
    }
    return {
        
        'address': json['address'] == null ? undefined : json['address'],
        'azureSubscription': json['azureSubscription'] == null ? undefined : json['azureSubscription'],
        'base': json['base'] == null ? undefined : json['base'],
        'googleProject': json['googleProject'] == null ? undefined : json['googleProject'],
        'helmfileRef': json['helmfileRef'] == null ? undefined : json['helmfileRef'],
        'location': json['location'] == null ? undefined : json['location'],
        'name': json['name'] == null ? undefined : json['name'],
        'provider': json['provider'] == null ? undefined : json['provider'],
        'requiredRole': json['requiredRole'] == null ? undefined : json['requiredRole'],
        'requiresSuitability': json['requiresSuitability'] == null ? undefined : json['requiresSuitability'],
    };
}

export function SherlockClusterV3CreateToJSON(json: any): SherlockClusterV3Create {
    return SherlockClusterV3CreateToJSONTyped(json, false);
}

export function SherlockClusterV3CreateToJSONTyped(value?: SherlockClusterV3Create | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'address': value['address'],
        'azureSubscription': value['azureSubscription'],
        'base': value['base'],
        'googleProject': value['googleProject'],
        'helmfileRef': value['helmfileRef'],
        'location': value['location'],
        'name': value['name'],
        'provider': value['provider'],
        'requiredRole': value['requiredRole'],
        'requiresSuitability': value['requiresSuitability'],
    };
}

