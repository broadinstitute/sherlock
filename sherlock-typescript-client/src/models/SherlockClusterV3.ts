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
import type { SherlockCiIdentifierV3 } from './SherlockCiIdentifierV3';
import {
    SherlockCiIdentifierV3FromJSON,
    SherlockCiIdentifierV3FromJSONTyped,
    SherlockCiIdentifierV3ToJSON,
} from './SherlockCiIdentifierV3';

/**
 * 
 * @export
 * @interface SherlockClusterV3
 */
export interface SherlockClusterV3 {
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockClusterV3
     */
    address?: string;
    /**
     * Required when creating if provider is 'azure'
     * @type {string}
     * @memberof SherlockClusterV3
     */
    azureSubscription?: string;
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockClusterV3
     */
    base?: string;
    /**
     * 
     * @type {SherlockCiIdentifierV3}
     * @memberof SherlockClusterV3
     */
    ciIdentifier?: SherlockCiIdentifierV3;
    /**
     * 
     * @type {Date}
     * @memberof SherlockClusterV3
     */
    createdAt?: Date;
    /**
     * Required when creating if provider is 'google'
     * @type {string}
     * @memberof SherlockClusterV3
     */
    googleProject?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockClusterV3
     */
    helmfileRef?: string;
    /**
     * 
     * @type {number}
     * @memberof SherlockClusterV3
     */
    id?: number;
    /**
     * 
     * @type {string}
     * @memberof SherlockClusterV3
     */
    location?: string;
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockClusterV3
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockClusterV3
     */
    provider?: SherlockClusterV3ProviderEnum;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockClusterV3
     */
    requiresSuitability?: boolean;
    /**
     * 
     * @type {Date}
     * @memberof SherlockClusterV3
     */
    updatedAt?: Date;
}


/**
 * @export
 */
export const SherlockClusterV3ProviderEnum = {
    Google: 'google',
    Azure: 'azure'
} as const;
export type SherlockClusterV3ProviderEnum = typeof SherlockClusterV3ProviderEnum[keyof typeof SherlockClusterV3ProviderEnum];


/**
 * Check if a given object implements the SherlockClusterV3 interface.
 */
export function instanceOfSherlockClusterV3(value: object): boolean {
    return true;
}

export function SherlockClusterV3FromJSON(json: any): SherlockClusterV3 {
    return SherlockClusterV3FromJSONTyped(json, false);
}

export function SherlockClusterV3FromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockClusterV3 {
    if (json == null) {
        return json;
    }
    return {
        
        'address': json['address'] == null ? undefined : json['address'],
        'azureSubscription': json['azureSubscription'] == null ? undefined : json['azureSubscription'],
        'base': json['base'] == null ? undefined : json['base'],
        'ciIdentifier': json['ciIdentifier'] == null ? undefined : SherlockCiIdentifierV3FromJSON(json['ciIdentifier']),
        'createdAt': json['createdAt'] == null ? undefined : (new Date(json['createdAt'])),
        'googleProject': json['googleProject'] == null ? undefined : json['googleProject'],
        'helmfileRef': json['helmfileRef'] == null ? undefined : json['helmfileRef'],
        'id': json['id'] == null ? undefined : json['id'],
        'location': json['location'] == null ? undefined : json['location'],
        'name': json['name'] == null ? undefined : json['name'],
        'provider': json['provider'] == null ? undefined : json['provider'],
        'requiresSuitability': json['requiresSuitability'] == null ? undefined : json['requiresSuitability'],
        'updatedAt': json['updatedAt'] == null ? undefined : (new Date(json['updatedAt'])),
    };
}

export function SherlockClusterV3ToJSON(value?: SherlockClusterV3 | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'address': value['address'],
        'azureSubscription': value['azureSubscription'],
        'base': value['base'],
        'ciIdentifier': SherlockCiIdentifierV3ToJSON(value['ciIdentifier']),
        'createdAt': value['createdAt'] == null ? undefined : ((value['createdAt']).toISOString()),
        'googleProject': value['googleProject'],
        'helmfileRef': value['helmfileRef'],
        'id': value['id'],
        'location': value['location'],
        'name': value['name'],
        'provider': value['provider'],
        'requiresSuitability': value['requiresSuitability'],
        'updatedAt': value['updatedAt'] == null ? undefined : ((value['updatedAt']).toISOString()),
    };
}

