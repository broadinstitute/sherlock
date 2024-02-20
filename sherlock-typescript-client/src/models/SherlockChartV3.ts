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
import type { SherlockCiIdentifierV3 } from './SherlockCiIdentifierV3';
import {
    SherlockCiIdentifierV3FromJSON,
    SherlockCiIdentifierV3FromJSONTyped,
    SherlockCiIdentifierV3ToJSON,
} from './SherlockCiIdentifierV3';

/**
 * 
 * @export
 * @interface SherlockChartV3
 */
export interface SherlockChartV3 {
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3
     */
    appImageGitMainBranch?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3
     */
    appImageGitRepo?: string;
    /**
     * Indicates if the default subdomain, protocol, and port fields are relevant for this chart
     * @type {boolean}
     * @memberof SherlockChartV3
     */
    chartExposesEndpoint?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3
     */
    chartRepo?: string;
    /**
     * 
     * @type {SherlockCiIdentifierV3}
     * @memberof SherlockChartV3
     */
    ciIdentifier?: SherlockCiIdentifierV3;
    /**
     * 
     * @type {Date}
     * @memberof SherlockChartV3
     */
    createdAt?: Date;
    /**
     * 
     * @type {number}
     * @memberof SherlockChartV3
     */
    defaultPort?: number;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3
     */
    defaultProtocol?: string;
    /**
     * When creating, will default to the name of the chart
     * @type {string}
     * @memberof SherlockChartV3
     */
    defaultSubdomain?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3
     */
    description?: string;
    /**
     * 
     * @type {number}
     * @memberof SherlockChartV3
     */
    id?: number;
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockChartV3
     */
    name?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockChartV3
     */
    pactParticipant?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartV3
     */
    playbookURL?: string;
    /**
     * 
     * @type {Date}
     * @memberof SherlockChartV3
     */
    updatedAt?: Date;
}

/**
 * Check if a given object implements the SherlockChartV3 interface.
 */
export function instanceOfSherlockChartV3(value: object): boolean {
    return true;
}

export function SherlockChartV3FromJSON(json: any): SherlockChartV3 {
    return SherlockChartV3FromJSONTyped(json, false);
}

export function SherlockChartV3FromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockChartV3 {
    if (json === undefined || json === null) {
        return json;
    }
    return {
        
        'appImageGitMainBranch': !exists(json, 'appImageGitMainBranch') ? undefined : json['appImageGitMainBranch'],
        'appImageGitRepo': !exists(json, 'appImageGitRepo') ? undefined : json['appImageGitRepo'],
        'chartExposesEndpoint': !exists(json, 'chartExposesEndpoint') ? undefined : json['chartExposesEndpoint'],
        'chartRepo': !exists(json, 'chartRepo') ? undefined : json['chartRepo'],
        'ciIdentifier': !exists(json, 'ciIdentifier') ? undefined : SherlockCiIdentifierV3FromJSON(json['ciIdentifier']),
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'defaultPort': !exists(json, 'defaultPort') ? undefined : json['defaultPort'],
        'defaultProtocol': !exists(json, 'defaultProtocol') ? undefined : json['defaultProtocol'],
        'defaultSubdomain': !exists(json, 'defaultSubdomain') ? undefined : json['defaultSubdomain'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'pactParticipant': !exists(json, 'pactParticipant') ? undefined : json['pactParticipant'],
        'playbookURL': !exists(json, 'playbookURL') ? undefined : json['playbookURL'],
        'updatedAt': !exists(json, 'updatedAt') ? undefined : (new Date(json['updatedAt'])),
    };
}

export function SherlockChartV3ToJSON(value?: SherlockChartV3 | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'appImageGitMainBranch': value['appImageGitMainBranch'],
        'appImageGitRepo': value['appImageGitRepo'],
        'chartExposesEndpoint': value['chartExposesEndpoint'],
        'chartRepo': value['chartRepo'],
        'ciIdentifier': SherlockCiIdentifierV3ToJSON(value['ciIdentifier']),
        'createdAt': !exists(value, 'createdAt') ? undefined : ((value['createdAt']).toISOString()),
        'defaultPort': value['defaultPort'],
        'defaultProtocol': value['defaultProtocol'],
        'defaultSubdomain': value['defaultSubdomain'],
        'description': value['description'],
        'id': value['id'],
        'name': value['name'],
        'pactParticipant': value['pactParticipant'],
        'playbookURL': value['playbookURL'],
        'updatedAt': !exists(value, 'updatedAt') ? undefined : ((value['updatedAt']).toISOString()),
    };
}

