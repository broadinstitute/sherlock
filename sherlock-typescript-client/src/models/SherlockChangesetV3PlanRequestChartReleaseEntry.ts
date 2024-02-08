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
 * @interface SherlockChangesetV3PlanRequestChartReleaseEntry
 */
export interface SherlockChangesetV3PlanRequestChartReleaseEntry {
    /**
     * 
     * @type {string}
     * @memberof SherlockChangesetV3PlanRequestChartReleaseEntry
     */
    chartRelease?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChangesetV3PlanRequestChartReleaseEntry
     */
    followVersionsFromOtherChartRelease?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChangesetV3PlanRequestChartReleaseEntry
     */
    toAppVersionBranch?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChangesetV3PlanRequestChartReleaseEntry
     */
    toAppVersionCommit?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChangesetV3PlanRequestChartReleaseEntry
     */
    toAppVersionExact?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChangesetV3PlanRequestChartReleaseEntry
     */
    toAppVersionFollowChartRelease?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChangesetV3PlanRequestChartReleaseEntry
     */
    toAppVersionResolver?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChangesetV3PlanRequestChartReleaseEntry
     */
    toChartVersionExact?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChangesetV3PlanRequestChartReleaseEntry
     */
    toChartVersionFollowChartRelease?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChangesetV3PlanRequestChartReleaseEntry
     */
    toChartVersionResolver?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChangesetV3PlanRequestChartReleaseEntry
     */
    toFirecloudDevelopRef?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChangesetV3PlanRequestChartReleaseEntry
     */
    toHelmfileRef?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockChangesetV3PlanRequestChartReleaseEntry
     */
    toHelmfileRefEnabled?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockChangesetV3PlanRequestChartReleaseEntry
     */
    useExactVersionsFromOtherChartRelease?: string;
}

/**
 * Check if a given object implements the SherlockChangesetV3PlanRequestChartReleaseEntry interface.
 */
export function instanceOfSherlockChangesetV3PlanRequestChartReleaseEntry(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SherlockChangesetV3PlanRequestChartReleaseEntryFromJSON(json: any): SherlockChangesetV3PlanRequestChartReleaseEntry {
    return SherlockChangesetV3PlanRequestChartReleaseEntryFromJSONTyped(json, false);
}

export function SherlockChangesetV3PlanRequestChartReleaseEntryFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockChangesetV3PlanRequestChartReleaseEntry {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'chartRelease': !exists(json, 'chartRelease') ? undefined : json['chartRelease'],
        'followVersionsFromOtherChartRelease': !exists(json, 'followVersionsFromOtherChartRelease') ? undefined : json['followVersionsFromOtherChartRelease'],
        'toAppVersionBranch': !exists(json, 'toAppVersionBranch') ? undefined : json['toAppVersionBranch'],
        'toAppVersionCommit': !exists(json, 'toAppVersionCommit') ? undefined : json['toAppVersionCommit'],
        'toAppVersionExact': !exists(json, 'toAppVersionExact') ? undefined : json['toAppVersionExact'],
        'toAppVersionFollowChartRelease': !exists(json, 'toAppVersionFollowChartRelease') ? undefined : json['toAppVersionFollowChartRelease'],
        'toAppVersionResolver': !exists(json, 'toAppVersionResolver') ? undefined : json['toAppVersionResolver'],
        'toChartVersionExact': !exists(json, 'toChartVersionExact') ? undefined : json['toChartVersionExact'],
        'toChartVersionFollowChartRelease': !exists(json, 'toChartVersionFollowChartRelease') ? undefined : json['toChartVersionFollowChartRelease'],
        'toChartVersionResolver': !exists(json, 'toChartVersionResolver') ? undefined : json['toChartVersionResolver'],
        'toFirecloudDevelopRef': !exists(json, 'toFirecloudDevelopRef') ? undefined : json['toFirecloudDevelopRef'],
        'toHelmfileRef': !exists(json, 'toHelmfileRef') ? undefined : json['toHelmfileRef'],
        'toHelmfileRefEnabled': !exists(json, 'toHelmfileRefEnabled') ? undefined : json['toHelmfileRefEnabled'],
        'useExactVersionsFromOtherChartRelease': !exists(json, 'useExactVersionsFromOtherChartRelease') ? undefined : json['useExactVersionsFromOtherChartRelease'],
    };
}

export function SherlockChangesetV3PlanRequestChartReleaseEntryToJSON(value?: SherlockChangesetV3PlanRequestChartReleaseEntry | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'chartRelease': value.chartRelease,
        'followVersionsFromOtherChartRelease': value.followVersionsFromOtherChartRelease,
        'toAppVersionBranch': value.toAppVersionBranch,
        'toAppVersionCommit': value.toAppVersionCommit,
        'toAppVersionExact': value.toAppVersionExact,
        'toAppVersionFollowChartRelease': value.toAppVersionFollowChartRelease,
        'toAppVersionResolver': value.toAppVersionResolver,
        'toChartVersionExact': value.toChartVersionExact,
        'toChartVersionFollowChartRelease': value.toChartVersionFollowChartRelease,
        'toChartVersionResolver': value.toChartVersionResolver,
        'toFirecloudDevelopRef': value.toFirecloudDevelopRef,
        'toHelmfileRef': value.toHelmfileRef,
        'toHelmfileRefEnabled': value.toHelmfileRefEnabled,
        'useExactVersionsFromOtherChartRelease': value.useExactVersionsFromOtherChartRelease,
    };
}
