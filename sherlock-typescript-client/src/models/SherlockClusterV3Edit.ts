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
 * @interface SherlockClusterV3Edit
 */
export interface SherlockClusterV3Edit {
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockClusterV3Edit
     */
    address?: string;
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockClusterV3Edit
     */
    base?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockClusterV3Edit
     */
    helmfileRef?: string;
    /**
     * If present, requires membership in the given role for mutations. Set to an empty string to clear.
     * @type {string}
     * @memberof SherlockClusterV3Edit
     */
    requiredRole?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockClusterV3Edit
     */
    requiresSuitability?: boolean;
}

/**
 * Check if a given object implements the SherlockClusterV3Edit interface.
 */
export function instanceOfSherlockClusterV3Edit(value: object): value is SherlockClusterV3Edit {
    return true;
}

export function SherlockClusterV3EditFromJSON(json: any): SherlockClusterV3Edit {
    return SherlockClusterV3EditFromJSONTyped(json, false);
}

export function SherlockClusterV3EditFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockClusterV3Edit {
    if (json == null) {
        return json;
    }
    return {
        
        'address': json['address'] == null ? undefined : json['address'],
        'base': json['base'] == null ? undefined : json['base'],
        'helmfileRef': json['helmfileRef'] == null ? undefined : json['helmfileRef'],
        'requiredRole': json['requiredRole'] == null ? undefined : json['requiredRole'],
        'requiresSuitability': json['requiresSuitability'] == null ? undefined : json['requiresSuitability'],
    };
}

  export function SherlockClusterV3EditToJSON(json: any): SherlockClusterV3Edit {
      return SherlockClusterV3EditToJSONTyped(json, false);
  }

  export function SherlockClusterV3EditToJSONTyped(value?: SherlockClusterV3Edit | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'address': value['address'],
        'base': value['base'],
        'helmfileRef': value['helmfileRef'],
        'requiredRole': value['requiredRole'],
        'requiresSuitability': value['requiresSuitability'],
    };
}

