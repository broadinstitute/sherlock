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

import { exists, mapValues } from '../runtime';
import type { SherlockChartReleaseV3 } from './SherlockChartReleaseV3';
import {
    SherlockChartReleaseV3FromJSON,
    SherlockChartReleaseV3FromJSONTyped,
    SherlockChartReleaseV3ToJSON,
} from './SherlockChartReleaseV3';

/**
 * 
 * @export
 * @interface SherlockDatabaseInstanceV3
 */
export interface SherlockDatabaseInstanceV3 {
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockDatabaseInstanceV3
     */
    chartRelease?: string;
    /**
     * 
     * @type {SherlockChartReleaseV3}
     * @memberof SherlockDatabaseInstanceV3
     */
    chartReleaseInfo?: SherlockChartReleaseV3;
    /**
     * 
     * @type {Date}
     * @memberof SherlockDatabaseInstanceV3
     */
    createdAt?: Date;
    /**
     * When creating, defaults to the chart name
     * @type {string}
     * @memberof SherlockDatabaseInstanceV3
     */
    defaultDatabase?: string;
    /**
     * Required if platform is 'google'
     * @type {string}
     * @memberof SherlockDatabaseInstanceV3
     */
    googleProject?: string;
    /**
     * 
     * @type {number}
     * @memberof SherlockDatabaseInstanceV3
     */
    id?: number;
    /**
     * Required if platform is 'google' or 'azure'
     * @type {string}
     * @memberof SherlockDatabaseInstanceV3
     */
    instanceName?: string;
    /**
     * 'google', 'azure', or default 'kubernetes'
     * @type {string}
     * @memberof SherlockDatabaseInstanceV3
     */
    platform?: string;
    /**
     * 
     * @type {Date}
     * @memberof SherlockDatabaseInstanceV3
     */
    updatedAt?: Date;
}

/**
 * Check if a given object implements the SherlockDatabaseInstanceV3 interface.
 */
export function instanceOfSherlockDatabaseInstanceV3(value: object): boolean {
    return true;
}

export function SherlockDatabaseInstanceV3FromJSON(json: any): SherlockDatabaseInstanceV3 {
    return SherlockDatabaseInstanceV3FromJSONTyped(json, false);
}

export function SherlockDatabaseInstanceV3FromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockDatabaseInstanceV3 {
    if (json === undefined || json === null) {
        return json;
    }
    return {
        
        'chartRelease': !exists(json, 'chartRelease') ? undefined : json['chartRelease'],
        'chartReleaseInfo': !exists(json, 'chartReleaseInfo') ? undefined : SherlockChartReleaseV3FromJSON(json['chartReleaseInfo']),
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'defaultDatabase': !exists(json, 'defaultDatabase') ? undefined : json['defaultDatabase'],
        'googleProject': !exists(json, 'googleProject') ? undefined : json['googleProject'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'instanceName': !exists(json, 'instanceName') ? undefined : json['instanceName'],
        'platform': !exists(json, 'platform') ? undefined : json['platform'],
        'updatedAt': !exists(json, 'updatedAt') ? undefined : (new Date(json['updatedAt'])),
    };
}

export function SherlockDatabaseInstanceV3ToJSON(value?: SherlockDatabaseInstanceV3 | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'chartRelease': value['chartRelease'],
        'chartReleaseInfo': SherlockChartReleaseV3ToJSON(value['chartReleaseInfo']),
        'createdAt': !exists(value, 'createdAt') ? undefined : ((value['createdAt']).toISOString()),
        'defaultDatabase': value['defaultDatabase'],
        'googleProject': value['googleProject'],
        'id': value['id'],
        'instanceName': value['instanceName'],
        'platform': value['platform'],
        'updatedAt': !exists(value, 'updatedAt') ? undefined : ((value['updatedAt']).toISOString()),
    };
}

