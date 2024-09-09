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
export function instanceOfSherlockChangesetV3PlanRequestChartReleaseEntry(value: object): value is SherlockChangesetV3PlanRequestChartReleaseEntry {
    return true;
}

export function SherlockChangesetV3PlanRequestChartReleaseEntryFromJSON(json: any): SherlockChangesetV3PlanRequestChartReleaseEntry {
    return SherlockChangesetV3PlanRequestChartReleaseEntryFromJSONTyped(json, false);
}

export function SherlockChangesetV3PlanRequestChartReleaseEntryFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockChangesetV3PlanRequestChartReleaseEntry {
    if (json == null) {
        return json;
    }
    return {
        
        'chartRelease': json['chartRelease'] == null ? undefined : json['chartRelease'],
        'followVersionsFromOtherChartRelease': json['followVersionsFromOtherChartRelease'] == null ? undefined : json['followVersionsFromOtherChartRelease'],
        'toAppVersionBranch': json['toAppVersionBranch'] == null ? undefined : json['toAppVersionBranch'],
        'toAppVersionCommit': json['toAppVersionCommit'] == null ? undefined : json['toAppVersionCommit'],
        'toAppVersionExact': json['toAppVersionExact'] == null ? undefined : json['toAppVersionExact'],
        'toAppVersionFollowChartRelease': json['toAppVersionFollowChartRelease'] == null ? undefined : json['toAppVersionFollowChartRelease'],
        'toAppVersionResolver': json['toAppVersionResolver'] == null ? undefined : json['toAppVersionResolver'],
        'toChartVersionExact': json['toChartVersionExact'] == null ? undefined : json['toChartVersionExact'],
        'toChartVersionFollowChartRelease': json['toChartVersionFollowChartRelease'] == null ? undefined : json['toChartVersionFollowChartRelease'],
        'toChartVersionResolver': json['toChartVersionResolver'] == null ? undefined : json['toChartVersionResolver'],
        'toHelmfileRef': json['toHelmfileRef'] == null ? undefined : json['toHelmfileRef'],
        'toHelmfileRefEnabled': json['toHelmfileRefEnabled'] == null ? undefined : json['toHelmfileRefEnabled'],
        'useExactVersionsFromOtherChartRelease': json['useExactVersionsFromOtherChartRelease'] == null ? undefined : json['useExactVersionsFromOtherChartRelease'],
    };
}

  export function SherlockChangesetV3PlanRequestChartReleaseEntryToJSON(json: any): SherlockChangesetV3PlanRequestChartReleaseEntry {
      return SherlockChangesetV3PlanRequestChartReleaseEntryToJSONTyped(json, false);
  }

  export function SherlockChangesetV3PlanRequestChartReleaseEntryToJSONTyped(value?: SherlockChangesetV3PlanRequestChartReleaseEntry | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'chartRelease': value['chartRelease'],
        'followVersionsFromOtherChartRelease': value['followVersionsFromOtherChartRelease'],
        'toAppVersionBranch': value['toAppVersionBranch'],
        'toAppVersionCommit': value['toAppVersionCommit'],
        'toAppVersionExact': value['toAppVersionExact'],
        'toAppVersionFollowChartRelease': value['toAppVersionFollowChartRelease'],
        'toAppVersionResolver': value['toAppVersionResolver'],
        'toChartVersionExact': value['toChartVersionExact'],
        'toChartVersionFollowChartRelease': value['toChartVersionFollowChartRelease'],
        'toChartVersionResolver': value['toChartVersionResolver'],
        'toHelmfileRef': value['toHelmfileRef'],
        'toHelmfileRefEnabled': value['toHelmfileRefEnabled'],
        'useExactVersionsFromOtherChartRelease': value['useExactVersionsFromOtherChartRelease'],
    };
}

