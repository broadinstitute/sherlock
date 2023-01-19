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
/**
 * 
 * @export
 * @interface PagerdutyAlertSummary
 */
export interface PagerdutyAlertSummary {
    /**
     * 
     * @type {string}
     * @memberof PagerdutyAlertSummary
     */
    summary?: string;
}

/**
 * Check if a given object implements the PagerdutyAlertSummary interface.
 */
export function instanceOfPagerdutyAlertSummary(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PagerdutyAlertSummaryFromJSON(json: any): PagerdutyAlertSummary {
    return PagerdutyAlertSummaryFromJSONTyped(json, false);
}

export function PagerdutyAlertSummaryFromJSONTyped(json: any, ignoreDiscriminator: boolean): PagerdutyAlertSummary {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'summary': !exists(json, 'summary') ? undefined : json['summary'],
    };
}

export function PagerdutyAlertSummaryToJSON(value?: PagerdutyAlertSummary | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'summary': value.summary,
    };
}

