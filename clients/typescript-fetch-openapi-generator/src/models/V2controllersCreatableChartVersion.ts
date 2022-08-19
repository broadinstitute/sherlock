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
/**
 * 
 * @export
 * @interface V2controllersCreatableChartVersion
 */
export interface V2controllersCreatableChartVersion {
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersCreatableChartVersion
     */
    chart?: string;
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersCreatableChartVersion
     */
    chartVersion?: string;
}

/**
 * Check if a given object implements the V2controllersCreatableChartVersion interface.
 */
export function instanceOfV2controllersCreatableChartVersion(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersCreatableChartVersionFromJSON(json: any): V2controllersCreatableChartVersion {
    return V2controllersCreatableChartVersionFromJSONTyped(json, false);
}

export function V2controllersCreatableChartVersionFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersCreatableChartVersion {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'chart': !exists(json, 'chart') ? undefined : json['chart'],
        'chartVersion': !exists(json, 'chartVersion') ? undefined : json['chartVersion'],
    };
}

export function V2controllersCreatableChartVersionToJSON(value?: V2controllersCreatableChartVersion | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'chart': value.chart,
        'chartVersion': value.chartVersion,
    };
}

