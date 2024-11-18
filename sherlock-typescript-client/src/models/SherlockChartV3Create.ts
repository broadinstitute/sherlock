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
 * @interface SherlockChartV3Create
 */
export interface SherlockChartV3Create {
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3Create
     */
    appImageGitMainBranch?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3Create
     */
    appImageGitRepo?: string;
    /**
     * Indicates if the default subdomain, protocol, and port fields are relevant for this chart
     * @type {boolean}
     * @memberof SherlockChartV3Create
     */
    chartExposesEndpoint?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3Create
     */
    chartRepo?: string;
    /**
     * 
     * @type {number}
     * @memberof SherlockChartV3Create
     */
    defaultPort?: number;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3Create
     */
    defaultProtocol?: string;
    /**
     * When creating, will default to the name of the chart
     * @type {string}
     * @memberof SherlockChartV3Create
     */
    defaultSubdomain?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3Create
     */
    description?: string;
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockChartV3Create
     */
    name?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockChartV3Create
     */
    pactParticipant?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3Create
     */
    playbookURL?: string;
}

/**
 * Check if a given object implements the SherlockChartV3Create interface.
 */
export function instanceOfSherlockChartV3Create(value: object): value is SherlockChartV3Create {
    return true;
}

export function SherlockChartV3CreateFromJSON(json: any): SherlockChartV3Create {
    return SherlockChartV3CreateFromJSONTyped(json, false);
}

export function SherlockChartV3CreateFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockChartV3Create {
    if (json == null) {
        return json;
    }
    return {
        
        'appImageGitMainBranch': json['appImageGitMainBranch'] == null ? undefined : json['appImageGitMainBranch'],
        'appImageGitRepo': json['appImageGitRepo'] == null ? undefined : json['appImageGitRepo'],
        'chartExposesEndpoint': json['chartExposesEndpoint'] == null ? undefined : json['chartExposesEndpoint'],
        'chartRepo': json['chartRepo'] == null ? undefined : json['chartRepo'],
        'defaultPort': json['defaultPort'] == null ? undefined : json['defaultPort'],
        'defaultProtocol': json['defaultProtocol'] == null ? undefined : json['defaultProtocol'],
        'defaultSubdomain': json['defaultSubdomain'] == null ? undefined : json['defaultSubdomain'],
        'description': json['description'] == null ? undefined : json['description'],
        'name': json['name'] == null ? undefined : json['name'],
        'pactParticipant': json['pactParticipant'] == null ? undefined : json['pactParticipant'],
        'playbookURL': json['playbookURL'] == null ? undefined : json['playbookURL'],
    };
}

export function SherlockChartV3CreateToJSON(json: any): SherlockChartV3Create {
    return SherlockChartV3CreateToJSONTyped(json, false);
}

export function SherlockChartV3CreateToJSONTyped(value?: SherlockChartV3Create | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'appImageGitMainBranch': value['appImageGitMainBranch'],
        'appImageGitRepo': value['appImageGitRepo'],
        'chartExposesEndpoint': value['chartExposesEndpoint'],
        'chartRepo': value['chartRepo'],
        'defaultPort': value['defaultPort'],
        'defaultProtocol': value['defaultProtocol'],
        'defaultSubdomain': value['defaultSubdomain'],
        'description': value['description'],
        'name': value['name'],
        'pactParticipant': value['pactParticipant'],
        'playbookURL': value['playbookURL'],
    };
}

