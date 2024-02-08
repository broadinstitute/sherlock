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
 * @interface SherlockChangesetV3PlanRequestEnvironmentEntry
 */
export interface SherlockChangesetV3PlanRequestEnvironmentEntry {
    /**
     * 
     * @type {string}
     * @memberof SherlockChangesetV3PlanRequestEnvironmentEntry
     */
    environment?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof SherlockChangesetV3PlanRequestEnvironmentEntry
     */
    excludeCharts?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof SherlockChangesetV3PlanRequestEnvironmentEntry
     */
    followVersionsFromOtherEnvironment?: string;
    /**
     * If omitted, will include all chart releases that haven't opted out of bulk updates
     * @type {Array<string>}
     * @memberof SherlockChangesetV3PlanRequestEnvironmentEntry
     */
    includeCharts?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof SherlockChangesetV3PlanRequestEnvironmentEntry
     */
    useExactVersionsFromOtherEnvironment?: string;
}

/**
 * Check if a given object implements the SherlockChangesetV3PlanRequestEnvironmentEntry interface.
 */
export function instanceOfSherlockChangesetV3PlanRequestEnvironmentEntry(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SherlockChangesetV3PlanRequestEnvironmentEntryFromJSON(json: any): SherlockChangesetV3PlanRequestEnvironmentEntry {
    return SherlockChangesetV3PlanRequestEnvironmentEntryFromJSONTyped(json, false);
}

export function SherlockChangesetV3PlanRequestEnvironmentEntryFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockChangesetV3PlanRequestEnvironmentEntry {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'environment': !exists(json, 'environment') ? undefined : json['environment'],
        'excludeCharts': !exists(json, 'excludeCharts') ? undefined : json['excludeCharts'],
        'followVersionsFromOtherEnvironment': !exists(json, 'followVersionsFromOtherEnvironment') ? undefined : json['followVersionsFromOtherEnvironment'],
        'includeCharts': !exists(json, 'includeCharts') ? undefined : json['includeCharts'],
        'useExactVersionsFromOtherEnvironment': !exists(json, 'useExactVersionsFromOtherEnvironment') ? undefined : json['useExactVersionsFromOtherEnvironment'],
    };
}

export function SherlockChangesetV3PlanRequestEnvironmentEntryToJSON(value?: SherlockChangesetV3PlanRequestEnvironmentEntry | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'environment': value.environment,
        'excludeCharts': value.excludeCharts,
        'followVersionsFromOtherEnvironment': value.followVersionsFromOtherEnvironment,
        'includeCharts': value.includeCharts,
        'useExactVersionsFromOtherEnvironment': value.useExactVersionsFromOtherEnvironment,
    };
}

