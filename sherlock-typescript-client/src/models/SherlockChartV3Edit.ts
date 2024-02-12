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
 * @interface SherlockChartV3Edit
 */
export interface SherlockChartV3Edit {
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3Edit
     */
    appImageGitMainBranch?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3Edit
     */
    appImageGitRepo?: string;
    /**
     * Indicates if the default subdomain, protocol, and port fields are relevant for this chart
     * @type {boolean}
     * @memberof SherlockChartV3Edit
     */
    chartExposesEndpoint?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3Edit
     */
    chartRepo?: string;
    /**
     * 
     * @type {number}
     * @memberof SherlockChartV3Edit
     */
    defaultPort?: number;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3Edit
     */
    defaultProtocol?: string;
    /**
     * When creating, will default to the name of the chart
     * @type {string}
     * @memberof SherlockChartV3Edit
     */
    defaultSubdomain?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3Edit
     */
    description?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockChartV3Edit
     */
    pactParticipant?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3Edit
     */
    playbookURL?: string;
}

/**
 * Check if a given object implements the SherlockChartV3Edit interface.
 */
export function instanceOfSherlockChartV3Edit(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SherlockChartV3EditFromJSON(json: any): SherlockChartV3Edit {
    return SherlockChartV3EditFromJSONTyped(json, false);
}

export function SherlockChartV3EditFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockChartV3Edit {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'appImageGitMainBranch': !exists(json, 'appImageGitMainBranch') ? undefined : json['appImageGitMainBranch'],
        'appImageGitRepo': !exists(json, 'appImageGitRepo') ? undefined : json['appImageGitRepo'],
        'chartExposesEndpoint': !exists(json, 'chartExposesEndpoint') ? undefined : json['chartExposesEndpoint'],
        'chartRepo': !exists(json, 'chartRepo') ? undefined : json['chartRepo'],
        'defaultPort': !exists(json, 'defaultPort') ? undefined : json['defaultPort'],
        'defaultProtocol': !exists(json, 'defaultProtocol') ? undefined : json['defaultProtocol'],
        'defaultSubdomain': !exists(json, 'defaultSubdomain') ? undefined : json['defaultSubdomain'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'pactParticipant': !exists(json, 'pactParticipant') ? undefined : json['pactParticipant'],
        'playbookURL': !exists(json, 'playbookURL') ? undefined : json['playbookURL'],
    };
}

export function SherlockChartV3EditToJSON(value?: SherlockChartV3Edit | null): any {
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
        'defaultPort': value.defaultPort,
        'defaultProtocol': value.defaultProtocol,
        'defaultSubdomain': value.defaultSubdomain,
        'description': value.description,
        'pactParticipant': value.pactParticipant,
        'playbookURL': value.playbookURL,
    };
}

