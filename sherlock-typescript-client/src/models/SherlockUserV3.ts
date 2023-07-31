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
 * @interface SherlockUserV3
 */
export interface SherlockUserV3 {
    /**
     * 
     * @type {Date}
     * @memberof SherlockUserV3
     */
    createdAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof SherlockUserV3
     */
    email?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockUserV3
     */
    githubID?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockUserV3
     */
    githubUsername?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockUserV3
     */
    googleID?: string;
    /**
     * 
     * @type {number}
     * @memberof SherlockUserV3
     */
    id?: number;
    /**
     * 
     * @type {string}
     * @memberof SherlockUserV3
     */
    name?: string;
    /**
     * Controls whether Sherlock should automatically update the user's name based on a connected GitHub identity.
     * Will be set to true if the user account has no name and a GitHub account is linked.
     * @type {boolean}
     * @memberof SherlockUserV3
     */
    nameInferredFromGithub?: boolean;
    /**
     * Available only in responses; describes the user's production-suitability
     * @type {string}
     * @memberof SherlockUserV3
     */
    suitabilityDescription?: string;
    /**
     * Available only in responses; indicates whether the user is production-suitable
     * @type {boolean}
     * @memberof SherlockUserV3
     */
    suitable?: boolean;
    /**
     * 
     * @type {Date}
     * @memberof SherlockUserV3
     */
    updatedAt?: Date;
}

/**
 * Check if a given object implements the SherlockUserV3 interface.
 */
export function instanceOfSherlockUserV3(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SherlockUserV3FromJSON(json: any): SherlockUserV3 {
    return SherlockUserV3FromJSONTyped(json, false);
}

export function SherlockUserV3FromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockUserV3 {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'email': !exists(json, 'email') ? undefined : json['email'],
        'githubID': !exists(json, 'githubID') ? undefined : json['githubID'],
        'githubUsername': !exists(json, 'githubUsername') ? undefined : json['githubUsername'],
        'googleID': !exists(json, 'googleID') ? undefined : json['googleID'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'nameInferredFromGithub': !exists(json, 'nameInferredFromGithub') ? undefined : json['nameInferredFromGithub'],
        'suitabilityDescription': !exists(json, 'suitabilityDescription') ? undefined : json['suitabilityDescription'],
        'suitable': !exists(json, 'suitable') ? undefined : json['suitable'],
        'updatedAt': !exists(json, 'updatedAt') ? undefined : (new Date(json['updatedAt'])),
    };
}

export function SherlockUserV3ToJSON(value?: SherlockUserV3 | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'email': value.email,
        'githubID': value.githubID,
        'githubUsername': value.githubUsername,
        'googleID': value.googleID,
        'id': value.id,
        'name': value.name,
        'nameInferredFromGithub': value.nameInferredFromGithub,
        'suitabilityDescription': value.suitabilityDescription,
        'suitable': value.suitable,
        'updatedAt': value.updatedAt === undefined ? undefined : (value.updatedAt.toISOString()),
    };
}
