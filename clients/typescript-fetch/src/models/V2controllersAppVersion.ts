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
import type { V2controllersChart } from './V2controllersChart';
import {
    V2controllersChartFromJSON,
    V2controllersChartFromJSONTyped,
    V2controllersChartToJSON,
} from './V2controllersChart';

/**
 * 
 * @export
 * @interface V2controllersAppVersion
 */
export interface V2controllersAppVersion {
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersAppVersion
     */
    appVersion?: string;
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersAppVersion
     */
    chart?: string;
    /**
     * 
     * @type {V2controllersChart}
     * @memberof V2controllersAppVersion
     */
    chartInfo?: V2controllersChart;
    /**
     * 
     * @type {string}
     * @memberof V2controllersAppVersion
     */
    createdAt?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersAppVersion
     */
    gitBranch?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersAppVersion
     */
    gitCommit?: string;
    /**
     * 
     * @type {number}
     * @memberof V2controllersAppVersion
     */
    id?: number;
    /**
     * 
     * @type {string}
     * @memberof V2controllersAppVersion
     */
    parentAppVersion?: string;
    /**
     * 
     * @type {object}
     * @memberof V2controllersAppVersion
     */
    parentAppVersionInfo?: object;
    /**
     * 
     * @type {string}
     * @memberof V2controllersAppVersion
     */
    updatedAt?: string;
}

/**
 * Check if a given object implements the V2controllersAppVersion interface.
 */
export function instanceOfV2controllersAppVersion(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersAppVersionFromJSON(json: any): V2controllersAppVersion {
    return V2controllersAppVersionFromJSONTyped(json, false);
}

export function V2controllersAppVersionFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersAppVersion {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'appVersion': !exists(json, 'appVersion') ? undefined : json['appVersion'],
        'chart': !exists(json, 'chart') ? undefined : json['chart'],
        'chartInfo': !exists(json, 'chartInfo') ? undefined : V2controllersChartFromJSON(json['chartInfo']),
        'createdAt': !exists(json, 'createdAt') ? undefined : json['createdAt'],
        'gitBranch': !exists(json, 'gitBranch') ? undefined : json['gitBranch'],
        'gitCommit': !exists(json, 'gitCommit') ? undefined : json['gitCommit'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'parentAppVersion': !exists(json, 'parentAppVersion') ? undefined : json['parentAppVersion'],
        'parentAppVersionInfo': !exists(json, 'parentAppVersionInfo') ? undefined : json['parentAppVersionInfo'],
        'updatedAt': !exists(json, 'updatedAt') ? undefined : json['updatedAt'],
    };
}

export function V2controllersAppVersionToJSON(value?: V2controllersAppVersion | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'appVersion': value.appVersion,
        'chart': value.chart,
        'chartInfo': V2controllersChartToJSON(value.chartInfo),
        'createdAt': value.createdAt,
        'gitBranch': value.gitBranch,
        'gitCommit': value.gitCommit,
        'id': value.id,
        'parentAppVersion': value.parentAppVersion,
        'parentAppVersionInfo': value.parentAppVersionInfo,
        'updatedAt': value.updatedAt,
    };
}

