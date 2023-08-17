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
 * @interface V2controllersChangesetPlanRequestEnvironmentEntry
 */
export interface V2controllersChangesetPlanRequestEnvironmentEntry {
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangesetPlanRequestEnvironmentEntry
     */
    environment?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof V2controllersChangesetPlanRequestEnvironmentEntry
     */
    excludeCharts?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangesetPlanRequestEnvironmentEntry
     */
    followVersionsFromOtherEnvironment?: string;
    /**
     * If omitted, will include all chart releases that haven't opted out of bulk updates
     * @type {Array<string>}
     * @memberof V2controllersChangesetPlanRequestEnvironmentEntry
     */
    includeCharts?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangesetPlanRequestEnvironmentEntry
     */
    useExactVersionsFromOtherEnvironment?: string;
    /**
     * If this is set, also copy the fc-dev ref from an OtherEnvironment
     * @type {boolean}
     * @memberof V2controllersChangesetPlanRequestEnvironmentEntry
     */
    useOthersFirecloudDevelopRef?: boolean;
    /**
     * If this is set, also copy the helmfile ref from an OtherEnvironment
     * @type {boolean}
     * @memberof V2controllersChangesetPlanRequestEnvironmentEntry
     */
    useOthersHelmfileRef?: boolean;
}

/**
 * Check if a given object implements the V2controllersChangesetPlanRequestEnvironmentEntry interface.
 */
export function instanceOfV2controllersChangesetPlanRequestEnvironmentEntry(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersChangesetPlanRequestEnvironmentEntryFromJSON(json: any): V2controllersChangesetPlanRequestEnvironmentEntry {
    return V2controllersChangesetPlanRequestEnvironmentEntryFromJSONTyped(json, false);
}

export function V2controllersChangesetPlanRequestEnvironmentEntryFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersChangesetPlanRequestEnvironmentEntry {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'environment': !exists(json, 'environment') ? undefined : json['environment'],
        'excludeCharts': !exists(json, 'excludeCharts') ? undefined : json['excludeCharts'],
        'followVersionsFromOtherEnvironment': !exists(json, 'followVersionsFromOtherEnvironment') ? undefined : json['followVersionsFromOtherEnvironment'],
        'includeCharts': !exists(json, 'includeCharts') ? undefined : json['includeCharts'],
        'useExactVersionsFromOtherEnvironment': !exists(json, 'useExactVersionsFromOtherEnvironment') ? undefined : json['useExactVersionsFromOtherEnvironment'],
        'useOthersFirecloudDevelopRef': !exists(json, 'useOthersFirecloudDevelopRef') ? undefined : json['useOthersFirecloudDevelopRef'],
        'useOthersHelmfileRef': !exists(json, 'useOthersHelmfileRef') ? undefined : json['useOthersHelmfileRef'],
    };
}

export function V2controllersChangesetPlanRequestEnvironmentEntryToJSON(value?: V2controllersChangesetPlanRequestEnvironmentEntry | null): any {
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
        'useOthersFirecloudDevelopRef': value.useOthersFirecloudDevelopRef,
        'useOthersHelmfileRef': value.useOthersHelmfileRef,
    };
}

