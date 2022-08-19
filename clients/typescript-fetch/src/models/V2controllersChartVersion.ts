/* tslint:disable */
/* eslint-disable */
/**
 * Sherlock
 * The Data Science Platform\'s source-of-truth service
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
 * @interface V2controllersChartVersion
 */
export interface V2controllersChartVersion {
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersChartVersion
     */
    chart?: string;
    /**
     * 
     * @type {V2controllersChart}
     * @memberof V2controllersChartVersion
     */
    chartInfo?: V2controllersChart;
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersChartVersion
     */
    chartVersion?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChartVersion
     */
    createdAt?: string;
    /**
     * 
     * @type {number}
     * @memberof V2controllersChartVersion
     */
    id?: number;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChartVersion
     */
    updatedAt?: string;
}

/**
 * Check if a given object implements the V2controllersChartVersion interface.
 */
export function instanceOfV2controllersChartVersion(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersChartVersionFromJSON(json: any): V2controllersChartVersion {
    return V2controllersChartVersionFromJSONTyped(json, false);
}

export function V2controllersChartVersionFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersChartVersion {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'chart': !exists(json, 'chart') ? undefined : json['chart'],
        'chartInfo': !exists(json, 'chartInfo') ? undefined : V2controllersChartFromJSON(json['chartInfo']),
        'chartVersion': !exists(json, 'chartVersion') ? undefined : json['chartVersion'],
        'createdAt': !exists(json, 'createdAt') ? undefined : json['createdAt'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'updatedAt': !exists(json, 'updatedAt') ? undefined : json['updatedAt'],
    };
}

export function V2controllersChartVersionToJSON(value?: V2controllersChartVersion | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'chart': value.chart,
        'chartInfo': V2controllersChartToJSON(value.chartInfo),
        'chartVersion': value.chartVersion,
        'createdAt': value.createdAt,
        'id': value.id,
        'updatedAt': value.updatedAt,
    };
}

