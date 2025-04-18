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
 * @interface SherlockChartReleaseV3Edit
 */
export interface SherlockChartReleaseV3Edit {
    /**
     * 
     * @type {boolean}
     * @memberof SherlockChartReleaseV3Edit
     */
    includedInBulkChangesets?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3Edit
     */
    pagerdutyIntegration?: string;
    /**
     * When creating, will use the chart's default if left empty
     * @type {number}
     * @memberof SherlockChartReleaseV3Edit
     */
    port?: number;
    /**
     * When creating, will use the chart's default if left empty
     * @type {string}
     * @memberof SherlockChartReleaseV3Edit
     */
    protocol?: string;
    /**
     * When creating, will use the chart's default if left empty
     * @type {string}
     * @memberof SherlockChartReleaseV3Edit
     */
    subdomain?: string;
}

/**
 * Check if a given object implements the SherlockChartReleaseV3Edit interface.
 */
export function instanceOfSherlockChartReleaseV3Edit(value: object): value is SherlockChartReleaseV3Edit {
    return true;
}

export function SherlockChartReleaseV3EditFromJSON(json: any): SherlockChartReleaseV3Edit {
    return SherlockChartReleaseV3EditFromJSONTyped(json, false);
}

export function SherlockChartReleaseV3EditFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockChartReleaseV3Edit {
    if (json == null) {
        return json;
    }
    return {
        
        'includedInBulkChangesets': json['includedInBulkChangesets'] == null ? undefined : json['includedInBulkChangesets'],
        'pagerdutyIntegration': json['pagerdutyIntegration'] == null ? undefined : json['pagerdutyIntegration'],
        'port': json['port'] == null ? undefined : json['port'],
        'protocol': json['protocol'] == null ? undefined : json['protocol'],
        'subdomain': json['subdomain'] == null ? undefined : json['subdomain'],
    };
}

export function SherlockChartReleaseV3EditToJSON(json: any): SherlockChartReleaseV3Edit {
    return SherlockChartReleaseV3EditToJSONTyped(json, false);
}

export function SherlockChartReleaseV3EditToJSONTyped(value?: SherlockChartReleaseV3Edit | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'includedInBulkChangesets': value['includedInBulkChangesets'],
        'pagerdutyIntegration': value['pagerdutyIntegration'],
        'port': value['port'],
        'protocol': value['protocol'],
        'subdomain': value['subdomain'],
    };
}

