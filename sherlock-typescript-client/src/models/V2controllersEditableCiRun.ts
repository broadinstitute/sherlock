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
 * @interface V2controllersEditableCiRun
 */
export interface V2controllersEditableCiRun {
    /**
     * Always appends; will eliminate duplicates.
     * @type {Array<string>}
     * @memberof V2controllersEditableCiRun
     */
    appVersions?: Array<string>;
    /**
     * Always appends; will eliminate duplicates. Spreads to associated chart releases (and environments and clusters) and new app/chart versions.
     * @type {Array<string>}
     * @memberof V2controllersEditableCiRun
     */
    changesets?: Array<string>;
    /**
     * Always appends; will eliminate duplicates. Spreads to associated environments and clusters.
     * @type {Array<string>}
     * @memberof V2controllersEditableCiRun
     */
    chartReleases?: Array<string>;
    /**
     * Always appends; will eliminate duplicates.
     * @type {Array<string>}
     * @memberof V2controllersEditableCiRun
     */
    chartVersions?: Array<string>;
    /**
     * Always appends; will eliminate duplicates.
     * @type {Array<string>}
     * @memberof V2controllersEditableCiRun
     */
    charts?: Array<string>;
    /**
     * Always appends; will eliminate duplicates.
     * @type {Array<string>}
     * @memberof V2controllersEditableCiRun
     */
    clusters?: Array<string>;
    /**
     * Always appends; will eliminate duplicates.
     * @type {Array<string>}
     * @memberof V2controllersEditableCiRun
     */
    environments?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof V2controllersEditableCiRun
     */
    startedAt?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersEditableCiRun
     */
    status?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersEditableCiRun
     */
    terminalAt?: string;
}

/**
 * Check if a given object implements the V2controllersEditableCiRun interface.
 */
export function instanceOfV2controllersEditableCiRun(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersEditableCiRunFromJSON(json: any): V2controllersEditableCiRun {
    return V2controllersEditableCiRunFromJSONTyped(json, false);
}

export function V2controllersEditableCiRunFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersEditableCiRun {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'appVersions': !exists(json, 'appVersions') ? undefined : json['appVersions'],
        'changesets': !exists(json, 'changesets') ? undefined : json['changesets'],
        'chartReleases': !exists(json, 'chartReleases') ? undefined : json['chartReleases'],
        'chartVersions': !exists(json, 'chartVersions') ? undefined : json['chartVersions'],
        'charts': !exists(json, 'charts') ? undefined : json['charts'],
        'clusters': !exists(json, 'clusters') ? undefined : json['clusters'],
        'environments': !exists(json, 'environments') ? undefined : json['environments'],
        'startedAt': !exists(json, 'startedAt') ? undefined : json['startedAt'],
        'status': !exists(json, 'status') ? undefined : json['status'],
        'terminalAt': !exists(json, 'terminalAt') ? undefined : json['terminalAt'],
    };
}

export function V2controllersEditableCiRunToJSON(value?: V2controllersEditableCiRun | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'appVersions': value.appVersions,
        'changesets': value.changesets,
        'chartReleases': value.chartReleases,
        'chartVersions': value.chartVersions,
        'charts': value.charts,
        'clusters': value.clusters,
        'environments': value.environments,
        'startedAt': value.startedAt,
        'status': value.status,
        'terminalAt': value.terminalAt,
    };
}
