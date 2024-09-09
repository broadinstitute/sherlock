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
 * @interface SherlockDatabaseInstanceV3Create
 */
export interface SherlockDatabaseInstanceV3Create {
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockDatabaseInstanceV3Create
     */
    chartRelease?: string;
    /**
     * When creating, defaults to the chart name
     * @type {string}
     * @memberof SherlockDatabaseInstanceV3Create
     */
    defaultDatabase?: string;
    /**
     * Required if platform is 'google'
     * @type {string}
     * @memberof SherlockDatabaseInstanceV3Create
     */
    googleProject?: string;
    /**
     * Required if platform is 'google' or 'azure'
     * @type {string}
     * @memberof SherlockDatabaseInstanceV3Create
     */
    instanceName?: string;
    /**
     * 'google', 'azure', or default 'kubernetes'
     * @type {string}
     * @memberof SherlockDatabaseInstanceV3Create
     */
    platform?: string;
}

/**
 * Check if a given object implements the SherlockDatabaseInstanceV3Create interface.
 */
export function instanceOfSherlockDatabaseInstanceV3Create(value: object): value is SherlockDatabaseInstanceV3Create {
    return true;
}

export function SherlockDatabaseInstanceV3CreateFromJSON(json: any): SherlockDatabaseInstanceV3Create {
    return SherlockDatabaseInstanceV3CreateFromJSONTyped(json, false);
}

export function SherlockDatabaseInstanceV3CreateFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockDatabaseInstanceV3Create {
    if (json == null) {
        return json;
    }
    return {
        
        'chartRelease': json['chartRelease'] == null ? undefined : json['chartRelease'],
        'defaultDatabase': json['defaultDatabase'] == null ? undefined : json['defaultDatabase'],
        'googleProject': json['googleProject'] == null ? undefined : json['googleProject'],
        'instanceName': json['instanceName'] == null ? undefined : json['instanceName'],
        'platform': json['platform'] == null ? undefined : json['platform'],
    };
}

  export function SherlockDatabaseInstanceV3CreateToJSON(json: any): SherlockDatabaseInstanceV3Create {
      return SherlockDatabaseInstanceV3CreateToJSONTyped(json, false);
  }

  export function SherlockDatabaseInstanceV3CreateToJSONTyped(value?: SherlockDatabaseInstanceV3Create | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'chartRelease': value['chartRelease'],
        'defaultDatabase': value['defaultDatabase'],
        'googleProject': value['googleProject'],
        'instanceName': value['instanceName'],
        'platform': value['platform'],
    };
}

