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
import type { SherlockChangesetV3PlanRequestEnvironmentEntry } from './SherlockChangesetV3PlanRequestEnvironmentEntry';
import {
    SherlockChangesetV3PlanRequestEnvironmentEntryFromJSON,
    SherlockChangesetV3PlanRequestEnvironmentEntryFromJSONTyped,
    SherlockChangesetV3PlanRequestEnvironmentEntryToJSON,
    SherlockChangesetV3PlanRequestEnvironmentEntryToJSONTyped,
} from './SherlockChangesetV3PlanRequestEnvironmentEntry';
import type { SherlockChangesetV3PlanRequestChartReleaseEntry } from './SherlockChangesetV3PlanRequestChartReleaseEntry';
import {
    SherlockChangesetV3PlanRequestChartReleaseEntryFromJSON,
    SherlockChangesetV3PlanRequestChartReleaseEntryFromJSONTyped,
    SherlockChangesetV3PlanRequestChartReleaseEntryToJSON,
    SherlockChangesetV3PlanRequestChartReleaseEntryToJSONTyped,
} from './SherlockChangesetV3PlanRequestChartReleaseEntry';

/**
 * 
 * @export
 * @interface SherlockChangesetV3PlanRequest
 */
export interface SherlockChangesetV3PlanRequest {
    /**
     * 
     * @type {Array<SherlockChangesetV3PlanRequestChartReleaseEntry>}
     * @memberof SherlockChangesetV3PlanRequest
     */
    chartReleases?: Array<SherlockChangesetV3PlanRequestChartReleaseEntry>;
    /**
     * 
     * @type {Array<SherlockChangesetV3PlanRequestEnvironmentEntry>}
     * @memberof SherlockChangesetV3PlanRequest
     */
    environments?: Array<SherlockChangesetV3PlanRequestEnvironmentEntry>;
    /**
     * 
     * @type {Array<number>}
     * @memberof SherlockChangesetV3PlanRequest
     */
    recreateChangesets?: Array<number>;
}

/**
 * Check if a given object implements the SherlockChangesetV3PlanRequest interface.
 */
export function instanceOfSherlockChangesetV3PlanRequest(value: object): value is SherlockChangesetV3PlanRequest {
    return true;
}

export function SherlockChangesetV3PlanRequestFromJSON(json: any): SherlockChangesetV3PlanRequest {
    return SherlockChangesetV3PlanRequestFromJSONTyped(json, false);
}

export function SherlockChangesetV3PlanRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockChangesetV3PlanRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'chartReleases': json['chartReleases'] == null ? undefined : ((json['chartReleases'] as Array<any>).map(SherlockChangesetV3PlanRequestChartReleaseEntryFromJSON)),
        'environments': json['environments'] == null ? undefined : ((json['environments'] as Array<any>).map(SherlockChangesetV3PlanRequestEnvironmentEntryFromJSON)),
        'recreateChangesets': json['recreateChangesets'] == null ? undefined : json['recreateChangesets'],
    };
}

export function SherlockChangesetV3PlanRequestToJSON(json: any): SherlockChangesetV3PlanRequest {
    return SherlockChangesetV3PlanRequestToJSONTyped(json, false);
}

export function SherlockChangesetV3PlanRequestToJSONTyped(value?: SherlockChangesetV3PlanRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'chartReleases': value['chartReleases'] == null ? undefined : ((value['chartReleases'] as Array<any>).map(SherlockChangesetV3PlanRequestChartReleaseEntryToJSON)),
        'environments': value['environments'] == null ? undefined : ((value['environments'] as Array<any>).map(SherlockChangesetV3PlanRequestEnvironmentEntryToJSON)),
        'recreateChangesets': value['recreateChangesets'],
    };
}

