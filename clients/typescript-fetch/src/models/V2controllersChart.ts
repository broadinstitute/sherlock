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
 * @interface V2controllersChart
 */
export interface V2controllersChart {
    /**
     * 
     * @type {string}
     * @memberof V2controllersChart
     */
    appImageGitMainBranch?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChart
     */
    appImageGitRepo?: string;
    /**
     * Indicates if the default subdomain, protocol, and port fields are relevant for this chart
     * @type {boolean}
     * @memberof V2controllersChart
     */
    chartExposesEndpoint?: boolean;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChart
     */
    chartRepo?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChart
     */
    createdAt?: string;
    /**
     * 
     * @type {number}
     * @memberof V2controllersChart
     */
    defaultPort?: number;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChart
     */
    defaultProtocol?: string;
    /**
     * When creating, will default to the name of the chart
     * @type {string}
     * @memberof V2controllersChart
     */
    defaultSubdomain?: string;
    /**
     * 
     * @type {number}
     * @memberof V2controllersChart
     */
    id?: number;
    /**
     * Indicates whether a chart requires config rendering from firecloud-develop
     * @type {boolean}
     * @memberof V2controllersChart
     */
    legacyConfigsEnabled?: boolean;
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersChart
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChart
     */
    updatedAt?: string;
}

/**
 * Check if a given object implements the V2controllersChart interface.
 */
export function instanceOfV2controllersChart(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersChartFromJSON(json: any): V2controllersChart {
    return V2controllersChartFromJSONTyped(json, false);
}

export function V2controllersChartFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersChart {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'appImageGitMainBranch': !exists(json, 'appImageGitMainBranch') ? undefined : json['appImageGitMainBranch'],
        'appImageGitRepo': !exists(json, 'appImageGitRepo') ? undefined : json['appImageGitRepo'],
        'chartExposesEndpoint': !exists(json, 'chartExposesEndpoint') ? undefined : json['chartExposesEndpoint'],
        'chartRepo': !exists(json, 'chartRepo') ? undefined : json['chartRepo'],
        'createdAt': !exists(json, 'createdAt') ? undefined : json['createdAt'],
        'defaultPort': !exists(json, 'defaultPort') ? undefined : json['defaultPort'],
        'defaultProtocol': !exists(json, 'defaultProtocol') ? undefined : json['defaultProtocol'],
        'defaultSubdomain': !exists(json, 'defaultSubdomain') ? undefined : json['defaultSubdomain'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'legacyConfigsEnabled': !exists(json, 'legacyConfigsEnabled') ? undefined : json['legacyConfigsEnabled'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'updatedAt': !exists(json, 'updatedAt') ? undefined : json['updatedAt'],
    };
}

export function V2controllersChartToJSON(value?: V2controllersChart | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'appImageGitMainBranch': value.appImageGitMainBranch,
        'appImageGitRepo': value.appImageGitRepo,
        'chartExposesEndpoint': value.chartExposesEndpoint,
        'chartRepo': value.chartRepo,
        'createdAt': value.createdAt,
        'defaultPort': value.defaultPort,
        'defaultProtocol': value.defaultProtocol,
        'defaultSubdomain': value.defaultSubdomain,
        'id': value.id,
        'legacyConfigsEnabled': value.legacyConfigsEnabled,
        'name': value.name,
        'updatedAt': value.updatedAt,
    };
}

